import { Component } from '@angular/core';
import { fadeAnimation } from 'src/app/animations/fade';

@Component({
  selector: 'app-rating-dashboard',
  templateUrl: './rating-dashboard.component.html',
  styleUrls: ['./rating-dashboard.component.scss'],
  animations: [
    fadeAnimation
  ],
})
export class RatingDashboardComponent {
}
