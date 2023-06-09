// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: rating_service/rating_service.proto

package rating

import (
	context "context"
	common "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RatingService_GetAllRatingsMadeByCustomer_FullMethodName   = "/rating.RatingService/GetAllRatingsMadeByCustomer"
	RatingService_GetAllRatingsForHost_FullMethodName          = "/rating.RatingService/GetAllRatingsForHost"
	RatingService_UpdateRating_FullMethodName                  = "/rating.RatingService/UpdateRating"
	RatingService_DeleteRating_FullMethodName                  = "/rating.RatingService/DeleteRating"
	RatingService_RateHost_FullMethodName                      = "/rating.RatingService/RateHost"
	RatingService_RateAccommodation_FullMethodName             = "/rating.RatingService/RateAccommodation"
	RatingService_GetAverageRatingForHost_FullMethodName       = "/rating.RatingService/GetAverageRatingForHost"
	RatingService_GetRatingForAccommodation_FullMethodName     = "/rating.RatingService/GetRatingForAccommodation"
	RatingService_GetAllRatingsForAccommodation_FullMethodName = "/rating.RatingService/GetAllRatingsForAccommodation"
)

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	GetAllRatingsMadeByCustomer(ctx context.Context, in *GetAllRatingsMadeByCustomerRequest, opts ...grpc.CallOption) (*GetAllRatingsResponse, error)
	GetAllRatingsForHost(ctx context.Context, in *GetAllRatingsForHostRequest, opts ...grpc.CallOption) (*GetAllRatingsResponse, error)
	UpdateRating(ctx context.Context, in *UpdateRatingRequest, opts ...grpc.CallOption) (*common.RequestResult, error)
	DeleteRating(ctx context.Context, in *DeleteRatingRequest, opts ...grpc.CallOption) (*common.RequestResult, error)
	RateHost(ctx context.Context, in *RateHostRequest, opts ...grpc.CallOption) (*common.RequestResult, error)
	RateAccommodation(ctx context.Context, in *RateAccommodationRequest, opts ...grpc.CallOption) (*common.RequestResult, error)
	GetAverageRatingForHost(ctx context.Context, in *GetAverageRatingForHostRequest, opts ...grpc.CallOption) (*GetAverageRatingForHostResponse, error)
	GetRatingForAccommodation(ctx context.Context, in *GetRatingForAccommodationRequest, opts ...grpc.CallOption) (*GetRatingForAccommodationResponse, error)
	GetAllRatingsForAccommodation(ctx context.Context, in *GetAllRatingsForAccommodationRequest, opts ...grpc.CallOption) (*GetAllRatingsResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) GetAllRatingsMadeByCustomer(ctx context.Context, in *GetAllRatingsMadeByCustomerRequest, opts ...grpc.CallOption) (*GetAllRatingsResponse, error) {
	out := new(GetAllRatingsResponse)
	err := c.cc.Invoke(ctx, RatingService_GetAllRatingsMadeByCustomer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetAllRatingsForHost(ctx context.Context, in *GetAllRatingsForHostRequest, opts ...grpc.CallOption) (*GetAllRatingsResponse, error) {
	out := new(GetAllRatingsResponse)
	err := c.cc.Invoke(ctx, RatingService_GetAllRatingsForHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) UpdateRating(ctx context.Context, in *UpdateRatingRequest, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, RatingService_UpdateRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) DeleteRating(ctx context.Context, in *DeleteRatingRequest, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, RatingService_DeleteRating_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) RateHost(ctx context.Context, in *RateHostRequest, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, RatingService_RateHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) RateAccommodation(ctx context.Context, in *RateAccommodationRequest, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, RatingService_RateAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetAverageRatingForHost(ctx context.Context, in *GetAverageRatingForHostRequest, opts ...grpc.CallOption) (*GetAverageRatingForHostResponse, error) {
	out := new(GetAverageRatingForHostResponse)
	err := c.cc.Invoke(ctx, RatingService_GetAverageRatingForHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetRatingForAccommodation(ctx context.Context, in *GetRatingForAccommodationRequest, opts ...grpc.CallOption) (*GetRatingForAccommodationResponse, error) {
	out := new(GetRatingForAccommodationResponse)
	err := c.cc.Invoke(ctx, RatingService_GetRatingForAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetAllRatingsForAccommodation(ctx context.Context, in *GetAllRatingsForAccommodationRequest, opts ...grpc.CallOption) (*GetAllRatingsResponse, error) {
	out := new(GetAllRatingsResponse)
	err := c.cc.Invoke(ctx, RatingService_GetAllRatingsForAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	GetAllRatingsMadeByCustomer(context.Context, *GetAllRatingsMadeByCustomerRequest) (*GetAllRatingsResponse, error)
	GetAllRatingsForHost(context.Context, *GetAllRatingsForHostRequest) (*GetAllRatingsResponse, error)
	UpdateRating(context.Context, *UpdateRatingRequest) (*common.RequestResult, error)
	DeleteRating(context.Context, *DeleteRatingRequest) (*common.RequestResult, error)
	RateHost(context.Context, *RateHostRequest) (*common.RequestResult, error)
	RateAccommodation(context.Context, *RateAccommodationRequest) (*common.RequestResult, error)
	GetAverageRatingForHost(context.Context, *GetAverageRatingForHostRequest) (*GetAverageRatingForHostResponse, error)
	GetRatingForAccommodation(context.Context, *GetRatingForAccommodationRequest) (*GetRatingForAccommodationResponse, error)
	GetAllRatingsForAccommodation(context.Context, *GetAllRatingsForAccommodationRequest) (*GetAllRatingsResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) GetAllRatingsMadeByCustomer(context.Context, *GetAllRatingsMadeByCustomerRequest) (*GetAllRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRatingsMadeByCustomer not implemented")
}
func (UnimplementedRatingServiceServer) GetAllRatingsForHost(context.Context, *GetAllRatingsForHostRequest) (*GetAllRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRatingsForHost not implemented")
}
func (UnimplementedRatingServiceServer) UpdateRating(context.Context, *UpdateRatingRequest) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRating not implemented")
}
func (UnimplementedRatingServiceServer) DeleteRating(context.Context, *DeleteRatingRequest) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRating not implemented")
}
func (UnimplementedRatingServiceServer) RateHost(context.Context, *RateHostRequest) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateHost not implemented")
}
func (UnimplementedRatingServiceServer) RateAccommodation(context.Context, *RateAccommodationRequest) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateAccommodation not implemented")
}
func (UnimplementedRatingServiceServer) GetAverageRatingForHost(context.Context, *GetAverageRatingForHostRequest) (*GetAverageRatingForHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAverageRatingForHost not implemented")
}
func (UnimplementedRatingServiceServer) GetRatingForAccommodation(context.Context, *GetRatingForAccommodationRequest) (*GetRatingForAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatingForAccommodation not implemented")
}
func (UnimplementedRatingServiceServer) GetAllRatingsForAccommodation(context.Context, *GetAllRatingsForAccommodationRequest) (*GetAllRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRatingsForAccommodation not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_GetAllRatingsMadeByCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRatingsMadeByCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAllRatingsMadeByCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetAllRatingsMadeByCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAllRatingsMadeByCustomer(ctx, req.(*GetAllRatingsMadeByCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetAllRatingsForHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRatingsForHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAllRatingsForHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetAllRatingsForHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAllRatingsForHost(ctx, req.(*GetAllRatingsForHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_UpdateRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).UpdateRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_UpdateRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).UpdateRating(ctx, req.(*UpdateRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_DeleteRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).DeleteRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_DeleteRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).DeleteRating(ctx, req.(*DeleteRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_RateHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).RateHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_RateHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).RateHost(ctx, req.(*RateHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_RateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).RateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_RateAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).RateAccommodation(ctx, req.(*RateAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetAverageRatingForHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAverageRatingForHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAverageRatingForHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetAverageRatingForHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAverageRatingForHost(ctx, req.(*GetAverageRatingForHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetRatingForAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingForAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRatingForAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetRatingForAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRatingForAccommodation(ctx, req.(*GetRatingForAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetAllRatingsForAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRatingsForAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAllRatingsForAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetAllRatingsForAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAllRatingsForAccommodation(ctx, req.(*GetAllRatingsForAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllRatingsMadeByCustomer",
			Handler:    _RatingService_GetAllRatingsMadeByCustomer_Handler,
		},
		{
			MethodName: "GetAllRatingsForHost",
			Handler:    _RatingService_GetAllRatingsForHost_Handler,
		},
		{
			MethodName: "UpdateRating",
			Handler:    _RatingService_UpdateRating_Handler,
		},
		{
			MethodName: "DeleteRating",
			Handler:    _RatingService_DeleteRating_Handler,
		},
		{
			MethodName: "RateHost",
			Handler:    _RatingService_RateHost_Handler,
		},
		{
			MethodName: "RateAccommodation",
			Handler:    _RatingService_RateAccommodation_Handler,
		},
		{
			MethodName: "GetAverageRatingForHost",
			Handler:    _RatingService_GetAverageRatingForHost_Handler,
		},
		{
			MethodName: "GetRatingForAccommodation",
			Handler:    _RatingService_GetRatingForAccommodation_Handler,
		},
		{
			MethodName: "GetAllRatingsForAccommodation",
			Handler:    _RatingService_GetAllRatingsForAccommodation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rating_service/rating_service.proto",
}
