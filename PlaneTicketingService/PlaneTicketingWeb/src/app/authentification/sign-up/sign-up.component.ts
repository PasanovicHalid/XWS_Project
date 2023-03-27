import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { SignUpRequest } from '../model/sign-up-request';
import { AuthentificationService } from '../services/authentification.service';

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent {
  public signUpRequest : SignUpRequest = new SignUpRequest();

  constructor(private authetificationService: AuthentificationService, 
    private toastr: ToastrService,
    private router: Router) {}

  public SignUp() {
    this.authetificationService.SignUpCustomer(this.signUpRequest).subscribe({
      next: () => {
        this.toastr.success("Successfull Sign Up!")
        this.router.navigate(['/'])
      },
      error: (e) => {
        this.toastr.error("Invalid Sign Up!\n" + e)
      }
    })
  }
}
