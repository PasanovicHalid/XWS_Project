import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Observable, catchError, throwError } from 'rxjs';
import { LoginRequest } from '../contracts/requests/login-request';
import { SignUpRequest } from '../contracts/requests/sign-up-request';
import { ChangePasswordRequest } from '../contracts/requests/change-password-request';
import { ChangeUsernameRequest } from '../contracts/requests/change-username-request';

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
    ).pipe(catchError(this.handleError))
  }

  public Logout(): void {
    localStorage.clear()
  }

  public SignUp(signUpRequest: SignUpRequest): Observable<any> {
    return this.http.post(
      this.basePath + 'register',
      signUpRequest,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public ChangePassword(changePasswordRequest: ChangePasswordRequest): Observable<any> {
    return this.http.put(
      this.basePath + 'changePassword',
      changePasswordRequest,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public ChangeUsername(changeUsernameRequest : ChangeUsernameRequest) : Observable<any> {
    return this.http.put(
      this.basePath + 'changeUsername',
      changeUsernameRequest,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public DeleteAccount(identityId : string) : Observable<any> {
    return this.http.delete(
      this.basePath + 'remove/' + identityId,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public UpdateApiKey(apiKey: string) : Observable<any> {
    return this.http.put(
      this.basePath + 'updateApiKey',
      {
        apiKey : apiKey
      },
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  public AddTokenWithInfoToLocalStorage(token: string): void | null {
    let tokenContent = this.DecodeToken(token)

    if (tokenContent == null) {
      return null
    }

    localStorage.setItem("jwt", token)
    localStorage.setItem("identityId", tokenContent.Id)
    localStorage.setItem("role", tokenContent.Role)
    localStorage.setItem("usename", tokenContent.Username)
    localStorage.setItem("apiKey", tokenContent.ApiKey)
  }

  public IsHost(): boolean {
    return localStorage.getItem("role") == "Host"
  }

  public IsGuest(): boolean {
    return localStorage.getItem("role") == "Guest"
  }

  public GetIdentityId(): string {
    return localStorage.getItem("identityId") ?? ""
  }

  GetApiKey(): string {
    return localStorage.getItem("apiKey") ?? ""
  }

  public IsLoggedIn(): boolean {
    let jwt = localStorage.getItem("jwt")
    if (!(typeof jwt != 'undefined' && jwt))
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
