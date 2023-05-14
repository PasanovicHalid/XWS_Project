import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ReservationsService } from '../services/reservations.service';
import { CreateReservationRequest } from '../contracts/create-reservation-request.model';
import { UserService } from 'src/app/authentification/services/user.service';
import { AuthentificationService } from 'src/app/authentification/services/authentification.service';
import { UpdateUserRequest } from 'src/app/authentification/contracts/requests/update-user-request';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-create-reservation',
  templateUrl: './create-reservation.component.html',
  styleUrls: ['./create-reservation.component.scss']
})
export class CreateReservationComponent implements OnInit{
  reservation:CreateReservationRequest = new CreateReservationRequest();
  

  userInfo: UpdateUserRequest = new UpdateUserRequest();
  constructor(private reservationService : ReservationsService,
              private router: Router,
              private userService : UserService,
              private authService : AuthentificationService,
              private toastr: ToastrService) {}

  ngOnInit(): void {
    this.authService.GetIdentityId()

    this.userService.GetUser(this.authService.GetIdentityId()).subscribe({
      next: (response) => {
        if(response.requestResult.code != 200){
        }
        this.userInfo = response.user
      },
      error: () => {
      }
    });
  }
  createReservation(){
    this.reservation.accommodationOfferId = "12";
    this.reservation.customerId = this.userInfo.identityId;
    this.reservation.hostId = "24";
    console.log(this.reservation)
    this.reservationService.CreateReservation(this.reservation).subscribe({
      next: (response) => {
        if(response.requestResult.code != 200){
          this.toastr.error(response.requestResult.message)
          return
        }
        this.toastr.success("Successfully created reservation")
      },
      error: (err) => {
        this.toastr.error("Something went wrong.")
      }
    });

  }
}
