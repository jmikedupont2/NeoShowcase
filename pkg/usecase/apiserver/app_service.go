package apiserver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"

	"github.com/traPtitech/neoshowcase/pkg/domain"
	"github.com/traPtitech/neoshowcase/pkg/domain/web"
	"github.com/traPtitech/neoshowcase/pkg/util/ds"
	"github.com/traPtitech/neoshowcase/pkg/util/optional"
	"github.com/traPtitech/neoshowcase/pkg/util/random"
)

func (s *Service) validateApp(ctx context.Context, app *domain.Application) error {
	// Validate app fields and conflict
	existingApps, err := s.appRepo.GetApplications(ctx, domain.GetApplicationCondition{})
	if err != nil {
		return errors.Wrap(err, "getting existing applications")
	}
	si, err := s.systemInfo(ctx)
	if err != nil {
		return errors.Wrap(err, "getting system info")
	}
	err = app.Validate(web.GetUser(ctx), existingApps, si.AvailableDomains, si.AvailablePorts)
	if err != nil {
		return newError(ErrorTypeBadRequest, "invalid application", err)
	}

	// Validate repository owner
	user := web.GetUser(ctx)
	repo, err := s.gitRepo.GetRepository(ctx, app.RepositoryID)
	if err != nil {
		return err
	}
	if !repo.CanCreateApp(user) {
		return newError(ErrorTypeBadRequest, "you cannot create application from this repository", nil)
	}

	// Validate ref by making request
	refMap, err := repo.ResolveRefs(ctx, s.fallbackKey)
	if err != nil {
		return newError(ErrorTypeBadRequest, "cannot fetch repository, check auth setting", err)
	}
	if _, ok := refMap[app.RefName]; !ok {
		return newError(ErrorTypeBadRequest, fmt.Sprintf("ref %v not found", app.RefName), nil)
	}
	return nil
}

func (s *Service) CreateApplication(ctx context.Context, app *domain.Application) (*domain.Application, error) {
	// Validate
	err := s.validateApp(ctx, app)
	if err != nil {
		return nil, err
	}

	// Create
	err = s.appRepo.CreateApplication(ctx, app)
	if err != nil {
		return nil, err
	}

	err = s.createApplicationDatabase(ctx, app)
	if err != nil {
		return nil, err
	}

	err = s.controller.FetchRepository(ctx, app.RepositoryID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request to fetch repository")
	}

	return handleRepoError(s.appRepo.GetApplication(ctx, app.ID))
}

