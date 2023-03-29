package contracts

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightContract struct {
	Id                       primitive.ObjectID
	Start                    time.Time
	End                      time.Time
	DepartureLocation        string
	DestinationLocation      string
	PriceOfTicket            float64
	MaxNumberOfTickets       int
	AvailableNumberOfTickets int
}

func (f *FlightContract) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(f)
}
