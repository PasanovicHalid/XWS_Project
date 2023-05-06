import { OnInit } from '@angular/core';
import { Component } from '@angular/core';
import { AuthentificationService } from '../services/authentification.service';
import { SignUpRequest } from '../contracts/requests/sign-up-request';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {
  hide : boolean = true;
  signUpRequest : SignUpRequest = new SignUpRequest();

  constructor(private authService : AuthentificationService,
    private toastr: ToastrService,
    private router: Router) {}

  public Login() {
    this.authService.Login(this.signUpRequest).subscribe({
      next: (response) => {
        if(response.requestResult.code != 200){
          this.toastr.error(response.requestResult.message)
          return
        }
        this.authService.AddTokenWithInfoToLocalStorage(response.token)
        this.router.navigate(['/dashboard']);
      },
      error: (err) => {
        this.toastr.error("Something went wrong.")
      }
    });
  }

}
