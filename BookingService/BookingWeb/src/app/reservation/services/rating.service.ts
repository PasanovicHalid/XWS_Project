import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CreateReservationRequest } from '../contracts/create-reservation-request.model';
import { Reservations } from '../contracts/reservations.model';

@Injectable({
  providedIn: 'root'
})
export class RatingService {
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  basePath: string = "/booking/api/rating/"
  constructor(private http: HttpClient) { }

  public GetAllRatingsForAccommodation(id: any): Observable<any> {
    return this.http.get<any>(this.basePath + 'get-ratings-for-accommodation/'+ id,  {headers: this.headers,})
  }
}
