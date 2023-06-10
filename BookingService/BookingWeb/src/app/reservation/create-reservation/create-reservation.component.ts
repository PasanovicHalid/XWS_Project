import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ReservationsService } from '../services/reservations.service';
import { CreateReservationRequest } from '../contracts/create-reservation-request.model';
import { UserService } from 'src/app/authentification/services/user.service';
import { AuthentificationService } from 'src/app/authentification/services/authentification.service';
import { UpdateUserRequest } from 'src/app/authentification/contracts/requests/update-user-request';
import { ToastrService } from 'ngx-toastr';
import { AccomodationService } from 'src/app/authentification/services/accommodation.service';

@Component({
  selector: 'app-create-reservation',
  templateUrl: './create-reservation.component.html',
  styleUrls: ['./create-reservation.component.scss']
})
export class CreateReservationComponent implements OnInit {
  reservation: CreateReservationRequest = new CreateReservationRequest();

  offerId: any;


  userInfo: UpdateUserRequest = new UpdateUserRequest();
  constructor(private reservationService: ReservationsService,
    private router: Router,
    private route: ActivatedRoute,
    private userService: UserService,
    private authService: AuthentificationService,
    private toastr: ToastrService,
    private accommodationService: AccomodationService) { }

  ngOnInit(): void {
    this.authService.GetIdentityId()

    this.offerId = this.route.snapshot.paramMap.get('id');

    this.userService.GetUser(this.authService.GetIdentityId()).subscribe({
      next: (response) => {
        if (response.requestResult.code != 200) {
        }
        this.userInfo = response.user
      },
      error: () => {
      }
    });
  }
  createReservation() {
    this.reservation.accommodationOfferId = this.offerId;
    this.reservation.customerId = this.userInfo.identityId;
    this.accommodationService.GetOwnerIdByAccommodationId(this.reservation.accommodationOfferId).subscribe(res => {
      this.reservation.hostId = res.id;
      console.log(this.reservation)
      this.reservationService.CreateReservation(this.reservation).subscribe({
        next: (response) => {
          if (response.requestResult.code != 200) {
            this.toastr.error(response.requestResult.message)
            return
          }
          this.toastr.success("Successfully created reservation")
        },
        error: (err) => {
          this.toastr.error("Something went wrong.")
        }
      });
    })
  }
}
