import { Component } from '@angular/core';
import { ChangePasswordRequest } from '../contracts/requests/change-password-request';
import { AuthentificationService } from '../services/authentification.service';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.scss']
})
export class ChangePasswordComponent {
  hide = true;
  changePasswordRequest : ChangePasswordRequest = new ChangePasswordRequest();

  constructor(private authService : AuthentificationService,
              private toastr: ToastrService) { }
  
  public ChangePassword() {
    this.authService.ChangePassword(this.changePasswordRequest).subscribe({
      next: (response) => {
        if(response.requestResult.code != 200){
          this.toastr.error(response.requestResult.message)
          return
        }
        this.toastr.success("Successfully changed password")
      },
      error: (err) => {
        this.toastr.error("Something went wrong.")
      }
    });
  }
}
