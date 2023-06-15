package presentation

import (
	"context"

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

/*
https://www.courier.com/guides/golang-send-email/
*/
func (handler *EmailHandler) SendEmail(ctx context.Context, request *email_pb.EmailRequest) (*common_pb.RequestResult, error) {

	from := "bezbednost.projekat.2023@gmail.com"
	password := "oyyucjesoapbgwlv"

	//change address
	toEmailAddress := "bezbednost.projekat.2023@gmail.com"
	to := []string{toEmailAddress}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	//change
	subject := request.Subject
	body := request.Body
	message := []byte(subject + body)
	auth := smtp.PlainAuth("", from, password, host)

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

func (handler *EmailHandler) CreatedReservationNotification() (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   "bezbednost.projekat.2023@gmail.com",
		Subject: "Reservation Created",
		Body:    "Someone created reservation. Please check your application.",
	}

	return handler.SendEmail(context.Background(), emailReq)
}

func (handler *EmailHandler) AccommodationRatingGivenNotification() (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   "bezbednost.projekat.2023@gmail.com",
		Subject: "Accommodation rating given",
		Body:    "Someone rated your accommodation.",
	}

	return handler.SendEmail(context.Background(), emailReq)
}

func (handler *EmailHandler) CanceledReservationNotification() (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   "bezbednost.projekat.2023@gmail.com",
		Subject: "Reservation Canceled",
		Body:    "Someone canceled reservation. Please check your application.",
	}

	return handler.SendEmail(context.Background(), emailReq)
}

func (handler *EmailHandler) ProminentHostStatusNotification() (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   "bezbednost.projekat.2023@gmail.com",
		Subject: "Prominent host status changed",
		Body:    "Your host status is changed.",
	}

	return handler.SendEmail(context.Background(), emailReq)
}

func (handler *EmailHandler) HostRatingGivenNotification() (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   "bezbednost.projekat.2023@gmail.com",
		Subject: "Host Rating Given",
		Body:    "Someone gave you rating. Please check your application.",
	}

	return handler.SendEmail(context.Background(), emailReq)
}

func (handler *EmailHandler) HostResponseOnAccommodationRequestNotification() (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   "bezbednost.projekat.2023@gmail.com",
		Subject: "Host response on accommodaton request",
		Body:    "Host gave response on your accomodation request.",
	}

	return handler.SendEmail(context.Background(), emailReq)
}
