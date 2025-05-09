// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: proto/mchost_spot_instance_payload.proto

package pb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_mchost_spot_instance_payload_proto_rawDescGZIP(), []int{0}
}

type DefaultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DefaultResponse) Reset() {
	*x = DefaultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultResponse) ProtoMessage() {}

func (x *DefaultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultResponse.ProtoReflect.Descriptor instead.
func (*DefaultResponse) Descriptor() ([]byte, []int) {
	return file_proto_mchost_spot_instance_payload_proto_rawDescGZIP(), []int{1}
}

func (x *DefaultResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *DefaultResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DefaultResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IpId uint64 `protobuf:"varint,2,opt,name=ipId,proto3" json:"ipId,omitempty"`
}

func (x *CreateTemplateRequest) Reset() {
	*x = CreateTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateRequest) ProtoMessage() {}

func (x *CreateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_proto_mchost_spot_instance_payload_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTemplateRequest) GetIpId() uint64 {
	if x != nil {
		return x.IpId
	}
	return 0
}

type GetTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpotInstanceTemplateId int64 `protobuf:"varint,1,opt,name=spotInstanceTemplateId,proto3" json:"spotInstanceTemplateId,omitempty"`
}

func (x *GetTemplateRequest) Reset() {
	*x = GetTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateRequest) ProtoMessage() {}

func (x *GetTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateRequest.ProtoReflect.Descriptor instead.
func (*GetTemplateRequest) Descriptor() ([]byte, []int) {
	return file_proto_mchost_spot_instance_payload_proto_rawDescGZIP(), []int{3}
}

func (x *GetTemplateRequest) GetSpotInstanceTemplateId() int64 {
	if x != nil {
		return x.SpotInstanceTemplateId
	}
	return 0
}

type GetTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    bool                  `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code     uint32                `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message  string                `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Template *SpotInstanceTemplate `protobuf:"bytes,4,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *GetTemplateResponse) Reset() {
	*x = GetTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateResponse) ProtoMessage() {}

func (x *GetTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateResponse.ProtoReflect.Descriptor instead.
func (*GetTemplateResponse) Descriptor() ([]byte, []int) {
	return file_proto_mchost_spot_instance_payload_proto_rawDescGZIP(), []int{4}
}

func (x *GetTemplateResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *GetTemplateResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetTemplateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetTemplateResponse) GetTemplate() *SpotInstanceTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

type LaunchTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpotInstanceTemplateId int64 `protobuf:"varint,1,opt,name=spotInstanceTemplateId,proto3" json:"spotInstanceTemplateId,omitempty"`
}

func (x *LaunchTemplateRequest) Reset() {
	*x = LaunchTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaunchTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaunchTemplateRequest) ProtoMessage() {}

func (x *LaunchTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaunchTemplateRequest.ProtoReflect.Descriptor instead.
func (*LaunchTemplateRequest) Descriptor() ([]byte, []int) {
	return file_proto_mchost_spot_instance_payload_proto_rawDescGZIP(), []int{5}
}

func (x *LaunchTemplateRequest) GetSpotInstanceTemplateId() int64 {
	if x != nil {
		return x.SpotInstanceTemplateId
	}
	return 0
}

type LaunchTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LaunchTemplateResponse) Reset() {
	*x = LaunchTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaunchTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaunchTemplateResponse) ProtoMessage() {}

