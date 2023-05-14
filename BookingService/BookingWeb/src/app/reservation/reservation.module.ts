import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateReservationComponent } from './create-reservation/create-reservation.component';
import { AngularMaterialModule } from '../library-modules/angular-material.module';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { PendingHostReservationsComponent } from './pending-host-reservations/pending-host-reservations.component';
import { CancelReservationComponent } from './cancel-reservation/cancel-reservation.component';


@NgModule({
  declarations: [
    CreateReservationComponent,
    PendingHostReservationsComponent,
    CancelReservationComponent,
  ],
  imports: [
    CommonModule,
    AngularMaterialModule,
    FormsModule,
    HttpClientModule,
  ]
})
export class ReservationModule { }
