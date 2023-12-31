// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: driver.proto

package driver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DriverAuthService_Create_FullMethodName      = "/driver.DriverAuthService/Create"
	DriverAuthService_Login_FullMethodName       = "/driver.DriverAuthService/Login"
	DriverAuthService_VerifyToken_FullMethodName = "/driver.DriverAuthService/VerifyToken"
)

// DriverAuthServiceClient is the client API for DriverAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverAuthServiceClient interface {
	Create(ctx context.Context, in *CreateDriverRequestProto, opts ...grpc.CallOption) (*CreateDriverResponseProto, error)
	Login(ctx context.Context, in *LoginRequestProto, opts ...grpc.CallOption) (*LoginResponseProto, error)
	VerifyToken(ctx context.Context, in *VerifyTokenRequestProto, opts ...grpc.CallOption) (*VerifyTokenResponseProto, error)
}

type driverAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverAuthServiceClient(cc grpc.ClientConnInterface) DriverAuthServiceClient {
	return &driverAuthServiceClient{cc}
}

func (c *driverAuthServiceClient) Create(ctx context.Context, in *CreateDriverRequestProto, opts ...grpc.CallOption) (*CreateDriverResponseProto, error) {
	out := new(CreateDriverResponseProto)
	err := c.cc.Invoke(ctx, DriverAuthService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverAuthServiceClient) Login(ctx context.Context, in *LoginRequestProto, opts ...grpc.CallOption) (*LoginResponseProto, error) {
	out := new(LoginResponseProto)
	err := c.cc.Invoke(ctx, DriverAuthService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverAuthServiceClient) VerifyToken(ctx context.Context, in *VerifyTokenRequestProto, opts ...grpc.CallOption) (*VerifyTokenResponseProto, error) {
	out := new(VerifyTokenResponseProto)
	err := c.cc.Invoke(ctx, DriverAuthService_VerifyToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverAuthServiceServer is the server API for DriverAuthService service.
// All implementations must embed UnimplementedDriverAuthServiceServer
// for forward compatibility
type DriverAuthServiceServer interface {
	Create(context.Context, *CreateDriverRequestProto) (*CreateDriverResponseProto, error)
	Login(context.Context, *LoginRequestProto) (*LoginResponseProto, error)
	VerifyToken(context.Context, *VerifyTokenRequestProto) (*VerifyTokenResponseProto, error)
	mustEmbedUnimplementedDriverAuthServiceServer()
}

// UnimplementedDriverAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDriverAuthServiceServer struct {
}

func (UnimplementedDriverAuthServiceServer) Create(context.Context, *CreateDriverRequestProto) (*CreateDriverResponseProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDriverAuthServiceServer) Login(context.Context, *LoginRequestProto) (*LoginResponseProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedDriverAuthServiceServer) VerifyToken(context.Context, *VerifyTokenRequestProto) (*VerifyTokenResponseProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedDriverAuthServiceServer) mustEmbedUnimplementedDriverAuthServiceServer() {}

// UnsafeDriverAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverAuthServiceServer will
// result in compilation errors.
type UnsafeDriverAuthServiceServer interface {
	mustEmbedUnimplementedDriverAuthServiceServer()
}

func RegisterDriverAuthServiceServer(s grpc.ServiceRegistrar, srv DriverAuthServiceServer) {
	s.RegisterService(&DriverAuthService_ServiceDesc, srv)
}

func _DriverAuthService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDriverRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverAuthServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverAuthService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverAuthServiceServer).Create(ctx, req.(*CreateDriverRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverAuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverAuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverAuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverAuthServiceServer).Login(ctx, req.(*LoginRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverAuthService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverAuthServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverAuthService_VerifyToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverAuthServiceServer).VerifyToken(ctx, req.(*VerifyTokenRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverAuthService_ServiceDesc is the grpc.ServiceDesc for DriverAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "driver.DriverAuthService",
	HandlerType: (*DriverAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DriverAuthService_Create_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _DriverAuthService_Login_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _DriverAuthService_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driver.proto",
}
