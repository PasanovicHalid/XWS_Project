import { Component, OnInit } from '@angular/core';
import { FlightService } from '../service/flight.service';
import { City } from '../model/cityResponse.model';
import { FlightFilter } from '../model/flightFilterRequest.model';
import { Flight } from '../model/flight.model';

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

  constructor(private service : FlightService ){}
  ngOnInit() : void{
     this.service.getCities().subscribe(res => {
      this.cities = res;
    });
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
  
}
