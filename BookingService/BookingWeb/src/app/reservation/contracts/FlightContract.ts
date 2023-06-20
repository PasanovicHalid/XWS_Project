export class FlightContract {
    Id: string;
    Start: Date;
    End: Date;
    DepartureLocation: string;
    DestinationLocation: string;
    PriceOfTicket: number;
    MaxNumberOfTickets: number;
    AvailableNumberOfTickets: number;
  
    constructor(
      Id: string,
      Start: Date,
      End: Date,
      DepartureLocation: string,
      DestinationLocation: string,
      PriceOfTicket: number,
      MaxNumberOfTickets: number,
      AvailableNumberOfTickets: number
    ) {
      this.Id = Id;
      this.Start = Start;
      this.End = End;
      this.DepartureLocation = DepartureLocation;
      this.DestinationLocation = DestinationLocation;
      this.PriceOfTicket = PriceOfTicket;
      this.MaxNumberOfTickets = MaxNumberOfTickets;
      this.AvailableNumberOfTickets = AvailableNumberOfTickets;
    }
  }