func (x *LaunchTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaunchTemplateResponse.ProtoReflect.Descriptor instead.
func (*LaunchTemplateResponse) Descriptor() ([]byte, []int) {
	return file_proto_mchost_spot_instance_payload_proto_rawDescGZIP(), []int{6}
}

func (x *LaunchTemplateResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *LaunchTemplateResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LaunchTemplateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StopTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpotInstanceTemplateId int64 `protobuf:"varint,1,opt,name=spotInstanceTemplateId,proto3" json:"spotInstanceTemplateId,omitempty"`
}

func (x *StopTemplateRequest) Reset() {
	*x = StopTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTemplateRequest) ProtoMessage() {}

func (x *StopTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTemplateRequest.ProtoReflect.Descriptor instead.
func (*StopTemplateRequest) Descriptor() ([]byte, []int) {
	return file_proto_mchost_spot_instance_payload_proto_rawDescGZIP(), []int{7}
}

func (x *StopTemplateRequest) GetSpotInstanceTemplateId() int64 {
	if x != nil {
		return x.SpotInstanceTemplateId
	}
	return 0
}

type StopTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StopTemplateResponse) Reset() {
	*x = StopTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTemplateResponse) ProtoMessage() {}

func (x *StopTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mchost_spot_instance_payload_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTemplateResponse.ProtoReflect.Descriptor instead.
func (*StopTemplateResponse) Descriptor() ([]byte, []int) {
	return file_proto_mchost_spot_instance_payload_proto_rawDescGZIP(), []int{8}
}

func (x *StopTemplateResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *StopTemplateResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StopTemplateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_mchost_spot_instance_payload_proto protoreflect.FileDescriptor

var file_proto_mchost_spot_instance_payload_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x63, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73,
	0x70, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6d, 0x63, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6d, 0x63, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x55, 0x0a, 0x0f, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x3f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x69, 0x70, 0x49,
	0x64, 0x22, 0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x70, 0x6f, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x73, 0x70, 0x6f, 0x74, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22,
	0xac, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x6d, 0x63, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x70, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x4f,
	0x0a, 0x15, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x70, 0x6f, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x73, 0x70, 0x6f, 0x74, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22,
	0x5c, 0x0a, 0x16, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a,
	0x13, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x70, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x73, 0x70, 0x6f, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x14,
	0x53, 0x74, 0x6f, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mchost_spot_instance_payload_proto_rawDescOnce sync.Once
	file_proto_mchost_spot_instance_payload_proto_rawDescData = file_proto_mchost_spot_instance_payload_proto_rawDesc
)

func file_proto_mchost_spot_instance_payload_proto_rawDescGZIP() []byte {
	file_proto_mchost_spot_instance_payload_proto_rawDescOnce.Do(func() {
		file_proto_mchost_spot_instance_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mchost_spot_instance_payload_proto_rawDescData)
	})
	return file_proto_mchost_spot_instance_payload_proto_rawDescData
}

var file_proto_mchost_spot_instance_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_mchost_spot_instance_payload_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: mchost_spot_instance.service.v1.Empty
	(*DefaultResponse)(nil),        // 1: mchost_spot_instance.service.v1.DefaultResponse
	(*CreateTemplateRequest)(nil),  // 2: mchost_spot_instance.service.v1.CreateTemplateRequest
	(*GetTemplateRequest)(nil),     // 3: mchost_spot_instance.service.v1.GetTemplateRequest
	(*GetTemplateResponse)(nil),    // 4: mchost_spot_instance.service.v1.GetTemplateResponse
	(*LaunchTemplateRequest)(nil),  // 5: mchost_spot_instance.service.v1.LaunchTemplateRequest
	(*LaunchTemplateResponse)(nil), // 6: mchost_spot_instance.service.v1.LaunchTemplateResponse
	(*StopTemplateRequest)(nil),    // 7: mchost_spot_instance.service.v1.StopTemplateRequest
	(*StopTemplateResponse)(nil),   // 8: mchost_spot_instance.service.v1.StopTemplateResponse
	(*SpotInstanceTemplate)(nil),   // 9: mchost_spot_instance.service.v1.SpotInstanceTemplate
}
var file_proto_mchost_spot_instance_payload_proto_depIdxs = []int32{
	9, // 0: mchost_spot_instance.service.v1.GetTemplateResponse.template:type_name -> mchost_spot_instance.service.v1.SpotInstanceTemplate
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_mchost_spot_instance_payload_proto_init() }
func file_proto_mchost_spot_instance_payload_proto_init() {
	if File_proto_mchost_spot_instance_payload_proto != nil {
		return
	}
	file_proto_mchost_spot_instance_db_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_mchost_spot_instance_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_mchost_spot_instance_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultResponse); i {
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
		file_proto_mchost_spot_instance_payload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateRequest); i {
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
		file_proto_mchost_spot_instance_payload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateRequest); i {
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
		file_proto_mchost_spot_instance_payload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateResponse); i {
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
		file_proto_mchost_spot_instance_payload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaunchTemplateRequest); i {
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
		file_proto_mchost_spot_instance_payload_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaunchTemplateResponse); i {
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
		file_proto_mchost_spot_instance_payload_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopTemplateRequest); i {
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
		file_proto_mchost_spot_instance_payload_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopTemplateResponse); i {
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
			RawDescriptor: file_proto_mchost_spot_instance_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_mchost_spot_instance_payload_proto_goTypes,
		DependencyIndexes: file_proto_mchost_spot_instance_payload_proto_depIdxs,
		MessageInfos:      file_proto_mchost_spot_instance_payload_proto_msgTypes,
	}.Build()
	File_proto_mchost_spot_instance_payload_proto = out.File
	file_proto_mchost_spot_instance_payload_proto_rawDesc = nil
	file_proto_mchost_spot_instance_payload_proto_goTypes = nil
	file_proto_mchost_spot_instance_payload_proto_depIdxs = nil
}
