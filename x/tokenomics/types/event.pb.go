// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/tokenomics/event.proto

package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/pokt-network/poktroll/x/proof/types"
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

// EventClaimExpired is an event emitted during settlement whenever a claim requiring
// an on-chain proof doesn't have one. The claim cannot be settled, leading to that work
// never being rewarded.
type EventClaimExpired struct {
	Claim           *types.Claim `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim"`
	NumRelays       uint64       `protobuf:"varint,2,opt,name=num_relays,json=numRelays,proto3" json:"num_relays"`
	NumComputeUnits uint64       `protobuf:"varint,3,opt,name=num_compute_units,json=numComputeUnits,proto3" json:"num_compute_units"`
}

func (m *EventClaimExpired) Reset()         { *m = EventClaimExpired{} }
func (m *EventClaimExpired) String() string { return proto.CompactTextString(m) }
func (*EventClaimExpired) ProtoMessage()    {}
func (*EventClaimExpired) Descriptor() ([]byte, []int) {
	return fileDescriptor_a78874bbf91a58c7, []int{0}
}
func (m *EventClaimExpired) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventClaimExpired) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventClaimExpired.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventClaimExpired) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventClaimExpired.Merge(m, src)
}
func (m *EventClaimExpired) XXX_Size() int {
	return m.Size()
}
func (m *EventClaimExpired) XXX_DiscardUnknown() {
	xxx_messageInfo_EventClaimExpired.DiscardUnknown(m)
}

var xxx_messageInfo_EventClaimExpired proto.InternalMessageInfo

func (m *EventClaimExpired) GetClaim() *types.Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

func (m *EventClaimExpired) GetNumRelays() uint64 {
	if m != nil {
		return m.NumRelays
	}
	return 0
}

func (m *EventClaimExpired) GetNumComputeUnits() uint64 {
	if m != nil {
		return m.NumComputeUnits
	}
	return 0
}

// EventClaimSettled is an event emitted whenever a claim is settled.
// The proof_required determines whether the claim requires a proof that has been submitted or not
type EventClaimSettled struct {
	Claim            *types.Claim                 `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim"`
	NumRelays        uint64                       `protobuf:"varint,2,opt,name=num_relays,json=numRelays,proto3" json:"num_relays"`
	NumComputeUnits  uint64                       `protobuf:"varint,3,opt,name=num_compute_units,json=numComputeUnits,proto3" json:"num_compute_units"`
	ProofRequirement types.ProofRequirementReason `protobuf:"varint,4,opt,name=proof_requirement,json=proofRequirement,proto3,enum=poktroll.proof.ProofRequirementReason" json:"proof_requirement"`
}

func (m *EventClaimSettled) Reset()         { *m = EventClaimSettled{} }
func (m *EventClaimSettled) String() string { return proto.CompactTextString(m) }
func (*EventClaimSettled) ProtoMessage()    {}
func (*EventClaimSettled) Descriptor() ([]byte, []int) {
	return fileDescriptor_a78874bbf91a58c7, []int{1}
}
func (m *EventClaimSettled) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventClaimSettled) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventClaimSettled.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventClaimSettled) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventClaimSettled.Merge(m, src)
}
func (m *EventClaimSettled) XXX_Size() int {
	return m.Size()
}
func (m *EventClaimSettled) XXX_DiscardUnknown() {
	xxx_messageInfo_EventClaimSettled.DiscardUnknown(m)
}

var xxx_messageInfo_EventClaimSettled proto.InternalMessageInfo

func (m *EventClaimSettled) GetClaim() *types.Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

func (m *EventClaimSettled) GetNumRelays() uint64 {
	if m != nil {
		return m.NumRelays
	}
	return 0
}

func (m *EventClaimSettled) GetNumComputeUnits() uint64 {
	if m != nil {
		return m.NumComputeUnits
	}
	return 0
}

func (m *EventClaimSettled) GetProofRequirement() types.ProofRequirementReason {
	if m != nil {
		return m.ProofRequirement
	}
	return types.ProofRequirementReason_NOT_REQUIRED
}

