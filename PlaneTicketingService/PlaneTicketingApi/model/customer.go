package model

type Customer struct {
	User          User     `bson:"user,inline" json:"user"`
	BoughtTickets []Ticket `bson:"boughtTickets,omitempty" json:"boughtTickets"`
}
