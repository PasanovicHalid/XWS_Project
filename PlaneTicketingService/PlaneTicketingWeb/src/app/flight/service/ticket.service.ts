import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { TicketPurchase } from '../model/ticket-purchase.model';
import { Observable } from 'rxjs';
import { TicketInfo } from '../model/ticket-info.model';

@Injectable({
  providedIn: 'root'
})
export class TicketService {

  apiHost: string = 'http://localhost:9000/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json' });
  
  constructor(private http: HttpClient) { }

  purchaseTicket(ticket: TicketPurchase): Observable<any>{
    return this.http.post<any>('/ticketing-api/purchase-ticket' , ticket, {headers: this.headers});
  }

  getTicketsForCustomer(id: any): Observable<TicketInfo[]> {
    return this.http.get<TicketInfo[]>('/ticketing-api/all-tickets-for-customer/'+id, {headers: this.headers});
  }
}
