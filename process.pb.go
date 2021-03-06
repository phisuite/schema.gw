// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: process.proto

package schema

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version    string              `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Status     Status              `protobuf:"varint,3,opt,name=status,proto3,enum=schema.Status" json:"status,omitempty"`
	Definition *Process_Definition `protobuf:"bytes,10,opt,name=definition,proto3" json:"definition,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_process_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_process_proto_rawDescGZIP(), []int{0}
}

func (x *Process) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Process) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Process) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNACTIVATED
}

func (x *Process) GetDefinition() *Process_Definition {
	if x != nil {
		return x.Definition
	}
	return nil
}

type Process_Definition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input  *Process_Data `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output *Process_Data `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Error  *Process_Data `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Process_Definition) Reset() {
	*x = Process_Definition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process_Definition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process_Definition) ProtoMessage() {}

func (x *Process_Definition) ProtoReflect() protoreflect.Message {
	mi := &file_process_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process_Definition.ProtoReflect.Descriptor instead.
func (*Process_Definition) Descriptor() ([]byte, []int) {
	return file_process_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Process_Definition) GetInput() *Process_Data {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Process_Definition) GetOutput() *Process_Data {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *Process_Definition) GetError() *Process_Data {
	if x != nil {
		return x.Error
	}
	return nil
}

type Process_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event  *Event  `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Entity *Entity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *Process_Data) Reset() {
	*x = Process_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process_Data) ProtoMessage() {}

func (x *Process_Data) ProtoReflect() protoreflect.Message {
	mi := &file_process_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process_Data.ProtoReflect.Descriptor instead.
func (*Process_Data) Descriptor() ([]byte, []int) {
	return file_process_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Process_Data) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *Process_Data) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

var File_process_proto protoreflect.FileDescriptor

var file_process_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85,
	0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3a, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x92, 0x01, 0x0a,
	0x0a, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x1a, 0x53, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xbe, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x61, 0x64, 0x41, 0x50, 0x49, 0x12, 0x4e, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x30, 0x2f,
	0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x30, 0x01, 0x12, 0x5c, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x69,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x32, 0x92, 0x03, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x41, 0x50, 0x49, 0x12, 0x50, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22,
	0x19, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x2f, 0x76, 0x30, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x61, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2f, 0x1a, 0x2a, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a,
	0x12, 0x62, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x0f, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0x34,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x32, 0x2c, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x7d, 0x12, 0x66, 0x0a, 0x0a, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x32, 0x2e, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x30,
	0x2f, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x2f, 0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x42, 0x41, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x68, 0x69, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x68, 0x69, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_process_proto_rawDescOnce sync.Once
	file_process_proto_rawDescData = file_process_proto_rawDesc
)

func file_process_proto_rawDescGZIP() []byte {
	file_process_proto_rawDescOnce.Do(func() {
		file_process_proto_rawDescData = protoimpl.X.CompressGZIP(file_process_proto_rawDescData)
	})
	return file_process_proto_rawDescData
}

var file_process_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_process_proto_goTypes = []interface{}{
	(*Process)(nil),            // 0: schema.Process
	(*Process_Definition)(nil), // 1: schema.Process.Definition
	(*Process_Data)(nil),       // 2: schema.Process.Data
	(Status)(0),                // 3: schema.Status
	(*Event)(nil),              // 4: schema.Event
	(*Entity)(nil),             // 5: schema.Entity
	(*Options)(nil),            // 6: schema.Options
}
var file_process_proto_depIdxs = []int32{
	3,  // 0: schema.Process.status:type_name -> schema.Status
	1,  // 1: schema.Process.definition:type_name -> schema.Process.Definition
	2,  // 2: schema.Process.Definition.input:type_name -> schema.Process.Data
	2,  // 3: schema.Process.Definition.output:type_name -> schema.Process.Data
	2,  // 4: schema.Process.Definition.error:type_name -> schema.Process.Data
	4,  // 5: schema.Process.Data.event:type_name -> schema.Event
	5,  // 6: schema.Process.Data.entity:type_name -> schema.Entity
	6,  // 7: schema.ProcessReadAPI.List:input_type -> schema.Options
	6,  // 8: schema.ProcessReadAPI.Get:input_type -> schema.Options
	0,  // 9: schema.ProcessWriteAPI.Create:input_type -> schema.Process
	0,  // 10: schema.ProcessWriteAPI.Update:input_type -> schema.Process
	6,  // 11: schema.ProcessWriteAPI.Activate:input_type -> schema.Options
	6,  // 12: schema.ProcessWriteAPI.Deactivate:input_type -> schema.Options
	0,  // 13: schema.ProcessReadAPI.List:output_type -> schema.Process
	0,  // 14: schema.ProcessReadAPI.Get:output_type -> schema.Process
	0,  // 15: schema.ProcessWriteAPI.Create:output_type -> schema.Process
	0,  // 16: schema.ProcessWriteAPI.Update:output_type -> schema.Process
	0,  // 17: schema.ProcessWriteAPI.Activate:output_type -> schema.Process
	0,  // 18: schema.ProcessWriteAPI.Deactivate:output_type -> schema.Process
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_process_proto_init() }
func file_process_proto_init() {
	if File_process_proto != nil {
		return
	}
	file_common_proto_init()
	file_event_proto_init()
	file_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_process_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process); i {
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
		file_process_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process_Definition); i {
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
		file_process_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process_Data); i {
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
			RawDescriptor: file_process_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_process_proto_goTypes,
		DependencyIndexes: file_process_proto_depIdxs,
		MessageInfos:      file_process_proto_msgTypes,
	}.Build()
	File_process_proto = out.File
	file_process_proto_rawDesc = nil
	file_process_proto_goTypes = nil
	file_process_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProcessReadAPIClient is the client API for ProcessReadAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessReadAPIClient interface {
	List(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessReadAPI_ListClient, error)
	Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Process, error)
}

type processReadAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessReadAPIClient(cc grpc.ClientConnInterface) ProcessReadAPIClient {
	return &processReadAPIClient{cc}
}

func (c *processReadAPIClient) List(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessReadAPI_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessReadAPI_serviceDesc.Streams[0], "/schema.ProcessReadAPI/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &processReadAPIListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessReadAPI_ListClient interface {
	Recv() (*Process, error)
	grpc.ClientStream
}

type processReadAPIListClient struct {
	grpc.ClientStream
}

func (x *processReadAPIListClient) Recv() (*Process, error) {
	m := new(Process)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processReadAPIClient) Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Process, error) {
	out := new(Process)
	err := c.cc.Invoke(ctx, "/schema.ProcessReadAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessReadAPIServer is the server API for ProcessReadAPI service.
type ProcessReadAPIServer interface {
	List(*Options, ProcessReadAPI_ListServer) error
	Get(context.Context, *Options) (*Process, error)
}

// UnimplementedProcessReadAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProcessReadAPIServer struct {
}

func (*UnimplementedProcessReadAPIServer) List(*Options, ProcessReadAPI_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedProcessReadAPIServer) Get(context.Context, *Options) (*Process, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterProcessReadAPIServer(s *grpc.Server, srv ProcessReadAPIServer) {
	s.RegisterService(&_ProcessReadAPI_serviceDesc, srv)
}

func _ProcessReadAPI_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Options)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessReadAPIServer).List(m, &processReadAPIListServer{stream})
}

type ProcessReadAPI_ListServer interface {
	Send(*Process) error
	grpc.ServerStream
}

type processReadAPIListServer struct {
	grpc.ServerStream
}

func (x *processReadAPIListServer) Send(m *Process) error {
	return x.ServerStream.SendMsg(m)
}

func _ProcessReadAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Options)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessReadAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.ProcessReadAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessReadAPIServer).Get(ctx, req.(*Options))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProcessReadAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schema.ProcessReadAPI",
	HandlerType: (*ProcessReadAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ProcessReadAPI_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ProcessReadAPI_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "process.proto",
}

// ProcessWriteAPIClient is the client API for ProcessWriteAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessWriteAPIClient interface {
	Create(ctx context.Context, in *Process, opts ...grpc.CallOption) (*Process, error)
	Update(ctx context.Context, in *Process, opts ...grpc.CallOption) (*Process, error)
	Activate(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Process, error)
	Deactivate(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Process, error)
}

type processWriteAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessWriteAPIClient(cc grpc.ClientConnInterface) ProcessWriteAPIClient {
	return &processWriteAPIClient{cc}
}

func (c *processWriteAPIClient) Create(ctx context.Context, in *Process, opts ...grpc.CallOption) (*Process, error) {
	out := new(Process)
	err := c.cc.Invoke(ctx, "/schema.ProcessWriteAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processWriteAPIClient) Update(ctx context.Context, in *Process, opts ...grpc.CallOption) (*Process, error) {
	out := new(Process)
	err := c.cc.Invoke(ctx, "/schema.ProcessWriteAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processWriteAPIClient) Activate(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Process, error) {
	out := new(Process)
	err := c.cc.Invoke(ctx, "/schema.ProcessWriteAPI/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processWriteAPIClient) Deactivate(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Process, error) {
	out := new(Process)
	err := c.cc.Invoke(ctx, "/schema.ProcessWriteAPI/Deactivate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessWriteAPIServer is the server API for ProcessWriteAPI service.
type ProcessWriteAPIServer interface {
	Create(context.Context, *Process) (*Process, error)
	Update(context.Context, *Process) (*Process, error)
	Activate(context.Context, *Options) (*Process, error)
	Deactivate(context.Context, *Options) (*Process, error)
}

// UnimplementedProcessWriteAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProcessWriteAPIServer struct {
}

func (*UnimplementedProcessWriteAPIServer) Create(context.Context, *Process) (*Process, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedProcessWriteAPIServer) Update(context.Context, *Process) (*Process, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedProcessWriteAPIServer) Activate(context.Context, *Options) (*Process, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (*UnimplementedProcessWriteAPIServer) Deactivate(context.Context, *Options) (*Process, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}

func RegisterProcessWriteAPIServer(s *grpc.Server, srv ProcessWriteAPIServer) {
	s.RegisterService(&_ProcessWriteAPI_serviceDesc, srv)
}

func _ProcessWriteAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Process)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessWriteAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.ProcessWriteAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessWriteAPIServer).Create(ctx, req.(*Process))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessWriteAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Process)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessWriteAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.ProcessWriteAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessWriteAPIServer).Update(ctx, req.(*Process))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessWriteAPI_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Options)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessWriteAPIServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.ProcessWriteAPI/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessWriteAPIServer).Activate(ctx, req.(*Options))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessWriteAPI_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Options)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessWriteAPIServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.ProcessWriteAPI/Deactivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessWriteAPIServer).Deactivate(ctx, req.(*Options))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProcessWriteAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schema.ProcessWriteAPI",
	HandlerType: (*ProcessWriteAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProcessWriteAPI_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProcessWriteAPI_Update_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _ProcessWriteAPI_Activate_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _ProcessWriteAPI_Deactivate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "process.proto",
}
