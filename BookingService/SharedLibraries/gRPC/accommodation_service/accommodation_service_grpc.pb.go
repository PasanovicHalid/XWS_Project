// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: accommodation_service/accommodation_service.proto

package accommodation

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

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	CreateAccomodation(ctx context.Context, in *NewAccomodation, opts ...grpc.CallOption) (*common.RequestResult, error)
	CreateAccomodationOffer(ctx context.Context, in *CreateOfferRequest, opts ...grpc.CallOption) (*common.RequestResult, error)
	UpdateAccomodationOffer(ctx context.Context, in *AccommodationOffer, opts ...grpc.CallOption) (*common.RequestResult, error)
	FilterAccommodations(ctx context.Context, in *AccommodationSearch, opts ...grpc.CallOption) (*GetFilteredAccommodationsResponse, error)
	GetOwnerIdByAccommodationId(ctx context.Context, in *GetOwnerIdRequest, opts ...grpc.CallOption) (*GetOwnerIdResponse, error)
	SetAutomaticAcception(ctx context.Context, in *SetAutomaticStatusRequest, opts ...grpc.CallOption) (*SetAutomaticStatusResponse, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) CreateAccomodation(ctx context.Context, in *NewAccomodation, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/CreateAccomodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreateAccomodationOffer(ctx context.Context, in *CreateOfferRequest, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/CreateAccomodationOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) UpdateAccomodationOffer(ctx context.Context, in *AccommodationOffer, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/UpdateAccomodationOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) FilterAccommodations(ctx context.Context, in *AccommodationSearch, opts ...grpc.CallOption) (*GetFilteredAccommodationsResponse, error) {
	out := new(GetFilteredAccommodationsResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/FilterAccommodations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetOwnerIdByAccommodationId(ctx context.Context, in *GetOwnerIdRequest, opts ...grpc.CallOption) (*GetOwnerIdResponse, error) {
	out := new(GetOwnerIdResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/GetOwnerIdByAccommodationId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) SetAutomaticAcception(ctx context.Context, in *SetAutomaticStatusRequest, opts ...grpc.CallOption) (*SetAutomaticStatusResponse, error) {
	out := new(SetAutomaticStatusResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/SetAutomaticAcception", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	CreateAccomodation(context.Context, *NewAccomodation) (*common.RequestResult, error)
	CreateAccomodationOffer(context.Context, *CreateOfferRequest) (*common.RequestResult, error)
	UpdateAccomodationOffer(context.Context, *AccommodationOffer) (*common.RequestResult, error)
	FilterAccommodations(context.Context, *AccommodationSearch) (*GetFilteredAccommodationsResponse, error)
	GetOwnerIdByAccommodationId(context.Context, *GetOwnerIdRequest) (*GetOwnerIdResponse, error)
	SetAutomaticAcception(context.Context, *SetAutomaticStatusRequest) (*SetAutomaticStatusResponse, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) CreateAccomodation(context.Context, *NewAccomodation) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccomodation not implemented")
}
func (UnimplementedAccommodationServiceServer) CreateAccomodationOffer(context.Context, *CreateOfferRequest) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccomodationOffer not implemented")
}
func (UnimplementedAccommodationServiceServer) UpdateAccomodationOffer(context.Context, *AccommodationOffer) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccomodationOffer not implemented")
}
func (UnimplementedAccommodationServiceServer) FilterAccommodations(context.Context, *AccommodationSearch) (*GetFilteredAccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterAccommodations not implemented")
}
func (UnimplementedAccommodationServiceServer) GetOwnerIdByAccommodationId(context.Context, *GetOwnerIdRequest) (*GetOwnerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwnerIdByAccommodationId not implemented")
}
func (UnimplementedAccommodationServiceServer) SetAutomaticAcception(context.Context, *SetAutomaticStatusRequest) (*SetAutomaticStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAutomaticAcception not implemented")
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_CreateAccomodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccomodation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAccomodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/CreateAccomodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAccomodation(ctx, req.(*NewAccomodation))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreateAccomodationOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAccomodationOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/CreateAccomodationOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAccomodationOffer(ctx, req.(*CreateOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_UpdateAccomodationOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationOffer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).UpdateAccomodationOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/UpdateAccomodationOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).UpdateAccomodationOffer(ctx, req.(*AccommodationOffer))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_FilterAccommodations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationSearch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).FilterAccommodations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/FilterAccommodations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).FilterAccommodations(ctx, req.(*AccommodationSearch))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetOwnerIdByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOwnerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetOwnerIdByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/GetOwnerIdByAccommodationId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetOwnerIdByAccommodationId(ctx, req.(*GetOwnerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_SetAutomaticAcception_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAutomaticStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).SetAutomaticAcception(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/SetAutomaticAcception",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).SetAutomaticAcception(ctx, req.(*SetAutomaticStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accommodation.AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccomodation",
			Handler:    _AccommodationService_CreateAccomodation_Handler,
		},
		{
			MethodName: "CreateAccomodationOffer",
			Handler:    _AccommodationService_CreateAccomodationOffer_Handler,
		},
		{
			MethodName: "UpdateAccomodationOffer",
			Handler:    _AccommodationService_UpdateAccomodationOffer_Handler,
		},
		{
			MethodName: "FilterAccommodations",
			Handler:    _AccommodationService_FilterAccommodations_Handler,
		},
		{
			MethodName: "GetOwnerIdByAccommodationId",
			Handler:    _AccommodationService_GetOwnerIdByAccommodationId_Handler,
		},
		{
			MethodName: "SetAutomaticAcception",
			Handler:    _AccommodationService_SetAutomaticAcception_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accommodation_service/accommodation_service.proto",
}
