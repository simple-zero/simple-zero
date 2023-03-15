// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config/v1/api.proto

package api

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PropCond struct {
	// @gotags: form:"key" uri:"key"
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" form:"key" uri:"key"`
	// @gotags: form:"prefix"
	Prefix               bool     `protobuf:"varint,2,opt,name=prefix,proto3" json:"prefix,omitempty" form:"prefix"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PropCond) Reset()         { *m = PropCond{} }
func (m *PropCond) String() string { return proto.CompactTextString(m) }
func (*PropCond) ProtoMessage()    {}
func (*PropCond) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3153a2604995399, []int{0}
}
func (m *PropCond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PropCond.Unmarshal(m, b)
}
func (m *PropCond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PropCond.Marshal(b, m, deterministic)
}
func (m *PropCond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PropCond.Merge(m, src)
}
func (m *PropCond) XXX_Size() int {
	return xxx_messageInfo_PropCond.Size(m)
}
func (m *PropCond) XXX_DiscardUnknown() {
	xxx_messageInfo_PropCond.DiscardUnknown(m)
}

var xxx_messageInfo_PropCond proto.InternalMessageInfo

func (m *PropCond) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PropCond) GetPrefix() bool {
	if m != nil {
		return m.Prefix
	}
	return false
}

type PropReply struct {
	// @gotags: json:"kvs"
	Kvs                  []*KvObj `protobuf:"bytes,1,rep,name=kvs,proto3" json:"kvs"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PropReply) Reset()         { *m = PropReply{} }
func (m *PropReply) String() string { return proto.CompactTextString(m) }
func (*PropReply) ProtoMessage()    {}
func (*PropReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3153a2604995399, []int{1}
}
func (m *PropReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PropReply.Unmarshal(m, b)
}
func (m *PropReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PropReply.Marshal(b, m, deterministic)
}
func (m *PropReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PropReply.Merge(m, src)
}
func (m *PropReply) XXX_Size() int {
	return xxx_messageInfo_PropReply.Size(m)
}
func (m *PropReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PropReply.DiscardUnknown(m)
}

var xxx_messageInfo_PropReply proto.InternalMessageInfo

func (m *PropReply) GetKvs() []*KvObj {
	if m != nil {
		return m.Kvs
	}
	return nil
}

type KvObj struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KvObj) Reset()         { *m = KvObj{} }
func (m *KvObj) String() string { return proto.CompactTextString(m) }
func (*KvObj) ProtoMessage()    {}
func (*KvObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3153a2604995399, []int{2}
}
func (m *KvObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KvObj.Unmarshal(m, b)
}
func (m *KvObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KvObj.Marshal(b, m, deterministic)
}
func (m *KvObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KvObj.Merge(m, src)
}
func (m *KvObj) XXX_Size() int {
	return xxx_messageInfo_KvObj.Size(m)
}
func (m *KvObj) XXX_DiscardUnknown() {
	xxx_messageInfo_KvObj.DiscardUnknown(m)
}

var xxx_messageInfo_KvObj proto.InternalMessageInfo

func (m *KvObj) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KvObj) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UIConfig struct {
	Serve                *UIServe  `protobuf:"bytes,1,opt,name=serve,proto3" json:"serve,omitempty"`
	Menus                []*UIMenu `protobuf:"bytes,2,rep,name=menus,proto3" json:"menus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UIConfig) Reset()         { *m = UIConfig{} }
func (m *UIConfig) String() string { return proto.CompactTextString(m) }
func (*UIConfig) ProtoMessage()    {}
func (*UIConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3153a2604995399, []int{3}
}
func (m *UIConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UIConfig.Unmarshal(m, b)
}
func (m *UIConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UIConfig.Marshal(b, m, deterministic)
}
func (m *UIConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UIConfig.Merge(m, src)
}
func (m *UIConfig) XXX_Size() int {
	return xxx_messageInfo_UIConfig.Size(m)
}
func (m *UIConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UIConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UIConfig proto.InternalMessageInfo

func (m *UIConfig) GetServe() *UIServe {
	if m != nil {
		return m.Serve
	}
	return nil
}

func (m *UIConfig) GetMenus() []*UIMenu {
	if m != nil {
		return m.Menus
	}
	return nil
}

type UIServe struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Domain               string   `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UIServe) Reset()         { *m = UIServe{} }
func (m *UIServe) String() string { return proto.CompactTextString(m) }
func (*UIServe) ProtoMessage()    {}
func (*UIServe) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3153a2604995399, []int{4}
}
func (m *UIServe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UIServe.Unmarshal(m, b)
}
func (m *UIServe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UIServe.Marshal(b, m, deterministic)
}
func (m *UIServe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UIServe.Merge(m, src)
}
func (m *UIServe) XXX_Size() int {
	return xxx_messageInfo_UIServe.Size(m)
}
func (m *UIServe) XXX_DiscardUnknown() {
	xxx_messageInfo_UIServe.DiscardUnknown(m)
}

