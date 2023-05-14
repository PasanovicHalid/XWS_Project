import { Component, OnInit } from '@angular/core';
import { Accommodation } from './model/accommodation.model';
import { AccomodationService } from 'src/app/authentification/services/accommodation.service';

@Component({
  selector: 'app-create-accommodation',
  templateUrl: './create-accommodation.component.html',
  styleUrls: ['./create-accommodation.component.scss']
})
export class CreateAccommodationComponent {
  temp : string = ""
  newAccommodation: Accommodation = new Accommodation()

  constructor(private accommodationService: AccomodationService) {}

  CreateAccommodation() : void {
    const identityId = localStorage.getItem("identityId");
    if (identityId) {
      this.newAccommodation.images = this.temp.split(",")
      this.newAccommodation.ownerId = identityId;
      console.log(this.newAccommodation)
      this.accommodationService.CreateAccommodation(this.newAccommodation).subscribe()
    }
  }
}
