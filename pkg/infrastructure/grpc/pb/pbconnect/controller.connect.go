// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: neoshowcase/protobuf/controller.proto

package pbconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	pb "github.com/traPtitech/neoshowcase/pkg/infrastructure/grpc/pb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ControllerServiceName is the fully-qualified name of the ControllerService service.
	ControllerServiceName = "neoshowcase.protobuf.ControllerService"
	// ControllerBuilderServiceName is the fully-qualified name of the ControllerBuilderService service.
	ControllerBuilderServiceName = "neoshowcase.protobuf.ControllerBuilderService"
	// ControllerSSGenServiceName is the fully-qualified name of the ControllerSSGenService service.
	ControllerSSGenServiceName = "neoshowcase.protobuf.ControllerSSGenService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ControllerServiceGetSystemInfoProcedure is the fully-qualified name of the ControllerService's
	// GetSystemInfo RPC.
	ControllerServiceGetSystemInfoProcedure = "/neoshowcase.protobuf.ControllerService/GetSystemInfo"
	// ControllerServiceFetchRepositoryProcedure is the fully-qualified name of the ControllerService's
	// FetchRepository RPC.
	ControllerServiceFetchRepositoryProcedure = "/neoshowcase.protobuf.ControllerService/FetchRepository"
	// ControllerServiceRegisterBuildProcedure is the fully-qualified name of the ControllerService's
	// RegisterBuild RPC.
	ControllerServiceRegisterBuildProcedure = "/neoshowcase.protobuf.ControllerService/RegisterBuild"
	// ControllerServiceSyncDeploymentsProcedure is the fully-qualified name of the ControllerService's
	// SyncDeployments RPC.
	ControllerServiceSyncDeploymentsProcedure = "/neoshowcase.protobuf.ControllerService/SyncDeployments"
	// ControllerServiceStreamBuildLogProcedure is the fully-qualified name of the ControllerService's
	// StreamBuildLog RPC.
	ControllerServiceStreamBuildLogProcedure = "/neoshowcase.protobuf.ControllerService/StreamBuildLog"
	// ControllerServiceCancelBuildProcedure is the fully-qualified name of the ControllerService's
	// CancelBuild RPC.
	ControllerServiceCancelBuildProcedure = "/neoshowcase.protobuf.ControllerService/CancelBuild"
	// ControllerBuilderServiceConnectBuilderProcedure is the fully-qualified name of the
	// ControllerBuilderService's ConnectBuilder RPC.
	ControllerBuilderServiceConnectBuilderProcedure = "/neoshowcase.protobuf.ControllerBuilderService/ConnectBuilder"
	// ControllerSSGenServiceConnectSSGenProcedure is the fully-qualified name of the
	// ControllerSSGenService's ConnectSSGen RPC.
	ControllerSSGenServiceConnectSSGenProcedure = "/neoshowcase.protobuf.ControllerSSGenService/ConnectSSGen"
)

