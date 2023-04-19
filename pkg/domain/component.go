package domain

import (
	"context"

	"github.com/traPtitech/neoshowcase/pkg/interface/grpc/pb"
	"github.com/traPtitech/neoshowcase/pkg/interface/grpc/pb/pbconnect"
)

type ControllerServiceClient interface {
	FetchRepository(ctx context.Context, repositoryID string) error
	RegisterBuilds(ctx context.Context) error
	SyncDeployments(ctx context.Context) error
	StreamBuildLog(ctx context.Context, buildID string) (<-chan *pb.BuildLog, error)
	CancelBuild(ctx context.Context, buildID string) error
}

type ControllerBuilderService interface {
	pbconnect.ControllerBuilderServiceHandler
	ListenBuilderIdle() (sub <-chan struct{}, unsub func())
	ListenBuildSettled() (sub <-chan struct{}, unsub func())
	BroadcastBuilder(req *pb.BuilderRequest)
}

type ControllerBuilderServiceClient interface {
	ConnectBuilder(ctx context.Context, onRequest func(req *pb.BuilderRequest), response <-chan *pb.BuilderResponse) error
}

type ControllerSSGenService interface {
	pbconnect.ControllerSSGenServiceHandler
	BroadcastSSGen(req *pb.SSGenRequest)
}

type ControllerSSGenServiceClient interface {
	ConnectSSGen(ctx context.Context, onRequest func(req *pb.SSGenRequest)) error
}
