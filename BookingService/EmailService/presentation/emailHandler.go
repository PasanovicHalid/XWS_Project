package presentation

import (
	"context"

	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	email_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/email_service"
)

type EmailHandler struct {
	email_pb.UnimplementedEmailServiceServer
}

func NewEmailHandler() *EmailHandler {
	return &EmailHandler{}
}

func (handler *EmailHandler) SendEmail(ctx context.Context, request *email_pb.EmailRequest) (*common_pb.RequestResult, error) {
	return &common_pb.RequestResult{
		Code:    200,
		Message: "Email sent successfully",
	}, nil
}
