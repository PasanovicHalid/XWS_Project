import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Flight } from '../model/flight.model';
import { FlightService } from '../service/flight.service';
import { TicketPurchase } from '../model/ticket-purchase.model';
import { TicketService } from '../service/ticket.service';

@Component({
  selector: 'app-ticket-purchase',
  templateUrl: './ticket-purchase.component.html',
  styleUrls: ['./ticket-purchase.component.css']
})
export class TicketPurchaseComponent implements OnInit{
  flightId: string = "";
  flight:Flight = new Flight();
  numOfTickets:number = 0;
  available:number = 0;
  ticket:TicketPurchase= new TicketPurchase();
  customerId: any;

  constructor(private router: Router,
    private flightService: FlightService,
    private ticketService: TicketService) {  
  }

  ngOnInit(): void {
    this.flightId = this.flightService.getFlightId();
    console.log(this.flightId);
    this.flightService.getFlightById(this.flightId).subscribe(res=>{
      this.flight = res;
      this.available = this.flight.AvailableNumberOfTickets;
    })
  }

  purchase(id: any){
    this.ticket.CustomerId = localStorage.getItem("userId") ?? ""
    this.ticket.FlightId=this.flightId;
    console.log(this.numOfTickets)
    this.ticket.NumberOfPurchasedTickets=this.numOfTickets;
    console.log(this.ticket.NumberOfPurchasedTickets)
    this.ticket.PriceOfTicket=this.flight.PriceOfTicket;
    this.ticketService.purchaseTicket(this.ticket).subscribe(res=>{
    })
  }
}
