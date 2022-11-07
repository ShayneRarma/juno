// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: juno/globalfee/whitelist.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// TODO:
type AccountRecord struct {
	Account         string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	AllowedBypasses uint64 `protobuf:"varint,2,opt,name=allowed_bypasses,json=allowedBypasses,proto3" json:"allowed_bypasses,omitempty"`
}

func (m *AccountRecord) Reset()         { *m = AccountRecord{} }
func (m *AccountRecord) String() string { return proto.CompactTextString(m) }
func (*AccountRecord) ProtoMessage()    {}
func (*AccountRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_112ab6bf3d141450, []int{0}
}
func (m *AccountRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRecord.Merge(m, src)
}
func (m *AccountRecord) XXX_Size() int {
	return m.Size()
}
func (m *AccountRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRecord proto.InternalMessageInfo

func (m *AccountRecord) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *AccountRecord) GetAllowedBypasses() uint64 {
	if m != nil {
		return m.AllowedBypasses
	}
	return 0
}

// WhitelistAccounts is the storage mechanism for many possible accounts & the
// number of times they can bypass the fee (given via governance)
type WhitelistAccounts struct {
	// TODO:
	Records []*AccountRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (m *WhitelistAccounts) Reset()         { *m = WhitelistAccounts{} }
func (m *WhitelistAccounts) String() string { return proto.CompactTextString(m) }
func (*WhitelistAccounts) ProtoMessage()    {}
func (*WhitelistAccounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_112ab6bf3d141450, []int{1}
}
func (m *WhitelistAccounts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhitelistAccounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhitelistAccounts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhitelistAccounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhitelistAccounts.Merge(m, src)
}
func (m *WhitelistAccounts) XXX_Size() int {
	return m.Size()
}
func (m *WhitelistAccounts) XXX_DiscardUnknown() {
	xxx_messageInfo_WhitelistAccounts.DiscardUnknown(m)
}

var xxx_messageInfo_WhitelistAccounts proto.InternalMessageInfo

func (m *WhitelistAccounts) GetRecords() []*AccountRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountRecord)(nil), "juno.globalfee.AccountRecord")
	proto.RegisterType((*WhitelistAccounts)(nil), "juno.globalfee.WhitelistAccounts")
}

func init() { proto.RegisterFile("juno/globalfee/whitelist.proto", fileDescriptor_112ab6bf3d141450) }

var fileDescriptor_112ab6bf3d141450 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x2a, 0xcd, 0xcb,
	0xd7, 0x4f, 0xcf, 0xc9, 0x4f, 0x4a, 0xcc, 0x49, 0x4b, 0x4d, 0xd5, 0x2f, 0xcf, 0xc8, 0x2c, 0x49,
	0xcd, 0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0xc9, 0xeb, 0xc1,
	0xe5, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x52, 0xfa, 0x20, 0x16, 0x44, 0x95, 0x52, 0x1c,
	0x17, 0xaf, 0x63, 0x72, 0x72, 0x7e, 0x69, 0x5e, 0x49, 0x50, 0x6a, 0x72, 0x7e, 0x51, 0x8a, 0x90,
	0x04, 0x17, 0x7b, 0x22, 0x44, 0x40, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0xd2,
	0xe4, 0x12, 0x48, 0xcc, 0xc9, 0xc9, 0x2f, 0x4f, 0x4d, 0x89, 0x4f, 0xaa, 0x2c, 0x48, 0x2c, 0x2e,
	0x4e, 0x2d, 0x96, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x09, 0xe2, 0x87, 0x8a, 0x3b, 0x41, 0x85, 0xad,
	0x38, 0x3a, 0x16, 0xc8, 0x33, 0xce, 0x58, 0x20, 0xcf, 0xa8, 0xe4, 0xc3, 0x25, 0x18, 0x0e, 0x73,
	0x18, 0xd4, 0xa2, 0x62, 0x21, 0x73, 0x2e, 0xf6, 0x22, 0xb0, 0x6d, 0xc5, 0x12, 0x8c, 0x0a, 0xcc,
	0x1a, 0xdc, 0x46, 0xb2, 0x7a, 0xa8, 0x8e, 0xd5, 0x43, 0x71, 0x53, 0x10, 0x4c, 0xb5, 0x93, 0xf7,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c,
	0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19, 0xa6, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x3b, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0x3b, 0xe7, 0xe7, 0x95,
	0x14, 0x25, 0x26, 0x97, 0x14, 0xeb, 0x83, 0x03, 0xaa, 0x02, 0x29, 0xa8, 0x4a, 0x2a, 0x0b, 0x52,
	0x8b, 0x93, 0xd8, 0xc0, 0x21, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x80, 0x21, 0xa7, 0x14,
	0x49, 0x01, 0x00, 0x00,
}

func (m *AccountRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AllowedBypasses != 0 {
		i = encodeVarintWhitelist(dAtA, i, uint64(m.AllowedBypasses))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintWhitelist(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WhitelistAccounts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhitelistAccounts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhitelistAccounts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWhitelist(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhitelist(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhitelist(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovWhitelist(uint64(l))
	}
	if m.AllowedBypasses != 0 {
		n += 1 + sovWhitelist(uint64(m.AllowedBypasses))
	}
	return n
}

func (m *WhitelistAccounts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovWhitelist(uint64(l))
		}
	}
	return n
}

func sovWhitelist(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhitelist(x uint64) (n int) {
	return sovWhitelist(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhitelist
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
			return fmt.Errorf("proto: AccountRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWhitelist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedBypasses", wireType)
			}
			m.AllowedBypasses = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllowedBypasses |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWhitelist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhitelist
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
func (m *WhitelistAccounts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhitelist
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
			return fmt.Errorf("proto: WhitelistAccounts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhitelistAccounts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWhitelist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, &AccountRecord{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhitelist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhitelist
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
func skipWhitelist(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhitelist
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
					return 0, ErrIntOverflowWhitelist
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
					return 0, ErrIntOverflowWhitelist
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
				return 0, ErrInvalidLengthWhitelist
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhitelist
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhitelist
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhitelist        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhitelist          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhitelist = fmt.Errorf("proto: unexpected end of group")
)
