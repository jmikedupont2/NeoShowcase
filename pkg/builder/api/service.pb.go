// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: neoshowcase/services/builder/service.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BuilderStatus int32

const (
	BuilderStatus_UNKNOWN     BuilderStatus = 0
	BuilderStatus_UNAVAILABLE BuilderStatus = 1
	BuilderStatus_WAITING     BuilderStatus = 2
	BuilderStatus_BUILDING    BuilderStatus = 3
)

// Enum value maps for BuilderStatus.
var (
	BuilderStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "UNAVAILABLE",
		2: "WAITING",
		3: "BUILDING",
	}
	BuilderStatus_value = map[string]int32{
		"UNKNOWN":     0,
		"UNAVAILABLE": 1,
		"WAITING":     2,
		"BUILDING":    3,
	}
)

func (x BuilderStatus) Enum() *BuilderStatus {
	p := new(BuilderStatus)
	*p = x
	return p
}

func (x BuilderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuilderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_neoshowcase_services_builder_service_proto_enumTypes[0].Descriptor()
}

func (BuilderStatus) Type() protoreflect.EnumType {
	return &file_neoshowcase_services_builder_service_proto_enumTypes[0]
}

func (x BuilderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuilderStatus.Descriptor instead.
func (BuilderStatus) EnumDescriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{0}
}

type Event_Type int32

const (
	Event_CONNECTED       Event_Type = 0
	Event_BUILD_STARTED   Event_Type = 1
	Event_BUILD_SUCCEEDED Event_Type = 2
	Event_BUILD_FAILED    Event_Type = 3
	Event_BUILD_CANCELED  Event_Type = 4
)

// Enum value maps for Event_Type.
var (
	Event_Type_name = map[int32]string{
		0: "CONNECTED",
		1: "BUILD_STARTED",
		2: "BUILD_SUCCEEDED",
		3: "BUILD_FAILED",
		4: "BUILD_CANCELED",
	}
	Event_Type_value = map[string]int32{
		"CONNECTED":       0,
		"BUILD_STARTED":   1,
		"BUILD_SUCCEEDED": 2,
		"BUILD_FAILED":    3,
		"BUILD_CANCELED":  4,
	}
)

func (x Event_Type) Enum() *Event_Type {
	p := new(Event_Type)
	*p = x
	return p
}

func (x Event_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_neoshowcase_services_builder_service_proto_enumTypes[1].Descriptor()
}

func (Event_Type) Type() protoreflect.EnumType {
	return &file_neoshowcase_services_builder_service_proto_enumTypes[1]
}

func (x Event_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_Type.Descriptor instead.
func (Event_Type) EnumDescriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{8, 0}
}

type GetStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  BuilderStatus `protobuf:"varint,1,opt,name=status,proto3,enum=neoshowcase.proto.services.builder.BuilderStatus" json:"status,omitempty"`
	BuildId string        `protobuf:"bytes,2,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *GetStatusResponse) Reset() {
	*x = GetStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusResponse) ProtoMessage() {}

func (x *GetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusResponse.ProtoReflect.Descriptor instead.
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetStatusResponse) GetStatus() BuilderStatus {
	if x != nil {
		return x.Status
	}
	return BuilderStatus_UNKNOWN
}

func (x *GetStatusResponse) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type BuildSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepositoryUrl string `protobuf:"bytes,1,opt,name=repository_url,json=repositoryUrl,proto3" json:"repository_url,omitempty"`
}

func (x *BuildSource) Reset() {
	*x = BuildSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildSource) ProtoMessage() {}

func (x *BuildSource) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildSource.ProtoReflect.Descriptor instead.
func (*BuildSource) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{1}
}

func (x *BuildSource) GetRepositoryUrl() string {
	if x != nil {
		return x.RepositoryUrl
	}
	return ""
}

type BuildOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseImageName string `protobuf:"bytes,1,opt,name=base_image_name,json=baseImageName,proto3" json:"base_image_name,omitempty"`
	EntrypointCmd string `protobuf:"bytes,2,opt,name=entrypoint_cmd,json=entrypointCmd,proto3" json:"entrypoint_cmd,omitempty"`
	StartupCmd    string `protobuf:"bytes,3,opt,name=startup_cmd,json=startupCmd,proto3" json:"startup_cmd,omitempty"`
	ArtifactPath  string `protobuf:"bytes,4,opt,name=artifact_path,json=artifactPath,proto3" json:"artifact_path,omitempty"`
}

func (x *BuildOptions) Reset() {
	*x = BuildOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildOptions) ProtoMessage() {}

func (x *BuildOptions) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildOptions.ProtoReflect.Descriptor instead.
func (*BuildOptions) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{2}
}

func (x *BuildOptions) GetBaseImageName() string {
	if x != nil {
		return x.BaseImageName
	}
	return ""
}

func (x *BuildOptions) GetEntrypointCmd() string {
	if x != nil {
		return x.EntrypointCmd
	}
	return ""
}

func (x *BuildOptions) GetStartupCmd() string {
	if x != nil {
		return x.StartupCmd
	}
	return ""
}

func (x *BuildOptions) GetArtifactPath() string {
	if x != nil {
		return x.ArtifactPath
	}
	return ""
}

type StartBuildImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageName     string        `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	Source        *BuildSource  `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Options       *BuildOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	ApplicationId string        `protobuf:"bytes,4,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
}

