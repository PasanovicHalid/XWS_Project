import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { LoginRequest } from '../model/login-request';
import { AuthentificationService } from '../services/authentification.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  public loginRequest : LoginRequest = new LoginRequest();

  constructor(private authetificationService: AuthentificationService, 
              private toastr: ToastrService,
              private router: Router) {}

  public Login(){
    this.authetificationService.Login(this.loginRequest).subscribe({
      next: (token) => {
        if(this.authetificationService.AddTokenWithInfoToLocalStorage(token) === null){
          this.toastr.error("Something failed with decoding the token")
          return
        }
        this.toastr.success("Successfull login!")

        if(this.authetificationService.IsAdmin()){
          this.router.navigate(['/admin-homepage']) 
        } else {
          this.router.navigate(['/customer-homepage'])
        }
      },
      error: (e) => {
        this.toastr.error("Invalid login")
      }
    })
  }
}
