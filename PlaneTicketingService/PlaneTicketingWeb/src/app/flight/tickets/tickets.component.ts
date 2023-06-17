import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TicketService } from '../service/ticket.service';
import { TicketInfo } from '../model/ticket-info.model';

@Component({
  selector: 'app-tickets',
  templateUrl: './tickets.component.html',
  styleUrls: ['./tickets.component.css']
})
export class TicketsComponent  implements OnInit{

  tickets: TicketInfo[] = [];
  customerId:any;
  
  constructor(private router: Router,
    private ticketService: TicketService) {  
  }

  ngOnInit(): void {
    this.customerId = localStorage.getItem("userId");
    this.ticketService.getTicketsForCustomer(this.customerId).subscribe(res=>{
      this.tickets = res;
    })
  }
}
