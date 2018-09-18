// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello_ack.proto

package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloAck struct {
	Guid                 int64    `protobuf:"varint,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Version              uint32   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Ip                   []byte   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 uint32   `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloAck) Reset()         { *m = HelloAck{} }
func (m *HelloAck) String() string { return proto.CompactTextString(m) }
func (*HelloAck) ProtoMessage()    {}
func (*HelloAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_ack_a2780f873a5bb2e2, []int{0}
}
func (m *HelloAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloAck.Unmarshal(m, b)
}
func (m *HelloAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloAck.Marshal(b, m, deterministic)
}
func (dst *HelloAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloAck.Merge(dst, src)
}
func (m *HelloAck) XXX_Size() int {
	return xxx_messageInfo_HelloAck.Size(m)
}
func (m *HelloAck) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloAck.DiscardUnknown(m)
}

var xxx_messageInfo_HelloAck proto.InternalMessageInfo

func (m *HelloAck) GetGuid() int64 {
	if m != nil {
		return m.Guid
	}
	return 0
}

func (m *HelloAck) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HelloAck) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *HelloAck) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloAck)(nil), "messages.HelloAck")
}

func init() { proto.RegisterFile("hello_ack.proto", fileDescriptor_hello_ack_a2780f873a5bb2e2) }

var fileDescriptor_hello_ack_a2780f873a5bb2e2 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x48, 0xcd, 0xc9,
	0xc9, 0x8f, 0x4f, 0x4c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x56, 0x8a, 0xe1, 0xe2, 0xf0, 0x00, 0x49, 0x3a, 0x26, 0x67, 0x0b,
	0x09, 0x71, 0xb1, 0xa4, 0x97, 0x66, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x81, 0xd9,
	0x42, 0x12, 0x5c, 0xec, 0x65, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0xbc, 0x41, 0x30, 0xae, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x81, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4f,
	0x10, 0x53, 0x66, 0x01, 0x48, 0x77, 0x41, 0x7e, 0x51, 0x89, 0x04, 0x0b, 0x58, 0x19, 0x98, 0x9d,
	0xc4, 0x06, 0xb6, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x98, 0xbc, 0x23, 0x4f, 0x81, 0x00,
	0x00, 0x00,
}
