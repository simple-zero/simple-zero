// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/v1/common.proto

package common

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/descriptorpb"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RespCode int32

const (
	RespCode_OK        RespCode = 0
	RespCode_FAIL      RespCode = 1
	RespCode_PARAM_ERR RespCode = 2
)

var RespCode_name = map[int32]string{
	0: "OK",
	1: "FAIL",
	2: "PARAM_ERR",
}

var RespCode_value = map[string]int32{
	"OK":        0,
	"FAIL":      1,
	"PARAM_ERR": 2,
}

func (x RespCode) String() string {
	return proto.EnumName(RespCode_name, int32(x))
}

func (RespCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_92d5df4519b8f2e3, []int{0}
}

type RespObj struct {
	// @gotags: json:"Code"
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"Code"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *AnyObj  `protobuf:"bytes,3,opt,name=Data,proto3,customtype=AnyObj" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespObj) Reset()         { *m = RespObj{} }
func (m *RespObj) String() string { return proto.CompactTextString(m) }
func (*RespObj) ProtoMessage()    {}
func (*RespObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d5df4519b8f2e3, []int{0}
}
func (m *RespObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespObj.Unmarshal(m, b)
}
func (m *RespObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespObj.Marshal(b, m, deterministic)
}
func (m *RespObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespObj.Merge(m, src)
}
func (m *RespObj) XXX_Size() int {
	return xxx_messageInfo_RespObj.Size(m)
}
func (m *RespObj) XXX_DiscardUnknown() {
	xxx_messageInfo_RespObj.DiscardUnknown(m)
}

var xxx_messageInfo_RespObj proto.InternalMessageInfo

func (m *RespObj) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespObj) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.common.v1.RespCode", RespCode_name, RespCode_value)
	proto.RegisterType((*RespObj)(nil), "api.common.v1.RespObj")
}

func init() { proto.RegisterFile("common/v1/common.proto", fileDescriptor_92d5df4519b8f2e3) }

var fileDescriptor_92d5df4519b8f2e3 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x14, 0xc4, 0x4d, 0x8c, 0x69, 0xfb, 0xb0, 0x10, 0xf6, 0x20, 0xa1, 0xa0, 0x06, 0x4f, 0x41, 0xb1,
	0xdb, 0x52, 0x3c, 0x79, 0x4a, 0xfd, 0x03, 0xa2, 0x12, 0xd9, 0x8b, 0xe0, 0x45, 0x76, 0xd3, 0x75,
	0xdd, 0x92, 0xf4, 0x2d, 0xd9, 0xb4, 0xa8, 0x1f, 0xd5, 0x83, 0x9f, 0x45, 0xb2, 0xd1, 0xde, 0x7e,
	0x33, 0xc3, 0x1b, 0x98, 0x07, 0x07, 0x05, 0x56, 0x15, 0xae, 0xe8, 0x66, 0x4a, 0x3b, 0x1a, 0x9b,
	0x1a, 0x1b, 0x24, 0x43, 0x6e, 0xf4, 0xf8, 0xcf, 0xd9, 0x4c, 0x47, 0x89, 0x42, 0x54, 0xa5, 0xa4,
	0x2e, 0x14, 0xeb, 0x37, 0xba, 0x90, 0xb6, 0xa8, 0xb5, 0x69, 0xb0, 0xee, 0x0e, 0x46, 0x87, 0xdb,
	0x48, 0xa1, 0x42, 0x27, 0x1c, 0x75, 0xf1, 0xc9, 0x33, 0xf4, 0x98, 0xb4, 0x26, 0x17, 0x4b, 0x42,
	0x20, 0x28, 0x70, 0x21, 0x63, 0x2f, 0xf1, 0xd2, 0x3d, 0xe6, 0x98, 0xc4, 0xd0, 0xab, 0xa4, 0xb5,
	0x5c, 0xc9, 0xd8, 0x4f, 0xbc, 0x74, 0xc0, 0xfe, 0x25, 0x39, 0x82, 0xe0, 0x9a, 0x37, 0x3c, 0xde,
	0x4d, 0xbc, 0x74, 0x7f, 0x0e, 0xdf, 0x3f, 0xc7, 0x61, 0xb6, 0xfa, 0xcc, 0xc5, 0x92, 0x39, 0xff,
	0xf4, 0x0c, 0xfa, 0x6d, 0xf1, 0x55, 0xdb, 0x12, 0x82, 0x9f, 0xdf, 0x47, 0x3b, 0xa4, 0x0f, 0xc1,
	0x6d, 0x76, 0xf7, 0x10, 0x79, 0x64, 0x08, 0x83, 0xa7, 0x8c, 0x65, 0x8f, 0xaf, 0x37, 0x8c, 0x45,
	0xfe, 0xfc, 0xe2, 0x65, 0xa6, 0x74, 0xf3, 0xbe, 0x16, 0xed, 0x34, 0x5a, 0x4a, 0xcd, 0x3f, 0x26,
	0x13, 0x6a, 0x75, 0x65, 0x4a, 0x79, 0xfe, 0x25, 0x6b, 0xa4, 0xdc, 0x68, 0xba, 0x7d, 0xc9, 0x65,
	0x47, 0x22, 0x74, 0x1b, 0x66, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79, 0xe1, 0x71, 0xa5, 0x2d,
	0x01, 0x00, 0x00,
}
