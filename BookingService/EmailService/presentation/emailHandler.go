package presentation

import (
	"context"
	"fmt"

	"net/smtp"

	"github.com/PasanovicHalid/XWS_Project/BookingService/EmailService/application"

	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	email_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/email_service"
)

type EmailHandler struct {
	email_pb.UnimplementedEmailServiceServer
	emailService *application.EmailService
}

func NewEmailHandler(emailService *application.EmailService) *EmailHandler {
	return &EmailHandler{
		emailService: emailService,
	}
}

func (handler *EmailHandler) SendEmail(ctx context.Context, request *email_pb.EmailRequest) (*common_pb.RequestResult, error) {

	from := "bezbednost.projekat.2023@gmail.com"
	password := "oyyucjesoapbgwlv"

	//change address
	toEmailAddress := "bezbednost.projekat.2023@gmail.com"
	to := []string{toEmailAddress}
	fmt.Print("\ngggggg\n\n")
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	//change
	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)
	fmt.Print("123\n\n")
	auth := smtp.PlainAuth("", from, password, host)

	fmt.Print("AAAAAA\n\n")
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}

	return &common_pb.RequestResult{
		Code:    200,
		Message: "Email sent successfully",
	}, nil
}

func (handler *EmailHandler) UpdateWantedNotifications(ctx context.Context, request *email_pb.UpdateWantedNotificationsRequest) (*common_pb.RequestResult, error) {
	_, err := handler.emailService.UpdateWantedNotifications(request)
	if err != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: "Error",
		}, nil
	}
	return &common_pb.RequestResult{
		Code:    200,
		Message: "Email sent successfully",
	}, nil
}
