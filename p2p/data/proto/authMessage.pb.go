// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authMessage.proto

package protobuf

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

type AuthMessagePb struct {
	Message              []byte   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Sig                  []byte   `protobuf:"bytes,2,opt,name=Sig,proto3" json:"Sig,omitempty"`
	Pubkey               []byte   `protobuf:"bytes,3,opt,name=Pubkey,proto3" json:"Pubkey,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthMessagePb) Reset()         { *m = AuthMessagePb{} }
func (m *AuthMessagePb) String() string { return proto.CompactTextString(m) }
func (*AuthMessagePb) ProtoMessage()    {}
func (*AuthMessagePb) Descriptor() ([]byte, []int) {
	return fileDescriptor_e29ffa46745b2d90, []int{0}
}

func (m *AuthMessagePb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthMessagePb.Unmarshal(m, b)
}
func (m *AuthMessagePb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthMessagePb.Marshal(b, m, deterministic)
}
func (m *AuthMessagePb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthMessagePb.Merge(m, src)
}
func (m *AuthMessagePb) XXX_Size() int {
	return xxx_messageInfo_AuthMessagePb.Size(m)
}
func (m *AuthMessagePb) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthMessagePb.DiscardUnknown(m)
}

var xxx_messageInfo_AuthMessagePb proto.InternalMessageInfo

func (m *AuthMessagePb) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *AuthMessagePb) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *AuthMessagePb) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *AuthMessagePb) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthMessagePb)(nil), "protobuf.AuthMessagePb")
}

func init() { proto.RegisterFile("authMessage.proto", fileDescriptor_e29ffa46745b2d90) }

var fileDescriptor_e29ffa46745b2d90 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2c, 0x2d, 0xc9,
	0xf0, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00,
	0x53, 0x49, 0xa5, 0x69, 0x4a, 0x85, 0x5c, 0xbc, 0x8e, 0x08, 0xe9, 0x80, 0x24, 0x21, 0x09, 0x2e,
	0x76, 0x28, 0x47, 0x82, 0x51, 0x81, 0x51, 0x83, 0x27, 0x08, 0xc6, 0x15, 0x12, 0xe0, 0x62, 0x0e,
	0xce, 0x4c, 0x97, 0x60, 0x02, 0x8b, 0x82, 0x98, 0x42, 0x62, 0x5c, 0x6c, 0x01, 0xa5, 0x49, 0xd9,
	0xa9, 0x95, 0x12, 0xcc, 0x60, 0x41, 0x28, 0x4f, 0x48, 0x86, 0x8b, 0x33, 0x24, 0x33, 0x37, 0xb5,
	0xb8, 0x24, 0x31, 0xb7, 0x40, 0x82, 0x05, 0x28, 0xc5, 0x1c, 0x84, 0x10, 0x48, 0x62, 0x03, 0x5b,
	0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xbc, 0x63, 0x5b, 0x98, 0x00, 0x00, 0x00,
}
