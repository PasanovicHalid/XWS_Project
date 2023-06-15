package notifications

type NotifyUserNotificationType int8

const (
	CreatedReservation NotifyUserNotificationType = iota
	ReservationCancelled
	RatedHost
	RatedAccommodation
	DistinguishedChanged
	ReservationResponded
)

type NotifyUserNotification struct {
	Type     NotifyUserNotificationType
	UserInfo NotifyUserEventInfo
}

type NotifyUserEventInfo struct {
	UserId              string
	Role                string
	Rating              float64
	Distinguished       bool
	ReservationAccepted bool
}
