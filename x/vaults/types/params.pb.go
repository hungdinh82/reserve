// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reserve/vaults/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/onomyprotocol/reserve/x/oracle/types"
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

// Params defines the parameters for the module.
type Params struct {
	MintingFee            cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=minting_fee,json=mintingFee,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"minting_fee"`
	StabilityFee          cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=stability_fee,json=stabilityFee,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"stability_fee"`
	LiquidationPenalty    cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=liquidation_penalty,json=liquidationPenalty,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"liquidation_penalty"`
	MinInitialDebt        cosmossdk_io_math.Int       `protobuf:"bytes,4,opt,name=min_initial_debt,json=minInitialDebt,proto3,customtype=cosmossdk.io/math.Int" json:"min_initial_debt"`
	RecalculateDebtPeriod uint64                      `protobuf:"varint,5,opt,name=recalculate_debt_period,json=recalculateDebtPeriod,proto3" json:"recalculate_debt_period,omitempty"`
	LiquidatePeriod       uint64                      `protobuf:"varint,6,opt,name=liquidate_period,json=liquidatePeriod,proto3" json:"liquidate_period,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f12ab0d072f9f7c, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetRecalculateDebtPeriod() uint64 {
	if m != nil {
		return m.RecalculateDebtPeriod
	}
	return 0
}

func (m *Params) GetLiquidatePeriod() uint64 {
	if m != nil {
		return m.LiquidatePeriod
	}
	return 0
}

// VaultParams defines the parameters for each collateral vault type.
type VaultMamagerParams struct {
	MinCollateralRatio cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=min_collateral_ratio,json=minCollateralRatio,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"min_collateral_ratio"`
	LiquidationRatio   cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=liquidation_ratio,json=liquidationRatio,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"liquidation_ratio"`
	MaxDebt            cosmossdk_io_math.Int       `protobuf:"bytes,3,opt,name=max_debt,json=maxDebt,proto3,customtype=cosmossdk.io/math.Int" json:"max_debt"`
}

func (m *VaultMamagerParams) Reset()         { *m = VaultMamagerParams{} }
func (m *VaultMamagerParams) String() string { return proto.CompactTextString(m) }
func (*VaultMamagerParams) ProtoMessage()    {}
func (*VaultMamagerParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f12ab0d072f9f7c, []int{1}
}
func (m *VaultMamagerParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultMamagerParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultMamagerParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultMamagerParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultMamagerParams.Merge(m, src)
}
func (m *VaultMamagerParams) XXX_Size() int {
	return m.Size()
}
func (m *VaultMamagerParams) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultMamagerParams.DiscardUnknown(m)
}

var xxx_messageInfo_VaultMamagerParams proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "reserve.vaults.Params")
	proto.RegisterType((*VaultMamagerParams)(nil), "reserve.vaults.VaultMamagerParams")
}

func init() { proto.RegisterFile("reserve/vaults/params.proto", fileDescriptor_1f12ab0d072f9f7c) }