// EventRelayMiningDifficultyUpdated is an event emitted whenever the relay mining difficulty is updated
// for a given service.
type EventRelayMiningDifficultyUpdated struct {
	ServiceId                string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	PrevTargetHashHexEncoded string `protobuf:"bytes,2,opt,name=prev_target_hash_hex_encoded,json=prevTargetHashHexEncoded,proto3" json:"prev_target_hash_hex_encoded,omitempty"`
	NewTargetHashHexEncoded  string `protobuf:"bytes,3,opt,name=new_target_hash_hex_encoded,json=newTargetHashHexEncoded,proto3" json:"new_target_hash_hex_encoded,omitempty"`
	PrevNumRelaysEma         uint64 `protobuf:"varint,4,opt,name=prev_num_relays_ema,json=prevNumRelaysEma,proto3" json:"prev_num_relays_ema,omitempty"`
	NewNumRelaysEma          uint64 `protobuf:"varint,5,opt,name=new_num_relays_ema,json=newNumRelaysEma,proto3" json:"new_num_relays_ema,omitempty"`
}

func (m *EventRelayMiningDifficultyUpdated) Reset()         { *m = EventRelayMiningDifficultyUpdated{} }
func (m *EventRelayMiningDifficultyUpdated) String() string { return proto.CompactTextString(m) }
func (*EventRelayMiningDifficultyUpdated) ProtoMessage()    {}
func (*EventRelayMiningDifficultyUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_a78874bbf91a58c7, []int{2}
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRelayMiningDifficultyUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRelayMiningDifficultyUpdated.Merge(m, src)
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Size() int {
	return m.Size()
}
func (m *EventRelayMiningDifficultyUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRelayMiningDifficultyUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_EventRelayMiningDifficultyUpdated proto.InternalMessageInfo

func (m *EventRelayMiningDifficultyUpdated) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetPrevTargetHashHexEncoded() string {
	if m != nil {
		return m.PrevTargetHashHexEncoded
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetNewTargetHashHexEncoded() string {
	if m != nil {
		return m.NewTargetHashHexEncoded
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetPrevNumRelaysEma() uint64 {
	if m != nil {
		return m.PrevNumRelaysEma
	}
	return 0
}

func (m *EventRelayMiningDifficultyUpdated) GetNewNumRelaysEma() uint64 {
	if m != nil {
		return m.NewNumRelaysEma
	}
	return 0
}

// EventApplicationOverserviced is emitted when an application has less stake
// than the expected burn.
type EventApplicationOverserviced struct {
	ApplicationAddr string       `protobuf:"bytes,1,opt,name=application_addr,json=applicationAddr,proto3" json:"application_addr,omitempty"`
	ExpectedBurn    *types1.Coin `protobuf:"bytes,2,opt,name=expected_burn,json=expectedBurn,proto3" json:"expected_burn,omitempty"`
	EffectiveBurn   *types1.Coin `protobuf:"bytes,3,opt,name=effective_burn,json=effectiveBurn,proto3" json:"effective_burn,omitempty"`
}

func (m *EventApplicationOverserviced) Reset()         { *m = EventApplicationOverserviced{} }
func (m *EventApplicationOverserviced) String() string { return proto.CompactTextString(m) }
func (*EventApplicationOverserviced) ProtoMessage()    {}
func (*EventApplicationOverserviced) Descriptor() ([]byte, []int) {
	return fileDescriptor_a78874bbf91a58c7, []int{3}
}
func (m *EventApplicationOverserviced) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventApplicationOverserviced) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventApplicationOverserviced.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventApplicationOverserviced) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventApplicationOverserviced.Merge(m, src)
}
func (m *EventApplicationOverserviced) XXX_Size() int {
	return m.Size()
}
func (m *EventApplicationOverserviced) XXX_DiscardUnknown() {
	xxx_messageInfo_EventApplicationOverserviced.DiscardUnknown(m)
}

var xxx_messageInfo_EventApplicationOverserviced proto.InternalMessageInfo

func (m *EventApplicationOverserviced) GetApplicationAddr() string {
	if m != nil {
		return m.ApplicationAddr
	}
	return ""
}

func (m *EventApplicationOverserviced) GetExpectedBurn() *types1.Coin {
	if m != nil {
		return m.ExpectedBurn
	}
	return nil
}

func (m *EventApplicationOverserviced) GetEffectiveBurn() *types1.Coin {
	if m != nil {
		return m.EffectiveBurn
	}
	return nil
}

func init() {
	proto.RegisterType((*EventClaimExpired)(nil), "poktroll.tokenomics.EventClaimExpired")
	proto.RegisterType((*EventClaimSettled)(nil), "poktroll.tokenomics.EventClaimSettled")
	proto.RegisterType((*EventRelayMiningDifficultyUpdated)(nil), "poktroll.tokenomics.EventRelayMiningDifficultyUpdated")
	proto.RegisterType((*EventApplicationOverserviced)(nil), "poktroll.tokenomics.EventApplicationOverserviced")
}

func init() { proto.RegisterFile("poktroll/tokenomics/event.proto", fileDescriptor_a78874bbf91a58c7) }

var fileDescriptor_a78874bbf91a58c7 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0xcb, 0x6e, 0x13, 0x3d,
	0x14, 0xc7, 0x3b, 0xbd, 0x7c, 0x52, 0xdc, 0xaf, 0x37, 0x97, 0x8a, 0x50, 0xca, 0x24, 0x64, 0x81,
	0x8a, 0x50, 0x67, 0xd4, 0x56, 0x62, 0x85, 0x2a, 0x9a, 0x12, 0xa9, 0x2c, 0x0a, 0x68, 0xa0, 0x1b,
	0x36, 0x23, 0x67, 0x7c, 0x92, 0x98, 0xce, 0xd8, 0xc6, 0xe3, 0x99, 0xa4, 0x6f, 0xc1, 0x03, 0xf0,
	0x02, 0x3c, 0x04, 0x7b, 0x24, 0x36, 0x5d, 0xb2, 0xaa, 0x50, 0xbb, 0xeb, 0x53, 0x20, 0x7b, 0x72,
	0x19, 0xa5, 0x20, 0xd6, 0x6c, 0x12, 0xeb, 0xfc, 0x7f, 0xe7, 0x9c, 0xbf, 0x8f, 0xc7, 0x46, 0x35,
	0x29, 0xce, 0xb4, 0x12, 0x71, 0xec, 0x6b, 0x71, 0x06, 0x5c, 0x24, 0x2c, 0x4a, 0x7d, 0xc8, 0x81,
	0x6b, 0x4f, 0x2a, 0xa1, 0x05, 0x5e, 0x1f, 0x01, 0xde, 0x04, 0xd8, 0xbc, 0xd3, 0x15, 0x5d, 0x61,
	0x75, 0xdf, 0xac, 0x0a, 0x74, 0xd3, 0x8d, 0x44, 0x9a, 0x88, 0xd4, 0x6f, 0x93, 0x14, 0xfc, 0x7c,
	0xb7, 0x0d, 0x9a, 0xec, 0xfa, 0x91, 0x60, 0x7c, 0xa8, 0x6f, 0x8e, 0x7b, 0x49, 0x25, 0x44, 0xc7,
	0x8f, 0x62, 0xc2, 0x92, 0xa1, 0x56, 0x9f, 0xd2, 0x14, 0x7c, 0xcc, 0x98, 0x82, 0x64, 0x6c, 0xa4,
	0xf1, 0xd5, 0x41, 0x6b, 0x2d, 0x63, 0xec, 0xc8, 0xa4, 0xb5, 0x06, 0x92, 0x29, 0xa0, 0xf8, 0x29,
	0x5a, 0xb0, 0x65, 0xaa, 0x4e, 0xdd, 0xd9, 0x5e, 0xdc, 0xdb, 0xf0, 0xc6, 0x76, 0x6d, 0x1d, 0xcf,
	0xc2, 0xcd, 0xca, 0xcd, 0x65, 0xad, 0xe0, 0x82, 0xe2, 0x0f, 0xef, 0x20, 0xc4, 0xb3, 0x24, 0x54,
	0x10, 0x93, 0xf3, 0xb4, 0x3a, 0x5b, 0x77, 0xb6, 0xe7, 0x9b, 0xcb, 0x37, 0x97, 0xb5, 0x52, 0x34,
	0xa8, 0xf0, 0x2c, 0x09, 0xec, 0x12, 0x1f, 0xa2, 0x35, 0x23, 0x44, 0x22, 0x91, 0x99, 0x86, 0x30,
	0xe3, 0x4c, 0xa7, 0xd5, 0x39, 0x9b, 0xb5, 0x71, 0x73, 0x59, 0xbb, 0x2d, 0x06, 0x2b, 0x3c, 0x4b,
	0x8e, 0x8a, 0xc8, 0xa9, 0x09, 0x34, 0xbe, 0xcc, 0x96, 0xfd, 0xbf, 0x05, 0xad, 0xe3, 0x7f, 0xc9,
	0x3f, 0xfe, 0x80, 0xd6, 0xac, 0xa5, 0xb0, 0x74, 0x34, 0xd5, 0xf9, 0xba, 0xb3, 0xbd, 0xbc, 0xf7,
	0x68, 0xda, 0xf5, 0x1b, 0xf3, 0x1b, 0x4c, 0xb8, 0x00, 0x48, 0x2a, 0x78, 0xd1, 0xea, 0x56, 0x91,
	0x60, 0x55, 0x4e, 0xe1, 0x8d, 0xcf, 0xb3, 0xe8, 0xa1, 0x9d, 0x95, 0xb5, 0x7f, 0xc2, 0x38, 0xe3,
	0xdd, 0x17, 0xac, 0xd3, 0x61, 0x51, 0x16, 0xeb, 0xf3, 0x53, 0x49, 0x89, 0x06, 0x8a, 0x1f, 0x20,
	0x94, 0x82, 0xca, 0x59, 0x04, 0x21, 0xa3, 0x76, 0x80, 0x95, 0xa0, 0x32, 0x8c, 0xbc, 0xa4, 0xf8,
	0x00, 0x6d, 0x49, 0x05, 0x79, 0xa8, 0x89, 0xea, 0x82, 0x0e, 0x7b, 0x24, 0xed, 0x85, 0x3d, 0x18,
	0x84, 0xc0, 0x23, 0x41, 0x81, 0xda, 0xa1, 0x55, 0x82, 0xaa, 0x61, 0xde, 0x59, 0xe4, 0x98, 0xa4,
	0xbd, 0x63, 0x18, 0xb4, 0x0a, 0x1d, 0x3f, 0x43, 0xf7, 0x39, 0xf4, 0xff, 0x98, 0x3e, 0x67, 0xd3,
	0xef, 0x72, 0xe8, 0xff, 0x36, 0x7b, 0x07, 0xad, 0xdb, 0xee, 0x93, 0xf3, 0x08, 0x21, 0x21, 0x76,
	0x60, 0xf3, 0x66, 0xc7, 0x90, 0xbf, 0x1a, 0x9d, 0x4e, 0x2b, 0x21, 0xf8, 0x09, 0xc2, 0xa6, 0xd9,
	0x14, 0xbd, 0x60, 0xe9, 0x15, 0x0e, 0xfd, 0x32, 0xdc, 0xf8, 0xee, 0xa0, 0x2d, 0x3b, 0x9e, 0x43,
	0x29, 0x63, 0x16, 0x11, 0xcd, 0x04, 0x7f, 0x9d, 0x83, 0x1a, 0xee, 0x9d, 0xe2, 0xc7, 0x68, 0x95,
	0x4c, 0xa4, 0x90, 0x50, 0xaa, 0x86, 0xf3, 0x59, 0x29, 0xc5, 0x0f, 0x29, 0x55, 0xf8, 0x00, 0x2d,
	0xc1, 0x40, 0x42, 0xa4, 0x81, 0x86, 0xed, 0x4c, 0x71, 0x3b, 0x96, 0xc5, 0xbd, 0x7b, 0x5e, 0x71,
	0x99, 0x3d, 0x73, 0x99, 0xbd, 0xe1, 0x65, 0xf6, 0x8e, 0x04, 0xe3, 0xc1, 0xff, 0x23, 0xbe, 0x99,
	0x29, 0x8e, 0x9f, 0xa3, 0x65, 0xe8, 0x74, 0x20, 0xd2, 0x2c, 0x87, 0xa2, 0xc0, 0xdc, 0xdf, 0x0a,
	0x2c, 0x8d, 0x13, 0x4c, 0x85, 0xe6, 0xc9, 0xb7, 0x2b, 0xd7, 0xb9, 0xb8, 0x72, 0x9d, 0x9f, 0x57,
	0xae, 0xf3, 0xe9, 0xda, 0x9d, 0xb9, 0xb8, 0x76, 0x67, 0x7e, 0x5c, 0xbb, 0x33, 0xef, 0xf7, 0xbb,
	0x4c, 0xf7, 0xb2, 0xb6, 0x17, 0x89, 0xc4, 0x37, 0x5f, 0xd8, 0x0e, 0x07, 0xdd, 0x17, 0xea, 0xcc,
	0x1f, 0x3f, 0x16, 0x83, 0xf2, 0xb3, 0xa5, 0xcf, 0x25, 0xa4, 0xed, 0xff, 0xec, 0x73, 0xb1, 0xff,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x67, 0x30, 0x58, 0x9b, 0xda, 0x04, 0x00, 0x00,
}

func (m *EventClaimExpired) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventClaimExpired) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventClaimExpired) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumComputeUnits != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NumComputeUnits))
		i--
		dAtA[i] = 0x18
	}
	if m.NumRelays != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NumRelays))
		i--
		dAtA[i] = 0x10
	}
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventClaimSettled) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventClaimSettled) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventClaimSettled) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProofRequirement != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.ProofRequirement))
		i--
		dAtA[i] = 0x20
	}
	if m.NumComputeUnits != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NumComputeUnits))
		i--
		dAtA[i] = 0x18
	}
	if m.NumRelays != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NumRelays))
		i--
		dAtA[i] = 0x10
	}
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventRelayMiningDifficultyUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRelayMiningDifficultyUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRelayMiningDifficultyUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewNumRelaysEma != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NewNumRelaysEma))
		i--
		dAtA[i] = 0x28
	}
	if m.PrevNumRelaysEma != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.PrevNumRelaysEma))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NewTargetHashHexEncoded) > 0 {
		i -= len(m.NewTargetHashHexEncoded)
		copy(dAtA[i:], m.NewTargetHashHexEncoded)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.NewTargetHashHexEncoded)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PrevTargetHashHexEncoded) > 0 {
		i -= len(m.PrevTargetHashHexEncoded)
		copy(dAtA[i:], m.PrevTargetHashHexEncoded)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.PrevTargetHashHexEncoded)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServiceId) > 0 {
		i -= len(m.ServiceId)
		copy(dAtA[i:], m.ServiceId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ServiceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventApplicationOverserviced) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventApplicationOverserviced) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventApplicationOverserviced) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EffectiveBurn != nil {
		{
			size, err := m.EffectiveBurn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ExpectedBurn != nil {
		{
			size, err := m.ExpectedBurn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ApplicationAddr) > 0 {
		i -= len(m.ApplicationAddr)
		copy(dAtA[i:], m.ApplicationAddr)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ApplicationAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventClaimExpired) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.NumRelays != 0 {
		n += 1 + sovEvent(uint64(m.NumRelays))
	}
	if m.NumComputeUnits != 0 {
		n += 1 + sovEvent(uint64(m.NumComputeUnits))
	}
	return n
}

func (m *EventClaimSettled) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.NumRelays != 0 {
		n += 1 + sovEvent(uint64(m.NumRelays))
	}
	if m.NumComputeUnits != 0 {
		n += 1 + sovEvent(uint64(m.NumComputeUnits))
	}
	if m.ProofRequirement != 0 {
		n += 1 + sovEvent(uint64(m.ProofRequirement))
	}
	return n
}

func (m *EventRelayMiningDifficultyUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.PrevTargetHashHexEncoded)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.NewTargetHashHexEncoded)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.PrevNumRelaysEma != 0 {
		n += 1 + sovEvent(uint64(m.PrevNumRelaysEma))
	}
	if m.NewNumRelaysEma != 0 {
		n += 1 + sovEvent(uint64(m.NewNumRelaysEma))
	}
	return n
}

