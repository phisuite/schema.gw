// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package schema

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	Status_UNACTIVATED Status = 0
	Status_ACTIVATED   Status = 1
	Status_DEACTIVATED Status = 2
)

var Status_name = map[int32]string{
	0: "UNACTIVATED",
	1: "ACTIVATED",
	2: "DEACTIVATED",
}

var Status_value = map[string]int32{
	"UNACTIVATED": 0,
	"ACTIVATED":   1,
	"DEACTIVATED": 2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type Field_Type int32

const (
	Field_STRING  Field_Type = 0
	Field_NUMBER  Field_Type = 1
	Field_BOOLEAN Field_Type = 2
	Field_FILE    Field_Type = 10
	Field_FOLDER  Field_Type = 11
	Field_EMAIL   Field_Type = 20
	Field_DATE    Field_Type = 21
	Field_TIME    Field_Type = 22
	Field_COLOR   Field_Type = 23
)

var Field_Type_name = map[int32]string{
	0:  "STRING",
	1:  "NUMBER",
	2:  "BOOLEAN",
	10: "FILE",
	11: "FOLDER",
	20: "EMAIL",
	21: "DATE",
	22: "TIME",
	23: "COLOR",
}

var Field_Type_value = map[string]int32{
	"STRING":  0,
	"NUMBER":  1,
	"BOOLEAN": 2,
	"FILE":    10,
	"FOLDER":  11,
	"EMAIL":   20,
	"DATE":    21,
	"TIME":    22,
	"COLOR":   23,
}

func (x Field_Type) String() string {
	return proto.EnumName(Field_Type_name, int32(x))
}

func (Field_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1, 0}
}

type Field_Category int32

const (
	Field_REQUIRED Field_Category = 0
	Field_OPTIONAL Field_Category = 1
)

var Field_Category_name = map[int32]string{
	0: "REQUIRED",
	1: "OPTIONAL",
}

var Field_Category_value = map[string]int32{
	"REQUIRED": 0,
	"OPTIONAL": 1,
}

func (x Field_Category) String() string {
	return proto.EnumName(Field_Category_name, int32(x))
}

func (Field_Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1, 1}
}

type Options struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Status               Status   `protobuf:"varint,3,opt,name=status,proto3,enum=schema.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Options) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Options) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UNACTIVATED
}

type Field struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 Field_Type     `protobuf:"varint,2,opt,name=type,proto3,enum=schema.Field_Type" json:"type,omitempty"`
	Category             Field_Category `protobuf:"varint,3,opt,name=category,proto3,enum=schema.Field_Category" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetType() Field_Type {
	if m != nil {
		return m.Type
	}
	return Field_STRING
}

func (m *Field) GetCategory() Field_Category {
	if m != nil {
		return m.Category
	}
	return Field_REQUIRED
}

func init() {
	proto.RegisterEnum("schema.Status", Status_name, Status_value)
	proto.RegisterEnum("schema.Field_Type", Field_Type_name, Field_Type_value)
	proto.RegisterEnum("schema.Field_Category", Field_Category_name, Field_Category_value)
	proto.RegisterType((*Options)(nil), "schema.Options")
	proto.RegisterType((*Field)(nil), "schema.Field")
}

func init() {
	proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206)
}

