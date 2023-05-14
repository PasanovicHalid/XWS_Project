import { Component } from '@angular/core';
import { AccomodationService } from 'src/app/authentification/services/accommodation.service';
import { AccommodationFilterOffer } from './model/filterOffer.model';

@Component({
  selector: 'app-filter-accommodation-offers',
  templateUrl: './filter-accommodation-offers.component.html',
  styleUrls: ['./filter-accommodation-offers.component.scss']
})
export class FilterAccommodationOffersComponent {

  filter: AccommodationFilterOffer = new AccommodationFilterOffer()
  startDate: Date = new Date();
  endDate: Date = new Date();
  constructor(private accommodationService: AccomodationService) {}

  CreateOffer() : void {
    const temps = this.startDate.toISOString().slice(0, 10);
    const tempe = this.endDate.toISOString().slice(0, 10);
    this.filter.start_date_time_utc = temps + "T00:00:00.000Z";
    this.filter.end_date_time_utc = tempe + "T00:00:00.000Z";
    this.accommodationService.Filter(this.filter).subscribe()
  }
}
