import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { ReservationsService } from '../services/reservations.service';
import { Router } from '@angular/router';
import { AuthentificationService } from 'src/app/authentification/services/authentification.service';
import { UserService } from 'src/app/authentification/services/user.service';
import { Reservation } from '../contracts/reservation.model';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-cancel-reservation',
  templateUrl: './cancel-reservation.component.html',
  styleUrls: ['./cancel-reservation.component.scss']
})
export class CancelReservationComponent implements OnInit {
  id : string = ""

  today : Date = new Date();
  dataSource: MatTableDataSource<Reservation> = new MatTableDataSource<Reservation>();

  displayedColumns: string[] = ['offer', 'customer', 'number', 'start', 'end', 'cancel'];
  constructor(private reservationService: ReservationsService,
              private router: Router,
              private authService: AuthentificationService,
              private userService: UserService,
              private toastr: ToastrService,){}

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
    this.reservationService.GetAllGuestAcceptedReservations(this.id).subscribe({
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

  validateDates(date : string){
    this.today = new Date();
    this.today.setHours(0);
    this.today.setMinutes(0);
    this.today.setSeconds(0);
    this.today.setMilliseconds(0);
    var timestamp: number = Date.parse(date);
    var dateOfReservation: Date = new Date(timestamp);
    dateOfReservation.setHours(0);
    dateOfReservation.setMinutes(0);
    dateOfReservation.setSeconds(0);
    dateOfReservation.setMilliseconds(0);
  
    if(this.today >= dateOfReservation){
      return false;
      

    }
    else{
      return true;
      
    }
  }

  cancelReservation(element:any){
    console.log(element)
    var prom = this.validateDates(element.startDateTimeUtc);
    if(prom == true){
      
      this.reservationService.CancelReservation(element.id).subscribe(res=>{
        setTimeout(()=>{
          window.location.reload();
        }, 100);
      })
      this.showSuccess("Reservation is successfully canceled.", "Cancellation approved")
    }else{
      this.showError("Reservation cannot be cancelled now.", "Cancelling error");
    }
    
  }


  showSuccess(message: string, title: string) {
    this.toastr.success(message, title);
  }
  showError(message: string, title: string) {
    this.toastr.error(message, title);
  }
}
