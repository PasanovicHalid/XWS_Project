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
import { PendingHostReservationsComponent } from './reservation/pending-host-reservations/pending-host-reservations.component';

const routes: Routes = [
  {
    path: '',
    component: PublicLayoutComponent,
    children: [
      { path: '', component: LandingPageComponent },
      { path: 'login', component: LoginComponent },
      { path: 'register', component: RegisterComponent },
      { path: 'initial-user-info-entry', component: InitialUserInfoEntryComponent },
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
      { path: 'create-reservation', component: CreateReservationComponent, canActivate: [AuthGuard] },
      { path: 'pending-host-reservation', component: PendingHostReservationsComponent, canActivate: [AuthGuard] },
    ]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