func (x *StartBuildImageRequest) Reset() {
	*x = StartBuildImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartBuildImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBuildImageRequest) ProtoMessage() {}

func (x *StartBuildImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBuildImageRequest.ProtoReflect.Descriptor instead.
func (*StartBuildImageRequest) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{3}
}

func (x *StartBuildImageRequest) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *StartBuildImageRequest) GetSource() *BuildSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *StartBuildImageRequest) GetOptions() *BuildOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *StartBuildImageRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

type StartBuildImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *StartBuildImageResponse) Reset() {
	*x = StartBuildImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartBuildImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBuildImageResponse) ProtoMessage() {}

func (x *StartBuildImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBuildImageResponse.ProtoReflect.Descriptor instead.
func (*StartBuildImageResponse) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{4}
}

func (x *StartBuildImageResponse) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type StartBuildStaticRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source        *BuildSource  `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Options       *BuildOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	ApplicationId string        `protobuf:"bytes,3,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
}

func (x *StartBuildStaticRequest) Reset() {
	*x = StartBuildStaticRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartBuildStaticRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBuildStaticRequest) ProtoMessage() {}

func (x *StartBuildStaticRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBuildStaticRequest.ProtoReflect.Descriptor instead.
func (*StartBuildStaticRequest) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{5}
}

func (x *StartBuildStaticRequest) GetSource() *BuildSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *StartBuildStaticRequest) GetOptions() *BuildOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *StartBuildStaticRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

type StartBuildStaticResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *StartBuildStaticResponse) Reset() {
	*x = StartBuildStaticResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartBuildStaticResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBuildStaticResponse) ProtoMessage() {}

func (x *StartBuildStaticResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBuildStaticResponse.ProtoReflect.Descriptor instead.
func (*StartBuildStaticResponse) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{6}
}

func (x *StartBuildStaticResponse) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type CancelTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Canceled bool   `protobuf:"varint,1,opt,name=canceled,proto3" json:"canceled,omitempty"`
	BuildId  string `protobuf:"bytes,2,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *CancelTaskResponse) Reset() {
	*x = CancelTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTaskResponse) ProtoMessage() {}

func (x *CancelTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTaskResponse.ProtoReflect.Descriptor instead.
func (*CancelTaskResponse) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{7}
}

func (x *CancelTaskResponse) GetCanceled() bool {
	if x != nil {
		return x.Canceled
	}
	return false
}

func (x *CancelTaskResponse) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Event_Type `protobuf:"varint,1,opt,name=type,proto3,enum=neoshowcase.proto.services.builder.Event_Type" json:"type,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{8}
}

func (x *Event) GetType() Event_Type {
	if x != nil {
		return x.Type
	}
	return Event_CONNECTED
}

var File_neoshowcase_services_builder_service_proto protoreflect.FileDescriptor

var file_neoshowcase_services_builder_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x6e, 0x65,
	0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x31, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x0b, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x72, 0x6c, 0x22, 0xa3,
	0x01, 0x0a, 0x0c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6d, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x5f, 0x63, 0x6d, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x43, 0x6d, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x50, 0x61, 0x74, 0x68, 0x22, 0xf3, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x47,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68,
	0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x17, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64,
	0x22, 0xd5, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e,
	0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22,
	0x4b, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0xb0, 0x01, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x63, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x55, 0x49,
	0x4c, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x42,
	0x55, 0x49, 0x4c, 0x44, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x2a,
	0x48, 0x0a, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x42,
	0x55, 0x49, 0x4c, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x32, 0xc2, 0x04, 0x0a, 0x0e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x35, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x30, 0x01, 0x12, 0x8a, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f,
	0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x8d, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x12, 0x3b, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5c, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x36, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61,
	0x50, 0x74, 0x69, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_neoshowcase_services_builder_service_proto_rawDescOnce sync.Once
	file_neoshowcase_services_builder_service_proto_rawDescData = file_neoshowcase_services_builder_service_proto_rawDesc
)

func file_neoshowcase_services_builder_service_proto_rawDescGZIP() []byte {
	file_neoshowcase_services_builder_service_proto_rawDescOnce.Do(func() {
		file_neoshowcase_services_builder_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_neoshowcase_services_builder_service_proto_rawDescData)
	})
	return file_neoshowcase_services_builder_service_proto_rawDescData
}

