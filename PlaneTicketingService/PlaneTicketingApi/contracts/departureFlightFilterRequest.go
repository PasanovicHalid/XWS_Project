package contracts

import (
	"encoding/json"
	"io"
)

type DepartureFlightFilter struct {
	Date                string
	DepartureLocation   string
	DestinationLocation string
	NumberOfTickets     int
}

func (f *DepartureFlightFilter) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(f)
}
