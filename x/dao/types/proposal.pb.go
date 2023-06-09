// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lordverse/dao/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type ProposalStatus int32

const (
	ProposalStatus_PROPOSAL_STATUS_VOTING    ProposalStatus = 0
	ProposalStatus_PROPOSAL_STATUS_SUCCESS   ProposalStatus = 1
	ProposalStatus_PROPOSAL_STATUS_FAILED    ProposalStatus = 2
	ProposalStatus_PROPOSAL_STATUS_CANCELLED ProposalStatus = 3
)

var ProposalStatus_name = map[int32]string{
	0: "PROPOSAL_STATUS_VOTING",
	1: "PROPOSAL_STATUS_SUCCESS",
	2: "PROPOSAL_STATUS_FAILED",
	3: "PROPOSAL_STATUS_CANCELLED",
}

var ProposalStatus_value = map[string]int32{
	"PROPOSAL_STATUS_VOTING":    0,
	"PROPOSAL_STATUS_SUCCESS":   1,
	"PROPOSAL_STATUS_FAILED":    2,
	"PROPOSAL_STATUS_CANCELLED": 3,
}

func (x ProposalStatus) String() string {
	return proto.EnumName(ProposalStatus_name, int32(x))
}

func (ProposalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f96a18dd8ede786, []int{0}
}

type VoteType int32

const (
	VoteType_VOTE_TYPE_ABSTAIN VoteType = 0
	VoteType_VOTE_TYPE_YES     VoteType = 1
	VoteType_VOTE_TYPE_NO      VoteType = 2
)

var VoteType_name = map[int32]string{
	0: "VOTE_TYPE_ABSTAIN",
	1: "VOTE_TYPE_YES",
	2: "VOTE_TYPE_NO",
}

var VoteType_value = map[string]int32{
	"VOTE_TYPE_ABSTAIN": 0,
	"VOTE_TYPE_YES":     1,
	"VOTE_TYPE_NO":      2,
}

func (x VoteType) String() string {
	return proto.EnumName(VoteType_name, int32(x))
}

func (VoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f96a18dd8ede786, []int{1}
}

type Proposal struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// id of the warehouse
	Warehouse  uint64     `protobuf:"varint,4,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
	Expiration *time.Time `protobuf:"bytes,5,opt,name=expiration,proto3,stdtime" json:"expiration,omitempty"`
	// ids of the voted voters
	VotedVoters map[uint64]VoteType `protobuf:"bytes,6,rep,name=votedVoters,proto3" json:"votedVoters,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=lordverse.dao.VoteType"`
	// total votes of different vote type
	Votes   []uint64       `protobuf:"varint,7,rep,packed,name=votes,proto3" json:"votes,omitempty"`
	Status  ProposalStatus `protobuf:"varint,8,opt,name=status,proto3,enum=lordverse.dao.ProposalStatus" json:"status,omitempty"`
	Creator string         `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f96a18dd8ede786, []int{0}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return m.Size()
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Proposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Proposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Proposal) GetWarehouse() uint64 {
	if m != nil {
		return m.Warehouse
	}
	return 0
}

func (m *Proposal) GetExpiration() *time.Time {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *Proposal) GetVotedVoters() map[uint64]VoteType {
	if m != nil {
		return m.VotedVoters
	}
	return nil
}

func (m *Proposal) GetVotes() []uint64 {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *Proposal) GetStatus() ProposalStatus {
	if m != nil {
		return m.Status
	}
	return ProposalStatus_PROPOSAL_STATUS_VOTING
}

func (m *Proposal) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterEnum("lordverse.dao.ProposalStatus", ProposalStatus_name, ProposalStatus_value)
	proto.RegisterEnum("lordverse.dao.VoteType", VoteType_name, VoteType_value)
	proto.RegisterType((*Proposal)(nil), "lordverse.dao.Proposal")
	proto.RegisterMapType((map[uint64]VoteType)(nil), "lordverse.dao.Proposal.VotedVotersEntry")
}

func init() { proto.RegisterFile("lordverse/dao/proposal.proto", fileDescriptor_6f96a18dd8ede786) }

