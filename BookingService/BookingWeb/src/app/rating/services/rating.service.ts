import { HttpHeaders, HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RatingService {

  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  basePath: string = "/booking/api/rating"

  constructor(private http: HttpClient) { }

  public GetHostsForRating(): Observable<any> {
    return this.http.get(
      this.basePath + '/get-hosts-for-rating',
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public GetCustomerRatings(id : any): Observable<any> {
    return this.http.post(
      this.basePath + '/customer/ratings/' + id,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public GetAccommodationsForRating(): Observable<any> {
    return this.http.get(
      this.basePath + '/get-accommodations-for-rating',
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public CreateHostRating(obj: any, id: any): Observable<any> {
    return this.http.post(
      this.basePath + '/host/' + id,
      obj,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public CreateAccommodationRating(obj: any, id: any): Observable<any> {
    return this.http.post(
      this.basePath + '/accommodation/' + id,
      obj,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public UpdateRating(obj: any): Observable<any> {
    return this.http.put( 
      this.basePath,
      obj,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public DeleteRating(id: any): Observable<any> {
    return this.http.delete(
      this.basePath + "/" + id,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  private handleError(error: HttpErrorResponse) {
    return throwError(() => new Error(error.status + '\n' + error.error));
  }
}
