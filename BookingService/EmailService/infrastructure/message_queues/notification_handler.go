package message_queues

import (
	"log"

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
	switch notification.Type {
	case events.CreatedReservation:
		break
	case events.ReservationCancelled:
		break
	case events.RatedHost:
		break
	case events.RatedAccommodation:
		break
	case events.DistinguishedChanged:
		break
	case events.ReservationResponded:
		break
	}
	log.Println("Notification received")
	log.Println(notification)
}
