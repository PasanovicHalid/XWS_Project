import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateReservationComponent } from './create-reservation/create-reservation.component';
import { AngularMaterialModule } from '../library-modules/angular-material.module';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';


@NgModule({
  declarations: [
    CreateReservationComponent
  ],
  imports: [
    CommonModule,
    AngularMaterialModule,
    FormsModule,
    HttpClientModule,
  ]
})
export class ReservationModule { }
