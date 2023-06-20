package contracts

import (
	"encoding/json"
	"io"
)

type DestinationFlightFilter struct {
	Date                string
	DepartureLocation   string
	DestinationLocation string
	NumberOfTickets     int
}

func (f *DestinationFlightFilter) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(f)
}
