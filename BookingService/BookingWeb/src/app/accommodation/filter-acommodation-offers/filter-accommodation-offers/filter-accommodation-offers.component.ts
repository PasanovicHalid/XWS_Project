import { Component } from '@angular/core';
import { AccomodationService } from 'src/app/authentification/services/accommodation.service';
import { AccommodationFilterOffer } from './model/filterOffer.model';
import { Accommodation } from '../../create-accommodation/model/accommodation.model';
import { MatTableDataSource } from '@angular/material/table';
import { AccommodationTemp } from '../../create-accommodation/model/accommodation.temp.model';
import { Router } from '@angular/router';
import { ThemePalette } from '@angular/material/core';

export interface Task {
  name: string;
  completed: boolean;
  color: ThemePalette;
  subtasks?: Task[];
}

@Component({
  selector: 'app-filter-accommodation-offers',
  templateUrl: './filter-accommodation-offers.component.html',
  styleUrls: ['./filter-accommodation-offers.component.scss']
})
export class FilterAccommodationOffersComponent {
  task: Task = {
    name: 'All Benefits',
    completed: false,
    color: 'primary',
    subtasks: [
      {name: 'Wifi', completed: false, color: 'primary'},
      {name: 'Kitchen', completed: false, color: 'primary'},
      {name: 'Air-Conditioning', completed: false, color: 'primary'},
      {name: 'Parking', completed: false, color: 'primary'},
    ],
  };

  visible: boolean = false;
  filter: AccommodationFilterOffer = new AccommodationFilterOffer()
  displayedColumns: string[] = ['name', 'wifi', 'kitchen', 'air_conditioner', 'parking', 'min_number_of_guests', 'max_number_of_guests', 'make'];
  dataSource: MatTableDataSource<AccommodationTemp> = new MatTableDataSource<AccommodationTemp>();
  startDate: Date = new Date();
  endDate: Date = new Date();
  constructor(private accommodationService: AccomodationService, private router:Router) {}

  

  allComplete: boolean = false;

  updateAllComplete() {
    this.allComplete = this.task.subtasks != null && this.task.subtasks.every(t => t.completed);
  }

  someComplete(): boolean {
    if (this.task.subtasks == null) {
      return false;
    }
    return this.task.subtasks.filter(t => t.completed).length > 0 && !this.allComplete;
  }

  setAll(completed: boolean) {
    this.allComplete = completed;
    if (this.task.subtasks == null) {
      return;
    }
    this.task.subtasks.forEach(t => (t.completed = completed));
  }

  CreateOffer() : void {
    const temps = this.startDate.toISOString().slice(0, 10);
    const tempe = this.endDate.toISOString().slice(0, 10);
    this.filter.from = this.startDate.toISOString();
    this.filter.to = this.endDate.toISOString();
    this.filter.wifi = this.task.subtasks![0].completed;
    this.filter.kitchen = this.task.subtasks![1].completed;
    this.filter.airConditioner = this.task.subtasks![2].completed;
    this.filter.parking = this.task.subtasks![3].completed;
    if(this.filter.ratingBottom != 0 || this.filter.ratingTop != 0){
      this.filter.filterByRating = true;
    }
    console.log(this.filter);
    this.accommodationService.Filter(this.filter).subscribe({
      next: (response) => {
          this.dataSource = new MatTableDataSource<AccommodationTemp>(response.filtered_accommodations);
      },
      error: () => {
        // Handle error if needed
      }
    });
}

makeReservation(accommodation :any){
  this.router.navigate(['/create-reservation', accommodation.accommodationOfferId])
}
}
