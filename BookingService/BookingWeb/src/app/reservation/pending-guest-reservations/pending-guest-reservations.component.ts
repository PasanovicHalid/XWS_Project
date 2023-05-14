import { Component, OnInit } from '@angular/core';
import { Reservation } from '../contracts/reservation.model';
import { MatTableDataSource } from '@angular/material/table';
import { ReservationsService } from '../services/reservations.service';
import { Router } from '@angular/router';
import { AuthentificationService } from 'src/app/authentification/services/authentification.service';
import { UserService } from 'src/app/authentification/services/user.service';

@Component({
  selector: 'app-pending-guest-reservations',
  templateUrl: './pending-guest-reservations.component.html',
  styleUrls: ['./pending-guest-reservations.component.scss']
})
export class PendingGuestReservationsComponent implements OnInit{
  id:string = "";
  dataSource: MatTableDataSource<Reservation> = new MatTableDataSource<Reservation>();

  displayedColumns: string[] = ['offer', 'customer', 'number', 'start', 'end', 'accept'];

  constructor(private resetvationService: ReservationsService,
    private router: Router,
    private authService: AuthentificationService,
    private userService: UserService){}

  ngOnInit(): void {
    this.authService.GetIdentityId()

    this.userService.GetUser(this.authService.GetIdentityId()).subscribe({
      next: (response) => {
      if(response.requestResult.code != 200){
      }
      this.id = response.user.identityId
      this.fetchReservations();
      },
    error: () => {
    }
    });

  }

  fetchReservations(): void {
    this.resetvationService.GetAllGuestPendingReservations(this.id).subscribe({
    next: (response) => {

    if (response.hasOwnProperty('reservations')) { // Check if 'reservations' property exists
      this.dataSource = new MatTableDataSource<Reservation>(response.reservations);
    } else {
    // Handle error if 'reservations' property is missing
    }

    },
    error: () => {
    // Handle error if needed
    }
    });
  }


  cancelReservation(element:any){
    this.resetvationService.RejectReservation(element.id).subscribe(res=>{
      setTimeout(()=>{
        window.location.reload();
      }, 100);
    })
  }
}
