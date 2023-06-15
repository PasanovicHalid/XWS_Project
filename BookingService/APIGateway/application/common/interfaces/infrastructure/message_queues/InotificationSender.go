package message_queues

import "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/notifications"

type INotificationSender interface {
	SendNotification(notification *notifications.NotifyUserNotification) error
}
