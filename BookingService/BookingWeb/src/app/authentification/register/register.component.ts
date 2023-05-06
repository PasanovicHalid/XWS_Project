import { Component } from '@angular/core';
import { SignUpRequest } from '../contracts/requests/sign-up-request';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { AuthentificationService } from '../services/authentification.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent {
  hide = true;
  signUpRequest : SignUpRequest = new SignUpRequest();
  conformationPassword : string = "";

  constructor(private authService : AuthentificationService,
    private toastr: ToastrService,
    private router: Router) {}
  
  public SignUp() {
    if(this.signUpRequest.password != this.conformationPassword){
      this.toastr.error("Passwords do not match")
      return
    }

    this.authService.SignUp(this.signUpRequest).subscribe({
      next: (response) => {
        if(response.requestResult.code != 200){
          this.toastr.error(response.requestResult.message)
          return
        }
        this.authService.AddTokenWithInfoToLocalStorage(response.token)
        this.toastr.success("Successfully registered")
        this.router.navigate(['/initial-user-info-entry']);
      },
      error: (err) => {
        this.toastr.error("Something went wrong.")
      }
    });
  }
}
