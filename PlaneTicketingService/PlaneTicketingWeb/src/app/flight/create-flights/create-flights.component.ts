import { Component } from '@angular/core';
import { FormControl } from '@angular/forms';
import { FloatLabelType } from '@angular/material/form-field';
import { Router } from '@angular/router';
import { Flight } from '../model/flight.model';
import { FlightService } from '../service/flight.service';


@Component({
  selector: 'app-create-flights',
  templateUrl: './create-flights.component.html',
  styleUrls: ['./create-flights.component.css']
})
export class CreateFlightsComponent {

  flight:Flight = new Flight();

  floatLabelControl = new FormControl('auto' as FloatLabelType);

  formattedDate :String ='';
  constructor(private router: Router,
    private flightService: FlightService) {
      
  }
  getFloatLabelValue(): FloatLabelType {
    return this.floatLabelControl.value || 'auto';
  }

  create(){
    //this.formattedDate = this.dateAsYYYYMMDDHHNNSS(this.flight.Start);
    
    console.log(this.flight)
    this.flight.Start = this.convertDate(new Date());
    this.flight.End = this.convertDate(new Date());
    this.flightService.createFlight(this.flight).subscribe(res=>{
      console.log("ok ok")
    })
  }

  dateAsYYYYMMDDHHNNSS(date:any): string {
    return date.getFullYear()
              + '-' + this.leftpad(date.getMonth() + 1)
              + '-' + this.leftpad(date.getDate())
              + 'T' + this.leftpad(date.getHours())
              + ':' + this.leftpad(date.getMinutes())
              + ':' + this.leftpad(date.getSeconds());
  }
  leftpad(val:any, resultLength = 2, leftpadChar = '0'): string {
    return (String(leftpadChar).repeat(resultLength)
          + String(val)).slice(String(val).length);
  }

  convertDate(date: any):string{
    const year = "2023";
    const month = ("0" + (date.getMonth() + 1)).slice(-2);
    const day = ("0" + date.getDate()).slice(-2);
    const hours = ("0" + date.getHours()).slice(-2);
    const minutes = ("0" + date.getMinutes()).slice(-2);
    const seconds = ("0" + date.getSeconds()).slice(-2);
    
    const backendDateString = `${year}-${month}-${day}T${hours}:${minutes}:${seconds}Z`;
    return backendDateString;
  }
}
