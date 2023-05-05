import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Observable, catchError, throwError } from 'rxjs';
import { LoginRequest } from '../contracts/requests/login-request';
import { SignUpRequest } from '../contracts/requests/sign-up-request';

@Injectable({
  providedIn: 'root'
})
export class AuthentificationService {

  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  basePath: string = "/booking/api/authenticate/"

  constructor(private http: HttpClient, private jwtHelper: JwtHelperService) { }

  public Login(loginRequest: LoginRequest): Observable<any> {
    return this.http.post(
                      this.basePath + 'login',
                      loginRequest,
                      { 
                        headers: this.headers,
                      },
                      )
                      .pipe(catchError(this.handleError))
  }

  public SignUp(signUpRequest: SignUpRequest): Observable<any> {
    return this.http.post(
                      this.basePath + 'register',
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
      localStorage.setItem("identityId", tokenContent.Id)
      localStorage.setItem("role", tokenContent.Role)
      localStorage.setItem("usename", tokenContent.Username)
  }

  public IsHost() : boolean {
    return localStorage.getItem("role") == "Host"
  }

  public IsGuest() : boolean {
    return localStorage.getItem("role") == "Guest"
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
