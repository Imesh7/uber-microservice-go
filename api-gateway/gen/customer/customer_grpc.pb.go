// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: customer.proto

package customer

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
	CustomerAuthService_Create_FullMethodName      = "/customer.CustomerAuthService/Create"
	CustomerAuthService_Login_FullMethodName       = "/customer.CustomerAuthService/Login"
	CustomerAuthService_VerifyToken_FullMethodName = "/customer.CustomerAuthService/VerifyToken"
)

// CustomerAuthServiceClient is the client API for CustomerAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerAuthServiceClient interface {
	Create(ctx context.Context, in *CreateCustomerRequestProto, opts ...grpc.CallOption) (*CreateCustomerResponseProto, error)
	Login(ctx context.Context, in *LoginRequestProto, opts ...grpc.CallOption) (*LoginResponseProto, error)
	VerifyToken(ctx context.Context, in *VerifyTokenRequestProto, opts ...grpc.CallOption) (*VerifyTokenResponseProto, error)
}

type customerAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerAuthServiceClient(cc grpc.ClientConnInterface) CustomerAuthServiceClient {
	return &customerAuthServiceClient{cc}
}

func (c *customerAuthServiceClient) Create(ctx context.Context, in *CreateCustomerRequestProto, opts ...grpc.CallOption) (*CreateCustomerResponseProto, error) {
	out := new(CreateCustomerResponseProto)
	err := c.cc.Invoke(ctx, CustomerAuthService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthServiceClient) Login(ctx context.Context, in *LoginRequestProto, opts ...grpc.CallOption) (*LoginResponseProto, error) {
	out := new(LoginResponseProto)
	err := c.cc.Invoke(ctx, CustomerAuthService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthServiceClient) VerifyToken(ctx context.Context, in *VerifyTokenRequestProto, opts ...grpc.CallOption) (*VerifyTokenResponseProto, error) {
	out := new(VerifyTokenResponseProto)
	err := c.cc.Invoke(ctx, CustomerAuthService_VerifyToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerAuthServiceServer is the server API for CustomerAuthService service.
// All implementations must embed UnimplementedCustomerAuthServiceServer
// for forward compatibility
type CustomerAuthServiceServer interface {
	Create(context.Context, *CreateCustomerRequestProto) (*CreateCustomerResponseProto, error)
	Login(context.Context, *LoginRequestProto) (*LoginResponseProto, error)
	VerifyToken(context.Context, *VerifyTokenRequestProto) (*VerifyTokenResponseProto, error)
	mustEmbedUnimplementedCustomerAuthServiceServer()
}

// UnimplementedCustomerAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerAuthServiceServer struct {
}

func (UnimplementedCustomerAuthServiceServer) Create(context.Context, *CreateCustomerRequestProto) (*CreateCustomerResponseProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCustomerAuthServiceServer) Login(context.Context, *LoginRequestProto) (*LoginResponseProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedCustomerAuthServiceServer) VerifyToken(context.Context, *VerifyTokenRequestProto) (*VerifyTokenResponseProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedCustomerAuthServiceServer) mustEmbedUnimplementedCustomerAuthServiceServer() {}

// UnsafeCustomerAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerAuthServiceServer will
// result in compilation errors.
type UnsafeCustomerAuthServiceServer interface {
	mustEmbedUnimplementedCustomerAuthServiceServer()
}

func RegisterCustomerAuthServiceServer(s grpc.ServiceRegistrar, srv CustomerAuthServiceServer) {
	s.RegisterService(&CustomerAuthService_ServiceDesc, srv)
}

func _CustomerAuthService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerAuthService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServiceServer).Create(ctx, req.(*CreateCustomerRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerAuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServiceServer).Login(ctx, req.(*LoginRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuthService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerAuthService_VerifyToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServiceServer).VerifyToken(ctx, req.(*VerifyTokenRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerAuthService_ServiceDesc is the grpc.ServiceDesc for CustomerAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer.CustomerAuthService",
	HandlerType: (*CustomerAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CustomerAuthService_Create_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _CustomerAuthService_Login_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _CustomerAuthService_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}
