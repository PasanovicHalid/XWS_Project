import { Injectable } from '@angular/core';
import {
  HttpClient,
  HttpErrorResponse,
  HttpHeaders,
} from '@angular/common/http';
import { catchError, Observable, throwError } from 'rxjs';
import { LoginRequest } from '../model/login-request';
import { JwtHelperService } from '@auth0/angular-jwt';
import { SignUpRequest } from '../model/sign-up-request';

@Injectable({
  providedIn: 'root'
})
export class AuthentificationService {
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  constructor(private http: HttpClient, private jwtHelper: JwtHelperService) { }

  public Login(loginRequest: LoginRequest): Observable<string> {
    return this.http.post(
                      '/ticketing-api/users/login',
                      loginRequest,
                      { 
                        headers: this.headers,
                        responseType : 'text'
                      },
                      )
                      .pipe(catchError(this.handleError))
  }

  public GetUserInfo(): Observable<any> {
    return this.http.get(
                      '/ticketing-api/users/info',
                      {
                        headers: this.headers,
                      },
                      )
                      .pipe(catchError(this.handleError))
  }

  public GenerateApiKey(duration: any, forever: boolean): Observable<any> {
    return this.http.post(
                      '/ticketing-api/users/generate-api-key',
                      {
                        duration: duration,
                        durationForever: forever
                      },
                      {
                        headers: this.headers,
                      },
                      )
                      .pipe(catchError(this.handleError))
  }

  public SignUpCustomer(signUpRequest: SignUpRequest): Observable<unknown> {
    return this.http.post(
                      '/ticketing-api/users/signup/customer',
                      signUpRequest,
                      { 
                        headers: this.headers,
                      },
                      )
                      .pipe(catchError(this.handleError))
  }

  public AddTokenWithInfoToLocalStorage(token: string) : void | null {
      let tokenContent = this.DecodeToken(token)

      if(tokenContent == null){
        return null
      }
      
      localStorage.setItem("jwt", token)
      localStorage.setItem("userId", tokenContent.Uid)
      localStorage.setItem("userRole", tokenContent.Role)
      localStorage.setItem("user", tokenContent.Username)
      localStorage.setItem("userFirstName", tokenContent.FirstName)
      localStorage.setItem("userLastName", tokenContent.LastName)
  }

  public IsAdmin() : boolean {
    return localStorage.getItem("userRole") == "admin"
  }

  public IsCustomer() : boolean {
    return localStorage.getItem("userRole") == "customer"
  }

  public IsLoggedIn() : boolean {
    let jwt = localStorage.getItem("jwt")
    if(!(typeof jwt != 'undefined' && jwt))
      return false;
    return true;
  }

  public DecodeToken(token: string) {
    try {
      return this.jwtHelper.decodeToken(token)
    } catch (error) {
      return null
    }
  }

  private handleError(error: HttpErrorResponse) {
    return throwError(() => new Error(error.status + '\n' + error.error));
  }
}