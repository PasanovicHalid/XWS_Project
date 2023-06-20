package contracts

import (
	"bytes"
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

func (f *FlightFilter) ToJSON() ([]byte, error) {
	return json.Marshal(f)
}

func (f *FlightFilter) ToReader() (io.Reader, error) {
	jsonBytes, err := f.ToJSON()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(jsonBytes), nil
}
