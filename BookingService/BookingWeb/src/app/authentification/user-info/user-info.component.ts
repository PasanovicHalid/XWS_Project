import { Component, OnInit } from '@angular/core';
import { UpdateUserRequest } from '../contracts/requests/update-user-request';
import { UserService } from '../services/user.service';
import { AuthentificationService } from '../services/authentification.service';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';

@Component({
  selector: 'app-user-info',
  templateUrl: './user-info.component.html',
  styleUrls: ['./user-info.component.scss']
})
export class UserInfoComponent implements OnInit {

  userInfo : UpdateUserRequest = new UpdateUserRequest();
  identityId : string = "";

  constructor(private userService : UserService,
              private authService : AuthentificationService,
              private toastr: ToastrService,
              private router: Router) {}

  ngOnInit(): void {
    this.identityId = this.authService.GetIdentityId()

    this.userService.GetUser(this.identityId).subscribe({
      next: (response) => {
        if(response.requestResult.code != 200){
          this.toastr.error(response.requestResult.message)
        }
        this.userInfo = response.user
      },
      error: () => {
        this.toastr.error("Something went wrong.")
      }
    });
  }

  public Save() {
    this.userService.UpdateUser(this.userInfo).subscribe({
      next: (response) => {
        if(response.requestResult.code != 200){
          this.toastr.error(response.requestResult.message)
          return
        }
        this.toastr.success("Successfully updated user")
      },
      error: (err) => {
        this.toastr.error("Something went wrong.")
      }
    });
  }

  public Delete() {
    this.authService.DeleteAccount(this.identityId).subscribe({
      next: (response) => {
        this.toastr.success("Successfully deleted user")
        this.authService.Logout()
        this.router.navigate(['/']);
      },
      error: (err) => {
        this.toastr.error("Something went wrong.")
      }
    });
  }
}
