// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/flux-grpc/flux.proto

package flux_grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubmitResponse_ResultType int32

const (
	SubmitResponse_SUBMIT_SUCCESS SubmitResponse_ResultType = 0
	SubmitResponse_SUBMIT_ERROR   SubmitResponse_ResultType = 1
	SubmitResponse_SUBMIT_DENIED  SubmitResponse_ResultType = 2
)

// Enum value maps for SubmitResponse_ResultType.
var (
	SubmitResponse_ResultType_name = map[int32]string{
		0: "SUBMIT_SUCCESS",
		1: "SUBMIT_ERROR",
		2: "SUBMIT_DENIED",
	}
	SubmitResponse_ResultType_value = map[string]int32{
		"SUBMIT_SUCCESS": 0,
		"SUBMIT_ERROR":   1,
		"SUBMIT_DENIED":  2,
	}
)

func (x SubmitResponse_ResultType) Enum() *SubmitResponse_ResultType {
	p := new(SubmitResponse_ResultType)
	*p = x
	return p
}

func (x SubmitResponse_ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubmitResponse_ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_flux_grpc_flux_proto_enumTypes[0].Descriptor()
}

func (SubmitResponse_ResultType) Type() protoreflect.EnumType {
	return &file_pkg_flux_grpc_flux_proto_enumTypes[0]
}

func (x SubmitResponse_ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubmitResponse_ResultType.Descriptor instead.
func (SubmitResponse_ResultType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_flux_grpc_flux_proto_rawDescGZIP(), []int{1, 0}
}

type InfoResponse_ResultType int32

const (
	InfoResponse_INFO_SUCCESS InfoResponse_ResultType = 0
	InfoResponse_INFO_ERROR   InfoResponse_ResultType = 1
	InfoResponse_INFO_DENIED  InfoResponse_ResultType = 2
)

// Enum value maps for InfoResponse_ResultType.
var (
	InfoResponse_ResultType_name = map[int32]string{
		0: "INFO_SUCCESS",
		1: "INFO_ERROR",
		2: "INFO_DENIED",
	}
	InfoResponse_ResultType_value = map[string]int32{
		"INFO_SUCCESS": 0,
		"INFO_ERROR":   1,
		"INFO_DENIED":  2,
	}
)

func (x InfoResponse_ResultType) Enum() *InfoResponse_ResultType {
	p := new(InfoResponse_ResultType)
	*p = x
	return p
}

func (x InfoResponse_ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InfoResponse_ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_flux_grpc_flux_proto_enumTypes[1].Descriptor()
}

func (InfoResponse_ResultType) Type() protoreflect.EnumType {
	return &file_pkg_flux_grpc_flux_proto_enumTypes[1]
}

func (x InfoResponse_ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InfoResponse_ResultType.Descriptor instead.
func (InfoResponse_ResultType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_flux_grpc_flux_proto_rawDescGZIP(), []int{3, 0}
}

type JobsResponse_ResultType int32

const (
	JobsResponse_JOBS_SUCCESS JobsResponse_ResultType = 0
	JobsResponse_JOBS_ERROR   JobsResponse_ResultType = 1
	JobsResponse_JOBS_DENIED  JobsResponse_ResultType = 2
)

// Enum value maps for JobsResponse_ResultType.
var (
	JobsResponse_ResultType_name = map[int32]string{
		0: "JOBS_SUCCESS",
		1: "JOBS_ERROR",
		2: "JOBS_DENIED",
	}
	JobsResponse_ResultType_value = map[string]int32{
		"JOBS_SUCCESS": 0,
		"JOBS_ERROR":   1,
		"JOBS_DENIED":  2,
	}
)

func (x JobsResponse_ResultType) Enum() *JobsResponse_ResultType {
	p := new(JobsResponse_ResultType)
	*p = x
	return p
}

func (x JobsResponse_ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobsResponse_ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_flux_grpc_flux_proto_enumTypes[2].Descriptor()
}

func (JobsResponse_ResultType) Type() protoreflect.EnumType {
	return &file_pkg_flux_grpc_flux_proto_enumTypes[2]
}

func (x JobsResponse_ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobsResponse_ResultType.Descriptor instead.
func (JobsResponse_ResultType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_flux_grpc_flux_proto_rawDescGZIP(), []int{6, 0}
}

// Note that this isn't a jobspec, formally
// It just is a definition of resources
type SubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobspec string `protobuf:"bytes,1,opt,name=jobspec,proto3" json:"jobspec,omitempty"`
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flux_grpc_flux_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flux_grpc_flux_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flux_grpc_flux_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitRequest) GetJobspec() string {
	if x != nil {
		return x.Jobspec
	}
	return ""
}

type SubmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status SubmitResponse_ResultType `protobuf:"varint,1,opt,name=status,proto3,enum=flux.SubmitResponse_ResultType" json:"status,omitempty"`
}

func (x *SubmitResponse) Reset() {
	*x = SubmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flux_grpc_flux_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitResponse) ProtoMessage() {}

func (x *SubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flux_grpc_flux_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitResponse.ProtoReflect.Descriptor instead.
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flux_grpc_flux_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitResponse) GetStatus() SubmitResponse_ResultType {
	if x != nil {
		return x.Status
	}
	return SubmitResponse_SUBMIT_SUCCESS
}

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobid string `protobuf:"bytes,1,opt,name=jobid,proto3" json:"jobid,omitempty"`
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flux_grpc_flux_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flux_grpc_flux_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flux_grpc_flux_proto_rawDescGZIP(), []int{2}
}

