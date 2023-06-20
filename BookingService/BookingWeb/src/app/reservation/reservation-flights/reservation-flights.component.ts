import { Component } from '@angular/core';
import { Reservation } from '../contracts/reservation.model';
import { MatTableDataSource } from '@angular/material/table';
import { ReservationsService } from '../services/reservations.service';
import { Router } from '@angular/router';
import { AuthentificationService } from 'src/app/authentification/services/authentification.service';
import { UserService } from 'src/app/authentification/services/user.service';
import { ToastrService } from 'ngx-toastr';
import { FlightService } from '../services/flight.service';
import { FlightFilter } from '../contracts/FlightFilter';
import { FlightContract } from '../contracts/FlightContract';
import { TicketPurchase } from '../contracts/TicketPurchaseContract';

@Component({
  selector: 'app-reservation-flights',
  templateUrl: './reservation-flights.component.html',
  styleUrls: ['./reservation-flights.component.scss']
})
export class ReservationFlightsComponent {

  id : string = ""
  useremail : string = ""
  departure: string = ""
  destination: string = ""
  selectedElement: any | null = null;
  today : Date = new Date();
  location: string = ""
  selectedFlightOne: any | null = null;
  selectedFlightTwo: any | null = null;

  dataSource: MatTableDataSource<Reservation> = new MatTableDataSource<Reservation>();
  flights1: MatTableDataSource<FlightContract> = new MatTableDataSource<FlightContract>();
  flights2: MatTableDataSource<FlightContract> = new MatTableDataSource<FlightContract>();

  displayedColumns: string[] = ['offer', 'customer', 'number', 'start', 'end'];
  columns: string[] = ['start date', 'end date', 'departure', 'destination', 'price'];
  constructor(private reservationService: ReservationsService,
              private router: Router,
              private authService: AuthentificationService,
              private userService: UserService,
              private flightService: FlightService,
              private toastr: ToastrService,){}

  ngOnInit(): void {
    this.authService.GetIdentityId()

    this.userService.GetUser(this.authService.GetIdentityId()).subscribe({
      next: (response) => {
        if(response.requestResult.code != 200){
        }
        this.useremail = response.user.email
        this.id = response.user.identityId
        console.log(this.id)
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

  isSelected(element: any): boolean {
    return this.selectedElement === element;
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
      
      this.reservationService.CancelReservation(element.id).subscribe(()=>{
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

  selectRow(element : any): void {
    this.selectedElement = element;
    this.toastr.success("Reservation chosen successfully")
    this.reservationService.GetOfferLocation(element.accommodationOfferId).subscribe(res=> {this.location = res.location;
      let temp1 = this.selectedElement.startDateTimeUtc.split(" ")[0]
      let temp2 = this.selectedElement.endDateTimeUtc.split(" ")[0]
      let filter = new FlightFilter(temp1,this.departure, this.location,this.selectedElement.numberOfGuests)
      this.flightService.GetFilteredDepartures(filter).subscribe(res => {
        this.flights1 = new MatTableDataSource<FlightContract>(res);
      })

      let filter2 = new FlightFilter(temp2,this.location, this.destination,this.selectedElement.numberOfGuests)
      this.flightService.GetFilteredDestinations(filter2).subscribe(res => {
        this.flights2 = new MatTableDataSource<FlightContract>(res);
      })
    });
  }

  selectFlightOne(flight : any) : void{
    this.selectedFlightOne = flight;
    this.toastr.success("Departure flight chosen successfully")
  }

  selectFlightTwo(flight : any) : void{
    this.selectedFlightTwo = flight;
    this.toastr.success("Destination flight chosen successfully")
  }

  purchase() : void {
    this.flightService.UserExists(this.useremail).subscribe(res => {
      if(!res){
        this.toastr.error("Cannot purchase tickets, user with this email does not exist in ticketing service")
      }
      else{
        let purchase1 = new TicketPurchase(this.id, this.selectedFlightOne.Id, this.selectedElement.numberOfGuests, this.selectedFlightOne.PriceOfTicket)
        let purchase2 = new TicketPurchase(this.id, this.selectedFlightTwo.Id, this.selectedElement.numberOfGuests, this.selectedFlightTwo.PriceOfTicket)
        this.flightService.Purchase(purchase1).subscribe(res => {
          if(res.requestResult.code != 200){
            this.toastr.error("Error while purchasing departure flights")
          }
          else{
            this.toastr.success("Departure flights purchased successfully")
          }
        });
        this.flightService.Purchase(purchase2).subscribe(res => {
          if(res.requestResult.code != 200){
            this.toastr.error("Error while purchasing destination flights")
          }
          else{
            this.toastr.success("Destination flights purchased successfully")
          }
        });
      }
    })
  }

}
