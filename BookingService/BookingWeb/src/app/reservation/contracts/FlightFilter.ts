export class FlightFilter {
    date: string = "";
    departureLocation: string = "";
    destinationLocation: string = "";
    numberOfTickets: number = 0;
  
    constructor(
      date: string,
      departureLocation: string,
      destinationLocation: string,
      numberOfTickets: number
    ) {
      this.date = date;
      this.departureLocation = departureLocation;
      this.destinationLocation = destinationLocation;
      this.numberOfTickets = numberOfTickets;
    }
  }
  