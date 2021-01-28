// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/authz/v1beta1/authz.proto

package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SendAuthorization allows the grantee to spend up to spend_limit coins from
// the granter's account.
type SendAuthorization struct {
	SpendLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=spend_limit,json=spendLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"spend_limit"`
}

func (m *SendAuthorization) Reset()         { *m = SendAuthorization{} }
func (m *SendAuthorization) String() string { return proto.CompactTextString(m) }
func (*SendAuthorization) ProtoMessage()    {}
func (*SendAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_544dc2e84b61c637, []int{0}
}
func (m *SendAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendAuthorization.Merge(m, src)
}
func (m *SendAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *SendAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_SendAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_SendAuthorization proto.InternalMessageInfo

func (m *SendAuthorization) GetSpendLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.SpendLimit
	}
	return nil
}

// GenericAuthorization gives the grantee unrestricted permissions to execute
// the provided method on behalf of the granter's account.
type GenericAuthorization struct {
	// method name to grant unrestricted permissions to execute
	// Note: MethodName() is already a method on `GenericAuthorization` type,
	// we need some custom naming here so using `MessageName`
	MessageName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
}

func (m *GenericAuthorization) Reset()         { *m = GenericAuthorization{} }
func (m *GenericAuthorization) String() string { return proto.CompactTextString(m) }
func (*GenericAuthorization) ProtoMessage()    {}
func (*GenericAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_544dc2e84b61c637, []int{1}
}
func (m *GenericAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenericAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenericAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenericAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericAuthorization.Merge(m, src)
}
func (m *GenericAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *GenericAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_GenericAuthorization proto.InternalMessageInfo

func (m *GenericAuthorization) GetMessageName() string {
	if m != nil {
		return m.MessageName
	}
	return ""
}

// AuthorizationGrant gives permissions to execute
// the provide method with expiration time.
type AuthorizationGrant struct {
	Authorization *types1.Any `protobuf:"bytes,1,opt,name=authorization,proto3" json:"authorization,omitempty"`
	Expiration    time.Time   `protobuf:"bytes,2,opt,name=expiration,proto3,stdtime" json:"expiration"`
}

func (m *AuthorizationGrant) Reset()         { *m = AuthorizationGrant{} }
func (m *AuthorizationGrant) String() string { return proto.CompactTextString(m) }
func (*AuthorizationGrant) ProtoMessage()    {}
func (*AuthorizationGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_544dc2e84b61c637, []int{2}
}
func (m *AuthorizationGrant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthorizationGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthorizationGrant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthorizationGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationGrant.Merge(m, src)
}
func (m *AuthorizationGrant) XXX_Size() int {
	return m.Size()
}
func (m *AuthorizationGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationGrant.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationGrant proto.InternalMessageInfo

func (m *AuthorizationGrant) GetAuthorization() *types1.Any {
	if m != nil {
		return m.Authorization
	}
	return nil
}

func (m *AuthorizationGrant) GetExpiration() time.Time {
	if m != nil {
		return m.Expiration
	}
	return time.Time{}
}

// VoteAuthorization allows the grantee to vote on a proposal on behalf of the granter's account.
type VoteAuthorization struct {
}

func (m *VoteAuthorization) Reset()         { *m = VoteAuthorization{} }
func (m *VoteAuthorization) String() string { return proto.CompactTextString(m) }
func (*VoteAuthorization) ProtoMessage()    {}
func (*VoteAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_544dc2e84b61c637, []int{3}
}
func (m *VoteAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteAuthorization.Merge(m, src)
}
func (m *VoteAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *VoteAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_VoteAuthorization proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SendAuthorization)(nil), "cosmos.authz.v1beta1.SendAuthorization")
	proto.RegisterType((*GenericAuthorization)(nil), "cosmos.authz.v1beta1.GenericAuthorization")
	proto.RegisterType((*AuthorizationGrant)(nil), "cosmos.authz.v1beta1.AuthorizationGrant")
	proto.RegisterType((*VoteAuthorization)(nil), "cosmos.authz.v1beta1.VoteAuthorization")
}

func init() { proto.RegisterFile("cosmos/authz/v1beta1/authz.proto", fileDescriptor_544dc2e84b61c637) }

var fileDescriptor_544dc2e84b61c637 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0x6f, 0xd4, 0x30,
	0x18, 0xc6, 0xcf, 0x20, 0x21, 0x70, 0x54, 0xa1, 0x8b, 0x6e, 0x68, 0x6f, 0x48, 0x4e, 0x37, 0xa0,
	0x13, 0x52, 0x9d, 0xb6, 0x6c, 0x6c, 0x0d, 0x95, 0xba, 0x50, 0x86, 0x80, 0x18, 0x60, 0x38, 0x39,
	0xc9, 0x4b, 0x62, 0x71, 0xb6, 0xa3, 0xd8, 0x41, 0xbd, 0x7e, 0x8a, 0x0e, 0x7c, 0x09, 0x98, 0xf9,
	0x10, 0x15, 0x53, 0xc5, 0xc4, 0xd4, 0xa2, 0xdc, 0x17, 0x41, 0xb1, 0x1d, 0xd4, 0x94, 0xaa, 0x53,
	0xfc, 0xfe, 0x79, 0x7e, 0x79, 0xfc, 0xbe, 0xc6, 0xb3, 0x4c, 0x2a, 0x2e, 0x55, 0x44, 0x1b, 0x5d,
	0x9e, 0x45, 0x5f, 0xf6, 0x53, 0xd0, 0x74, 0xdf, 0x46, 0xa4, 0xaa, 0xa5, 0x96, 0xfe, 0xc4, 0x76,
	0x10, 0x9b, 0x73, 0x1d, 0xd3, 0xc0, 0xe9, 0x52, 0xaa, 0xe0, 0x9f, 0x2c, 0x93, 0x4c, 0x58, 0xd5,
	0x74, 0xc7, 0xd6, 0x97, 0x26, 0x8a, 0x1c, 0xc2, 0x96, 0xc2, 0x42, 0xca, 0x62, 0x05, 0x91, 0x89,
	0xd2, 0xe6, 0x53, 0xa4, 0x19, 0x07, 0xa5, 0x29, 0xaf, 0x5c, 0xc3, 0xa4, 0x90, 0x85, 0xb4, 0xc2,
	0xee, 0xd4, 0x13, 0x6f, 0xcb, 0xa8, 0x58, 0xdb, 0xd2, 0xfc, 0x2b, 0xc2, 0xe3, 0xb7, 0x20, 0xf2,
	0xc3, 0x46, 0x97, 0xb2, 0x66, 0x67, 0x54, 0x33, 0x29, 0xfc, 0x15, 0xf6, 0x54, 0x05, 0x22, 0x5f,
	0xae, 0x18, 0x67, 0x7a, 0x1b, 0xcd, 0x1e, 0x2e, 0xbc, 0x83, 0x1d, 0xe2, 0xbc, 0x74, 0xc6, 0xfb,
	0xdb, 0x90, 0x57, 0x92, 0x89, 0x78, 0xef, 0xe2, 0x2a, 0x1c, 0x7d, 0xbf, 0x0e, 0x17, 0x05, 0xd3,
	0x65, 0x93, 0x92, 0x4c, 0x72, 0x67, 0xdc, 0x7d, 0x76, 0x55, 0xfe, 0x39, 0xd2, 0xeb, 0x0a, 0x94,
	0x11, 0xa8, 0x04, 0x1b, 0xfe, 0xeb, 0x0e, 0xff, 0x72, 0xfc, 0xeb, 0xc7, 0xee, 0xd6, 0xc0, 0xc0,
	0xfc, 0x23, 0x9e, 0x1c, 0x83, 0x80, 0x9a, 0x65, 0x43, 0x63, 0x7b, 0xd8, 0xe3, 0xa0, 0x4b, 0x99,
	0x2f, 0x05, 0xe5, 0xb0, 0x8d, 0x66, 0x68, 0xf1, 0x24, 0x7e, 0xda, 0x5e, 0x85, 0xde, 0x09, 0x28,
	0x45, 0x0b, 0x78, 0x43, 0x39, 0x24, 0xd8, 0xf6, 0x74, 0xe7, 0xbb, 0xe0, 0xdf, 0x10, 0xf6, 0x07,
	0x99, 0xe3, 0x9a, 0x0a, 0xed, 0x9f, 0xe0, 0x2d, 0x7a, 0x33, 0x6b, 0xe8, 0xde, 0xc1, 0x84, 0xd8,
	0xe9, 0x91, 0x7e, 0x7a, 0xe4, 0x50, 0xac, 0xe3, 0xf1, 0xcf, 0xdb, 0xd8, 0x64, 0xa8, 0xf6, 0x8f,
	0x30, 0x86, 0xd3, 0x8a, 0xd5, 0x96, 0xf5, 0xc0, 0xb0, 0xa6, 0xff, 0xb1, 0xde, 0xf5, 0x0b, 0x8c,
	0x1f, 0x77, 0x33, 0x3c, 0xbf, 0x0e, 0x51, 0x72, 0x43, 0x37, 0x7f, 0x86, 0xc7, 0xef, 0xa5, 0x86,
	0xc1, 0x9f, 0xee, 0xb8, 0x53, 0x7c, 0x74, 0xd1, 0x06, 0xe8, 0xb2, 0x0d, 0xd0, 0x9f, 0x36, 0x40,
	0xe7, 0x9b, 0x60, 0x74, 0xb9, 0x09, 0x46, 0xbf, 0x37, 0xc1, 0xe8, 0xc3, 0xf3, 0x7b, 0x77, 0x72,
	0xea, 0x9e, 0xaf, 0xd9, 0x4d, 0xfa, 0xc8, 0xf8, 0x7a, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xcf,
	0x0b, 0x2e, 0x6f, 0xdb, 0x02, 0x00, 0x00,
}

func (m *SendAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for iNdEx := len(m.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthz(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GenericAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenericAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenericAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageName) > 0 {
		i -= len(m.MessageName)
		copy(dAtA[i:], m.MessageName)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.MessageName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthorizationGrant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthorizationGrant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthorizationGrant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintAuthz(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.Authorization != nil {
		{
			size, err := m.Authorization.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VoteAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for _, e := range m.SpendLimit {
			l = e.Size()
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func (m *GenericAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MessageName)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}

func (m *AuthorizationGrant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Authorization != nil {
		l = m.Authorization.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration)
	n += 1 + l + sovAuthz(uint64(l))
	return n
}

func (m *VoteAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: SendAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendLimit = append(m.SpendLimit, types.Coin{})
			if err := m.SpendLimit[len(m.SpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *GenericAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: GenericAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenericAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *AuthorizationGrant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: AuthorizationGrant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthorizationGrant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authorization == nil {
				m.Authorization = &types1.Any{}
			}
			if err := m.Authorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *VoteAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: VoteAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
