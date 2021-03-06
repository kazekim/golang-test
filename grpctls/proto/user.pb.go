// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

package test

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetByIDRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByIDRequest) Reset()         { *m = GetByIDRequest{} }
func (m *GetByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetByIDRequest) ProtoMessage()    {}
func (*GetByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{0}
}

func (m *GetByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIDRequest.Unmarshal(m, b)
}
func (m *GetByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIDRequest.Merge(m, src)
}
func (m *GetByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetByIDRequest.Size(m)
}
func (m *GetByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIDRequest proto.InternalMessageInfo

func (m *GetByIDRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Id                   uint32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*GetByIDRequest)(nil), "test.GetByIDRequest")
	proto.RegisterType((*User)(nil), "test.User")
}

func init() { proto.RegisterFile("proto/user.proto", fileDescriptor_d570e3e37e5899c5) }

var fileDescriptor_d570e3e37e5899c5 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0x33, 0x85, 0x58, 0x4a, 0x52, 0x8b, 0x4b, 0x94,
	0x14, 0xb8, 0xf8, 0xdc, 0x53, 0x4b, 0x9c, 0x2a, 0x3d, 0x5d, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0x98, 0x32,
	0x53, 0x94, 0x1c, 0xb8, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0xc1, 0x32, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x08, 0x17, 0x6b, 0x6a, 0x6e, 0x62, 0x66, 0x8e,
	0x04, 0x13, 0x58, 0x10, 0xc2, 0x81, 0x9a, 0xc0, 0x0c, 0x33, 0xc1, 0xc8, 0x84, 0x8b, 0x35, 0x1d,
	0x6c, 0x84, 0x36, 0x17, 0x3b, 0xd4, 0x32, 0x21, 0x11, 0x3d, 0x90, 0xf5, 0x7a, 0xa8, 0x76, 0x4b,
	0x71, 0x41, 0x44, 0x41, 0x8a, 0x93, 0xd8, 0xc0, 0xce, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0xc2, 0x5e, 0x30, 0xba, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GUserClient is the client API for GUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GUserClient interface {
	GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*User, error)
}

type gUserClient struct {
	cc *grpc.ClientConn
}

func NewGUserClient(cc *grpc.ClientConn) GUserClient {
	return &gUserClient{cc}
}

func (c *gUserClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/test.gUser/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GUserServer is the server API for GUser service.
type GUserServer interface {
	GetByID(context.Context, *GetByIDRequest) (*User, error)
}

// UnimplementedGUserServer can be embedded to have forward compatible implementations.
type UnimplementedGUserServer struct {
}

func (*UnimplementedGUserServer) GetByID(ctx context.Context, req *GetByIDRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}

func RegisterGUserServer(s *grpc.Server, srv GUserServer) {
	s.RegisterService(&_GUser_serviceDesc, srv)
}

func _GUser_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GUserServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.gUser/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GUserServer).GetByID(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.gUser",
	HandlerType: (*GUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _GUser_GetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
