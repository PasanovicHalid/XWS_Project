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

  public GetAllHostPendingReservations(id: any): Observable<CreateReservationRequest[]> {
    return this.http.get<CreateReservationRequest[]>(this.basePath + 'getHostPendingReservations/'+id,  {headers: this.headers,})
  }

  public GetAllGuestPendingReservations(id: any): Observable<CreateReservationRequest[]> {
    return this.http.get<CreateReservationRequest[]>(this.basePath + 'getGuestPendingReservations/'+id,  {headers: this.headers,})
  }

  public AcceptReservation(id: string): Observable<any> {
    return this.http.put(this.basePath + 'acceptReservation', id, {headers: this.headers,})
  }

  public RejectReservation(id: string): Observable<any> {
    return this.http.put(this.basePath + 'rejectReservation', id, {headers: this.headers,})
  }
}
