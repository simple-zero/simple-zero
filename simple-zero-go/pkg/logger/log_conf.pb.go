// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: logger/log_conf.proto

package logger

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type LogConf struct {
	Level                string   `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	LogFile              string   `protobuf:"bytes,2,opt,name=log_file,json=logFile,proto3" json:"log_file,omitempty"`
	MaxSize              int64    `protobuf:"varint,3,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	MaxAge               int64    `protobuf:"varint,4,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	MaxBackup            int64    `protobuf:"varint,5,opt,name=max_backup,json=maxBackup,proto3" json:"max_backup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogConf) Reset()         { *m = LogConf{} }
func (m *LogConf) String() string { return proto.CompactTextString(m) }
func (*LogConf) ProtoMessage()    {}
func (*LogConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee842ecf57a2e77d, []int{0}
}
func (m *LogConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogConf.Unmarshal(m, b)
}
func (m *LogConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogConf.Marshal(b, m, deterministic)
}
func (m *LogConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogConf.Merge(m, src)
}
func (m *LogConf) XXX_Size() int {
	return xxx_messageInfo_LogConf.Size(m)
}
func (m *LogConf) XXX_DiscardUnknown() {
	xxx_messageInfo_LogConf.DiscardUnknown(m)
}

var xxx_messageInfo_LogConf proto.InternalMessageInfo

func (m *LogConf) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *LogConf) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *LogConf) GetMaxSize() int64 {
	if m != nil {
		return m.MaxSize
	}
	return 0
}

func (m *LogConf) GetMaxAge() int64 {
	if m != nil {
		return m.MaxAge
	}
	return 0
}

func (m *LogConf) GetMaxBackup() int64 {
	if m != nil {
		return m.MaxBackup
	}
	return 0
}

func init() {
	proto.RegisterType((*LogConf)(nil), "sz.api.conf.LogConf")
}

func init() { proto.RegisterFile("logger/log_conf.proto", fileDescriptor_ee842ecf57a2e77d) }

var fileDescriptor_ee842ecf57a2e77d = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0xcf, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0x85, 0xd2, 0x86, 0x9a, 0xcd, 0x02, 0x61, 0x06, 0xa4, 0x8a, 0xa9, 0x0b, 0x71,
	0x04, 0x23, 0x13, 0x45, 0x62, 0x62, 0x2a, 0x1b, 0x4b, 0xe4, 0x44, 0x97, 0xe3, 0xd4, 0x73, 0xcf,
	0x4a, 0x5a, 0x64, 0xe5, 0x1d, 0x78, 0x67, 0x64, 0x67, 0xba, 0xbb, 0xff, 0x93, 0x4e, 0xfa, 0xd5,
	0x2d, 0x0b, 0x22, 0x0c, 0x96, 0x05, 0x9b, 0x4e, 0x8e, 0x7d, 0x15, 0x06, 0x39, 0x89, 0xbe, 0x1e,
	0xa7, 0xca, 0x05, 0xaa, 0x52, 0xf4, 0xf8, 0x57, 0xa8, 0xf2, 0x53, 0xf0, 0x5d, 0x8e, 0xbd, 0xbe,
	0x51, 0x4b, 0x86, 0x5f, 0x60, 0x53, 0x6c, 0x8a, 0xed, 0x7a, 0x3f, 0x1f, 0xfa, 0x5e, 0x5d, 0xa5,
	0x07, 0x3d, 0x31, 0x98, 0x8b, 0x0c, 0x25, 0x0b, 0x7e, 0x10, 0x43, 0x22, 0xef, 0x62, 0x33, 0xd2,
	0x04, 0x66, 0xb1, 0x29, 0xb6, 0x8b, 0x7d, 0xe9, 0x5d, 0xfc, 0xa2, 0x09, 0xf4, 0x9d, 0x4a, 0x6b,
	0xe3, 0x10, 0xcc, 0x65, 0x96, 0x95, 0x77, 0xf1, 0x0d, 0x41, 0x3f, 0x28, 0x95, 0xa0, 0x75, 0xdd,
	0xe1, 0x1c, 0xcc, 0x32, 0xdb, 0xda, 0xbb, 0xb8, 0xcb, 0xc1, 0xee, 0xf9, 0xbb, 0x46, 0x3a, 0xfd,
	0x9c, 0xdb, 0xaa, 0x13, 0x6f, 0x19, 0xc8, 0xc5, 0xba, 0xb6, 0x23, 0xf9, 0xc0, 0xf0, 0x34, 0xc1,
	0x20, 0x36, 0x1c, 0xd0, 0xce, 0xc5, 0x5e, 0xe7, 0xd1, 0xae, 0x72, 0xaf, 0x97, 0xff, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa1, 0x68, 0xc3, 0x39, 0xf0, 0x00, 0x00, 0x00,
}
