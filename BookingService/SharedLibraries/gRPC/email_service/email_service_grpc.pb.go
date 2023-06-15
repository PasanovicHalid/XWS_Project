// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: email_service/email_service.proto

package email

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
	EmailService_SendEmail_FullMethodName                                      = "/email.EmailService/SendEmail"
	EmailService_UpdateWantedNotifications_FullMethodName                      = "/email.EmailService/UpdateWantedNotifications"
	EmailService_CreateNotificationSettings_FullMethodName                     = "/email.EmailService/CreateNotificationSettings"
	EmailService_CreatedReservationNotification_FullMethodName                 = "/email.EmailService/CreatedReservationNotification"
	EmailService_CanceledReservationNotification_FullMethodName                = "/email.EmailService/CanceledReservationNotification"
	EmailService_HostRatingGivenNotification_FullMethodName                    = "/email.EmailService/HostRatingGivenNotification"
	EmailService_AccommodationRatingGivenNotification_FullMethodName           = "/email.EmailService/AccommodationRatingGivenNotification"
	EmailService_ProminentHostStatusNotification_FullMethodName                = "/email.EmailService/ProminentHostStatusNotification"
	EmailService_HostResponseOnAccommodationRequestNotification_FullMethodName = "/email.EmailService/HostResponseOnAccommodationRequestNotification"
)

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailServiceClient interface {
	SendEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*common.RequestResult, error)
	UpdateWantedNotifications(ctx context.Context, in *UpdateWantedNotificationsRequest, opts ...grpc.CallOption) (*common.RequestResult, error)
	CreateNotificationSettings(ctx context.Context, in *CreateNotificationSettingsRequest, opts ...grpc.CallOption) (*common.RequestResult, error)
	CreatedReservationNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error)
	CanceledReservationNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error)
	HostRatingGivenNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error)
	AccommodationRatingGivenNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error)
	ProminentHostStatusNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error)
	HostResponseOnAccommodationRequestNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) SendEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, EmailService_SendEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) UpdateWantedNotifications(ctx context.Context, in *UpdateWantedNotificationsRequest, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, EmailService_UpdateWantedNotifications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) CreateNotificationSettings(ctx context.Context, in *CreateNotificationSettingsRequest, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, EmailService_CreateNotificationSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) CreatedReservationNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, EmailService_CreatedReservationNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) CanceledReservationNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, EmailService_CanceledReservationNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) HostRatingGivenNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, EmailService_HostRatingGivenNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) AccommodationRatingGivenNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, EmailService_AccommodationRatingGivenNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) ProminentHostStatusNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, EmailService_ProminentHostStatusNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) HostResponseOnAccommodationRequestNotification(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.RequestResult, error) {
	out := new(common.RequestResult)
	err := c.cc.Invoke(ctx, EmailService_HostResponseOnAccommodationRequestNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
// All implementations must embed UnimplementedEmailServiceServer
// for forward compatibility
type EmailServiceServer interface {
	SendEmail(context.Context, *EmailRequest) (*common.RequestResult, error)
	UpdateWantedNotifications(context.Context, *UpdateWantedNotificationsRequest) (*common.RequestResult, error)
	CreateNotificationSettings(context.Context, *CreateNotificationSettingsRequest) (*common.RequestResult, error)
	CreatedReservationNotification(context.Context, *Empty) (*common.RequestResult, error)
	CanceledReservationNotification(context.Context, *Empty) (*common.RequestResult, error)
	HostRatingGivenNotification(context.Context, *Empty) (*common.RequestResult, error)
	AccommodationRatingGivenNotification(context.Context, *Empty) (*common.RequestResult, error)
	ProminentHostStatusNotification(context.Context, *Empty) (*common.RequestResult, error)
	HostResponseOnAccommodationRequestNotification(context.Context, *Empty) (*common.RequestResult, error)
	mustEmbedUnimplementedEmailServiceServer()
}

// UnimplementedEmailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (UnimplementedEmailServiceServer) SendEmail(context.Context, *EmailRequest) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedEmailServiceServer) UpdateWantedNotifications(context.Context, *UpdateWantedNotificationsRequest) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWantedNotifications not implemented")
}
func (UnimplementedEmailServiceServer) CreateNotificationSettings(context.Context, *CreateNotificationSettingsRequest) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotificationSettings not implemented")
}
func (UnimplementedEmailServiceServer) CreatedReservationNotification(context.Context, *Empty) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatedReservationNotification not implemented")
}
func (UnimplementedEmailServiceServer) CanceledReservationNotification(context.Context, *Empty) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanceledReservationNotification not implemented")
}
func (UnimplementedEmailServiceServer) HostRatingGivenNotification(context.Context, *Empty) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostRatingGivenNotification not implemented")
}
func (UnimplementedEmailServiceServer) AccommodationRatingGivenNotification(context.Context, *Empty) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccommodationRatingGivenNotification not implemented")
}
func (UnimplementedEmailServiceServer) ProminentHostStatusNotification(context.Context, *Empty) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProminentHostStatusNotification not implemented")
}
func (UnimplementedEmailServiceServer) HostResponseOnAccommodationRequestNotification(context.Context, *Empty) (*common.RequestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostResponseOnAccommodationRequestNotification not implemented")
}
func (UnimplementedEmailServiceServer) mustEmbedUnimplementedEmailServiceServer() {}

// UnsafeEmailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServiceServer will
// result in compilation errors.
type UnsafeEmailServiceServer interface {
	mustEmbedUnimplementedEmailServiceServer()
}

func RegisterEmailServiceServer(s grpc.ServiceRegistrar, srv EmailServiceServer) {
	s.RegisterService(&EmailService_ServiceDesc, srv)
}

func _EmailService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_SendEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_UpdateWantedNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWantedNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).UpdateWantedNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_UpdateWantedNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).UpdateWantedNotifications(ctx, req.(*UpdateWantedNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_CreateNotificationSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).CreateNotificationSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_CreateNotificationSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).CreateNotificationSettings(ctx, req.(*CreateNotificationSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_CreatedReservationNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).CreatedReservationNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_CreatedReservationNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).CreatedReservationNotification(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_CanceledReservationNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).CanceledReservationNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_CanceledReservationNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).CanceledReservationNotification(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_HostRatingGivenNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).HostRatingGivenNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_HostRatingGivenNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).HostRatingGivenNotification(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_AccommodationRatingGivenNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).AccommodationRatingGivenNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_AccommodationRatingGivenNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).AccommodationRatingGivenNotification(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_ProminentHostStatusNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).ProminentHostStatusNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_ProminentHostStatusNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).ProminentHostStatusNotification(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_HostResponseOnAccommodationRequestNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).HostResponseOnAccommodationRequestNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailService_HostResponseOnAccommodationRequestNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).HostResponseOnAccommodationRequestNotification(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailService_ServiceDesc is the grpc.ServiceDesc for EmailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "email.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _EmailService_SendEmail_Handler,
		},
		{
			MethodName: "UpdateWantedNotifications",
			Handler:    _EmailService_UpdateWantedNotifications_Handler,
		},
		{
			MethodName: "CreateNotificationSettings",
			Handler:    _EmailService_CreateNotificationSettings_Handler,
		},
		{
			MethodName: "CreatedReservationNotification",
			Handler:    _EmailService_CreatedReservationNotification_Handler,
		},
		{
			MethodName: "CanceledReservationNotification",
			Handler:    _EmailService_CanceledReservationNotification_Handler,
		},
		{
			MethodName: "HostRatingGivenNotification",
			Handler:    _EmailService_HostRatingGivenNotification_Handler,
		},
		{
			MethodName: "AccommodationRatingGivenNotification",
			Handler:    _EmailService_AccommodationRatingGivenNotification_Handler,
		},
		{
			MethodName: "ProminentHostStatusNotification",
			Handler:    _EmailService_ProminentHostStatusNotification_Handler,
		},
		{
			MethodName: "HostResponseOnAccommodationRequestNotification",
			Handler:    _EmailService_HostResponseOnAccommodationRequestNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email_service/email_service.proto",
}
