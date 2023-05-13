package usecase

import (
	"context"

	"github.com/samber/lo"

	"github.com/traPtitech/neoshowcase/pkg/domain"
	"github.com/traPtitech/neoshowcase/pkg/domain/builder"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/grpc/pb"
	"github.com/traPtitech/neoshowcase/pkg/util/ds"
	"github.com/traPtitech/neoshowcase/pkg/util/optional"
)

type AppDeployHelper struct {
	backend   domain.Backend
	appRepo   domain.ApplicationRepository
	buildRepo domain.BuildRepository
	envRepo   domain.EnvironmentRepository
	ssgen     domain.ControllerSSGenService
	image     builder.ImageConfig
}

func NewAppDeployHelper(
	backend domain.Backend,
	appRepo domain.ApplicationRepository,
	buildRepo domain.BuildRepository,
	envRepo domain.EnvironmentRepository,
	ssgen domain.ControllerSSGenService,
	imageConfig builder.ImageConfig,
) *AppDeployHelper {
	return &AppDeployHelper{
		backend:   backend,
		appRepo:   appRepo,
		buildRepo: buildRepo,
		envRepo:   envRepo,
		ssgen:     ssgen,
		image:     imageConfig,
	}
}

func (s *AppDeployHelper) _getEnv(ctx context.Context, apps []*domain.Application) (map[string]map[string]string, error) {
	appIDs := ds.Map(apps, func(app *domain.Application) string { return app.ID })
	envs, err := s.envRepo.GetEnv(ctx, domain.GetEnvCondition{ApplicationIDIn: optional.From(appIDs)})
	if err != nil {
		return nil, err
	}
	ret := make(map[string]map[string]string, len(appIDs))
	for _, env := range envs {
		if _, ok := ret[env.ApplicationID]; !ok {
			ret[env.ApplicationID] = make(map[string]string)
		}
		ret[env.ApplicationID][env.Key] = env.Value
	}
	return ret, nil
}

func (s *AppDeployHelper) _runtimeDesiredStates(ctx context.Context) ([]*domain.RuntimeDesiredState, error) {
	// Get all 'running' state applications
	apps, err := s.appRepo.GetApplications(ctx, domain.GetApplicationCondition{
		DeployType: optional.From(domain.DeployTypeRuntime),
		Running:    optional.From(true),
	})
	if err != nil {
		return nil, err
	}

	// Calculate deploy-able applications
	commits := lo.SliceToMap(apps, func(app *domain.Application) (string, struct{}) { return app.CurrentCommit, struct{}{} })
	builds, err := s.buildRepo.GetBuilds(ctx, domain.GetBuildCondition{
		CommitIn: optional.From(lo.Keys(commits)),
		Status:   optional.From(domain.BuildStatusSucceeded),
	})
	if err != nil {
		return nil, err
	}
	buildExists := lo.SliceToMap(builds, func(b *domain.Build) (string, bool) { return b.ApplicationID + b.Commit, true })
	syncableApps := lo.Filter(apps, func(app *domain.Application, _ int) bool { return buildExists[app.ID+app.WantCommit] })

	envs, err := s._getEnv(ctx, syncableApps)
	if err != nil {
		return nil, err
	}
	desiredStates := ds.Map(syncableApps, func(app *domain.Application) *domain.RuntimeDesiredState {
		return &domain.RuntimeDesiredState{
			App:       app,
			ImageName: s.image.FullImageName(app.ID),
			ImageTag:  app.CurrentCommit,
			Envs:      envs[app.ID],
		}
	})
	return desiredStates, nil
}

func (s *AppDeployHelper) synchronize(ctx context.Context) error {
	var st domain.DesiredState
	var err error
	st.Runtime, err = s._runtimeDesiredStates(ctx)
	if err != nil {
		return err
	}
	st.StaticSites, err = domain.GetActiveStaticSites(ctx, s.appRepo, s.buildRepo)
	if err != nil {
		return err
	}

	// Synchronize
	s.ssgen.BroadcastSSGen(&pb.SSGenRequest{Type: pb.SSGenRequest_RELOAD})
	return s.backend.Synchronize(ctx, &st)
}
