// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: authMessage.proto

package data

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type AuthMessagePb struct {
	Message   []byte `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Sig       []byte `protobuf:"bytes,2,opt,name=Sig,proto3" json:"Sig,omitempty"`
	Pubkey    []byte `protobuf:"bytes,3,opt,name=Pubkey,proto3" json:"Pubkey,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (m *AuthMessagePb) Reset()      { *m = AuthMessagePb{} }
func (*AuthMessagePb) ProtoMessage() {}
func (*AuthMessagePb) Descriptor() ([]byte, []int) {
	return fileDescriptor_e29ffa46745b2d90, []int{0}
}
func (m *AuthMessagePb) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthMessagePb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *AuthMessagePb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthMessagePb.Merge(m, src)
}
func (m *AuthMessagePb) XXX_Size() int {
	return m.Size()
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
	proto.RegisterType((*AuthMessagePb)(nil), "proto.AuthMessagePb")
}

func init() { proto.RegisterFile("authMessage.proto", fileDescriptor_e29ffa46745b2d90) }

var fileDescriptor_e29ffa46745b2d90 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2c, 0x2d, 0xc9,
	0xf0, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x53, 0x52, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9,
	0xe9, 0xf9, 0xfa, 0x60, 0xe1, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc, 0x01, 0xb3, 0x20, 0xba, 0x94,
	0x0a, 0xb9, 0x78, 0x1d, 0x11, 0x46, 0x05, 0x24, 0x09, 0x49, 0x70, 0xb1, 0x43, 0x39, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x30, 0xae, 0x90, 0x00, 0x17, 0x73, 0x70, 0x66, 0xba, 0x04, 0x13,
	0x58, 0x14, 0xc4, 0x14, 0x12, 0xe3, 0x62, 0x0b, 0x28, 0x4d, 0xca, 0x4e, 0xad, 0x94, 0x60, 0x06,
	0x0b, 0x42, 0x79, 0x42, 0x32, 0x5c, 0x9c, 0x21, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05,
	0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x08, 0x01, 0x27, 0xbb, 0x0b, 0x0f, 0xe5, 0x18, 0x6e,
	0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50, 0x8e, 0xb1, 0xe1, 0x91, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0x78, 0xe3, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72,
	0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xb1, 0xa4, 0x24, 0x96, 0x24, 0x26, 0xb1, 0x81, 0x5d,
	0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x78, 0x6f, 0xdd, 0x30, 0x04, 0x01, 0x00, 0x00,
}

func (this *AuthMessagePb) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuthMessagePb)
	if !ok {
		that2, ok := that.(AuthMessagePb)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Message, that1.Message) {
		return false
	}
	if !bytes.Equal(this.Sig, that1.Sig) {
		return false
	}
	if !bytes.Equal(this.Pubkey, that1.Pubkey) {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	return true
}
func (this *AuthMessagePb) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&data.AuthMessagePb{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "Sig: "+fmt.Sprintf("%#v", this.Sig)+",\n")
	s = append(s, "Pubkey: "+fmt.Sprintf("%#v", this.Pubkey)+",\n")
	s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAuthMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AuthMessagePb) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthMessagePb) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthMessagePb) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintAuthMessage(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintAuthMessage(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVarintAuthMessage(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintAuthMessage(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuthMessagePb) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovAuthMessage(uint64(l))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovAuthMessage(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovAuthMessage(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovAuthMessage(uint64(m.Timestamp))
	}
	return n
}

func sovAuthMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthMessage(x uint64) (n int) {
	return sovAuthMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AuthMessagePb) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AuthMessagePb{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`Sig:` + fmt.Sprintf("%v", this.Sig) + `,`,
		`Pubkey:` + fmt.Sprintf("%v", this.Pubkey) + `,`,
		`Timestamp:` + fmt.Sprintf("%v", this.Timestamp) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAuthMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AuthMessagePb) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthMessagePb: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthMessagePb: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAuthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message[:0], dAtA[iNdEx:postIndex]...)
			if m.Message == nil {
				m.Message = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAuthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAuthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = append(m.Pubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.Pubkey == nil {
				m.Pubkey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuthMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAuthMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuthMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuthMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAuthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthMessage = fmt.Errorf("proto: unexpected end of group")
)
