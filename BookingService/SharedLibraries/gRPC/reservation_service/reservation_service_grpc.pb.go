// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: reservation_service/reservation_service.proto

package reservation

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

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	GetReservationById(ctx context.Context, in *GetReservationByIdRequest, opts ...grpc.CallOption) (*GetReservationByIdResponse, error)
	GetAllReservations(ctx context.Context, in *GetAllReservationsRequest, opts ...grpc.CallOption) (*GetAllReservationsResponse, error)
	CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error)
	DeleteReservation(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error)
	UpdateReservation(ctx context.Context, in *UpdateReservationRequest, opts ...grpc.CallOption) (*UpdateReservationResponse, error)
	CheckGuestActiveReservations(ctx context.Context, in *CheckUserActiveReservationsRequest, opts ...grpc.CallOption) (*CheckUserActiveReservationsResponse, error)
	CheckHostActiveReservations(ctx context.Context, in *CheckUserActiveReservationsRequest, opts ...grpc.CallOption) (*CheckUserActiveReservationsResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) GetReservationById(ctx context.Context, in *GetReservationByIdRequest, opts ...grpc.CallOption) (*GetReservationByIdResponse, error) {
	out := new(GetReservationByIdResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetReservationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllReservations(ctx context.Context, in *GetAllReservationsRequest, opts ...grpc.CallOption) (*GetAllReservationsResponse, error) {
	out := new(GetAllReservationsResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetAllReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error) {
	out := new(CreateReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/CreateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteReservation(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error) {
	out := new(DeleteReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/DeleteReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) UpdateReservation(ctx context.Context, in *UpdateReservationRequest, opts ...grpc.CallOption) (*UpdateReservationResponse, error) {
	out := new(UpdateReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/UpdateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckGuestActiveReservations(ctx context.Context, in *CheckUserActiveReservationsRequest, opts ...grpc.CallOption) (*CheckUserActiveReservationsResponse, error) {
	out := new(CheckUserActiveReservationsResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/CheckGuestActiveReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckHostActiveReservations(ctx context.Context, in *CheckUserActiveReservationsRequest, opts ...grpc.CallOption) (*CheckUserActiveReservationsResponse, error) {
	out := new(CheckUserActiveReservationsResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/CheckHostActiveReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	GetReservationById(context.Context, *GetReservationByIdRequest) (*GetReservationByIdResponse, error)
	GetAllReservations(context.Context, *GetAllReservationsRequest) (*GetAllReservationsResponse, error)
	CreateReservation(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error)
	DeleteReservation(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error)
	UpdateReservation(context.Context, *UpdateReservationRequest) (*UpdateReservationResponse, error)
	CheckGuestActiveReservations(context.Context, *CheckUserActiveReservationsRequest) (*CheckUserActiveReservationsResponse, error)
	CheckHostActiveReservations(context.Context, *CheckUserActiveReservationsRequest) (*CheckUserActiveReservationsResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) GetReservationById(context.Context, *GetReservationByIdRequest) (*GetReservationByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationById not implemented")
}
func (UnimplementedReservationServiceServer) GetAllReservations(context.Context, *GetAllReservationsRequest) (*GetAllReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReservations not implemented")
}
func (UnimplementedReservationServiceServer) CreateReservation(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReservation not implemented")
}
func (UnimplementedReservationServiceServer) DeleteReservation(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservation not implemented")
}
func (UnimplementedReservationServiceServer) UpdateReservation(context.Context, *UpdateReservationRequest) (*UpdateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservation not implemented")
}
func (UnimplementedReservationServiceServer) CheckGuestActiveReservations(context.Context, *CheckUserActiveReservationsRequest) (*CheckUserActiveReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckGuestActiveReservations not implemented")
}
func (UnimplementedReservationServiceServer) CheckHostActiveReservations(context.Context, *CheckUserActiveReservationsRequest) (*CheckUserActiveReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHostActiveReservations not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_GetReservationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetReservationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetReservationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetReservationById(ctx, req.(*GetReservationByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetAllReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllReservations(ctx, req.(*GetAllReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/CreateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CreateReservation(ctx, req.(*CreateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeleteReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeleteReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/DeleteReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeleteReservation(ctx, req.(*DeleteReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_UpdateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).UpdateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/UpdateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).UpdateReservation(ctx, req.(*UpdateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckGuestActiveReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserActiveReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckGuestActiveReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/CheckGuestActiveReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckGuestActiveReservations(ctx, req.(*CheckUserActiveReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckHostActiveReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserActiveReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckHostActiveReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/CheckHostActiveReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckHostActiveReservations(ctx, req.(*CheckUserActiveReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReservationById",
			Handler:    _ReservationService_GetReservationById_Handler,
		},
		{
			MethodName: "GetAllReservations",
			Handler:    _ReservationService_GetAllReservations_Handler,
		},
		{
			MethodName: "CreateReservation",
			Handler:    _ReservationService_CreateReservation_Handler,
		},
		{
			MethodName: "DeleteReservation",
			Handler:    _ReservationService_DeleteReservation_Handler,
		},
		{
			MethodName: "UpdateReservation",
			Handler:    _ReservationService_UpdateReservation_Handler,
		},
		{
			MethodName: "CheckGuestActiveReservations",
			Handler:    _ReservationService_CheckGuestActiveReservations_Handler,
		},
		{
			MethodName: "CheckHostActiveReservations",
			Handler:    _ReservationService_CheckHostActiveReservations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation_service/reservation_service.proto",
}
