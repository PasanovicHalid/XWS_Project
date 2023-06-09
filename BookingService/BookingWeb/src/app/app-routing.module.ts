import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { PublicLayoutComponent } from './layouts/public-layout/public-layout.component';
import { PrivateLayoutComponent } from './layouts/private-layout/private-layout.component';
import { LoginComponent } from './authentification/login/login.component';
import { LandingPageComponent } from './landing-pages/landing-page/landing-page.component';
import { DashboardComponent } from './landing-pages/dashboard/dashboard.component';
import { RegisterComponent } from './authentification/register/register.component';
import { InitialUserInfoEntryComponent } from './authentification/initial-user-info-entry/initial-user-info-entry.component';
import { UserPanelComponent } from './authentification/user-panel/user-panel.component';
import { AuthGuard } from './guards/auth.gard';
import { CreateAccommodationComponent } from './accommodation/create-accommodation/create-accommodation.component';
import { CreateAccommodationOfferComponent } from './accommodation/create-accommodation-offer/create-accommodation-offer.component';
import { CreateReservationComponent } from './reservation/create-reservation/create-reservation.component';
import { FilterAccommodationOffersComponent } from './accommodation/filter-acommodation-offers/filter-accommodation-offers/filter-accommodation-offers.component';
import { PendingHostReservationsComponent } from './reservation/pending-host-reservations/pending-host-reservations.component';
import { UpdateAccommodationOfferComponent } from './accommodation/update-accommodation-offer/update-accommodation-offer/update-accommodation-offer.component';
import { PendingGuestReservationsComponent } from './reservation/pending-guest-reservations/pending-guest-reservations.component';
import { CancelReservationComponent } from './reservation/cancel-reservation/cancel-reservation.component';
import { RatingDashboardComponent } from './rating/rating-dashboard/rating-dashboard.component';
import { HostRatingsComponent } from './rating/host-ratings/host-ratings.component';
import { AccommodationRatingsComponent } from './rating/accommodation-ratings/accommodation-ratings.component';
import { AllRatingsComponent } from './rating/all-ratings/all-ratings.component';
import { ReservationFlightsComponent } from './reservation/reservation-flights/reservation-flights.component';


const routes: Routes = [
  {
    path: '',
    component: PublicLayoutComponent,
    children: [
      { path: '', component: LandingPageComponent },
      { path: 'login', component: LoginComponent },
      { path: 'register', component: RegisterComponent },
      { path: 'initial-user-info-entry', component: InitialUserInfoEntryComponent },
      { path: 'filter-offers', component: FilterAccommodationOffersComponent},
    ]
  },
  {
    path: '',
    component: PrivateLayoutComponent,
    children: [
      { path: 'dashboard', component: DashboardComponent, canActivate: [AuthGuard] },
      { path: 'user-panel', component: UserPanelComponent, canActivate: [AuthGuard] },
      { path: 'create-accommodation', component: CreateAccommodationComponent, canActivate: [AuthGuard]},
      { path: 'create-accommodation-offer', component: CreateAccommodationOfferComponent, canActivate: [AuthGuard]},
      { path: 'update-accommodation-offer', component: UpdateAccommodationOfferComponent, canActivate: [AuthGuard]},
      { path: 'create-reservation/:id', component: CreateReservationComponent, canActivate: [AuthGuard] },
      { path: 'filter-accommodation-offers', component: FilterAccommodationOffersComponent, canActivate: [AuthGuard] },
      { path: 'pending-host-reservation', component: PendingHostReservationsComponent, canActivate: [AuthGuard] },
      { path: 'pending-guest-reservation', component: PendingGuestReservationsComponent, canActivate: [AuthGuard] },
      { path: 'reservation-flights', component: ReservationFlightsComponent, canActivate: [AuthGuard]},
      { path: 'cancel-reservation', component: CancelReservationComponent, canActivate: [AuthGuard] },
      { 
        path: 'rating-dashboard', 
        component: RatingDashboardComponent,
        children: [
          { path: 'hosts', component: HostRatingsComponent },
          { path: 'accommodations', component: AccommodationRatingsComponent },
          { path: 'all', component: AllRatingsComponent },
        ], 
        canActivate: [AuthGuard]
      },
    ]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
