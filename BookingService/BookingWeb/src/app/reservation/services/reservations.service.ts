import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CreateReservationRequest } from '../contracts/create-reservation-request.model';
import { Reservation } from '../contracts/reservation.model';
import { Reservations } from '../contracts/reservations.model';
import { ReservationId } from '../contracts/reservation-id.model';

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

  public GetAllHostPendingReservations(id: any): Observable<Reservations> {
    return this.http.get<Reservations>(this.basePath + 'getHostPendingReservations/'+id,  {headers: this.headers,})
  }

  public GetAllGuestPendingReservations(id: any): Observable<Reservations> {
    return this.http.get<Reservations>(this.basePath + 'getGuestPendingReservations/'+id,  {headers: this.headers,})
  }

  public AcceptReservation(id: string): Observable<any> {
    const requestBody = { "id": id };
    console.log(requestBody)
    return this.http.put(this.basePath + 'acceptReservation', requestBody, { headers: this.headers });
  }

  public RejectReservation(id: string): Observable<any> {
    const requestBody = { "id": id };
    console.log(requestBody)
    return this.http.put(this.basePath + 'rejectReservation', requestBody, {headers: this.headers,})
  }

  public CancelReservation(id: string): Observable<any> {
    const requestBody = { "id": id };
    console.log(requestBody)
    return this.http.put(this.basePath + 'cancelReservation', requestBody, { headers: this.headers });
  }

  public GetAllGuestAcceptedReservations(id: any): Observable<Reservations> {
    return this.http.get<Reservations>(this.basePath + 'getGuestAcceptedReservations/'+id,  {headers: this.headers,})
  }
}
