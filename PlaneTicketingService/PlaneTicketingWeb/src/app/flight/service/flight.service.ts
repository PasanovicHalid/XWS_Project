import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { EMPTY, Observable, empty } from 'rxjs';
import { Flight } from '../model/flight.model';
import { City } from '../model/cityResponse.model';
import { FlightFilter } from '../model/flightFilterRequest.model';

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
  
  getCities() : Observable<City[]>{
    return this.http.get<City[]>('/ticketing-api/flights/cities', {headers: this.headers});
  }

  filterFlights(flightFilter : FlightFilter) : Observable<any>{
    if(flightFilter.NumberOfTickets == 0 || flightFilter.Date == '' || flightFilter.DepartureLocation == '' || flightFilter.Destinationlocation == ''){
      alert("Please enter all informations")
      return new Observable<any>;
    }
    else{
      return this.http.post<any>('/ticketing-api/flights/filter' , flightFilter, {headers: this.headers});
    }
  }
}