func (m *EventApplicationOverserviced) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ApplicationAddr)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.ExpectedBurn != nil {
		l = m.ExpectedBurn.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.EffectiveBurn != nil {
		l = m.EffectiveBurn.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventClaimExpired) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventClaimExpired: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventClaimExpired: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &types.Claim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRelays", wireType)
			}
			m.NumRelays = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRelays |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumComputeUnits", wireType)
			}
			m.NumComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventClaimSettled) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventClaimSettled: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventClaimSettled: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &types.Claim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRelays", wireType)
			}
			m.NumRelays = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRelays |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumComputeUnits", wireType)
			}
			m.NumComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofRequirement", wireType)
			}
			m.ProofRequirement = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProofRequirement |= types.ProofRequirementReason(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventRelayMiningDifficultyUpdated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventRelayMiningDifficultyUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRelayMiningDifficultyUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevTargetHashHexEncoded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevTargetHashHexEncoded = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewTargetHashHexEncoded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewTargetHashHexEncoded = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevNumRelaysEma", wireType)
			}
			m.PrevNumRelaysEma = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrevNumRelaysEma |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewNumRelaysEma", wireType)
			}
			m.NewNumRelaysEma = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewNumRelaysEma |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventApplicationOverserviced) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventApplicationOverserviced: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventApplicationOverserviced: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedBurn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpectedBurn == nil {
				m.ExpectedBurn = &types1.Coin{}
			}
			if err := m.ExpectedBurn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EffectiveBurn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EffectiveBurn == nil {
				m.EffectiveBurn = &types1.Coin{}
			}
			if err := m.EffectiveBurn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