func (x *InfoRequest) GetJobid() string {
	if x != nil {
		return x.Jobid
	}
	return ""
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status InfoResponse_ResultType `protobuf:"varint,1,opt,name=status,proto3,enum=flux.InfoResponse_ResultType" json:"status,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flux_grpc_flux_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flux_grpc_flux_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flux_grpc_flux_proto_rawDescGZIP(), []int{3}
}

func (x *InfoResponse) GetStatus() InfoResponse_ResultType {
	if x != nil {
		return x.Status
	}
	return InfoResponse_INFO_SUCCESS
}

type JobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JobsRequest) Reset() {
	*x = JobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flux_grpc_flux_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobsRequest) ProtoMessage() {}

func (x *JobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flux_grpc_flux_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobsRequest.ProtoReflect.Descriptor instead.
func (*JobsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flux_grpc_flux_proto_rawDescGZIP(), []int{4}
}

type JobInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobid  string `protobuf:"bytes,1,opt,name=jobid,proto3" json:"jobid,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *JobInfo) Reset() {
	*x = JobInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flux_grpc_flux_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobInfo) ProtoMessage() {}

func (x *JobInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flux_grpc_flux_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobInfo.ProtoReflect.Descriptor instead.
func (*JobInfo) Descriptor() ([]byte, []int) {
	return file_pkg_flux_grpc_flux_proto_rawDescGZIP(), []int{5}
}

func (x *JobInfo) GetJobid() string {
	if x != nil {
		return x.Jobid
	}
	return ""
}

func (x *JobInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// TODO turn into JobInfo
type JobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status JobsResponse_ResultType `protobuf:"varint,1,opt,name=status,proto3,enum=flux.JobsResponse_ResultType" json:"status,omitempty"`
	Jobs   string                  `protobuf:"bytes,2,opt,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *JobsResponse) Reset() {
	*x = JobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flux_grpc_flux_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobsResponse) ProtoMessage() {}

func (x *JobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flux_grpc_flux_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobsResponse.ProtoReflect.Descriptor instead.
func (*JobsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flux_grpc_flux_proto_rawDescGZIP(), []int{6}
}

func (x *JobsResponse) GetStatus() JobsResponse_ResultType {
	if x != nil {
		return x.Status
	}
	return JobsResponse_JOBS_SUCCESS
}

func (x *JobsResponse) GetJobs() string {
	if x != nil {
		return x.Jobs
	}
	return ""
}

var File_pkg_flux_grpc_flux_proto protoreflect.FileDescriptor

var file_pkg_flux_grpc_flux_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x66, 0x6c, 0x75, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6c, 0x75, 0x78,
	0x22, 0x29, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x73, 0x70, 0x65, 0x63, 0x22, 0x90, 0x01, 0x0a, 0x0e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x45, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x55, 0x42,
	0x4d, 0x49, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53,
	0x55, 0x42, 0x4d, 0x49, 0x54, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x02, 0x22, 0x23,
	0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6a, 0x6f, 0x62, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f,
	0x62, 0x69, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3f, 0x0a, 0x0a, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x46,
	0x4f, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49,
	0x4e, 0x46, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49,
	0x4e, 0x46, 0x4f, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x02, 0x22, 0x0d, 0x0a, 0x0b,
	0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x07, 0x4a,
	0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x4a, 0x6f, 0x62,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6a, 0x6f, 0x62, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73,
	0x22, 0x3f, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x0c, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10,
	0x02, 0x32, 0x75, 0x0a, 0x0b, 0x46, 0x6c, 0x75, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x35, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x13, 0x2e, 0x66, 0x6c, 0x75,
	0x78, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x04, 0x4a, 0x6f, 0x62, 0x73, 0x12,
	0x11, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x2d, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_flux_grpc_flux_proto_rawDescOnce sync.Once
	file_pkg_flux_grpc_flux_proto_rawDescData = file_pkg_flux_grpc_flux_proto_rawDesc
)

func file_pkg_flux_grpc_flux_proto_rawDescGZIP() []byte {
	file_pkg_flux_grpc_flux_proto_rawDescOnce.Do(func() {
		file_pkg_flux_grpc_flux_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_flux_grpc_flux_proto_rawDescData)
	})
	return file_pkg_flux_grpc_flux_proto_rawDescData
}

var file_pkg_flux_grpc_flux_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pkg_flux_grpc_flux_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_flux_grpc_flux_proto_goTypes = []interface{}{
	(SubmitResponse_ResultType)(0), // 0: flux.SubmitResponse.ResultType
	(InfoResponse_ResultType)(0),   // 1: flux.InfoResponse.ResultType
	(JobsResponse_ResultType)(0),   // 2: flux.JobsResponse.ResultType
	(*SubmitRequest)(nil),          // 3: flux.SubmitRequest
	(*SubmitResponse)(nil),         // 4: flux.SubmitResponse
	(*InfoRequest)(nil),            // 5: flux.InfoRequest
	(*InfoResponse)(nil),           // 6: flux.InfoResponse
	(*JobsRequest)(nil),            // 7: flux.JobsRequest
	(*JobInfo)(nil),                // 8: flux.JobInfo
	(*JobsResponse)(nil),           // 9: flux.JobsResponse
}
var file_pkg_flux_grpc_flux_proto_depIdxs = []int32{
	0, // 0: flux.SubmitResponse.status:type_name -> flux.SubmitResponse.ResultType
	1, // 1: flux.InfoResponse.status:type_name -> flux.InfoResponse.ResultType
	2, // 2: flux.JobsResponse.status:type_name -> flux.JobsResponse.ResultType
	3, // 3: flux.FluxService.Submit:input_type -> flux.SubmitRequest
	7, // 4: flux.FluxService.Jobs:input_type -> flux.JobsRequest
	4, // 5: flux.FluxService.Submit:output_type -> flux.SubmitResponse
	9, // 6: flux.FluxService.Jobs:output_type -> flux.JobsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_flux_grpc_flux_proto_init() }
func file_pkg_flux_grpc_flux_proto_init() {
	if File_pkg_flux_grpc_flux_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_flux_grpc_flux_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRequest); i {
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
		file_pkg_flux_grpc_flux_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitResponse); i {
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
		file_pkg_flux_grpc_flux_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_pkg_flux_grpc_flux_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
		file_pkg_flux_grpc_flux_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobsRequest); i {
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
		file_pkg_flux_grpc_flux_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobInfo); i {
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
		file_pkg_flux_grpc_flux_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobsResponse); i {
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
			RawDescriptor: file_pkg_flux_grpc_flux_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_flux_grpc_flux_proto_goTypes,
		DependencyIndexes: file_pkg_flux_grpc_flux_proto_depIdxs,
		EnumInfos:         file_pkg_flux_grpc_flux_proto_enumTypes,
		MessageInfos:      file_pkg_flux_grpc_flux_proto_msgTypes,
	}.Build()
	File_pkg_flux_grpc_flux_proto = out.File
	file_pkg_flux_grpc_flux_proto_rawDesc = nil
	file_pkg_flux_grpc_flux_proto_goTypes = nil
	file_pkg_flux_grpc_flux_proto_depIdxs = nil
}
