import { Component } from '@angular/core';
import { MatDatepickerInputEvent } from '@angular/material/datepicker';
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';
import { AccomodationService } from 'src/app/authentification/services/accommodation.service';
import { CreateOfferRequest } from './model/accommodationOffer.model';

@Component({
  selector: 'app-create-accommodation-offer',
  templateUrl: './create-accommodation-offer.component.html',
  styleUrls: ['./create-accommodation-offer.component.scss']
})
export class CreateAccommodationOfferComponent {

  newOffer : CreateOfferRequest = new CreateOfferRequest()
  startDate: Date = new Date();
  endDate: Date = new Date();
  automatic:boolean = false;
  constructor(private accommodationService: AccomodationService) {}

  CreateOffer() : void {
    this.startDate.setHours(12);
    this.endDate.setHours(12);
    const temps = this.startDate.toISOString().slice(0, 10);
    console.log(temps)
    console.log(this.startDate)
    const tempe = this.endDate.toISOString().slice(0, 10);
    console.log(tempe)
    console.log(this.endDate)
    this.newOffer.start_date_time_utc = temps + "T12:00:00.000Z";
    this.newOffer.end_date_time_utc = tempe + "T12:00:00.000Z";
    this.accommodationService.CreateOffer(this.newOffer).subscribe()
  }
  
}