// ControllerServiceClient is a client for the neoshowcase.protobuf.ControllerService service.
type ControllerServiceClient interface {
	GetSystemInfo(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[pb.SystemInfo], error)
	FetchRepository(context.Context, *connect_go.Request[pb.RepositoryIdRequest]) (*connect_go.Response[emptypb.Empty], error)
	RegisterBuild(context.Context, *connect_go.Request[pb.ApplicationIdRequest]) (*connect_go.Response[emptypb.Empty], error)
	SyncDeployments(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error)
	StreamBuildLog(context.Context, *connect_go.Request[pb.BuildIdRequest]) (*connect_go.ServerStreamForClient[pb.BuildLog], error)
	CancelBuild(context.Context, *connect_go.Request[pb.BuildIdRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewControllerServiceClient constructs a client for the neoshowcase.protobuf.ControllerService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewControllerServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ControllerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &controllerServiceClient{
		getSystemInfo: connect_go.NewClient[emptypb.Empty, pb.SystemInfo](
			httpClient,
			baseURL+ControllerServiceGetSystemInfoProcedure,
			opts...,
		),
		fetchRepository: connect_go.NewClient[pb.RepositoryIdRequest, emptypb.Empty](
			httpClient,
			baseURL+ControllerServiceFetchRepositoryProcedure,
			opts...,
		),
		registerBuild: connect_go.NewClient[pb.ApplicationIdRequest, emptypb.Empty](
			httpClient,
			baseURL+ControllerServiceRegisterBuildProcedure,
			opts...,
		),
		syncDeployments: connect_go.NewClient[emptypb.Empty, emptypb.Empty](
			httpClient,
			baseURL+ControllerServiceSyncDeploymentsProcedure,
			opts...,
		),
		streamBuildLog: connect_go.NewClient[pb.BuildIdRequest, pb.BuildLog](
			httpClient,
			baseURL+ControllerServiceStreamBuildLogProcedure,
			opts...,
		),
		cancelBuild: connect_go.NewClient[pb.BuildIdRequest, emptypb.Empty](
			httpClient,
			baseURL+ControllerServiceCancelBuildProcedure,
			opts...,
		),
	}
}

// controllerServiceClient implements ControllerServiceClient.
type controllerServiceClient struct {
	getSystemInfo   *connect_go.Client[emptypb.Empty, pb.SystemInfo]
	fetchRepository *connect_go.Client[pb.RepositoryIdRequest, emptypb.Empty]
	registerBuild   *connect_go.Client[pb.ApplicationIdRequest, emptypb.Empty]
	syncDeployments *connect_go.Client[emptypb.Empty, emptypb.Empty]
	streamBuildLog  *connect_go.Client[pb.BuildIdRequest, pb.BuildLog]
	cancelBuild     *connect_go.Client[pb.BuildIdRequest, emptypb.Empty]
}

// GetSystemInfo calls neoshowcase.protobuf.ControllerService.GetSystemInfo.
func (c *controllerServiceClient) GetSystemInfo(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[pb.SystemInfo], error) {
	return c.getSystemInfo.CallUnary(ctx, req)
}

// FetchRepository calls neoshowcase.protobuf.ControllerService.FetchRepository.
func (c *controllerServiceClient) FetchRepository(ctx context.Context, req *connect_go.Request[pb.RepositoryIdRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.fetchRepository.CallUnary(ctx, req)
}

// RegisterBuild calls neoshowcase.protobuf.ControllerService.RegisterBuild.
func (c *controllerServiceClient) RegisterBuild(ctx context.Context, req *connect_go.Request[pb.ApplicationIdRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.registerBuild.CallUnary(ctx, req)
}

// SyncDeployments calls neoshowcase.protobuf.ControllerService.SyncDeployments.
func (c *controllerServiceClient) SyncDeployments(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error) {
	return c.syncDeployments.CallUnary(ctx, req)
}

// StreamBuildLog calls neoshowcase.protobuf.ControllerService.StreamBuildLog.
func (c *controllerServiceClient) StreamBuildLog(ctx context.Context, req *connect_go.Request[pb.BuildIdRequest]) (*connect_go.ServerStreamForClient[pb.BuildLog], error) {
	return c.streamBuildLog.CallServerStream(ctx, req)
}

// CancelBuild calls neoshowcase.protobuf.ControllerService.CancelBuild.
func (c *controllerServiceClient) CancelBuild(ctx context.Context, req *connect_go.Request[pb.BuildIdRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.cancelBuild.CallUnary(ctx, req)
}

// ControllerServiceHandler is an implementation of the neoshowcase.protobuf.ControllerService
// service.
type ControllerServiceHandler interface {
	GetSystemInfo(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[pb.SystemInfo], error)
	FetchRepository(context.Context, *connect_go.Request[pb.RepositoryIdRequest]) (*connect_go.Response[emptypb.Empty], error)
	RegisterBuild(context.Context, *connect_go.Request[pb.ApplicationIdRequest]) (*connect_go.Response[emptypb.Empty], error)
	SyncDeployments(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error)
	StreamBuildLog(context.Context, *connect_go.Request[pb.BuildIdRequest], *connect_go.ServerStream[pb.BuildLog]) error
	CancelBuild(context.Context, *connect_go.Request[pb.BuildIdRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewControllerServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewControllerServiceHandler(svc ControllerServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ControllerServiceGetSystemInfoProcedure, connect_go.NewUnaryHandler(
		ControllerServiceGetSystemInfoProcedure,
		svc.GetSystemInfo,
		opts...,
	))
	mux.Handle(ControllerServiceFetchRepositoryProcedure, connect_go.NewUnaryHandler(
		ControllerServiceFetchRepositoryProcedure,
		svc.FetchRepository,
		opts...,
	))
	mux.Handle(ControllerServiceRegisterBuildProcedure, connect_go.NewUnaryHandler(
		ControllerServiceRegisterBuildProcedure,
		svc.RegisterBuild,
		opts...,
	))
	mux.Handle(ControllerServiceSyncDeploymentsProcedure, connect_go.NewUnaryHandler(
		ControllerServiceSyncDeploymentsProcedure,
		svc.SyncDeployments,
		opts...,
	))
	mux.Handle(ControllerServiceStreamBuildLogProcedure, connect_go.NewServerStreamHandler(
		ControllerServiceStreamBuildLogProcedure,
		svc.StreamBuildLog,
		opts...,
	))
	mux.Handle(ControllerServiceCancelBuildProcedure, connect_go.NewUnaryHandler(
		ControllerServiceCancelBuildProcedure,
		svc.CancelBuild,
		opts...,
	))
	return "/neoshowcase.protobuf.ControllerService/", mux
}

// UnimplementedControllerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedControllerServiceHandler struct{}

func (UnimplementedControllerServiceHandler) GetSystemInfo(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[pb.SystemInfo], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neoshowcase.protobuf.ControllerService.GetSystemInfo is not implemented"))
}

func (UnimplementedControllerServiceHandler) FetchRepository(context.Context, *connect_go.Request[pb.RepositoryIdRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neoshowcase.protobuf.ControllerService.FetchRepository is not implemented"))
}

func (UnimplementedControllerServiceHandler) RegisterBuild(context.Context, *connect_go.Request[pb.ApplicationIdRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neoshowcase.protobuf.ControllerService.RegisterBuild is not implemented"))
}

func (UnimplementedControllerServiceHandler) SyncDeployments(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neoshowcase.protobuf.ControllerService.SyncDeployments is not implemented"))
}

func (UnimplementedControllerServiceHandler) StreamBuildLog(context.Context, *connect_go.Request[pb.BuildIdRequest], *connect_go.ServerStream[pb.BuildLog]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neoshowcase.protobuf.ControllerService.StreamBuildLog is not implemented"))
}

func (UnimplementedControllerServiceHandler) CancelBuild(context.Context, *connect_go.Request[pb.BuildIdRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neoshowcase.protobuf.ControllerService.CancelBuild is not implemented"))
}

// ControllerBuilderServiceClient is a client for the neoshowcase.protobuf.ControllerBuilderService
// service.
type ControllerBuilderServiceClient interface {
	ConnectBuilder(context.Context) *connect_go.BidiStreamForClient[pb.BuilderResponse, pb.BuilderRequest]
}

// NewControllerBuilderServiceClient constructs a client for the
// neoshowcase.protobuf.ControllerBuilderService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewControllerBuilderServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ControllerBuilderServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &controllerBuilderServiceClient{
		connectBuilder: connect_go.NewClient[pb.BuilderResponse, pb.BuilderRequest](
			httpClient,
			baseURL+ControllerBuilderServiceConnectBuilderProcedure,
			opts...,
		),
	}
}

// controllerBuilderServiceClient implements ControllerBuilderServiceClient.
type controllerBuilderServiceClient struct {
	connectBuilder *connect_go.Client[pb.BuilderResponse, pb.BuilderRequest]
}

// ConnectBuilder calls neoshowcase.protobuf.ControllerBuilderService.ConnectBuilder.
func (c *controllerBuilderServiceClient) ConnectBuilder(ctx context.Context) *connect_go.BidiStreamForClient[pb.BuilderResponse, pb.BuilderRequest] {
	return c.connectBuilder.CallBidiStream(ctx)
}

// ControllerBuilderServiceHandler is an implementation of the
// neoshowcase.protobuf.ControllerBuilderService service.
type ControllerBuilderServiceHandler interface {
	ConnectBuilder(context.Context, *connect_go.BidiStream[pb.BuilderResponse, pb.BuilderRequest]) error
}

// NewControllerBuilderServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewControllerBuilderServiceHandler(svc ControllerBuilderServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ControllerBuilderServiceConnectBuilderProcedure, connect_go.NewBidiStreamHandler(
		ControllerBuilderServiceConnectBuilderProcedure,
		svc.ConnectBuilder,
		opts...,
	))
	return "/neoshowcase.protobuf.ControllerBuilderService/", mux
}

// UnimplementedControllerBuilderServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedControllerBuilderServiceHandler struct{}

func (UnimplementedControllerBuilderServiceHandler) ConnectBuilder(context.Context, *connect_go.BidiStream[pb.BuilderResponse, pb.BuilderRequest]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neoshowcase.protobuf.ControllerBuilderService.ConnectBuilder is not implemented"))
}

// ControllerSSGenServiceClient is a client for the neoshowcase.protobuf.ControllerSSGenService
// service.
type ControllerSSGenServiceClient interface {
	ConnectSSGen(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.ServerStreamForClient[pb.SSGenRequest], error)
}

// NewControllerSSGenServiceClient constructs a client for the
// neoshowcase.protobuf.ControllerSSGenService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewControllerSSGenServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ControllerSSGenServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &controllerSSGenServiceClient{
		connectSSGen: connect_go.NewClient[emptypb.Empty, pb.SSGenRequest](
			httpClient,
			baseURL+ControllerSSGenServiceConnectSSGenProcedure,
			opts...,
		),
	}
}

// controllerSSGenServiceClient implements ControllerSSGenServiceClient.
type controllerSSGenServiceClient struct {
	connectSSGen *connect_go.Client[emptypb.Empty, pb.SSGenRequest]
}

// ConnectSSGen calls neoshowcase.protobuf.ControllerSSGenService.ConnectSSGen.
func (c *controllerSSGenServiceClient) ConnectSSGen(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.ServerStreamForClient[pb.SSGenRequest], error) {
	return c.connectSSGen.CallServerStream(ctx, req)
}

// ControllerSSGenServiceHandler is an implementation of the
// neoshowcase.protobuf.ControllerSSGenService service.
type ControllerSSGenServiceHandler interface {
	ConnectSSGen(context.Context, *connect_go.Request[emptypb.Empty], *connect_go.ServerStream[pb.SSGenRequest]) error
}

// NewControllerSSGenServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewControllerSSGenServiceHandler(svc ControllerSSGenServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ControllerSSGenServiceConnectSSGenProcedure, connect_go.NewServerStreamHandler(
		ControllerSSGenServiceConnectSSGenProcedure,
		svc.ConnectSSGen,
		opts...,
	))
	return "/neoshowcase.protobuf.ControllerSSGenService/", mux
}

// UnimplementedControllerSSGenServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedControllerSSGenServiceHandler struct{}

func (UnimplementedControllerSSGenServiceHandler) ConnectSSGen(context.Context, *connect_go.Request[emptypb.Empty], *connect_go.ServerStream[pb.SSGenRequest]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("neoshowcase.protobuf.ControllerSSGenService.ConnectSSGen is not implemented"))
}
