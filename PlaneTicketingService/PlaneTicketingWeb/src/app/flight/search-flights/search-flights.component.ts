import { Component, OnInit } from '@angular/core';
import { FlightService } from '../service/flight.service';
import { City } from '../model/cityResponse.model';
import { FlightFilter } from '../model/flightFilterRequest.model';
import { Flight } from '../model/flight.model';
import { Route, Router } from '@angular/router';

@Component({
  selector: 'app-search-flights',
  templateUrl: './search-flights.component.html',
  styleUrls: ['./search-flights.component.css']
})
export class SearchFlightsComponent implements OnInit{
  
  cities : City[] = [];
  temp!: Date;
  flightFilter : FlightFilter = new FlightFilter();
  flights : Flight[] = [];
  visible : boolean = false;
  visiblePurchase : boolean = false;
  userId:any;
  constructor(private service : FlightService,
    private router: Router ){}
  ngOnInit() : void{
    this.userId = localStorage.getItem("userId");
     this.service.getCities().subscribe(res => {
      this.cities = res;
    });
    console.log(localStorage.getItem("userId"))
    if(this.userId == null){
      this.visiblePurchase=false;
    }else{
      this.visiblePurchase= true;
    }

    document.getElementById("date")?.setAttribute("min",new Date().toISOString().split('T')[0]);

  }

  onSelectDeparture(departure : String) : void{
    this.flightFilter.DepartureLocation = departure;
  }

  onSelectDestination(destination : String) : void{
    this.flightFilter.Destinationlocation = destination;
  }
  onSelectPassengers(number : String) : void{
    this.flightFilter.NumberOfTickets = Number(number);
  }

  filterFlights() : void{
    this.service.filterFlights(this.flightFilter).subscribe(res => {
      if(res == null)
        this.visible = false;
      else{
      this.flights = res;
      if(this.flights.length > 0)
        this.visible = true;
    }
    })
  }

  purchase(id:any){
    this.service.setFlightId(id);
    this.router.navigate(["/ticket-purchase"])
  }
}
