// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: authentification_service/authentification_service.proto

package authenticate

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
	AuthenticateService_Login_FullMethodName                 = "/authenticate.AuthenticateService/Login"
	AuthenticateService_Register_FullMethodName              = "/authenticate.AuthenticateService/Register"
	AuthenticateService_ChangePassword_FullMethodName        = "/authenticate.AuthenticateService/ChangePassword"
	AuthenticateService_ChangeUsername_FullMethodName        = "/authenticate.AuthenticateService/ChangeUsername"
	AuthenticateService_GetIdentityByUsername_FullMethodName = "/authenticate.AuthenticateService/GetIdentityByUsername"
	AuthenticateService_GetPublicKey_FullMethodName          = "/authenticate.AuthenticateService/GetPublicKey"
)

// AuthenticateServiceClient is the client API for AuthenticateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticateServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	ChangeUsername(ctx context.Context, in *ChangeUsernameRequest, opts ...grpc.CallOption) (*ChangeUsernameResponse, error)
	GetIdentityByUsername(ctx context.Context, in *GetIdentityByUsernameRequest, opts ...grpc.CallOption) (*GetIdentityByUsernameResponse, error)
	GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error)
}

type authenticateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticateServiceClient(cc grpc.ClientConnInterface) AuthenticateServiceClient {
	return &authenticateServiceClient{cc}
}

func (c *authenticateServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthenticateService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, AuthenticateService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, AuthenticateService_ChangePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateServiceClient) ChangeUsername(ctx context.Context, in *ChangeUsernameRequest, opts ...grpc.CallOption) (*ChangeUsernameResponse, error) {
	out := new(ChangeUsernameResponse)
	err := c.cc.Invoke(ctx, AuthenticateService_ChangeUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateServiceClient) GetIdentityByUsername(ctx context.Context, in *GetIdentityByUsernameRequest, opts ...grpc.CallOption) (*GetIdentityByUsernameResponse, error) {
	out := new(GetIdentityByUsernameResponse)
	err := c.cc.Invoke(ctx, AuthenticateService_GetIdentityByUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateServiceClient) GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error) {
	out := new(GetPublicKeyResponse)
	err := c.cc.Invoke(ctx, AuthenticateService_GetPublicKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticateServiceServer is the server API for AuthenticateService service.
// All implementations must embed UnimplementedAuthenticateServiceServer
// for forward compatibility
type AuthenticateServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	ChangeUsername(context.Context, *ChangeUsernameRequest) (*ChangeUsernameResponse, error)
	GetIdentityByUsername(context.Context, *GetIdentityByUsernameRequest) (*GetIdentityByUsernameResponse, error)
	GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error)
	mustEmbedUnimplementedAuthenticateServiceServer()
}

// UnimplementedAuthenticateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticateServiceServer struct {
}

func (UnimplementedAuthenticateServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthenticateServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthenticateServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAuthenticateServiceServer) ChangeUsername(context.Context, *ChangeUsernameRequest) (*ChangeUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUsername not implemented")
}
func (UnimplementedAuthenticateServiceServer) GetIdentityByUsername(context.Context, *GetIdentityByUsernameRequest) (*GetIdentityByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentityByUsername not implemented")
}
func (UnimplementedAuthenticateServiceServer) GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedAuthenticateServiceServer) mustEmbedUnimplementedAuthenticateServiceServer() {}

// UnsafeAuthenticateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticateServiceServer will
// result in compilation errors.
type UnsafeAuthenticateServiceServer interface {
	mustEmbedUnimplementedAuthenticateServiceServer()
}

func RegisterAuthenticateServiceServer(s grpc.ServiceRegistrar, srv AuthenticateServiceServer) {
	s.RegisterService(&AuthenticateService_ServiceDesc, srv)
}

func _AuthenticateService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateService_ChangeUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServiceServer).ChangeUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateService_ChangeUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServiceServer).ChangeUsername(ctx, req.(*ChangeUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateService_GetIdentityByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdentityByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServiceServer).GetIdentityByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateService_GetIdentityByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServiceServer).GetIdentityByUsername(ctx, req.(*GetIdentityByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticateService_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServiceServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticateService_GetPublicKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServiceServer).GetPublicKey(ctx, req.(*GetPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticateService_ServiceDesc is the grpc.ServiceDesc for AuthenticateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authenticate.AuthenticateService",
	HandlerType: (*AuthenticateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthenticateService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AuthenticateService_Register_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _AuthenticateService_ChangePassword_Handler,
		},
		{
			MethodName: "ChangeUsername",
			Handler:    _AuthenticateService_ChangeUsername_Handler,
		},
		{
			MethodName: "GetIdentityByUsername",
			Handler:    _AuthenticateService_GetIdentityByUsername_Handler,
		},
		{
			MethodName: "GetPublicKey",
			Handler:    _AuthenticateService_GetPublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentification_service/authentification_service.proto",
}