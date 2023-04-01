package contracts

import (
	"encoding/json"
	"io"
)

type FlightFilter struct {
	Date                string
	DepartureLocation   string
	DestinationLocation string
	NumberOfTickets     int
}

func (f *FlightFilter) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(f)
}