var file_neoshowcase_services_builder_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_neoshowcase_services_builder_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_neoshowcase_services_builder_service_proto_goTypes = []interface{}{
	(BuilderStatus)(0),               // 0: neoshowcase.proto.services.builder.BuilderStatus
	(Event_Type)(0),                  // 1: neoshowcase.proto.services.builder.Event.Type
	(*GetStatusResponse)(nil),        // 2: neoshowcase.proto.services.builder.GetStatusResponse
	(*BuildSource)(nil),              // 3: neoshowcase.proto.services.builder.BuildSource
	(*BuildOptions)(nil),             // 4: neoshowcase.proto.services.builder.BuildOptions
	(*StartBuildImageRequest)(nil),   // 5: neoshowcase.proto.services.builder.StartBuildImageRequest
	(*StartBuildImageResponse)(nil),  // 6: neoshowcase.proto.services.builder.StartBuildImageResponse
	(*StartBuildStaticRequest)(nil),  // 7: neoshowcase.proto.services.builder.StartBuildStaticRequest
	(*StartBuildStaticResponse)(nil), // 8: neoshowcase.proto.services.builder.StartBuildStaticResponse
	(*CancelTaskResponse)(nil),       // 9: neoshowcase.proto.services.builder.CancelTaskResponse
	(*Event)(nil),                    // 10: neoshowcase.proto.services.builder.Event
	(*emptypb.Empty)(nil),            // 11: google.protobuf.Empty
}
var file_neoshowcase_services_builder_service_proto_depIdxs = []int32{
	0,  // 0: neoshowcase.proto.services.builder.GetStatusResponse.status:type_name -> neoshowcase.proto.services.builder.BuilderStatus
	3,  // 1: neoshowcase.proto.services.builder.StartBuildImageRequest.source:type_name -> neoshowcase.proto.services.builder.BuildSource
	4,  // 2: neoshowcase.proto.services.builder.StartBuildImageRequest.options:type_name -> neoshowcase.proto.services.builder.BuildOptions
	3,  // 3: neoshowcase.proto.services.builder.StartBuildStaticRequest.source:type_name -> neoshowcase.proto.services.builder.BuildSource
	4,  // 4: neoshowcase.proto.services.builder.StartBuildStaticRequest.options:type_name -> neoshowcase.proto.services.builder.BuildOptions
	1,  // 5: neoshowcase.proto.services.builder.Event.type:type_name -> neoshowcase.proto.services.builder.Event.Type
	11, // 6: neoshowcase.proto.services.builder.BuilderService.GetStatus:input_type -> google.protobuf.Empty
	11, // 7: neoshowcase.proto.services.builder.BuilderService.ConnectEventStream:input_type -> google.protobuf.Empty
	5,  // 8: neoshowcase.proto.services.builder.BuilderService.StartBuildImage:input_type -> neoshowcase.proto.services.builder.StartBuildImageRequest
	7,  // 9: neoshowcase.proto.services.builder.BuilderService.StartBuildStatic:input_type -> neoshowcase.proto.services.builder.StartBuildStaticRequest
	11, // 10: neoshowcase.proto.services.builder.BuilderService.CancelTask:input_type -> google.protobuf.Empty
	2,  // 11: neoshowcase.proto.services.builder.BuilderService.GetStatus:output_type -> neoshowcase.proto.services.builder.GetStatusResponse
	10, // 12: neoshowcase.proto.services.builder.BuilderService.ConnectEventStream:output_type -> neoshowcase.proto.services.builder.Event
	6,  // 13: neoshowcase.proto.services.builder.BuilderService.StartBuildImage:output_type -> neoshowcase.proto.services.builder.StartBuildImageResponse
	8,  // 14: neoshowcase.proto.services.builder.BuilderService.StartBuildStatic:output_type -> neoshowcase.proto.services.builder.StartBuildStaticResponse
	9,  // 15: neoshowcase.proto.services.builder.BuilderService.CancelTask:output_type -> neoshowcase.proto.services.builder.CancelTaskResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_neoshowcase_services_builder_service_proto_init() }
func file_neoshowcase_services_builder_service_proto_init() {
	if File_neoshowcase_services_builder_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_neoshowcase_services_builder_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_neoshowcase_services_builder_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildSource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_neoshowcase_services_builder_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_neoshowcase_services_builder_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartBuildImageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_neoshowcase_services_builder_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartBuildImageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_neoshowcase_services_builder_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartBuildStaticRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_neoshowcase_services_builder_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartBuildStaticResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_neoshowcase_services_builder_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTaskResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_neoshowcase_services_builder_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_neoshowcase_services_builder_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_neoshowcase_services_builder_service_proto_goTypes,
		DependencyIndexes: file_neoshowcase_services_builder_service_proto_depIdxs,
		EnumInfos:         file_neoshowcase_services_builder_service_proto_enumTypes,
		MessageInfos:      file_neoshowcase_services_builder_service_proto_msgTypes,
	}.Build()
	File_neoshowcase_services_builder_service_proto = out.File
	file_neoshowcase_services_builder_service_proto_rawDesc = nil
	file_neoshowcase_services_builder_service_proto_goTypes = nil
	file_neoshowcase_services_builder_service_proto_depIdxs = nil
}
