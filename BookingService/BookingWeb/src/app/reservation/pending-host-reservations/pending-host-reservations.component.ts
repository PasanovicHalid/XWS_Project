import { Component, OnInit } from '@angular/core';
import { CreateReservationRequest } from '../contracts/create-reservation-request.model';
import { ReservationsService } from '../services/reservations.service';
import { Router } from '@angular/router';
import { AuthentificationService } from 'src/app/authentification/services/authentification.service';
import { UserService } from 'src/app/authentification/services/user.service';
import { MatTableDataSource } from '@angular/material/table';

@Component({
  selector: 'app-pending-host-reservations',
  templateUrl: './pending-host-reservations.component.html',
  styleUrls: ['./pending-host-reservations.component.scss']
})
export class PendingHostReservationsComponent implements OnInit{

  id:string = "";
  dataSource: MatTableDataSource<CreateReservationRequest> = new MatTableDataSource<CreateReservationRequest>();

  displayedColumns: string[] = ['offer', 'customer', 'number', 'start', 'end', 'accept','reject'];
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
    this.resetvationService.GetAllHostPendingReservations(this.id).subscribe({
      next: (response) => {
        
          
       
      },
      error: () => {
        // Handle error if needed
      }
    });
  }

  acceptReservation(element:any){
    this.resetvationService.AcceptReservation(element.Id).subscribe(res=>{
      setTimeout(()=>{
        window.location.reload();
      }, 100);
    })
  }

  rejectReservation(element:any){
    this.resetvationService.RejectReservation(element.Id).subscribe(res=>{
      setTimeout(()=>{
        window.location.reload();
      }, 100);
    })
  }
}