var fileDescriptor_6f96a18dd8ede786 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0xaf, 0x93, 0xb6, 0x6b, 0x5f, 0x59, 0x95, 0x59, 0x1b, 0x0b, 0x65, 0xcb, 0x22, 0x4e, 0xd1,
	0x24, 0x12, 0xa9, 0x08, 0x09, 0x71, 0x22, 0x2d, 0x19, 0x2a, 0xaa, 0xda, 0x2a, 0xc9, 0x8a, 0xc6,
	0xa5, 0xca, 0x16, 0x53, 0x22, 0xba, 0x39, 0xb2, 0xdd, 0xb2, 0xde, 0xf9, 0x00, 0xfb, 0x44, 0x9c,
	0x39, 0xee, 0xc8, 0x0d, 0xd4, 0x7e, 0x11, 0x94, 0xa4, 0x25, 0x5d, 0x05, 0xb7, 0xf7, 0x7e, 0x7f,
	0x9e, 0x7f, 0x7e, 0xb2, 0xe1, 0x68, 0x42, 0x59, 0x38, 0x23, 0x8c, 0x13, 0x2b, 0x0c, 0xa8, 0x15,
	0x33, 0x1a, 0x53, 0x1e, 0x4c, 0xcc, 0x98, 0x51, 0x41, 0xf1, 0xee, 0x5f, 0xd6, 0x0c, 0x03, 0xda,
	0x38, 0x19, 0x53, 0x3a, 0x9e, 0x10, 0x2b, 0x25, 0x2f, 0xa7, 0x9f, 0x2c, 0x11, 0x5d, 0x13, 0x2e,
	0x82, 0xeb, 0x38, 0xd3, 0x37, 0xf6, 0xc7, 0x74, 0x4c, 0xd3, 0xd2, 0x4a, 0xaa, 0x0c, 0x7d, 0xf6,
	0x5d, 0x86, 0xca, 0x60, 0x35, 0x18, 0xd7, 0x41, 0x8a, 0x42, 0x15, 0xe9, 0xc8, 0x28, 0xba, 0x52,
	0x14, 0xe2, 0x7d, 0x28, 0x89, 0x48, 0x4c, 0x88, 0x2a, 0xe9, 0xc8, 0xa8, 0xba, 0x59, 0x83, 0x75,
	0xa8, 0x85, 0x84, 0x5f, 0xb1, 0x28, 0x16, 0x11, 0xbd, 0x51, 0xe5, 0x94, 0xdb, 0x84, 0xf0, 0x11,
	0x54, 0xbf, 0x06, 0x8c, 0x7c, 0xa6, 0x53, 0x4e, 0xd4, 0x62, 0x3a, 0x2e, 0x07, 0xf0, 0x1b, 0x00,
	0x72, 0x1b, 0x47, 0x2c, 0x48, 0xed, 0x25, 0x1d, 0x19, 0xb5, 0x66, 0xc3, 0xcc, 0xe2, 0x9b, 0xeb,
	0xf8, 0xa6, 0xbf, 0x8e, 0xdf, 0x2a, 0xde, 0xfd, 0x3a, 0x41, 0xee, 0x86, 0x07, 0xbf, 0x87, 0xda,
	0x8c, 0x0a, 0x12, 0x0e, 0xa9, 0x20, 0x8c, 0xab, 0x65, 0x5d, 0x36, 0x6a, 0x4d, 0xc3, 0x7c, 0xb0,
	0x10, 0x73, 0x7d, 0x2b, 0x73, 0x98, 0x4b, 0x9d, 0x1b, 0xc1, 0xe6, 0xee, 0xa6, 0x39, 0xb9, 0x63,
	0xd2, 0x72, 0x75, 0x47, 0x97, 0x8d, 0xa2, 0x9b, 0x35, 0xf8, 0x25, 0x94, 0xb9, 0x08, 0xc4, 0x94,
	0xab, 0x15, 0x1d, 0x19, 0xf5, 0xe6, 0xf1, 0x7f, 0x86, 0x7b, 0xa9, 0xc8, 0x5d, 0x89, 0xb1, 0x0a,
	0x3b, 0x57, 0x8c, 0x04, 0x82, 0x32, 0xb5, 0x9a, 0xae, 0x65, 0xdd, 0x36, 0x3e, 0x80, 0xb2, 0x9d,
	0x03, 0x2b, 0x20, 0x7f, 0x21, 0xf3, 0xd5, 0xbe, 0x93, 0x12, 0x3f, 0x87, 0xd2, 0x2c, 0x98, 0x4c,
	0xb3, 0x85, 0xd7, 0x9b, 0x87, 0x5b, 0xa7, 0x26, 0x66, 0x7f, 0x1e, 0x13, 0x37, 0x53, 0xbd, 0x96,
	0x5e, 0xa1, 0xd3, 0x6f, 0x08, 0xea, 0x0f, 0xd3, 0xe0, 0x06, 0x3c, 0x1e, 0xb8, 0xfd, 0x41, 0xdf,
	0xb3, 0xbb, 0x23, 0xcf, 0xb7, 0xfd, 0x73, 0x6f, 0x34, 0xec, 0xfb, 0x9d, 0xde, 0x3b, 0xa5, 0x80,
	0x9f, 0xc2, 0xe1, 0x36, 0xe7, 0x9d, 0xb7, 0xdb, 0x8e, 0xe7, 0x29, 0xe8, 0x5f, 0xc6, 0x33, 0xbb,
	0xd3, 0x75, 0xde, 0x2a, 0x12, 0x3e, 0x86, 0x27, 0xdb, 0x5c, 0xdb, 0xee, 0xb5, 0x9d, 0x6e, 0x42,
	0xcb, 0xa7, 0x67, 0x50, 0x59, 0xa7, 0xc3, 0x07, 0xb0, 0x37, 0xec, 0xfb, 0xce, 0xc8, 0xbf, 0x18,
	0x38, 0x23, 0xbb, 0xe5, 0xf9, 0x76, 0xa7, 0xa7, 0x14, 0xf0, 0x1e, 0xec, 0xe6, 0xf0, 0x85, 0x93,
	0x1c, 0xa8, 0xc0, 0xa3, 0x1c, 0xea, 0xf5, 0x15, 0xa9, 0x65, 0xfd, 0x58, 0x68, 0xe8, 0x7e, 0xa1,
	0xa1, 0xdf, 0x0b, 0x0d, 0xdd, 0x2d, 0xb5, 0xc2, 0xfd, 0x52, 0x2b, 0xfc, 0x5c, 0x6a, 0x85, 0x8f,
	0x07, 0xf9, 0x6f, 0xb8, 0x4d, 0xff, 0x83, 0x98, 0xc7, 0x84, 0x5f, 0x96, 0xd3, 0x17, 0xf3, 0xe2,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x46, 0x65, 0x7e, 0x2d, 0x03, 0x00, 0x00,
}

