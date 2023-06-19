package message_queues

import (
	"context"
	"log"
	"time"

	email_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/email_service"

	"github.com/PasanovicHalid/XWS_Project/BookingService/EmailService/application"
	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
	events "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/notifications"
)

type NotificationHandler struct {
	emailService           *application.EmailService
	notificationSubscriber saga.Subscriber
}

func NewNotificationHandler(emailService *application.EmailService, notificationSubscriber saga.Subscriber) (*NotificationHandler, error) {
	handler := &NotificationHandler{
		emailService:           emailService,
		notificationSubscriber: notificationSubscriber,
	}

	err := handler.notificationSubscriber.Subscribe(handler.handle)

	if err != nil {
		return nil, err
	}

	return handler, nil
}

func (handler *NotificationHandler) handle(notification *events.NotifyUserNotification) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	request := email_pb.Empty{
		Id: notification.UserInfo.UserId,
	}

	switch notification.Type {
	case events.CreatedReservation:
		_, err := handler.emailService.CreatedReservationNotification(ctx, &request)
		if err != nil {
			return
		}
		break
	case events.ReservationCancelled:
		_, err := handler.emailService.CanceledReservationNotification(ctx, &request)
		if err != nil {
			return
		}
		break
	case events.RatedHost:
		_, err := handler.emailService.HostRatingGivenNotification(ctx, &request)
		if err != nil {
			return
		}
		break
	case events.RatedAccommodation:
		_, err := handler.emailService.AccommodationRatingGivenNotification(ctx, &request)
		if err != nil {
			return
		}
		break
	case events.DistinguishedChanged:
		_, err := handler.emailService.ProminentHostStatusNotification(ctx, &request)
		if err != nil {
			return
		}
		break
	case events.ReservationResponded:
		_, err := handler.emailService.HostResponseOnAccommodationRequestNotification(ctx, &request)
		if err != nil {
			return
		}
		break
	}
	log.Println("Notification received")
	log.Println(notification)
}
