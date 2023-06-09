import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateReservationComponent } from './create-reservation/create-reservation.component';
import { AngularMaterialModule } from '../library-modules/angular-material.module';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { PendingHostReservationsComponent } from './pending-host-reservations/pending-host-reservations.component';
import { PendingGuestReservationsComponent } from './pending-guest-reservations/pending-guest-reservations.component';
import { CancelReservationComponent } from './cancel-reservation/cancel-reservation.component';
import { ReservationFlightsComponent } from './reservation-flights/reservation-flights.component';


@NgModule({
  declarations: [
    CreateReservationComponent,
    PendingHostReservationsComponent,
    ReservationFlightsComponent,
    PendingGuestReservationsComponent,
    CancelReservationComponent,
    ReservationFlightsComponent,
  ],
  imports: [
    CommonModule,
    AngularMaterialModule,
    FormsModule,
    HttpClientModule,
  ]
})
export class ReservationModule { }
