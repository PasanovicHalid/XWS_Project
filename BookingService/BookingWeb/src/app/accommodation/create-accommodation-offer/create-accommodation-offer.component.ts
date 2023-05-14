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
  constructor(private accommodationService: AccomodationService) {}

  CreateOffer() : void {
    const temps = new Date().toISOString().slice(0, 10);
    const tempe = new Date().toISOString().slice(0, 10);
    this.newOffer.start_date_time_utc = temps + "T00:00:00.000Z";
    this.newOffer.end_date_time_utc = tempe + "T00:00:00.000Z";
    this.accommodationService.CreateOffer(this.newOffer).subscribe()
  }
  
}
