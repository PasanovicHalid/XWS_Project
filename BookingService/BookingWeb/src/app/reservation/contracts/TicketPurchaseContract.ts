export class TicketPurchase {
    customerId: string = "";
    flightId: string = "";
    numberOfPurchasedTickets: number = 0;
    priceOfTicket: number = 0;

  constructor(
    public tcustomerId: string,
    public tflightId: string,
    public tnumberOfPurchasedTickets: number,
    public tpriceOfTicket: number
  ) {this.customerId = tcustomerId;
    this.flightId = tflightId;
    this.numberOfPurchasedTickets = tnumberOfPurchasedTickets;
    this.priceOfTicket = tpriceOfTicket;}
}