package contracts
import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketPurchase struct {
	CustomerId primitive.ObjectID
	FlightId primitive.ObjectID
	NumberOfPurchasedTickets int
	PriceOfTicket float64
}

func (f *TicketPurchase) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(f)
}