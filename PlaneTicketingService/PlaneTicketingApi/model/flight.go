package model

type Flight struct {
	StartDateTimeUTC    string
	EndDateTimeUTC      string
	DepartureLocation   string
	DestinationLocation string
	Price               float64
	AvailableTickets    []Ticket
}
