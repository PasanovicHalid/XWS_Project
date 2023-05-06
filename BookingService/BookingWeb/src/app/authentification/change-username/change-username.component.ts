import { Component } from '@angular/core';
import { ChangeUsernameRequest } from '../contracts/requests/change-username-request';
import { ToastrService } from 'ngx-toastr';
import { AuthentificationService } from '../services/authentification.service';

@Component({
  selector: 'app-change-username',
  templateUrl: './change-username.component.html',
  styleUrls: ['./change-username.component.scss']
})
export class ChangeUsernameComponent {
  hide = true;
  changeUsernameRequest : ChangeUsernameRequest = new ChangeUsernameRequest();

  constructor(private authService : AuthentificationService,
    private toastr: ToastrService) { }
  
  public ChangeUsername() {
    this.authService.ChangeUsername(this.changeUsernameRequest).subscribe({
      next: (response) => {
        if(response.requestResult.code != 200){
          this.toastr.error(response.requestResult.message)
          return
        }
        this.toastr.success("Successfully changed username")
      },
      error: (err) => {
        this.toastr.error("Something went wrong.")
      }
    });
  }
}
