// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_block.proto

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

type GetBlock struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlock) Reset()         { *m = GetBlock{} }
func (m *GetBlock) String() string { return proto.CompactTextString(m) }
func (*GetBlock) ProtoMessage()    {}
func (*GetBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_block_987f408defdef5b4, []int{0}
}
func (m *GetBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlock.Unmarshal(m, b)
}
func (m *GetBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlock.Marshal(b, m, deterministic)
}
func (dst *GetBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlock.Merge(dst, src)
}
func (m *GetBlock) XXX_Size() int {
	return xxx_messageInfo_GetBlock.Size(m)
}
func (m *GetBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlock.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlock proto.InternalMessageInfo

func (m *GetBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBlock)(nil), "messages.GetBlock")
}

func init() { proto.RegisterFile("get_block.proto", fileDescriptor_get_block_987f408defdef5b4) }

var fileDescriptor_get_block_987f408defdef5b4 = []byte{
	// 83 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4f, 0x2d, 0x89,
	0x4f, 0xca, 0xc9, 0x4f, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x56, 0x92, 0xe3, 0xe2, 0x70, 0x4f, 0x2d, 0x71, 0x02, 0xc9, 0x09,
	0x09, 0x71, 0xb1, 0x64, 0x24, 0x16, 0x67, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9,
	0x49, 0x6c, 0x60, 0x0d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x00, 0xb8, 0xca, 0x43,
	0x00, 0x00, 0x00,
}
