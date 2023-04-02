package contracts
import (
	"encoding/json"
	"io"
	"time"
	
)

type TicketInfoContract struct {
	Start                    time.Time
	End                      time.Time
	DepartureLocation        string
	DestinationLocation      string
	PriceOfTicket            float64
}

func (f *TicketInfoContract) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(f)
}