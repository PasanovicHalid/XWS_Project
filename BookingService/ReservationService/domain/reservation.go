package domain

type Reservation struct {
	Id                   string            `bson:"_id,omitempty"`
	AccommodationOfferId string            `bson:"offerId,omitempty"`
	CustomerId           string            `bson:"customerId,omitempty"`
	Status               ReservationStatus `bson:"status,omitempty"`
}

type ReservationStatus string

const (
	Pennding ReservationStatus = "PENNDING"
	Accepted ReservationStatus = "ACCEPTED"
	Rejected ReservationStatus = "REJECTED"
)
