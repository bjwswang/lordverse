// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lordverse/dao/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the dao module's genesis state.
type GenesisState struct {
	Params         Params      `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	WarehouseList  []Warehouse `protobuf:"bytes,2,rep,name=warehouseList,proto3" json:"warehouseList"`
	WarehouseCount uint64      `protobuf:"varint,3,opt,name=warehouseCount,proto3" json:"warehouseCount,omitempty"`
	ProposalList   []Proposal  `protobuf:"bytes,6,rep,name=proposalList,proto3" json:"proposalList"`
	ProposalCount  uint64      `protobuf:"varint,7,opt,name=proposalCount,proto3" json:"proposalCount,omitempty"`
	VoterList      []Voter     `protobuf:"bytes,8,rep,name=voterList,proto3" json:"voterList"`
	VoterCount     uint64      `protobuf:"varint,9,opt,name=voterCount,proto3" json:"voterCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f147d1202479bed, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetWarehouseList() []Warehouse {
	if m != nil {
		return m.WarehouseList
	}
	return nil
}

func (m *GenesisState) GetWarehouseCount() uint64 {
	if m != nil {
		return m.WarehouseCount
	}
	return 0
}

func (m *GenesisState) GetProposalList() []Proposal {
	if m != nil {
		return m.ProposalList
	}
	return nil
}

func (m *GenesisState) GetProposalCount() uint64 {
	if m != nil {
		return m.ProposalCount
	}
	return 0
}

func (m *GenesisState) GetVoterList() []Voter {
	if m != nil {
		return m.VoterList
	}
	return nil
}

func (m *GenesisState) GetVoterCount() uint64 {
	if m != nil {
		return m.VoterCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lordverse.dao.GenesisState")
}

func init() { proto.RegisterFile("lordverse/dao/genesis.proto", fileDescriptor_2f147d1202479bed) }

var fileDescriptor_2f147d1202479bed = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4e, 0x02, 0x31,
	0x10, 0x86, 0xb7, 0x40, 0x50, 0x0a, 0x78, 0x68, 0x20, 0xae, 0xab, 0x56, 0x62, 0x8c, 0xe1, 0xb4,
	0x9b, 0xc0, 0xc5, 0xab, 0x68, 0xe2, 0xc5, 0x83, 0xc1, 0x44, 0x13, 0x6f, 0x35, 0x34, 0x48, 0x82,
	0xcc, 0xa6, 0x2d, 0xa8, 0x6f, 0xe1, 0x63, 0x71, 0xe4, 0xe8, 0xc9, 0x18, 0x78, 0x05, 0x1f, 0xc0,
	0x30, 0x2d, 0x8b, 0xdd, 0xdb, 0xce, 0xfc, 0xff, 0x3f, 0xdf, 0x6c, 0x87, 0x1e, 0x8e, 0x41, 0x0d,
	0x66, 0x52, 0x69, 0x99, 0x0c, 0x04, 0x24, 0x43, 0x39, 0x91, 0x7a, 0xa4, 0xe3, 0x54, 0x81, 0x01,
	0x56, 0xcf, 0xc4, 0x78, 0x20, 0x20, 0x6a, 0x0c, 0x61, 0x08, 0xa8, 0x24, 0xeb, 0x2f, 0x6b, 0x8a,
	0x22, 0x7f, 0x42, 0x2a, 0x94, 0x78, 0x75, 0x03, 0xa2, 0x63, 0x5f, 0x7b, 0x13, 0x4a, 0xbe, 0xc0,
	0x54, 0x4b, 0x27, 0x1f, 0xe5, 0xa2, 0x0a, 0x52, 0xd0, 0x62, 0xec, 0xd4, 0x03, 0x5f, 0x9d, 0x81,
	0x91, 0xca, 0x4a, 0xa7, 0xbf, 0x05, 0x5a, 0xbb, 0xb1, 0xab, 0xde, 0x1b, 0x61, 0x24, 0xeb, 0xd2,
	0xb2, 0x05, 0x87, 0xa4, 0x45, 0xda, 0xd5, 0x4e, 0x33, 0xf6, 0x56, 0x8f, 0xef, 0x50, 0xec, 0x95,
	0xe6, 0xdf, 0x27, 0x41, 0xdf, 0x59, 0xd9, 0x35, 0xad, 0x67, 0x1b, 0xdd, 0x8e, 0xb4, 0x09, 0x0b,
	0xad, 0x62, 0xbb, 0xda, 0x09, 0x73, 0xd9, 0xc7, 0x8d, 0xc7, 0xc5, 0xfd, 0x10, 0x3b, 0xa7, 0x7b,
	0x59, 0xe3, 0x0a, 0xa6, 0x13, 0x13, 0x16, 0x5b, 0xa4, 0x5d, 0xea, 0xe7, 0xba, 0xec, 0x92, 0xd6,
	0x36, 0x3f, 0x88, 0xb0, 0x32, 0xc2, 0xf6, 0xf3, 0x8b, 0x3a, 0x8b, 0x63, 0x79, 0x11, 0x76, 0x46,
	0xeb, 0x9b, 0xda, 0x92, 0x76, 0x90, 0xe4, 0x37, 0xd9, 0x05, 0xad, 0xe0, 0x5b, 0x21, 0x65, 0x17,
	0x29, 0x8d, 0x1c, 0xe5, 0x61, 0xad, 0x3b, 0xc4, 0xd6, 0xcc, 0x38, 0xa5, 0x58, 0xd8, 0xe1, 0x15,
	0x1c, 0xfe, 0xaf, 0xd3, 0x4b, 0xe6, 0x4b, 0x4e, 0x16, 0x4b, 0x4e, 0x7e, 0x96, 0x9c, 0x7c, 0xae,
	0x78, 0xb0, 0x58, 0xf1, 0xe0, 0x6b, 0xc5, 0x83, 0xa7, 0xe6, 0xf6, 0x56, 0xef, 0x78, 0x2d, 0xf3,
	0x91, 0x4a, 0xfd, 0x5c, 0xc6, 0x73, 0x75, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x50, 0xd4, 0xb5,
	0x48, 0x66, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VoterCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.VoterCount))
		i--
		dAtA[i] = 0x48
	}
	if len(m.VoterList) > 0 {
		for iNdEx := len(m.VoterList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VoterList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.ProposalCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ProposalCount))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ProposalList) > 0 {
		for iNdEx := len(m.ProposalList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProposalList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.WarehouseCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.WarehouseCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.WarehouseList) > 0 {
		for iNdEx := len(m.WarehouseList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WarehouseList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.WarehouseList) > 0 {
		for _, e := range m.WarehouseList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.WarehouseCount != 0 {
		n += 1 + sovGenesis(uint64(m.WarehouseCount))
	}
	if len(m.ProposalList) > 0 {
		for _, e := range m.ProposalList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ProposalCount != 0 {
		n += 1 + sovGenesis(uint64(m.ProposalCount))
	}
	if len(m.VoterList) > 0 {
		for _, e := range m.VoterList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.VoterCount != 0 {
		n += 1 + sovGenesis(uint64(m.VoterCount))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WarehouseList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WarehouseList = append(m.WarehouseList, Warehouse{})
			if err := m.WarehouseList[len(m.WarehouseList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WarehouseCount", wireType)
			}
			m.WarehouseCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WarehouseCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposalList = append(m.ProposalList, Proposal{})
			if err := m.ProposalList[len(m.ProposalList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalCount", wireType)
			}
			m.ProposalCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoterList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoterList = append(m.VoterList, Voter{})
			if err := m.VoterList[len(m.VoterList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoterCount", wireType)
			}
			m.VoterCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoterCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
