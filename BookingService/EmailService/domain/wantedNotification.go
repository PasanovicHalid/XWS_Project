package domain

type WantedNotification struct {
	UserId                   string `bson:"_id"`
	Deleted                  bool   `bson:"deleted"`
	SagaTimestamp            int64  `bson:"saga_timestamp"`
	CreatedRequest           bool   `bson:"created_request"`
	CanceledReservation      bool   `bson:"canceled_reservation"`
	HostRatingGiven          bool   ` bson:"host_rating_given"`
	AccommodationRatingGiven bool   `bson:"accommodation_rating_given"`
	ProminentHost            bool   `bson:"prominent_host"`
	HostResponded            bool   `bson:"hostResponded"`
}
