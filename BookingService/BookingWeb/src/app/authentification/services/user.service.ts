import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { CreateUserRequest } from '../contracts/requests/create-user-request';
import { Observable, catchError, throwError } from 'rxjs';
import { UpdateUserRequest } from '../contracts/requests/update-user-request';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  basePath: string = "/booking/api/user/"

  constructor(private http: HttpClient, private jwtHelper: JwtHelperService) { }

  public CreateUser(createUser: CreateUserRequest): Observable<any> {
    return this.http.post(
      this.basePath + 'createUser',
      createUser,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public GetUser(identityId: string): Observable<any> {
    return this.http.get(
      this.basePath + 'getUserById/' + identityId,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public UpdateUser(updateUser: UpdateUserRequest): Observable<any> {
    return this.http.put(
      this.basePath + 'updateUser',
      updateUser,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public DeleteUser(identityId: string): Observable<any> {
    return this.http.delete(
      this.basePath + 'deregister/' + identityId,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  private handleError(error: HttpErrorResponse) {
    return throwError(() => new Error(error.status + '\n' + error.error));
  }
}
