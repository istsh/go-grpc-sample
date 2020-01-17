// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/login/login.proto

package login

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30160f50bf2ceaad, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30160f50bf2ceaad, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "v1.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "v1.LoginResponse")
}

func init() { proto.RegisterFile("proto/v1/login/login.proto", fileDescriptor_30160f50bf2ceaad) }

var fileDescriptor_30160f50bf2ceaad = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x69, 0xc7, 0x8e, 0x33, 0x71, 0x94, 0x1a, 0x06, 0x2c, 0xc5, 0x85, 0x14, 0x44, 0x11,
	0xa6, 0xa1, 0xba, 0x73, 0x25, 0x5d, 0xbb, 0xaa, 0x3b, 0x5d, 0xa5, 0x9d, 0x90, 0x09, 0xb6, 0x79,
	0x31, 0xc9, 0xd4, 0xbd, 0x57, 0xf0, 0x36, 0xae, 0xbc, 0x83, 0x57, 0x70, 0xe3, 0x25, 0x44, 0x26,
	0x19, 0xcb, 0x6c, 0x42, 0xde, 0xfb, 0xff, 0xf7, 0xf1, 0xf3, 0xa3, 0x54, 0x69, 0xb0, 0x40, 0xfa,
	0x82, 0xb4, 0xc0, 0x85, 0xf4, 0x6f, 0xee, 0x96, 0x38, 0xec, 0x8b, 0xf4, 0x94, 0x03, 0xf0, 0x96,
	0x11, 0xaa, 0x04, 0xa1, 0x52, 0x82, 0xa5, 0x56, 0x80, 0x34, 0xde, 0x91, 0x9e, 0xf4, 0xb4, 0x15,
	0x4b, 0x6a, 0x19, 0xf9, 0xff, 0x78, 0x21, 0x7b, 0x42, 0xb3, 0xfb, 0x0d, 0xa9, 0x62, 0x2f, 0x6b,
	0x66, 0x2c, 0xce, 0x50, 0xc4, 0x3a, 0x2a, 0xda, 0x24, 0x38, 0x0b, 0x2e, 0xa7, 0xe5, 0xec, 0xe3,
	0xe7, 0x73, 0xb4, 0xaf, 0xa3, 0x78, 0x94, 0xfc, 0x06, 0x95, 0x97, 0xf0, 0x05, 0x9a, 0x28, 0x6a,
	0xcc, 0x2b, 0xe8, 0x65, 0x12, 0x3a, 0xdb, 0xc1, 0xc6, 0x36, 0xd6, 0x7b, 0xf1, 0x24, 0xb9, 0xab,
	0x06, 0x31, 0x3b, 0x47, 0x87, 0x5b, 0xb8, 0x51, 0x20, 0x0d, 0xc3, 0x73, 0x14, 0x59, 0x78, 0x66,
	0xd2, 0xd3, 0x2b, 0x3f, 0x5c, 0x57, 0xdb, 0x0c, 0x0f, 0x4c, 0xf7, 0xa2, 0x61, 0xb8, 0x44, 0x91,
	0x9b, 0x71, 0x9c, 0xf7, 0x45, 0xbe, 0x1b, 0x2f, 0x3d, 0xde, 0xd9, 0x78, 0x66, 0x36, 0x7f, 0xfb,
	0xfa, 0x7e, 0x0f, 0x8f, 0xb2, 0xe9, 0xd0, 0xcd, 0x6d, 0x70, 0x55, 0x16, 0x8f, 0x84, 0x0b, 0xbb,
	0x5a, 0xd7, 0x79, 0x03, 0x1d, 0x11, 0xc6, 0x9a, 0x15, 0xe1, 0xb0, 0xe0, 0x5a, 0x35, 0x0b, 0x43,
	0x3b, 0xe5, 0x9a, 0x52, 0x44, 0xd5, 0xc3, 0x55, 0x3d, 0x76, 0x8d, 0xdc, 0xfc, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x52, 0x8b, 0x7f, 0x20, 0x6a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type loginServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoginServiceClient(cc *grpc.ClientConn) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/v1.LoginService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
type LoginServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedLoginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (*UnimplementedLoginServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterLoginServiceServer(s *grpc.Server, srv LoginServiceServer) {
	s.RegisterService(&_LoginService_serviceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.LoginService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/login/login.proto",
}
