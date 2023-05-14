import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Observable, catchError, throwError } from 'rxjs';
import { LoginRequest } from '../contracts/requests/login-request';
import { SignUpRequest } from '../contracts/requests/sign-up-request';
import { ChangePasswordRequest } from '../contracts/requests/change-password-request';
import { ChangeUsernameRequest } from '../contracts/requests/change-username-request';
import { Accommodation } from 'src/app/accommodation/create-accommodation/model/accommodation.model';
import { CreateOfferRequest } from 'src/app/accommodation/create-accommodation-offer/model/accommodationOffer.model';
import { SetAutomaticStatusRequest } from '../contracts/requests/set-automatic-status-request.model';

@Injectable({
  providedIn: 'root'
})
export class AccomodationService {
  
  token = localStorage.getItem('jwt');
  temp : boolean = false;
  headers: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  basePath: string = "/booking/api/authenticate/"

  basePathTemp: string = "/booking/api/accommodation/"

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

  public CreateAccommodation(newAccommodation: Accommodation): Observable<any> {
    return this.http.post(
      this.basePathTemp + 'create',
      newAccommodation,
      {
        headers: this.headers,
      },
    ).pipe(catchError(this.handleError))
  }

  CreateOffer(newOffer: CreateOfferRequest) : Observable<any>{
    return this.http.post(
      this.basePathTemp + 'create-offer',
      newOffer,
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

  public AddTokenWithInfoToLocalStorage(token: string): void | null {
    let tokenContent = this.DecodeToken(token)

    if (tokenContent == null) {
      return null
    }

    localStorage.setItem("jwt", token)
    localStorage.setItem("identityId", tokenContent.Id)
    localStorage.setItem("role", tokenContent.Role)
    localStorage.setItem("usename", tokenContent.Username)
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

  public SetAutomaticAcception(request: SetAutomaticStatusRequest) :any{
    return this.http.post(this.basePath + 'setAutomaticAcception',request, {headers: this.headers,})
  }

  public GetAutomaticAcception(id:string) :any{
    return this.http.get(this.basePath + 'getAutomaticAcception/'+ id, {headers: this.headers,})
  }

  public GetOwnerIdByAccommodationId(id:string) :Observable<any>{
    return this.http.get<Observable<any>>(this.basePath + 'getOwnerId/'+ id, {headers: this.headers,})
  }
}
