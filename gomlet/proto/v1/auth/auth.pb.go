// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/auth/auth.proto

package auth

import (
	context "context"
	fmt "fmt"
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

func init() { proto.RegisterFile("proto/v1/auth/auth.proto", fileDescriptor_2f453af304448bbc) }

var fileDescriptor_2f453af304448bbc = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x33, 0xd4, 0x4f, 0x2c, 0x2d, 0xc9, 0x00, 0x13, 0x7a, 0x60, 0x21, 0x21, 0x2e,
	0x30, 0xa5, 0x07, 0x12, 0x91, 0x92, 0xc7, 0x54, 0xa5, 0x9b, 0x9c, 0x9f, 0x9b, 0x9b, 0x9f, 0x07,
	0x51, 0x2c, 0x25, 0x93, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x9f, 0x58, 0x90, 0xa9, 0x9f, 0x98,
	0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0x91, 0x35, 0xda, 0xcb, 0xc8, 0xc5,
	0xe2, 0x58, 0x5a, 0x92, 0x21, 0x14, 0xce, 0xc5, 0xea, 0x93, 0x9f, 0x9e, 0x99, 0x27, 0x24, 0xa1,
	0x87, 0x30, 0x5d, 0x0f, 0x2c, 0x14, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x25, 0x89, 0x45,
	0xa6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x49, 0xb2, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xc2, 0x4a,
	0x7c, 0x70, 0x87, 0xe4, 0x80, 0xe4, 0xad, 0x18, 0xb5, 0x84, 0x62, 0xb8, 0xd8, 0xfd, 0x52, 0xcb,
	0x43, 0x8b, 0x53, 0x8b, 0x84, 0xa4, 0x90, 0x0d, 0x80, 0x0a, 0xc2, 0x0c, 0x97, 0xc6, 0x2a, 0x07,
	0x35, 0x5e, 0x02, 0x6c, 0xbc, 0x90, 0x12, 0x2f, 0xdc, 0xf8, 0xd2, 0xe2, 0xd4, 0x22, 0x2b, 0x46,
	0x2d, 0x27, 0xb5, 0x28, 0x95, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd,
	0xf4, 0xfc, 0xdc, 0x1c, 0x30, 0x91, 0x5a, 0xa2, 0x8f, 0x12, 0x2e, 0x49, 0x6c, 0x60, 0xae, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xc5, 0x91, 0x9b, 0x55, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	//
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	//
	NewUser(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*NewUserResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.auth.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) NewUser(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*NewUserResponse, error) {
	out := new(NewUserResponse)
	err := c.cc.Invoke(ctx, "/proto.auth.Auth/NewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	//
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	//
	NewUser(context.Context, *NewUserRequest) (*NewUserResponse, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServer) NewUser(ctx context.Context, req *NewUserRequest) (*NewUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewUser not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.auth.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_NewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).NewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.auth.Auth/NewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).NewUser(ctx, req.(*NewUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "NewUser",
			Handler:    _Auth_NewUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/auth/auth.proto",
}
