package domain

type Reservation struct {
	Id                   string            `bson:"_id,omitempty"`
	AccommodationOfferId string            `bson:"offerId,omitempty"`
	CustomerId           string            `bson:"customerId,omitempty"`
	Status               ReservationStatus `bson:"status,omitempty"`
	NumberOfGuests       int               `bson:"numOfGuests,omitempty"`
	DateRange            DateRange         `bson:"dateRange,omitempty"`
}

type ReservationStatus int

const (
	Pending ReservationStatus = iota
	Accepted
	Rejected
)
