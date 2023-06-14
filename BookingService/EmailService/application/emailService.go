package application

import (
	"context"
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
