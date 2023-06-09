import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Flight } from '../model/flight.model';
import { FlightService } from '../service/flight.service';

@Component({
  selector: 'app-flights',
  templateUrl: './flights.component.html',
  styleUrls: ['./flights.component.css']
})
export class FlightsComponent  implements OnInit{
  
  flights:Flight[] = []

  constructor(private router: Router,
    private flightService: FlightService) {
      
  }

  ngOnInit(): void {
    this.flightService.getFlights().subscribe(res=>{
      this.flights = res;
      console.log(this.flights)
    })
  }
  delete(id: any){
    this.flightService.deleteFlight(id).subscribe(res=>{
      console.log("Flight deleted")
      setTimeout(()=>{
        window.location.reload();
      }, 100);
    })
  }

  logout(): void {
    localStorage.removeItem("jwt")
      localStorage.removeItem("userId")
      localStorage.removeItem("userRole")
      localStorage.removeItem("user")
      localStorage.removeItem("userFirstName")
      localStorage.removeItem("userLastName")
    // add any other local storage keys you want to remove here
  }
}