func (m *Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Status != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Votes) > 0 {
		dAtA2 := make([]byte, len(m.Votes)*10)
		var j1 int
		for _, num := range m.Votes {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintProposal(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.VotedVoters) > 0 {
		for k := range m.VotedVoters {
			v := m.VotedVoters[k]
			baseI := i
			i = encodeVarintProposal(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i = encodeVarintProposal(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintProposal(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Expiration != nil {
		n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.Expiration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiration):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintProposal(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x2a
	}
	if m.Warehouse != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Warehouse))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovProposal(uint64(m.Id))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Warehouse != 0 {
		n += 1 + sovProposal(uint64(m.Warehouse))
	}
	if m.Expiration != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiration)
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.VotedVoters) > 0 {
		for k, v := range m.VotedVoters {
			_ = k
			_ = v
			mapEntrySize := 1 + sovProposal(uint64(k)) + 1 + sovProposal(uint64(v))
			n += mapEntrySize + 1 + sovProposal(uint64(mapEntrySize))
		}
	}
	if len(m.Votes) > 0 {
		l = 0
		for _, e := range m.Votes {
			l += sovProposal(uint64(e))
		}
		n += 1 + sovProposal(uint64(l)) + l
	}
	if m.Status != 0 {
		n += 1 + sovProposal(uint64(m.Status))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Warehouse", wireType)
			}
			m.Warehouse = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Warehouse |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotedVoters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VotedVoters == nil {
				m.VotedVoters = make(map[uint64]VoteType)
			}
			var mapkey uint64
			var mapvalue VoteType
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProposal
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProposal
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProposal
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= VoteType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipProposal(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthProposal
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.VotedVoters[mapkey] = mapvalue
			iNdEx = postIndex
		case 7:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProposal
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Votes = append(m.Votes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProposal
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthProposal
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthProposal
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Votes) == 0 {
					m.Votes = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProposal
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Votes = append(m.Votes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ProposalStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
