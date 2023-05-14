import { Component } from '@angular/core';
import { AccomodationService } from 'src/app/authentification/services/accommodation.service';
import { AccommodationFilterOffer } from './model/filterOffer.model';
import { Accommodation } from '../../create-accommodation/model/accommodation.model';
import { MatTableDataSource } from '@angular/material/table';
import { AccommodationTemp } from '../../create-accommodation/model/accommodation.temp.model';

@Component({
  selector: 'app-filter-accommodation-offers',
  templateUrl: './filter-accommodation-offers.component.html',
  styleUrls: ['./filter-accommodation-offers.component.scss']
})
export class FilterAccommodationOffersComponent {

  visible: boolean = false;
  filter: AccommodationFilterOffer = new AccommodationFilterOffer()
  displayedColumns: string[] = ['name', 'wifi', 'kitchen', 'air_conditioner', 'parking', 'min_number_of_guests', 'max_number_of_guests'];
  dataSource: MatTableDataSource<AccommodationTemp> = new MatTableDataSource<AccommodationTemp>();
  startDate: Date = new Date();
  endDate: Date = new Date();
  constructor(private accommodationService: AccomodationService) {}

  CreateOffer() : void {
    const temps = this.startDate.toISOString().slice(0, 10);
    const tempe = this.endDate.toISOString().slice(0, 10);
    this.filter.start_date_time_utc = temps + "T00:00:00.000Z";
    this.filter.end_date_time_utc = tempe + "T00:00:00.000Z";
    this.accommodationService.Filter(this.filter).subscribe({
      next: (response) => {
        console.log(response.filteredAccommodations);
        if (response.hasOwnProperty('filteredAccommodations')) { // Check if 'reservations' property exists
          this.dataSource = new MatTableDataSource<AccommodationTemp>(response.filteredAccommodations);
        } else {
          // Handle error if 'reservations' property is missing
        }
       
      },
      error: () => {
        // Handle error if needed
      }
    });
}
}
