package contracts

import (
	"encoding/json"
	"io"
	"time"
)

type FlightContract struct {
	Start               time.Time
	End                 time.Time
	DepartureLocation   string
	DestinationLocation string
	PriceOfTicket       float64
	NumberOfTickets     int
}

func (f *FlightContract) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(f)
}
