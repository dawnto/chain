// Code generated by protoc-gen-go. DO NOT EDIT.
// source: new_transaction.proto

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

type NewTransaction struct {
	Transaction          []byte   `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewTransaction) Reset()         { *m = NewTransaction{} }
func (m *NewTransaction) String() string { return proto.CompactTextString(m) }
func (*NewTransaction) ProtoMessage()    {}
func (*NewTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_new_transaction_1e80528064d56e6b, []int{0}
}
func (m *NewTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTransaction.Unmarshal(m, b)
}
func (m *NewTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTransaction.Marshal(b, m, deterministic)
}
func (dst *NewTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTransaction.Merge(dst, src)
}
func (m *NewTransaction) XXX_Size() int {
	return xxx_messageInfo_NewTransaction.Size(m)
}
func (m *NewTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_NewTransaction proto.InternalMessageInfo

func (m *NewTransaction) GetTransaction() []byte {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func init() {
	proto.RegisterType((*NewTransaction)(nil), "messages.NewTransaction")
}

func init() {
	proto.RegisterFile("new_transaction.proto", fileDescriptor_new_transaction_1e80528064d56e6b)
}

var fileDescriptor_new_transaction_1e80528064d56e6b = []byte{
	// 87 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4b, 0x2d, 0x8f,
	0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0xc8, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x56, 0x32, 0xe2, 0xe2, 0xf3,
	0x4b, 0x2d, 0x0f, 0x41, 0xa8, 0x10, 0x52, 0xe0, 0xe2, 0x46, 0xd2, 0x20, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x13, 0x84, 0x2c, 0x94, 0xc4, 0x06, 0x36, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x70,
	0x2c, 0x10, 0x1d, 0x5d, 0x00, 0x00, 0x00,
}
