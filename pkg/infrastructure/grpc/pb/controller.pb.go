// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: neoshowcase/protobuf/controller.proto

package pb

import (
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

type BuilderRequest_Type int32

const (
	BuilderRequest_START_BUILD  BuilderRequest_Type = 0
	BuilderRequest_CANCEL_BUILD BuilderRequest_Type = 1
)

// Enum value maps for BuilderRequest_Type.
var (
	BuilderRequest_Type_name = map[int32]string{
		0: "START_BUILD",
		1: "CANCEL_BUILD",
	}
	BuilderRequest_Type_value = map[string]int32{
		"START_BUILD":  0,
		"CANCEL_BUILD": 1,
	}
)

func (x BuilderRequest_Type) Enum() *BuilderRequest_Type {
	p := new(BuilderRequest_Type)
	*p = x
	return p
}

func (x BuilderRequest_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuilderRequest_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_neoshowcase_protobuf_controller_proto_enumTypes[0].Descriptor()
}

func (BuilderRequest_Type) Type() protoreflect.EnumType {
	return &file_neoshowcase_protobuf_controller_proto_enumTypes[0]
}

func (x BuilderRequest_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuilderRequest_Type.Descriptor instead.
func (BuilderRequest_Type) EnumDescriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_controller_proto_rawDescGZIP(), []int{1, 0}
}

type BuildSettled_Reason int32

const (
	BuildSettled_SUCCESS   BuildSettled_Reason = 0
	BuildSettled_FAILED    BuildSettled_Reason = 1
	BuildSettled_CANCELLED BuildSettled_Reason = 2
)

// Enum value maps for BuildSettled_Reason.
var (
	BuildSettled_Reason_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
		2: "CANCELLED",
	}
	BuildSettled_Reason_value = map[string]int32{
		"SUCCESS":   0,
		"FAILED":    1,
		"CANCELLED": 2,
	}
)

func (x BuildSettled_Reason) Enum() *BuildSettled_Reason {
	p := new(BuildSettled_Reason)
	*p = x
	return p
}

func (x BuildSettled_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuildSettled_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_neoshowcase_protobuf_controller_proto_enumTypes[1].Descriptor()
}

func (BuildSettled_Reason) Type() protoreflect.EnumType {
	return &file_neoshowcase_protobuf_controller_proto_enumTypes[1]
}

func (x BuildSettled_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuildSettled_Reason.Descriptor instead.
func (BuildSettled_Reason) EnumDescriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_controller_proto_rawDescGZIP(), []int{3, 0}
}

type BuilderResponse_Type int32

const (
	BuilderResponse_CONNECTED     BuilderResponse_Type = 0
	BuilderResponse_BUILD_STARTED BuilderResponse_Type = 1
	BuilderResponse_BUILD_SETTLED BuilderResponse_Type = 2
	BuilderResponse_BUILD_LOG     BuilderResponse_Type = 3
)

// Enum value maps for BuilderResponse_Type.
var (
	BuilderResponse_Type_name = map[int32]string{
		0: "CONNECTED",
		1: "BUILD_STARTED",
		2: "BUILD_SETTLED",
		3: "BUILD_LOG",
	}
	BuilderResponse_Type_value = map[string]int32{
		"CONNECTED":     0,
		"BUILD_STARTED": 1,
		"BUILD_SETTLED": 2,
		"BUILD_LOG":     3,
	}
)

func (x BuilderResponse_Type) Enum() *BuilderResponse_Type {
	p := new(BuilderResponse_Type)
	*p = x
	return p
}

func (x BuilderResponse_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuilderResponse_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_neoshowcase_protobuf_controller_proto_enumTypes[2].Descriptor()
}

func (BuilderResponse_Type) Type() protoreflect.EnumType {
	return &file_neoshowcase_protobuf_controller_proto_enumTypes[2]
}

func (x BuilderResponse_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuilderResponse_Type.Descriptor instead.
func (BuilderResponse_Type) EnumDescriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_controller_proto_rawDescGZIP(), []int{5, 0}
}

type SSGenRequest_Type int32

const (
	SSGenRequest_RELOAD SSGenRequest_Type = 0
)

// Enum value maps for SSGenRequest_Type.
var (
	SSGenRequest_Type_name = map[int32]string{
		0: "RELOAD",
	}
	SSGenRequest_Type_value = map[string]int32{
		"RELOAD": 0,
	}
)

