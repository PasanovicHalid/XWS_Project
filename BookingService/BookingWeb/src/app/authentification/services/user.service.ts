import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { CreateUserRequest } from '../contracts/requests/create-user-request';
import { Observable, catchError, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  basePath: string = "/booking/api/user/"

  constructor(private http: HttpClient, private jwtHelper: JwtHelperService) { }

  public CreateUser(createUser : CreateUserRequest): Observable<any> {
    return this.http.post(
                      this.basePath + 'createUser',
                      createUser,
                      { 
                        headers: this.headers,
                      },
                      ).pipe(catchError(this.handleError))
  }

  private handleError(error: HttpErrorResponse) {
    return throwError(() => new Error(error.status + '\n' + error.error));
  }
}