var fileDescriptor_1f12ab0d072f9f7c = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xbf, 0x6e, 0xd3, 0x50,
	0x14, 0xc6, 0xe3, 0x36, 0x04, 0xb8, 0x40, 0x49, 0x4d, 0xab, 0x86, 0x54, 0x72, 0xa2, 0x4e, 0xa5,
	0x12, 0xb9, 0x54, 0x48, 0x1d, 0x18, 0x4b, 0x84, 0x14, 0x01, 0x52, 0x94, 0x01, 0xa4, 0x32, 0x58,
	0xc7, 0x37, 0x07, 0xe7, 0x8a, 0xfb, 0x27, 0xd8, 0x37, 0x51, 0xf2, 0x06, 0x88, 0x89, 0x47, 0x60,
	0x64, 0xcc, 0xc0, 0x43, 0x74, 0xac, 0x98, 0x10, 0x43, 0x85, 0x92, 0xa1, 0x3c, 0x06, 0xf2, 0xbd,
	0x4e, 0x64, 0xd4, 0x8d, 0x2c, 0x96, 0x7d, 0xbe, 0x73, 0x7e, 0x9f, 0x7d, 0xcf, 0x67, 0xb2, 0x9f,
	0x60, 0x8a, 0xc9, 0x18, 0xe9, 0x18, 0x46, 0xc2, 0xa4, 0x74, 0x08, 0x09, 0xc8, 0xb4, 0x35, 0x4c,
	0xb4, 0xd1, 0xfe, 0x56, 0x2e, 0xb6, 0x9c, 0x58, 0xdf, 0x89, 0x75, 0xac, 0xad, 0x44, 0xb3, 0x3b,
	0xd7, 0x55, 0x5f, 0x21, 0x74, 0x02, 0x4c, 0xe0, 0x3f, 0x88, 0xfa, 0x36, 0x48, 0xae, 0x34, 0xb5,
	0xd7, 0xbc, 0x14, 0x30, 0x9d, 0x4a, 0x9d, 0xd2, 0x08, 0x52, 0xa4, 0xe3, 0xe3, 0x08, 0x0d, 0x1c,
	0x53, 0xa6, 0xb9, 0xca, 0xf5, 0x87, 0x4e, 0x0f, 0x9d, 0x91, 0x7b, 0x70, 0xd2, 0xc1, 0xa7, 0x32,
	0xa9, 0x74, 0x2d, 0xde, 0x7f, 0x4b, 0xee, 0x48, 0xae, 0x0c, 0x57, 0x71, 0xf8, 0x1e, 0xb1, 0xe6,
	0x35, 0xbd, 0xc3, 0xdb, 0xa7, 0x27, 0xe7, 0x97, 0x8d, 0xd2, 0xaf, 0xcb, 0xc6, 0xbe, 0x9b, 0x4a,
	0xfb, 0x1f, 0x5a, 0x5c, 0x53, 0x09, 0x66, 0xd0, 0x7a, 0x85, 0x31, 0xb0, 0x69, 0x1b, 0xd9, 0x8f,
	0xef, 0x8f, 0x49, 0x0e, 0x6d, 0x23, 0xfb, 0x76, 0x35, 0x3b, 0xf2, 0x7a, 0x24, 0x47, 0xbd, 0x40,
	0xf4, 0xdf, 0x91, 0x7b, 0xa9, 0x81, 0x88, 0x0b, 0x6e, 0xa6, 0x16, 0xbd, 0xb1, 0x16, 0xfa, 0xee,
	0x0a, 0x96, 0xc1, 0x63, 0xf2, 0x40, 0xf0, 0x8f, 0x23, 0xde, 0x07, 0xc3, 0xb5, 0x0a, 0x87, 0xa8,
	0x40, 0x98, 0x69, 0x6d, 0x73, 0x2d, 0x0b, 0xbf, 0x80, 0xec, 0x3a, 0xa2, 0x7f, 0x46, 0xaa, 0x92,
	0xab, 0x90, 0x2b, 0x6e, 0x38, 0x88, 0xb0, 0x8f, 0x91, 0xa9, 0x95, 0xad, 0xcb, 0x93, 0xdc, 0x65,
	0xf7, 0xba, 0x4b, 0x47, 0x99, 0x02, 0xbf, 0xa3, 0x8c, 0xe3, 0x6f, 0x49, 0xae, 0x3a, 0x0e, 0xd4,
	0xc6, 0xc8, 0xf8, 0x27, 0x64, 0x2f, 0x41, 0x06, 0x82, 0x8d, 0x04, 0x18, 0xb4, 0xec, 0x70, 0x88,
	0x09, 0xd7, 0xfd, 0xda, 0x8d, 0xa6, 0x77, 0x58, 0xee, 0xed, 0x16, 0xe4, 0x6c, 0xa2, 0x6b, 0x45,
	0xff, 0x11, 0xa9, 0x2e, 0xdf, 0x14, 0x97, 0x03, 0x15, 0x3b, 0x70, 0x7f, 0x55, 0x77, 0xad, 0xcf,
	0x9a, 0x7f, 0xbe, 0x36, 0xbc, 0xcf, 0x57, 0xb3, 0xa3, 0xbd, 0x65, 0xb8, 0x26, 0xcb, 0x84, 0xba,
	0xfd, 0x1f, 0xcc, 0x36, 0x88, 0xff, 0x26, 0xab, 0xbc, 0x06, 0x09, 0x31, 0x26, 0x79, 0x2c, 0x06,
	0x64, 0x27, 0xfb, 0x6e, 0xa6, 0x45, 0x66, 0x9e, 0x80, 0x08, 0x93, 0xec, 0x5c, 0xd6, 0xcc, 0x87,
	0x2f, 0xb9, 0x7a, 0xbe, 0x42, 0xf6, 0x32, 0xa2, 0xcf, 0xc8, 0x76, 0x71, 0x95, 0xce, 0x66, 0xbd,
	0xac, 0x54, 0x0b, 0x40, 0x67, 0xf2, 0x92, 0xdc, 0x92, 0x30, 0x71, 0xeb, 0xdb, 0xfc, 0xcf, 0xf5,
	0xdd, 0x94, 0x30, 0xc9, 0xb6, 0x70, 0xda, 0x39, 0x9f, 0x07, 0xde, 0xc5, 0x3c, 0xf0, 0x7e, 0xcf,
	0x03, 0xef, 0xcb, 0x22, 0x28, 0x5d, 0x2c, 0x82, 0xd2, 0xcf, 0x45, 0x50, 0x3a, 0xa3, 0x31, 0x37,
	0x83, 0x51, 0xd4, 0x62, 0x5a, 0x52, 0xad, 0xb4, 0x9c, 0xda, 0xdf, 0x8d, 0x69, 0x41, 0xaf, 0x1d,
	0xbf, 0x99, 0x0e, 0x31, 0x8d, 0x2a, 0xb6, 0xe1, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12,
	0xd4, 0xe0, 0xb6, 0x3f, 0x04, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if !this.MintingFee.Equal(that1.MintingFee) {
		return false
	}
	if !this.StabilityFee.Equal(that1.StabilityFee) {
		return false
	}
	if !this.LiquidationPenalty.Equal(that1.LiquidationPenalty) {
		return false
	}
	if !this.MinInitialDebt.Equal(that1.MinInitialDebt) {
		return false
	}
	if this.RecalculateDebtPeriod != that1.RecalculateDebtPeriod {
		return false
	}
	if this.LiquidatePeriod != that1.LiquidatePeriod {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LiquidatePeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.LiquidatePeriod))
		i--
		dAtA[i] = 0x30
	}
	if m.RecalculateDebtPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RecalculateDebtPeriod))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.MinInitialDebt.Size()
		i -= size
		if _, err := m.MinInitialDebt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.LiquidationPenalty.Size()
		i -= size
		if _, err := m.LiquidationPenalty.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.StabilityFee.Size()
		i -= size
		if _, err := m.StabilityFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MintingFee.Size()
		i -= size
		if _, err := m.MintingFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VaultMamagerParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultMamagerParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultMamagerParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxDebt.Size()
		i -= size
		if _, err := m.MaxDebt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.LiquidationRatio.Size()
		i -= size
		if _, err := m.LiquidationRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinCollateralRatio.Size()
		i -= size
		if _, err := m.MinCollateralRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MintingFee.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.StabilityFee.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.LiquidationPenalty.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinInitialDebt.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.RecalculateDebtPeriod != 0 {
		n += 1 + sovParams(uint64(m.RecalculateDebtPeriod))
	}
	if m.LiquidatePeriod != 0 {
		n += 1 + sovParams(uint64(m.LiquidatePeriod))
	}
	return n
}

func (m *VaultMamagerParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinCollateralRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.LiquidationRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxDebt.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintingFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MintingFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StabilityFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StabilityFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationPenalty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationPenalty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitialDebt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinInitialDebt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecalculateDebtPeriod", wireType)
			}
			m.RecalculateDebtPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecalculateDebtPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidatePeriod", wireType)
			}
			m.LiquidatePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LiquidatePeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *VaultMamagerParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: VaultMamagerParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultMamagerParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCollateralRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinCollateralRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDebt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxDebt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