var fileDescriptor_555bd8c177793206 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4f, 0x8f, 0x9a, 0x40,
	0x1c, 0x15, 0x8a, 0x88, 0x3f, 0xac, 0x9d, 0x4c, 0x5b, 0xcb, 0xd1, 0x90, 0xc6, 0x98, 0x1e, 0x38,
	0xd8, 0x53, 0x8f, 0xfc, 0x19, 0x1b, 0x12, 0x64, 0xec, 0x88, 0xbd, 0x36, 0x94, 0x9d, 0xac, 0x64,
	0x17, 0x86, 0x08, 0x6e, 0xe2, 0x77, 0xd8, 0x0f, 0xbd, 0x99, 0x11, 0xd7, 0x6c, 0xb2, 0xb7, 0xf7,
	0x8f, 0xc7, 0xcc, 0x3c, 0x98, 0x14, 0xa2, 0xaa, 0x44, 0xed, 0x35, 0x47, 0xd1, 0x09, 0x6c, 0xb6,
	0xc5, 0x81, 0x57, 0xb9, 0xfb, 0x0f, 0x46, 0xb4, 0xe9, 0x4a, 0x51, 0xb7, 0x18, 0x83, 0x51, 0xe7,
	0x15, 0x77, 0xb4, 0xb9, 0xb6, 0x1c, 0x33, 0x85, 0xb1, 0x03, 0xa3, 0x27, 0x7e, 0x6c, 0x4b, 0x51,
	0x3b, 0xba, 0x92, 0xaf, 0x14, 0x2f, 0xc0, 0x6c, 0xbb, 0xbc, 0x3b, 0xb5, 0xce, 0x87, 0xb9, 0xb6,
	0x9c, 0xae, 0xa6, 0xde, 0xa5, 0xd1, 0xdb, 0x29, 0x95, 0xf5, 0xae, 0xfb, 0xac, 0xc3, 0x70, 0x5d,
	0xf2, 0xc7, 0xbb, 0x77, 0xfb, 0x17, 0x60, 0x74, 0xe7, 0x86, 0xab, 0xf2, 0xe9, 0x0a, 0x5f, 0x3b,
	0xd4, 0x07, 0x5e, 0x76, 0x6e, 0x38, 0x53, 0x3e, 0x5e, 0x81, 0x55, 0xe4, 0x1d, 0xbf, 0x17, 0xc7,
	0x73, 0xff, 0xbf, 0xd9, 0xdb, 0x6c, 0xd8, 0xbb, 0xec, 0x35, 0xe7, 0x3e, 0x80, 0x21, 0x1b, 0x30,
	0x80, 0xb9, 0xcb, 0x58, 0x9c, 0xfe, 0x46, 0x03, 0x89, 0xd3, 0xfd, 0x26, 0x20, 0x0c, 0x69, 0xd8,
	0x86, 0x51, 0x40, 0x69, 0x42, 0xfc, 0x14, 0xe9, 0xd8, 0x02, 0x63, 0x1d, 0x27, 0x04, 0x81, 0x8c,
	0xac, 0x69, 0x12, 0x11, 0x86, 0x6c, 0x3c, 0x86, 0x21, 0xd9, 0xf8, 0x71, 0x82, 0xbe, 0xc8, 0x40,
	0xe4, 0x67, 0x04, 0x7d, 0x95, 0x28, 0x8b, 0x37, 0x04, 0xcd, 0xa4, 0x1d, 0xd2, 0x84, 0x32, 0xf4,
	0xcd, 0x5d, 0x80, 0x75, 0x3d, 0x02, 0x9e, 0x80, 0xc5, 0xc8, 0x9f, 0x7d, 0xcc, 0x48, 0x84, 0x06,
	0x92, 0xd1, 0x6d, 0x16, 0xd3, 0xd4, 0x4f, 0x90, 0xf6, 0xe3, 0x17, 0x98, 0x97, 0x07, 0xc2, 0x9f,
	0xc0, 0xde, 0xa7, 0x7e, 0x98, 0xc5, 0x7f, 0xfd, 0x4c, 0x05, 0x3f, 0xc2, 0xf8, 0x46, 0x35, 0xe9,
	0x47, 0xe4, 0x26, 0xe8, 0xc1, 0x77, 0xf8, 0x5c, 0x88, 0xca, 0x6b, 0x0e, 0x65, 0x7b, 0x2a, 0x3b,
	0xde, 0xdf, 0x3f, 0xb0, 0x43, 0xb5, 0xeb, 0x56, 0xce, 0xba, 0xd5, 0xfe, 0x9b, 0x6a, 0xdf, 0x9f,
	0x2f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xb5, 0x86, 0x0d, 0xef, 0x01, 0x00, 0x00,
}