package pbconvert

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/traPtitech/neoshowcase/pkg/domain"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/grpc/pb"
	"github.com/traPtitech/neoshowcase/pkg/util/ds"
	"github.com/traPtitech/neoshowcase/pkg/util/mapper"
)

var BuildStatusMapper = mapper.MustNewValueMapper(map[domain.BuildStatus]pb.Build_BuildStatus{
	domain.BuildStatusQueued:    pb.Build_QUEUED,
	domain.BuildStatusBuilding:  pb.Build_BUILDING,
	domain.BuildStatusSucceeded: pb.Build_SUCCEEDED,
	domain.BuildStatusFailed:    pb.Build_FAILED,
	domain.BuildStatusCanceled:  pb.Build_CANCELLED,
	domain.BuildStatusSkipped:   pb.Build_SKIPPED,
})

func ToPBBuild(build *domain.Build) *pb.Build {
	b := &pb.Build{
		Id:            build.ID,
		ApplicationId: build.ApplicationID,
		Commit:        build.Commit,
		Status:        BuildStatusMapper.IntoMust(build.Status),
		QueuedAt:      timestamppb.New(build.QueuedAt),
		StartedAt:     ToPBNullTimestamp(build.StartedAt),
		UpdatedAt:     ToPBNullTimestamp(build.UpdatedAt),
		FinishedAt:    ToPBNullTimestamp(build.FinishedAt),
		Retriable:     build.Retriable,
		Artifacts:     ds.Map(build.Artifacts, ToPBArtifact),
	}
	return b
}