var xxx_messageInfo_UIServe proto.InternalMessageInfo

func (m *UIServe) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UIServe) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UIServe) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *UIServe) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type UIMenu struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=Path,proto3" json:"Path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UIMenu) Reset()         { *m = UIMenu{} }
func (m *UIMenu) String() string { return proto.CompactTextString(m) }
func (*UIMenu) ProtoMessage()    {}
func (*UIMenu) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3153a2604995399, []int{5}
}
func (m *UIMenu) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UIMenu.Unmarshal(m, b)
}
func (m *UIMenu) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UIMenu.Marshal(b, m, deterministic)
}
func (m *UIMenu) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UIMenu.Merge(m, src)
}
func (m *UIMenu) XXX_Size() int {
	return xxx_messageInfo_UIMenu.Size(m)
}
func (m *UIMenu) XXX_DiscardUnknown() {
	xxx_messageInfo_UIMenu.DiscardUnknown(m)
}

var xxx_messageInfo_UIMenu proto.InternalMessageInfo

func (m *UIMenu) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UIMenu) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UIMenu) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type UIConfigReply struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UIConfigReply) Reset()         { *m = UIConfigReply{} }
func (m *UIConfigReply) String() string { return proto.CompactTextString(m) }
func (*UIConfigReply) ProtoMessage()    {}
func (*UIConfigReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3153a2604995399, []int{6}
}
func (m *UIConfigReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UIConfigReply.Unmarshal(m, b)
}
func (m *UIConfigReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UIConfigReply.Marshal(b, m, deterministic)
}
func (m *UIConfigReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UIConfigReply.Merge(m, src)
}
func (m *UIConfigReply) XXX_Size() int {
	return xxx_messageInfo_UIConfigReply.Size(m)
}
func (m *UIConfigReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UIConfigReply.DiscardUnknown(m)
}

var xxx_messageInfo_UIConfigReply proto.InternalMessageInfo

func (m *UIConfigReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*PropCond)(nil), "config.v1.PropCond")
	proto.RegisterType((*PropReply)(nil), "config.v1.PropReply")
	proto.RegisterType((*KvObj)(nil), "config.v1.KvObj")
	proto.RegisterType((*UIConfig)(nil), "config.v1.UIConfig")
	proto.RegisterType((*UIServe)(nil), "config.v1.UIServe")
	proto.RegisterType((*UIMenu)(nil), "config.v1.UIMenu")
	proto.RegisterType((*UIConfigReply)(nil), "config.v1.UIConfigReply")
}

func init() { proto.RegisterFile("config/v1/api.proto", fileDescriptor_b3153a2604995399) }

