// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poa/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgCreateValidator struct {
	// string name = 1 [(gogoproto.moretags) = "yaml:\"name\""];
	Description      types.Description `protobuf:"bytes,1,opt,name=description,proto3" json:"description"`
	OwnerAddress     string            `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty" yaml:"owner_address"`
	ValidatorAddress string            `protobuf:"bytes,3,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
	Pubkey           *types1.Any       `protobuf:"bytes,4,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
}

func (m *MsgCreateValidator) Reset()         { *m = MsgCreateValidator{} }
func (m *MsgCreateValidator) String() string { return proto.CompactTextString(m) }
func (*MsgCreateValidator) ProtoMessage()    {}
func (*MsgCreateValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac16a0aed23c2e2, []int{0}
}
func (m *MsgCreateValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateValidator.Merge(m, src)
}
func (m *MsgCreateValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateValidator proto.InternalMessageInfo

type MsgCreateValidatorResponse struct {
}

func (m *MsgCreateValidatorResponse) Reset()         { *m = MsgCreateValidatorResponse{} }
func (m *MsgCreateValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateValidatorResponse) ProtoMessage()    {}
func (*MsgCreateValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac16a0aed23c2e2, []int{1}
}
func (m *MsgCreateValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateValidatorResponse.Merge(m, src)
}
func (m *MsgCreateValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateValidatorResponse proto.InternalMessageInfo

type MsgVoteValidator struct {
	// string name = 1 [(gogoproto.moretags) = "yaml:\"name\""];
	VoterAddress string `protobuf:"bytes,1,opt,name=voter_address,json=voterAddress,proto3" json:"voter_address,omitempty" yaml:"voter_address"`
	InFavor      bool   `protobuf:"varint,2,opt,name=in_favor,json=inFavor,proto3" json:"in_favor,omitempty" yaml:"infavor"`
	OwnerAddress string `protobuf:"bytes,3,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty" yaml:"owner_address"`
}

func (m *MsgVoteValidator) Reset()         { *m = MsgVoteValidator{} }
func (m *MsgVoteValidator) String() string { return proto.CompactTextString(m) }
func (*MsgVoteValidator) ProtoMessage()    {}
func (*MsgVoteValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac16a0aed23c2e2, []int{2}
}
func (m *MsgVoteValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVoteValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVoteValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVoteValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVoteValidator.Merge(m, src)
}
func (m *MsgVoteValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgVoteValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVoteValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVoteValidator proto.InternalMessageInfo

func (m *MsgVoteValidator) GetVoterAddress() string {
	if m != nil {
		return m.VoterAddress
	}
	return ""
}

func (m *MsgVoteValidator) GetInFavor() bool {
	if m != nil {
		return m.InFavor
	}
	return false
}

func (m *MsgVoteValidator) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

type MsgVoteValidatorResponse struct {
}

func (m *MsgVoteValidatorResponse) Reset()         { *m = MsgVoteValidatorResponse{} }
func (m *MsgVoteValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgVoteValidatorResponse) ProtoMessage()    {}
func (*MsgVoteValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac16a0aed23c2e2, []int{3}
}
func (m *MsgVoteValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVoteValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVoteValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVoteValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVoteValidatorResponse.Merge(m, src)
}
func (m *MsgVoteValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgVoteValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVoteValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVoteValidatorResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateValidator)(nil), "noble.poa.MsgCreateValidator")
	proto.RegisterType((*MsgCreateValidatorResponse)(nil), "noble.poa.MsgCreateValidatorResponse")
	proto.RegisterType((*MsgVoteValidator)(nil), "noble.poa.MsgVoteValidator")
	proto.RegisterType((*MsgVoteValidatorResponse)(nil), "noble.poa.MsgVoteValidatorResponse")
}

func init() { proto.RegisterFile("poa/tx.proto", fileDescriptor_8ac16a0aed23c2e2) }

var fileDescriptor_8ac16a0aed23c2e2 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xde, 0x6d, 0x4a, 0x4d, 0xa7, 0x89, 0xd6, 0x25, 0x87, 0x75, 0xad, 0xbb, 0x65, 0xab, 0xd0,
	0x4b, 0x66, 0x49, 0xbd, 0x15, 0x7a, 0x68, 0x94, 0x82, 0x84, 0x80, 0xee, 0xa1, 0x82, 0x97, 0x30,
	0x9b, 0x4c, 0xd7, 0xa5, 0xc9, 0xbc, 0x65, 0x66, 0xb2, 0x76, 0xff, 0x03, 0x8f, 0xfe, 0x09, 0xfd,
	0x17, 0x04, 0xaf, 0xe2, 0xb5, 0x78, 0xea, 0xd1, 0x53, 0x90, 0xe4, 0xe2, 0x39, 0x7f, 0x81, 0x74,
	0x36, 0x13, 0xf3, 0x03, 0x0a, 0xde, 0xe6, 0x9b, 0xef, 0x7d, 0xdf, 0xbc, 0xf7, 0xf1, 0x06, 0x55,
	0x52, 0x20, 0x81, 0xbc, 0xc2, 0x29, 0x07, 0x09, 0xd6, 0x36, 0x83, 0xa8, 0x4f, 0x71, 0x0a, 0xc4,
	0x79, 0xde, 0x05, 0x31, 0x00, 0x11, 0x08, 0x49, 0x2e, 0x13, 0x16, 0x07, 0x59, 0x23, 0xa2, 0x92,
	0x34, 0x34, 0x2e, 0x04, 0xce, 0x93, 0xa2, 0xaa, 0xa3, 0x50, 0x50, 0x00, 0x4d, 0xc5, 0x00, 0x71,
	0x9f, 0x06, 0x0a, 0x45, 0xc3, 0x8b, 0x80, 0xb0, 0x7c, 0x46, 0xd5, 0x62, 0x88, 0xa1, 0x90, 0xdc,
	0x9d, 0x8a, 0x5b, 0xff, 0xfb, 0x06, 0xb2, 0xda, 0x22, 0x7e, 0xc5, 0x29, 0x91, 0xf4, 0x9c, 0xf4,
	0x93, 0x1e, 0x91, 0xc0, 0xad, 0x16, 0xda, 0xe9, 0x51, 0xd1, 0xe5, 0x49, 0x2a, 0x13, 0x60, 0xb6,
	0xb9, 0x6f, 0x1e, 0xee, 0x1c, 0x1d, 0xe0, 0xd9, 0x5b, 0xba, 0x9d, 0x59, 0x7b, 0xf8, 0xf5, 0xbf,
	0xd2, 0xe6, 0xe6, 0xcd, 0xc8, 0x33, 0xc2, 0x45, 0xb5, 0x75, 0x82, 0xaa, 0xf0, 0x89, 0x51, 0xde,
	0x21, 0xbd, 0x1e, 0xa7, 0x42, 0xd8, 0x1b, 0xfb, 0xe6, 0xe1, 0x76, 0xd3, 0x9e, 0x8e, 0xbc, 0x5a,
	0x4e, 0x06, 0xfd, 0x63, 0x7f, 0x89, 0xf6, 0xc3, 0x8a, 0xc2, 0xa7, 0x05, 0xb4, 0xde, 0xa0, 0xc7,
	0x99, 0x6e, 0x6c, 0x6e, 0x51, 0x52, 0x16, 0x7b, 0xd3, 0x91, 0x67, 0x17, 0x16, 0x6b, 0x25, 0x7e,
	0xb8, 0x3b, 0xbf, 0xd3, 0x56, 0x67, 0x68, 0x2b, 0x1d, 0x46, 0x97, 0x34, 0xb7, 0x37, 0xd5, 0x44,
	0x35, 0x5c, 0xe4, 0x85, 0x75, 0x5e, 0xf8, 0x94, 0xe5, 0x4d, 0xfb, 0xe7, 0xb7, 0x7a, 0x6d, 0x36,
	0x6a, 0x97, 0xe7, 0xa9, 0x04, 0xfc, 0x76, 0x18, 0xb5, 0x68, 0x1e, 0xce, 0xd4, 0xc7, 0xe5, 0xcf,
	0xd7, 0x9e, 0xf1, 0xe7, 0xda, 0x33, 0xfc, 0x3d, 0xe4, 0xac, 0xc7, 0x17, 0x52, 0x91, 0x02, 0x13,
	0xd4, 0xff, 0x61, 0xa2, 0xdd, 0xb6, 0x88, 0xcf, 0x61, 0x31, 0xdb, 0x13, 0x54, 0xcd, 0x40, 0x2e,
	0xc4, 0x61, 0xae, 0xc6, 0xb1, 0x44, 0xfb, 0x61, 0x45, 0x61, 0x3d, 0x43, 0x1d, 0x95, 0x13, 0xd6,
	0xb9, 0x20, 0x19, 0x70, 0x15, 0x64, 0xb9, 0x69, 0x4d, 0x47, 0xde, 0xc3, 0x42, 0x99, 0x30, 0x45,
	0xf8, 0xe1, 0x83, 0x84, 0x9d, 0xdd, 0x9d, 0xd6, 0xc3, 0x2f, 0xfd, 0x4f, 0xf8, 0xbe, 0x83, 0xec,
	0xd5, 0x01, 0xf4, 0x74, 0x47, 0x5f, 0x4d, 0x54, 0x6a, 0x8b, 0xd8, 0x7a, 0x8f, 0x1e, 0xad, 0xee,
	0xcf, 0x33, 0x3c, 0x5f, 0x6a, 0xbc, 0x9e, 0x8f, 0xf3, 0xe2, 0x5e, 0x5a, 0x3f, 0x60, 0xbd, 0x43,
	0xd5, 0xe5, 0xe8, 0x9e, 0x2e, 0xeb, 0x96, 0x48, 0xe7, 0xe0, 0x1e, 0x52, 0x5b, 0x36, 0x5b, 0x37,
	0x63, 0xd7, 0xbc, 0x1d, 0xbb, 0xe6, 0xef, 0xb1, 0x6b, 0x7e, 0x99, 0xb8, 0xc6, 0xed, 0xc4, 0x35,
	0x7e, 0x4d, 0x5c, 0xe3, 0x43, 0x23, 0x4e, 0xe4, 0xc7, 0x61, 0x84, 0xbb, 0x30, 0x08, 0x84, 0xe4,
	0x84, 0xc5, 0xb4, 0x0f, 0x19, 0xad, 0x67, 0x94, 0xc9, 0x21, 0xa7, 0x22, 0x50, 0xee, 0xc1, 0x55,
	0xa0, 0x3e, 0x6f, 0x9e, 0x52, 0x11, 0x6d, 0xa9, 0xb5, 0x79, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xfc, 0x19, 0xae, 0x10, 0xd0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateValidator(ctx context.Context, in *MsgCreateValidator, opts ...grpc.CallOption) (*MsgCreateValidatorResponse, error)
	VoteValidator(ctx context.Context, in *MsgVoteValidator, opts ...grpc.CallOption) (*MsgVoteValidatorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateValidator(ctx context.Context, in *MsgCreateValidator, opts ...grpc.CallOption) (*MsgCreateValidatorResponse, error) {
	out := new(MsgCreateValidatorResponse)
	err := c.cc.Invoke(ctx, "/noble.poa.Msg/CreateValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) VoteValidator(ctx context.Context, in *MsgVoteValidator, opts ...grpc.CallOption) (*MsgVoteValidatorResponse, error) {
	out := new(MsgVoteValidatorResponse)
	err := c.cc.Invoke(ctx, "/noble.poa.Msg/VoteValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateValidator(context.Context, *MsgCreateValidator) (*MsgCreateValidatorResponse, error)
	VoteValidator(context.Context, *MsgVoteValidator) (*MsgVoteValidatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateValidator(ctx context.Context, req *MsgCreateValidator) (*MsgCreateValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateValidator not implemented")
}
func (*UnimplementedMsgServer) VoteValidator(ctx context.Context, req *MsgVoteValidator) (*MsgVoteValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteValidator not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noble.poa.Msg/CreateValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateValidator(ctx, req.(*MsgCreateValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_VoteValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVoteValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VoteValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noble.poa.Msg/VoteValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VoteValidator(ctx, req.(*MsgVoteValidator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "noble.poa.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateValidator",
			Handler:    _Msg_CreateValidator_Handler,
		},
		{
			MethodName: "VoteValidator",
			Handler:    _Msg_VoteValidator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "poa/tx.proto",
}

func (m *MsgCreateValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pubkey != nil {
		{
			size, err := m.Pubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgCreateValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgVoteValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVoteValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVoteValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.InFavor {
		i--
		if m.InFavor {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.VoterAddress) > 0 {
		i -= len(m.VoterAddress)
		copy(dAtA[i:], m.VoterAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VoterAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgVoteValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVoteValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVoteValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Description.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Pubkey != nil {
		l = m.Pubkey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgVoteValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VoterAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.InFavor {
		n += 2
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgVoteValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pubkey == nil {
				m.Pubkey = &types1.Any{}
			}
			if err := m.Pubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgVoteValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgVoteValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVoteValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InFavor", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InFavor = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgVoteValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgVoteValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVoteValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
