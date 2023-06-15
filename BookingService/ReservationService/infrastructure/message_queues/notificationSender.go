package message_queues

import (
	mq "github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/application/common/interfaces/infrastructure/message_queues"
	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
	"github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/notifications"
)

type NotificationSender struct {
	mq.INotificationSender
	notificationPublisher saga.Publisher
}

func NewNotificationSender(notificationPublisher saga.Publisher) *NotificationSender {
	return &NotificationSender{notificationPublisher: notificationPublisher}
}

func (n *NotificationSender) SendNotification(notification *notifications.NotifyUserNotification) error {
	return n.notificationPublisher.Publish(notification)
}
