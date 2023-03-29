import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Flight } from '../model/flight.model';

@Injectable({
  providedIn: 'root'
})
export class FlightService {

  apiHost: string = 'http://localhost:9000/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json' });

  constructor(private http: HttpClient) { }

  getFlights(): Observable<Flight[]> {
    return this.http.get<Flight[]>('/ticketing-api/flights/all', {headers: this.headers});
  }

  createFlight(flight: Flight): Observable<any>{
    return this.http.post<any>('/ticketing-api/flight/create', flight, {headers: this.headers});
  }

  deleteFlight(id: any): Observable<any> {
    return this.http.delete<any>('/ticketing-api/flight/delete/' + id, {headers: this.headers});
  }
}