func (x SSGenRequest_Type) Enum() *SSGenRequest_Type {
	p := new(SSGenRequest_Type)
	*p = x
	return p
}

func (x SSGenRequest_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SSGenRequest_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_neoshowcase_protobuf_controller_proto_enumTypes[3].Descriptor()
}

func (SSGenRequest_Type) Type() protoreflect.EnumType {
	return &file_neoshowcase_protobuf_controller_proto_enumTypes[3]
}

func (x SSGenRequest_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SSGenRequest_Type.Descriptor instead.
func (SSGenRequest_Type) EnumDescriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_controller_proto_rawDescGZIP(), []int{6, 0}
}

type StartBuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *StartBuildRequest) Reset() {
	*x = StartBuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartBuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBuildRequest) ProtoMessage() {}

func (x *StartBuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBuildRequest.ProtoReflect.Descriptor instead.
func (*StartBuildRequest) Descriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_controller_proto_rawDescGZIP(), []int{0}
}

func (x *StartBuildRequest) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type BuilderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type BuilderRequest_Type `protobuf:"varint,1,opt,name=type,proto3,enum=neoshowcase.protobuf.BuilderRequest_Type" json:"type,omitempty"`
	// Types that are assignable to Body:
	//
	//	*BuilderRequest_StartBuild
	//	*BuilderRequest_CancelBuild
	Body isBuilderRequest_Body `protobuf_oneof:"body"`
}

func (x *BuilderRequest) Reset() {
	*x = BuilderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuilderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuilderRequest) ProtoMessage() {}

func (x *BuilderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuilderRequest.ProtoReflect.Descriptor instead.
func (*BuilderRequest) Descriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_controller_proto_rawDescGZIP(), []int{1}
}

func (x *BuilderRequest) GetType() BuilderRequest_Type {
	if x != nil {
		return x.Type
	}
	return BuilderRequest_START_BUILD
}

func (m *BuilderRequest) GetBody() isBuilderRequest_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *BuilderRequest) GetStartBuild() *StartBuildRequest {
	if x, ok := x.GetBody().(*BuilderRequest_StartBuild); ok {
		return x.StartBuild
	}
	return nil
}

func (x *BuilderRequest) GetCancelBuild() *BuildIdRequest {
	if x, ok := x.GetBody().(*BuilderRequest_CancelBuild); ok {
		return x.CancelBuild
	}
	return nil
}

type isBuilderRequest_Body interface {
	isBuilderRequest_Body()
}

type BuilderRequest_StartBuild struct {
	StartBuild *StartBuildRequest `protobuf:"bytes,2,opt,name=start_build,json=startBuild,proto3,oneof"`
}

type BuilderRequest_CancelBuild struct {
	CancelBuild *BuildIdRequest `protobuf:"bytes,3,opt,name=cancel_build,json=cancelBuild,proto3,oneof"`
}

func (*BuilderRequest_StartBuild) isBuilderRequest_Body() {}

func (*BuilderRequest_CancelBuild) isBuilderRequest_Body() {}

type BuildStarted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *BuildStarted) Reset() {
	*x = BuildStarted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildStarted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildStarted) ProtoMessage() {}

func (x *BuildStarted) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildStarted.ProtoReflect.Descriptor instead.
func (*BuildStarted) Descriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_controller_proto_rawDescGZIP(), []int{2}
}

func (x *BuildStarted) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type BuildSettled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string              `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	Reason  BuildSettled_Reason `protobuf:"varint,2,opt,name=reason,proto3,enum=neoshowcase.protobuf.BuildSettled_Reason" json:"reason,omitempty"`
}

func (x *BuildSettled) Reset() {
	*x = BuildSettled{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildSettled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildSettled) ProtoMessage() {}

func (x *BuildSettled) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildSettled.ProtoReflect.Descriptor instead.
func (*BuildSettled) Descriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_controller_proto_rawDescGZIP(), []int{3}
}

func (x *BuildSettled) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

func (x *BuildSettled) GetReason() BuildSettled_Reason {
	if x != nil {
		return x.Reason
	}
	return BuildSettled_SUCCESS
}

type BuildLogPortion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	Log     []byte `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *BuildLogPortion) Reset() {
	*x = BuildLogPortion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildLogPortion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildLogPortion) ProtoMessage() {}

