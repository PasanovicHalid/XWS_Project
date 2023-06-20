import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { FlightFilter } from "../contracts/FlightFilter";
import { Observable } from "rxjs";
import { FlightContract } from "../contracts/FlightContract";
import { TicketPurchase } from "../contracts/TicketPurchaseContract";

@Injectable({
    providedIn: 'root'
  })
export class FlightService {
  
    headers: HttpHeaders = new HttpHeaders({
        'Content-Type': 'application/json',
      });
    
      basePath: string = "/plane/flights/filter/"
      constructor(private http: HttpClient) { }

    public GetFilteredDepartures(filter : FlightFilter) : Observable<FlightContract[]> {
        return this.http.post<FlightContract[]>(this.basePath + "departures", filter, {headers : this.headers})
    }

    public GetFilteredDestinations(filter : FlightFilter) : Observable<FlightContract[]> {
        return this.http.post<FlightContract[]>(this.basePath + "destinations", filter, {headers : this.headers})
    }

    public UserExists(filter : string) : Observable<boolean> {
      return this.http.get<boolean>("/plane/users/exist/" +  filter, {headers : this.headers})
    }

    public Purchase(purchase: TicketPurchase) : Observable<any>{
      return this.http.post<any>("/plane/purchase-ticket", purchase, {headers: this.headers})
    }
}