var fileDescriptor_b3153a2604995399 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0x6e, 0x7e, 0x9a, 0x9d, 0x8a, 0xaa, 0xb8, 0x01, 0x56, 0x11, 0x88, 0xc8, 0x17, 0x72,
	0xe9, 0x3a, 0x0d, 0xdc, 0xb8, 0x20, 0x7a, 0x40, 0x15, 0x42, 0x54, 0x8e, 0xc2, 0x01, 0xa9, 0x07,
	0xa7, 0x71, 0xb7, 0x66, 0x77, 0x6d, 0x6b, 0xff, 0xd4, 0x80, 0xb8, 0xf0, 0x0a, 0x3c, 0x0b, 0x4f,
	0xc2, 0x2b, 0xf0, 0x20, 0xc8, 0xf6, 0x6e, 0xd9, 0xa2, 0x1c, 0x7a, 0x9b, 0xef, 0x9b, 0x99, 0xef,
	0x1b, 0x8f, 0x07, 0x8e, 0x2e, 0x95, 0xbc, 0x12, 0x31, 0xa9, 0x4f, 0x08, 0xd3, 0x22, 0xd2, 0xb9,
	0x2a, 0x15, 0x0a, 0x1c, 0x19, 0xd5, 0x27, 0x93, 0xa7, 0xb1, 0x52, 0x71, 0xca, 0x4d, 0x92, 0x30,
	0x29, 0x55, 0xc9, 0x4a, 0xa1, 0x64, 0xe1, 0x0a, 0xf1, 0x2b, 0x18, 0x9d, 0xe7, 0x4a, 0x9f, 0x2a,
	0xb9, 0x41, 0x87, 0xd0, 0x4b, 0xf8, 0x36, 0xf4, 0xa6, 0xde, 0x2c, 0xa0, 0x26, 0x44, 0x8f, 0x61,
	0xa8, 0x73, 0x7e, 0x25, 0x6e, 0x42, 0x7f, 0xea, 0xcd, 0x46, 0xb4, 0x41, 0x98, 0x40, 0x60, 0xba,
	0x28, 0xd7, 0xe9, 0x16, 0x61, 0xe8, 0x25, 0x75, 0x11, 0x7a, 0xd3, 0xde, 0x6c, 0x7f, 0x71, 0x18,
	0xdd, 0x3a, 0x47, 0xef, 0xeb, 0x8f, 0xeb, 0x2f, 0xd4, 0x24, 0x31, 0x81, 0x81, 0x45, 0x3b, 0x3c,
	0xc6, 0x30, 0xa8, 0x59, 0x5a, 0x71, 0x6b, 0x11, 0x50, 0x07, 0xf0, 0x05, 0x8c, 0x56, 0x67, 0xa7,
	0x56, 0x0a, 0xcd, 0x60, 0x50, 0xf0, 0xbc, 0xe6, 0xb6, 0x6b, 0x7f, 0x81, 0x3a, 0x16, 0xab, 0xb3,
	0xa5, 0xc9, 0x50, 0x57, 0x80, 0x5e, 0xc0, 0x20, 0xe3, 0xb2, 0x2a, 0x42, 0xdf, 0x0e, 0xf3, 0xf0,
	0x4e, 0xe5, 0x07, 0x2e, 0x2b, 0xea, 0xf2, 0xf8, 0x02, 0xf6, 0x9a, 0x56, 0x74, 0x00, 0xbe, 0xd8,
	0x34, 0x03, 0xf9, 0x62, 0x83, 0x10, 0xf4, 0x25, 0xcb, 0xda, 0x71, 0x6c, 0xdc, 0xd9, 0x43, 0xcf,
	0xb2, 0x0d, 0x32, 0xfc, 0x46, 0x65, 0x4c, 0xc8, 0xb0, 0xef, 0x78, 0x87, 0xf0, 0x1b, 0x18, 0x3a,
	0xbf, 0x7b, 0xa9, 0x23, 0xe8, 0x9f, 0xb3, 0xf2, 0xba, 0xd1, 0xb6, 0x31, 0x7e, 0x0e, 0x0f, 0xda,
	0xf7, 0xbb, 0x2d, 0x1f, 0x80, 0xaf, 0x12, 0x2b, 0x34, 0xa2, 0xbe, 0x4a, 0x16, 0xbf, 0x3c, 0x18,
	0x36, 0xfb, 0x59, 0xc2, 0xde, 0x3b, 0x5e, 0x9a, 0x0f, 0x41, 0x47, 0x9d, 0x17, 0xb7, 0xff, 0x3a,
	0x19, 0xff, 0x47, 0x5a, 0x41, 0xfc, 0xec, 0xc7, 0xef, 0x3f, 0x3f, 0xfd, 0x27, 0xe8, 0x11, 0xf9,
	0x77, 0x40, 0x3a, 0x57, 0x9a, 0x7c, 0x4b, 0xf8, 0xf6, 0x3b, 0xfa, 0x04, 0xc1, 0x92, 0x97, 0x2b,
	0x61, 0x3c, 0xee, 0xc8, 0xb6, 0x63, 0x4d, 0xc2, 0x1d, 0xa4, 0x93, 0x9e, 0x58, 0xe9, 0x31, 0x46,
	0x1d, 0xe9, 0x4a, 0x1c, 0x1b, 0xf0, 0x76, 0xf1, 0x79, 0x1e, 0x8b, 0xf2, 0xba, 0x5a, 0x47, 0x97,
	0x2a, 0x23, 0x29, 0x17, 0xec, 0x66, 0x3e, 0x27, 0x85, 0xc8, 0x74, 0xca, 0x8f, 0xbf, 0xf2, 0x5c,
	0xd9, 0x43, 0xbd, 0xed, 0x7b, 0xcd, 0xb4, 0x58, 0x0f, 0xed, 0xad, 0xbe, 0xfc, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x79, 0xf8, 0x5e, 0xd7, 0xeb, 0x02, 0x00, 0x00,
}
