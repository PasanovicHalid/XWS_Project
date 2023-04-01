import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FlightsComponent } from './flights/flights.component';
import { RouterModule, Routes } from '@angular/router';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { CreateFlightsComponent } from './create-flights/create-flights.component';
import { FormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatTableModule } from '@angular/material/table';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatFormFieldModule } from "@angular/material/form-field";
import { MatGridListModule } from '@angular/material/grid-list'; 
import { MatListModule } from '@angular/material/list';
import { MatDatepickerModule } from '@angular/material/datepicker'; // import MatDatepickerModule
import { MatNativeDateModule, MatOptionModule } from '@angular/material/core';
import { SearchFlightsComponent } from './search-flights/search-flights.component';
import { AdminGuard } from '../authentification/services/auth.guard';

const routes: Routes = [
  { path: 'flights', component: FlightsComponent, canActivate: [AdminGuard] },
  { path: 'create-flight', component: CreateFlightsComponent,  canActivate: [AdminGuard] },
];

@NgModule({
  declarations: [
    FlightsComponent,
    CreateFlightsComponent,
    SearchFlightsComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    BrowserAnimationsModule,
    MatOptionModule,
    MatCardModule,
    MatIconModule,
    FormsModule,
    MatInputModule,
    MatButtonModule,
    MatFormFieldModule,
    MatGridListModule,
    MatListModule,
    MatDatepickerModule,
    MatNativeDateModule,
  ]
})
export class FlightModule { }