func (s *Service) createApplicationDatabase(ctx context.Context, app *domain.Application) error {
	dbName := domain.DBName(app.ID)

	if app.Config.BuildConfig.MariaDB() {
		host, port := s.mariaDBManager.GetHost()
		dbPassword := random.SecureGeneratePassword(32)
		dbSetting := domain.CreateArgs{
			Database: dbName,
			Password: dbPassword,
		}
		err := s.mariaDBManager.Create(ctx, dbSetting)
		if err != nil {
			return err
		}

		envs := []*domain.Environment{
			{ApplicationID: app.ID, Key: domain.EnvMariaDBHostnameKey, Value: host, System: true},
			{ApplicationID: app.ID, Key: domain.EnvMariaDBPortKey, Value: strconv.Itoa(port), System: true},
			{ApplicationID: app.ID, Key: domain.EnvMariaDBUserKey, Value: dbName, System: true},
			{ApplicationID: app.ID, Key: domain.EnvMariaDBPasswordKey, Value: dbPassword, System: true},
			{ApplicationID: app.ID, Key: domain.EnvMariaDBDatabaseKey, Value: dbName, System: true},
		}
		for _, env := range envs {
			err = s.envRepo.SetEnv(ctx, env)
			if err != nil {
				return err
			}
		}
	}

	if app.Config.BuildConfig.MongoDB() {
		host, port := s.mongoDBManager.GetHost()
		dbPassword := random.SecureGeneratePassword(32)
		dbSetting := domain.CreateArgs{
			Database: dbName,
			Password: dbPassword,
		}
		err := s.mongoDBManager.Create(ctx, dbSetting)
		if err != nil {
			return err
		}

		envs := []*domain.Environment{
			{ApplicationID: app.ID, Key: domain.EnvMongoDBHostnameKey, Value: host, System: true},
			{ApplicationID: app.ID, Key: domain.EnvMongoDBPortKey, Value: strconv.Itoa(port), System: true},
			{ApplicationID: app.ID, Key: domain.EnvMongoDBUserKey, Value: dbName, System: true},
			{ApplicationID: app.ID, Key: domain.EnvMongoDBPasswordKey, Value: dbPassword, System: true},
			{ApplicationID: app.ID, Key: domain.EnvMongoDBDatabaseKey, Value: dbName, System: true},
		}
		for _, env := range envs {
			err = s.envRepo.SetEnv(ctx, env)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Service) deleteApplicationDatabase(ctx context.Context, app *domain.Application, envs []*domain.Environment) error {
	if app.Config.BuildConfig.MariaDB() {
		dbKey, ok := lo.Find(envs, func(e *domain.Environment) bool { return e.Key == domain.EnvMariaDBDatabaseKey })
		if !ok {
			return errors.New("failed to find mariadb name from env key")
		}
		err := s.mariaDBManager.Delete(ctx, domain.DeleteArgs{Database: dbKey.Value})
		if err != nil {
			return err
		}
	}

	if app.Config.BuildConfig.MongoDB() {
		dbKey, ok := lo.Find(envs, func(e *domain.Environment) bool { return e.Key == domain.EnvMongoDBDatabaseKey })
		if !ok {
			return errors.New("failed to find mongodb name from env key")
		}
		err := s.mongoDBManager.Delete(ctx, domain.DeleteArgs{Database: dbKey.Value})
		if err != nil {
			return err
		}
	}

	return nil
}

type TopAppInfo struct {
	App         *domain.Application
	LatestBuild *domain.Build
}

func (s *Service) GetApplications(ctx context.Context, all bool) ([]*TopAppInfo, error) {
	// Fetch apps
	var cond domain.GetApplicationCondition
	if !all {
		cond.UserID = optional.From(web.GetUser(ctx).ID)
	}
	apps, err := s.appRepo.GetApplications(ctx, cond)
	if err != nil {
		return nil, err
	}

	// Fetch latest builds
	appIDs := ds.Map(apps, func(app *domain.Application) string { return app.ID })
	builds, err := s.buildRepo.GetLatestBuilds(ctx, appIDs)
	if err != nil {
		return nil, err
	}
	buildsMap := lo.SliceToMap(builds, func(b *domain.Build) (string, *domain.Build) { return b.ApplicationID, b })

	// Construct
	return ds.Map(apps, func(app *domain.Application) *TopAppInfo {
		return &TopAppInfo{
			App:         app,
			LatestBuild: buildsMap[app.ID],
		}
	}), nil
}

func (s *Service) GetApplication(ctx context.Context, id string) (*TopAppInfo, error) {
	// Fetch app
	app, err := s.appRepo.GetApplication(ctx, id)
	if err != nil {
		return nil, err
	}
	// Fetch latest build
	builds, err := s.buildRepo.GetLatestBuilds(ctx, []string{app.ID})
	if err != nil {
		return nil, err
	}
	// Construct
	info := &TopAppInfo{
		App: app,
	}
	if len(builds) > 0 {
		info.LatestBuild = builds[0]
	}
	return info, nil
}

func (s *Service) UpdateApplication(ctx context.Context, id string, args *domain.UpdateApplicationArgs) error {
	err := s.isApplicationOwner(ctx, id)
	if err != nil {
		return err
	}

	app, err := s.appRepo.GetApplication(ctx, id)
	if err != nil {
		return err
	}
	app.Apply(args)

	// Validate
	if err = s.validateApp(ctx, app); err != nil {
		return err
	}
	// Validate immutable fields
	{
		appBefore, err := s.appRepo.GetApplication(ctx, id)
		if err != nil {
			return err
		}
		if appBefore.Config.BuildConfig.MariaDB() != app.Config.BuildConfig.MariaDB() {
			return newError(ErrorTypeBadRequest, "use_mariadb is immutable", nil)
		}
		if appBefore.Config.BuildConfig.MongoDB() != app.Config.BuildConfig.MongoDB() {
			return newError(ErrorTypeBadRequest, "use_mongodb is immutable", nil)
		}
	}

	// Update
	err = s.appRepo.UpdateApplication(ctx, id, args)
	if err != nil {
		return errors.Wrap(err, "updating application")
	}

	// Sync
	err = s.controller.FetchRepository(ctx, app.RepositoryID)
	if err != nil {
		return errors.Wrap(err, "requesting fetch repository")
	}
	err = s.controller.SyncDeployments(ctx)
	if err != nil {
		return errors.Wrap(err, "requesting sync deployments")
	}

	return nil
}

func (s *Service) DeleteApplication(ctx context.Context, id string) error {
	err := s.isApplicationOwner(ctx, id)
	if err != nil {
		return err
	}

	app, err := s.appRepo.GetApplication(ctx, id)
	if err != nil {
		return err
	}
	if app.Running {
		return newError(ErrorTypeBadRequest, "stop the application first before deleting", nil)
	}

	env, err := s.envRepo.GetEnv(ctx, domain.GetEnvCondition{ApplicationID: optional.From(id)})
	if err != nil {
		return err
	}
	err = s.deleteApplicationDatabase(ctx, app, env)
	if err != nil {
		return err
	}

	// delete artifacts
	artifacts, err := s.artifactRepo.GetArtifacts(ctx, domain.GetArtifactCondition{ApplicationID: optional.From(app.ID)})
	if err != nil {
		return err
	}
	for _, artifact := range artifacts {
		if artifact.DeletedAt.Valid {
			continue
		}
		err = domain.DeleteArtifact(s.storage, artifact.ID)
		if err != nil {
			log.Errorf("failed to delete artifact: %+v", err) // fail-safe
		}
	}
	err = s.artifactRepo.HardDeleteArtifacts(ctx, domain.GetArtifactCondition{ApplicationID: optional.From(app.ID)})
	if err != nil {
		return err
	}
	// delete builds
	builds, err := s.buildRepo.GetBuilds(ctx, domain.GetBuildCondition{ApplicationID: optional.From(app.ID)})
	if err != nil {
		return err
	}
	for _, build := range builds {
		err = domain.DeleteBuildLog(s.storage, build.ID)
		if err != nil {
			log.Errorf("failed to delete build log: %+v", err) // fail-safe
		}
	}
	err = s.buildRepo.DeleteBuilds(ctx, domain.GetBuildCondition{ApplicationID: optional.From(app.ID)})
	if err != nil {
		return err
	}
	// delete environments
	err = s.envRepo.DeleteEnv(ctx, domain.GetEnvCondition{ApplicationID: optional.From(app.ID)})
	if err != nil {
		return err
	}
	// delete websites, owners, application
	err = s.appRepo.DeleteApplication(ctx, app.ID)
	if err != nil {
		return err
	}

	return nil
}