func (x *BuildLogPortion) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildLogPortion.ProtoReflect.Descriptor instead.
func (*BuildLogPortion) Descriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_controller_proto_rawDescGZIP(), []int{4}
}

func (x *BuildLogPortion) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

func (x *BuildLogPortion) GetLog() []byte {
	if x != nil {
		return x.Log
	}
	return nil
}

type BuilderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type BuilderResponse_Type `protobuf:"varint,1,opt,name=type,proto3,enum=neoshowcase.protobuf.BuilderResponse_Type" json:"type,omitempty"`
	// Types that are assignable to Body:
	//
	//	*BuilderResponse_Started
	//	*BuilderResponse_Settled
	//	*BuilderResponse_Log
	Body isBuilderResponse_Body `protobuf_oneof:"body"`
}

func (x *BuilderResponse) Reset() {
	*x = BuilderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuilderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuilderResponse) ProtoMessage() {}

func (x *BuilderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuilderResponse.ProtoReflect.Descriptor instead.
func (*BuilderResponse) Descriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_controller_proto_rawDescGZIP(), []int{5}
}

func (x *BuilderResponse) GetType() BuilderResponse_Type {
	if x != nil {
		return x.Type
	}
	return BuilderResponse_CONNECTED
}

func (m *BuilderResponse) GetBody() isBuilderResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *BuilderResponse) GetStarted() *BuildStarted {
	if x, ok := x.GetBody().(*BuilderResponse_Started); ok {
		return x.Started
	}
	return nil
}

func (x *BuilderResponse) GetSettled() *BuildSettled {
	if x, ok := x.GetBody().(*BuilderResponse_Settled); ok {
		return x.Settled
	}
	return nil
}

func (x *BuilderResponse) GetLog() *BuildLogPortion {
	if x, ok := x.GetBody().(*BuilderResponse_Log); ok {
		return x.Log
	}
	return nil
}

type isBuilderResponse_Body interface {
	isBuilderResponse_Body()
}

type BuilderResponse_Started struct {
	Started *BuildStarted `protobuf:"bytes,2,opt,name=started,proto3,oneof"`
}

type BuilderResponse_Settled struct {
	Settled *BuildSettled `protobuf:"bytes,3,opt,name=settled,proto3,oneof"`
}

type BuilderResponse_Log struct {
	Log *BuildLogPortion `protobuf:"bytes,4,opt,name=log,proto3,oneof"`
}

func (*BuilderResponse_Started) isBuilderResponse_Body() {}

func (*BuilderResponse_Settled) isBuilderResponse_Body() {}

func (*BuilderResponse_Log) isBuilderResponse_Body() {}

type SSGenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type SSGenRequest_Type `protobuf:"varint,1,opt,name=type,proto3,enum=neoshowcase.protobuf.SSGenRequest_Type" json:"type,omitempty"`
}

func (x *SSGenRequest) Reset() {
	*x = SSGenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSGenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSGenRequest) ProtoMessage() {}

func (x *SSGenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_protobuf_controller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSGenRequest.ProtoReflect.Descriptor instead.
func (*SSGenRequest) Descriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_controller_proto_rawDescGZIP(), []int{6}
}

func (x *SSGenRequest) GetType() SSGenRequest_Type {
	if x != nil {
		return x.Type
	}
	return SSGenRequest_RELOAD
}

var File_neoshowcase_protobuf_controller_proto protoreflect.FileDescriptor

var file_neoshowcase_protobuf_controller_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6e, 0x65, 0x6f, 0x73,
	0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e,
	0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x99,
	0x02, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x29, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x4a, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x49, 0x0a, 0x0c,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x22, 0x29, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44,
	0x10, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x29, 0x0a, 0x0c, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x9e, 0x01, 0x0a, 0x0c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x64, 0x12, 0x41, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x64, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22, 0x3e, 0x0a, 0x0f, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4c,
	0x6f, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0xe0, 0x02, 0x0a, 0x0f, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68,
	0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65,
	0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x07, 0x73, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65,
	0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x07, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x03, 0x6c, 0x6f,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f,
	0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0x4a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x4c, 0x4f, 0x47, 0x10,
	0x03, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x5f, 0x0a, 0x0c, 0x53, 0x53, 0x47,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f,
	0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x53, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x52, 0x45, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x00, 0x32, 0xf3, 0x03, 0x0a, 0x11, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x6e, 0x65, 0x6f, 0x73,
	0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x54, 0x0a, 0x0f, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x29,
	0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x53, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x12, 0x2a, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x58, 0x0a, 0x0e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x24, 0x2e, 0x6e, 0x65,
	0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x6f,
	0x67, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x12, 0x24, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0x7d, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x25,
	0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x24, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x01, 0x30, 0x01, 0x32,
	0x66, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x53, 0x47,
	0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x53, 0x53, 0x47, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x22, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x53, 0x47, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x30, 0x01, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x50, 0x74, 0x69, 0x74, 0x65, 0x63, 0x68,
	0x2f, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_neoshowcase_protobuf_controller_proto_rawDescOnce sync.Once
	file_neoshowcase_protobuf_controller_proto_rawDescData = file_neoshowcase_protobuf_controller_proto_rawDesc
)

