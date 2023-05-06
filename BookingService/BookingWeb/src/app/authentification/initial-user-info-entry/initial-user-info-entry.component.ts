import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { UserService } from '../services/user.service';
import { CreateUserRequest } from '../contracts/requests/create-user-request';

@Component({
  selector: 'app-initial-user-info-entry',
  templateUrl: './initial-user-info-entry.component.html',
  styleUrls: ['./initial-user-info-entry.component.scss']
})
export class InitialUserInfoEntryComponent {

  initialUserInfoRequest : CreateUserRequest = new CreateUserRequest();

  constructor(private userService : UserService,
    private toastr: ToastrService,
    private router: Router) {}

  public Submit() {
    this.initialUserInfoRequest.identityId = localStorage.getItem("identityId") ?? ""
    
    this.userService.CreateUser(this.initialUserInfoRequest).subscribe({
      next: (response) => {
        if(response.requestResult.code != 200){
          this.toastr.error(response.requestResult.message)
          return
        }
        this.toastr.success("Successfully created user")
        this.router.navigate(['/dashboard']);
      },
      error: (err) => {
        this.toastr.error(err.error)
      }
    });
  }
}
