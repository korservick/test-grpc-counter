// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/counter.proto

package counter

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CurrentNumber struct {
	Number               uint32   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentNumber) Reset()         { *m = CurrentNumber{} }
func (m *CurrentNumber) String() string { return proto.CompactTextString(m) }
func (*CurrentNumber) ProtoMessage()    {}
func (*CurrentNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a6284c8cb82ab12, []int{0}
}

func (m *CurrentNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentNumber.Unmarshal(m, b)
}
func (m *CurrentNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentNumber.Marshal(b, m, deterministic)
}
func (m *CurrentNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentNumber.Merge(m, src)
}
func (m *CurrentNumber) XXX_Size() int {
	return xxx_messageInfo_CurrentNumber.Size(m)
}
func (m *CurrentNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentNumber.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentNumber proto.InternalMessageInfo

func (m *CurrentNumber) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type Settings struct {
	Increment            uint32   `protobuf:"varint,1,opt,name=increment,proto3" json:"increment,omitempty"`
	MaximumNumber        uint32   `protobuf:"varint,2,opt,name=maximum_number,json=maximumNumber,proto3" json:"maximum_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a6284c8cb82ab12, []int{1}
}

func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetIncrement() uint32 {
	if m != nil {
		return m.Increment
	}
	return 0
}

func (m *Settings) GetMaximumNumber() uint32 {
	if m != nil {
		return m.MaximumNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*CurrentNumber)(nil), "counter.CurrentNumber")
	proto.RegisterType((*Settings)(nil), "counter.Settings")
}

func init() { proto.RegisterFile("pb/counter.proto", fileDescriptor_1a6284c8cb82ab12) }

var fileDescriptor_1a6284c8cb82ab12 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0xd2, 0x4f,
	0xce, 0x2f, 0xcd, 0x2b, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72,
	0xa5, 0xa4, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xc2, 0x49, 0xa5, 0x69, 0xfa, 0xa9,
	0xb9, 0x05, 0x25, 0x95, 0x10, 0x55, 0x4a, 0xea, 0x5c, 0xbc, 0xce, 0xa5, 0x45, 0x45, 0xa9, 0x79,
	0x25, 0x7e, 0xa5, 0xb9, 0x49, 0xa9, 0x45, 0x42, 0x62, 0x5c, 0x6c, 0x79, 0x60, 0x96, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x6f, 0x10, 0x94, 0xa7, 0xe4, 0xcf, 0xc5, 0x11, 0x9c, 0x5a, 0x52, 0x92, 0x99,
	0x97, 0x5e, 0x2c, 0x24, 0xc3, 0xc5, 0x99, 0x99, 0x97, 0x5c, 0x94, 0x9a, 0x9b, 0x9a, 0x57, 0x02,
	0x55, 0x86, 0x10, 0x10, 0x52, 0xe5, 0xe2, 0xcb, 0x4d, 0xac, 0xc8, 0xcc, 0x2d, 0xcd, 0x8d, 0x87,
	0x9a, 0xc4, 0x04, 0x56, 0xc2, 0x0b, 0x15, 0x85, 0x58, 0x64, 0x74, 0x9c, 0x91, 0x8b, 0xdb, 0x13,
	0xa6, 0x29, 0xbf, 0x48, 0xc8, 0x9a, 0x8b, 0xd3, 0x3d, 0x15, 0xee, 0x0a, 0x3d, 0x88, 0xa3, 0xf5,
	0x60, 0x8e, 0xd6, 0x73, 0x05, 0x39, 0x5a, 0x4a, 0x4c, 0x0f, 0xe6, 0x49, 0x54, 0x57, 0x3b, 0x72,
	0xf1, 0xc3, 0xcd, 0x22, 0x68, 0x04, 0x56, 0x71, 0x21, 0x0b, 0x2e, 0xee, 0xe0, 0xd4, 0x12, 0xb8,
	0x1f, 0x05, 0xe1, 0x36, 0xc1, 0x84, 0x70, 0xe9, 0x4c, 0x62, 0x03, 0xf3, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x0b, 0x8a, 0x38, 0x14, 0x84, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IncrementorClient is the client API for Incrementor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IncrementorClient interface {
	GetNumber(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CurrentNumber, error)
	IncrementNumber(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	SetSettings(ctx context.Context, in *Settings, opts ...grpc.CallOption) (*empty.Empty, error)
}

type incrementorClient struct {
	cc *grpc.ClientConn
}

func NewIncrementorClient(cc *grpc.ClientConn) IncrementorClient {
	return &incrementorClient{cc}
}

func (c *incrementorClient) GetNumber(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CurrentNumber, error) {
	out := new(CurrentNumber)
	err := c.cc.Invoke(ctx, "/counter.Incrementor/GetNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementorClient) IncrementNumber(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/counter.Incrementor/IncrementNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementorClient) SetSettings(ctx context.Context, in *Settings, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/counter.Incrementor/SetSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IncrementorServer is the server API for Incrementor service.
type IncrementorServer interface {
	GetNumber(context.Context, *empty.Empty) (*CurrentNumber, error)
	IncrementNumber(context.Context, *empty.Empty) (*empty.Empty, error)
	SetSettings(context.Context, *Settings) (*empty.Empty, error)
}

// UnimplementedIncrementorServer can be embedded to have forward compatible implementations.
type UnimplementedIncrementorServer struct {
}

func (*UnimplementedIncrementorServer) GetNumber(ctx context.Context, req *empty.Empty) (*CurrentNumber, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumber not implemented")
}
func (*UnimplementedIncrementorServer) IncrementNumber(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementNumber not implemented")
}
func (*UnimplementedIncrementorServer) SetSettings(ctx context.Context, req *Settings) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSettings not implemented")
}

func RegisterIncrementorServer(s *grpc.Server, srv IncrementorServer) {
	s.RegisterService(&_Incrementor_serviceDesc, srv)
}

func _Incrementor_GetNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementorServer).GetNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.Incrementor/GetNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementorServer).GetNumber(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incrementor_IncrementNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementorServer).IncrementNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.Incrementor/IncrementNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementorServer).IncrementNumber(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incrementor_SetSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Settings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementorServer).SetSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.Incrementor/SetSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementorServer).SetSettings(ctx, req.(*Settings))
	}
	return interceptor(ctx, in, info, handler)
}

var _Incrementor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "counter.Incrementor",
	HandlerType: (*IncrementorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNumber",
			Handler:    _Incrementor_GetNumber_Handler,
		},
		{
			MethodName: "IncrementNumber",
			Handler:    _Incrementor_IncrementNumber_Handler,
		},
		{
			MethodName: "SetSettings",
			Handler:    _Incrementor_SetSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/counter.proto",
}
