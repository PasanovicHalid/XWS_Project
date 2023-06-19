package application

import (
	"context"
	"fmt"
	"net/smtp"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/EmailService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/EmailService/domain"
	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	email_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/email_service"
)

type EmailService struct {
	emailRepository persistance.IEmailRepository
}

func NewEmailService(emailRepository persistance.IEmailRepository) *EmailService {
	return &EmailService{
		emailRepository: emailRepository,
	}
}

func (service *EmailService) UpdateWantedNotifications(request *email_pb.UpdateWantedNotificationsRequest) (*common_pb.RequestResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	wantedNotifications, err := convert(request)

	if err != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: "NEUSPESNO UNETI PODACI",
		}, nil
	}
	err2 := service.emailRepository.UpdateWantedNotifications(&ctx, wantedNotifications)
	if err2 != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: err2.Error(),
		}, nil
	}

	return &common_pb.RequestResult{
		Code:    200,
		Message: "USPESNO",
	}, nil
}

func convert(req *email_pb.UpdateWantedNotificationsRequest) (*domain.WantedNotification, error) {

	return &domain.WantedNotification{
		UserId:                   req.Id,
		CreatedRequest:           req.CreatedRequest,
		CanceledReservation:      req.CanceledReservation,
		HostRatingGiven:          req.HostRatingGiven,
		AccommodationRatingGiven: req.AccommodationRatingGiven,
		ProminentHost:            req.ProminentHost,
		HostResponded:            req.HostResponded,
	}, nil
}

func (service *EmailService) SetWantedNotifications(request *email_pb.UpdateWantedNotificationsRequest) (*common_pb.RequestResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	wantedNotifications, err := convert(request)

	fmt.Print("PPPPPP\n\n")
	if err != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: "NEUSPESNO UNETI PODACI",
		}, nil
	}
	err2 := service.emailRepository.SetWantedNotifications(&ctx, wantedNotifications)
	fmt.Print("LLLLLL\n")
	if err2 != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: err2.Error(),
		}, nil
	}

	return &common_pb.RequestResult{
		Code:    200,
		Message: "USPESNO",
	}, nil
}

/*
https://www.courier.com/guides/golang-send-email/
*/
func (service *EmailService) SendEmail(ctx context.Context, request *email_pb.EmailRequest) (*common_pb.RequestResult, error) {

	from := "bezbednost.projekat.2023@gmail.com"
	password := "oyyucjesoapbgwlv"

	//change address
	toEmailAddress := request.Email
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

func (service *EmailService) CreatedReservationNotification(ctx context.Context, request *email_pb.Empty) (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   request.Id,
		Subject: "Reservation Created\n",
		Body:    "Someone created reservation. Please check your application.",
	}

	return service.SendEmail(context.Background(), emailReq)
}

func (service *EmailService) AccommodationRatingGivenNotification(ctx context.Context, request *email_pb.Empty) (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   request.Id,
		Subject: "Accommodation rating given\n",
		Body:    "Someone rated your accommodation.",
	}

	return service.SendEmail(context.Background(), emailReq)
}

func (service *EmailService) CanceledReservationNotification(ctx context.Context, request *email_pb.Empty) (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   request.Id,
		Subject: "Reservation Canceled\n",
		Body:    "Someone canceled reservation. Please check your application.",
	}

	return service.SendEmail(context.Background(), emailReq)
}

func (service *EmailService) ProminentHostStatusNotification(ctx context.Context, request *email_pb.Empty) (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   request.Id,
		Subject: "Prominent host status changed\n",
		Body:    "Your host status is changed.",
	}

	return service.SendEmail(context.Background(), emailReq)
}

func (service *EmailService) HostRatingGivenNotification(ctx context.Context, request *email_pb.Empty) (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   request.Id,
		Subject: "Host Rating Given\n",
		Body:    "Someone gave you rating. Please check your application.",
	}

	return service.SendEmail(context.Background(), emailReq)
}

func (service *EmailService) HostResponseOnAccommodationRequestNotification(ctx context.Context, request *email_pb.Empty) (*common_pb.RequestResult, error) {
	emailReq := &email_pb.EmailRequest{
		Email:   request.Id,
		Subject: "Host response on accommodaton request\n",
		Body:    "Host gave response on your accomodation request.",
	}

	return service.SendEmail(context.Background(), emailReq)
}