func file_neoshowcase_protobuf_controller_proto_rawDescGZIP() []byte {
	file_neoshowcase_protobuf_controller_proto_rawDescOnce.Do(func() {
		file_neoshowcase_protobuf_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_neoshowcase_protobuf_controller_proto_rawDescData)
	})
	return file_neoshowcase_protobuf_controller_proto_rawDescData
}

var file_neoshowcase_protobuf_controller_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_neoshowcase_protobuf_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_neoshowcase_protobuf_controller_proto_goTypes = []interface{}{
	(BuilderRequest_Type)(0),     // 0: neoshowcase.protobuf.BuilderRequest.Type
	(BuildSettled_Reason)(0),     // 1: neoshowcase.protobuf.BuildSettled.Reason
	(BuilderResponse_Type)(0),    // 2: neoshowcase.protobuf.BuilderResponse.Type
	(SSGenRequest_Type)(0),       // 3: neoshowcase.protobuf.SSGenRequest.Type
	(*StartBuildRequest)(nil),    // 4: neoshowcase.protobuf.StartBuildRequest
	(*BuilderRequest)(nil),       // 5: neoshowcase.protobuf.BuilderRequest
	(*BuildStarted)(nil),         // 6: neoshowcase.protobuf.BuildStarted
	(*BuildSettled)(nil),         // 7: neoshowcase.protobuf.BuildSettled
	(*BuildLogPortion)(nil),      // 8: neoshowcase.protobuf.BuildLogPortion
	(*BuilderResponse)(nil),      // 9: neoshowcase.protobuf.BuilderResponse
	(*SSGenRequest)(nil),         // 10: neoshowcase.protobuf.SSGenRequest
	(*BuildIdRequest)(nil),       // 11: neoshowcase.protobuf.BuildIdRequest
	(*emptypb.Empty)(nil),        // 12: google.protobuf.Empty
	(*RepositoryIdRequest)(nil),  // 13: neoshowcase.protobuf.RepositoryIdRequest
	(*ApplicationIdRequest)(nil), // 14: neoshowcase.protobuf.ApplicationIdRequest
	(*SystemInfo)(nil),           // 15: neoshowcase.protobuf.SystemInfo
	(*BuildLog)(nil),             // 16: neoshowcase.protobuf.BuildLog
}
var file_neoshowcase_protobuf_controller_proto_depIdxs = []int32{
	0,  // 0: neoshowcase.protobuf.BuilderRequest.type:type_name -> neoshowcase.protobuf.BuilderRequest.Type
	4,  // 1: neoshowcase.protobuf.BuilderRequest.start_build:type_name -> neoshowcase.protobuf.StartBuildRequest
	11, // 2: neoshowcase.protobuf.BuilderRequest.cancel_build:type_name -> neoshowcase.protobuf.BuildIdRequest
	1,  // 3: neoshowcase.protobuf.BuildSettled.reason:type_name -> neoshowcase.protobuf.BuildSettled.Reason
	2,  // 4: neoshowcase.protobuf.BuilderResponse.type:type_name -> neoshowcase.protobuf.BuilderResponse.Type
	6,  // 5: neoshowcase.protobuf.BuilderResponse.started:type_name -> neoshowcase.protobuf.BuildStarted
	7,  // 6: neoshowcase.protobuf.BuilderResponse.settled:type_name -> neoshowcase.protobuf.BuildSettled
	8,  // 7: neoshowcase.protobuf.BuilderResponse.log:type_name -> neoshowcase.protobuf.BuildLogPortion
	3,  // 8: neoshowcase.protobuf.SSGenRequest.type:type_name -> neoshowcase.protobuf.SSGenRequest.Type
	12, // 9: neoshowcase.protobuf.ControllerService.GetSystemInfo:input_type -> google.protobuf.Empty
	13, // 10: neoshowcase.protobuf.ControllerService.FetchRepository:input_type -> neoshowcase.protobuf.RepositoryIdRequest
	14, // 11: neoshowcase.protobuf.ControllerService.RegisterBuild:input_type -> neoshowcase.protobuf.ApplicationIdRequest
	12, // 12: neoshowcase.protobuf.ControllerService.SyncDeployments:input_type -> google.protobuf.Empty
	11, // 13: neoshowcase.protobuf.ControllerService.StreamBuildLog:input_type -> neoshowcase.protobuf.BuildIdRequest
	11, // 14: neoshowcase.protobuf.ControllerService.CancelBuild:input_type -> neoshowcase.protobuf.BuildIdRequest
	9,  // 15: neoshowcase.protobuf.ControllerBuilderService.ConnectBuilder:input_type -> neoshowcase.protobuf.BuilderResponse
	12, // 16: neoshowcase.protobuf.ControllerSSGenService.ConnectSSGen:input_type -> google.protobuf.Empty
	15, // 17: neoshowcase.protobuf.ControllerService.GetSystemInfo:output_type -> neoshowcase.protobuf.SystemInfo
	12, // 18: neoshowcase.protobuf.ControllerService.FetchRepository:output_type -> google.protobuf.Empty
	12, // 19: neoshowcase.protobuf.ControllerService.RegisterBuild:output_type -> google.protobuf.Empty
	12, // 20: neoshowcase.protobuf.ControllerService.SyncDeployments:output_type -> google.protobuf.Empty
	16, // 21: neoshowcase.protobuf.ControllerService.StreamBuildLog:output_type -> neoshowcase.protobuf.BuildLog
	12, // 22: neoshowcase.protobuf.ControllerService.CancelBuild:output_type -> google.protobuf.Empty
	5,  // 23: neoshowcase.protobuf.ControllerBuilderService.ConnectBuilder:output_type -> neoshowcase.protobuf.BuilderRequest
	10, // 24: neoshowcase.protobuf.ControllerSSGenService.ConnectSSGen:output_type -> neoshowcase.protobuf.SSGenRequest
	17, // [17:25] is the sub-list for method output_type
	9,  // [9:17] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_neoshowcase_protobuf_controller_proto_init() }
func file_neoshowcase_protobuf_controller_proto_init() {
	if File_neoshowcase_protobuf_controller_proto != nil {
		return
	}
	file_neoshowcase_protobuf_gateway_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_neoshowcase_protobuf_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartBuildRequest); i {
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
		file_neoshowcase_protobuf_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuilderRequest); i {
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
		file_neoshowcase_protobuf_controller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildStarted); i {
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
		file_neoshowcase_protobuf_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildSettled); i {
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
		file_neoshowcase_protobuf_controller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildLogPortion); i {
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
		file_neoshowcase_protobuf_controller_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuilderResponse); i {
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
		file_neoshowcase_protobuf_controller_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSGenRequest); i {
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
	file_neoshowcase_protobuf_controller_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*BuilderRequest_StartBuild)(nil),
		(*BuilderRequest_CancelBuild)(nil),
	}
	file_neoshowcase_protobuf_controller_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*BuilderResponse_Started)(nil),
		(*BuilderResponse_Settled)(nil),
		(*BuilderResponse_Log)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_neoshowcase_protobuf_controller_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_neoshowcase_protobuf_controller_proto_goTypes,
		DependencyIndexes: file_neoshowcase_protobuf_controller_proto_depIdxs,
		EnumInfos:         file_neoshowcase_protobuf_controller_proto_enumTypes,
		MessageInfos:      file_neoshowcase_protobuf_controller_proto_msgTypes,
	}.Build()
	File_neoshowcase_protobuf_controller_proto = out.File
	file_neoshowcase_protobuf_controller_proto_rawDesc = nil
	file_neoshowcase_protobuf_controller_proto_goTypes = nil
	file_neoshowcase_protobuf_controller_proto_depIdxs = nil
}
