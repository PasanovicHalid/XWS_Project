import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CreateReservationRequest } from '../contracts/create-reservation-request.model';
@Injectable({
  providedIn: 'root'
})
export class ReservationsService {

  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  basePath: string = "/booking/api/reservation/"
  constructor(private http: HttpClient) { }

  public CreateReservation(reservation: CreateReservationRequest): Observable<any> {
    return this.http.post(this.basePath + 'createReservation', reservation, {headers: this.headers,})
  }
}
