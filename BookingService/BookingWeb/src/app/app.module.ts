import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AngularMaterialModule } from './library-modules/angular-material.module';
import { AuthentificationModule } from './authentification/authentification.module';
import { PublicLayoutComponent } from './layouts/public-layout/public-layout.component';
import { PrivateLayoutComponent } from './layouts/private-layout/private-layout.component';
import { DashboardComponent } from './landing-pages/dashboard/dashboard.component';
import { LandingPageComponent } from './landing-pages/landing-page/landing-page.component';
import { NavsModule } from './navs/navs.module';
import { SideNavComponent } from './navs/side-nav/side-nav.component';
import { JwtHelperService, JwtModule } from '@auth0/angular-jwt';
import { ToastrModule } from 'ngx-toastr';
import { fadeAnimation } from './animations/fade';
import { CreateAccommodationComponent } from './accommodation/create-accommodation/create-accommodation.component';
import { FormsModule } from '@angular/forms';
import { CreateAccommodationOfferComponent } from './accommodation/create-accommodation-offer/create-accommodation-offer.component';
import { ReservationModule } from './reservation/reservation.module';
import { FilterAccommodationOffersComponent } from './accommodation/filter-acommodation-offers/filter-accommodation-offers/filter-accommodation-offers.component';
import { UpdateAccommodationOfferComponent } from './accommodation/update-accommodation-offer/update-accommodation-offer/update-accommodation-offer.component';
import { CancelReservationComponent } from './reservation/cancel-reservation/cancel-reservation.component';
import { RatingDashboardComponent } from './rating/rating-dashboard/rating-dashboard.component';
import { HostRatingsComponent } from './rating/host-ratings/host-ratings.component';
import { AccommodationRatingsComponent } from './rating/accommodation-ratings/accommodation-ratings.component';
import { AllRatingsComponent } from './rating/all-ratings/all-ratings.component';
import { RateDialogComponent } from './rating/rate-dialog/rate-dialog.component';
import { CommonCodeModule } from './common-code/common-code.module';

export function tokenGetter() {
  return localStorage.getItem("jwt")
}

@NgModule({
  declarations: [
    AppComponent,
    PublicLayoutComponent,
    PrivateLayoutComponent,
    DashboardComponent,
    LandingPageComponent,
    CreateAccommodationComponent,
    CreateAccommodationOfferComponent,
    FilterAccommodationOffersComponent,
    UpdateAccommodationOfferComponent,
    RatingDashboardComponent,
    HostRatingsComponent,
    AccommodationRatingsComponent,
    AllRatingsComponent,
    RateDialogComponent,

  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    AngularMaterialModule,
    AuthentificationModule,
    FormsModule,
    NavsModule,
    CommonCodeModule,
    ReservationModule,
    JwtModule.forRoot({
      config: {
        tokenGetter: tokenGetter,
      }
    }),
    ToastrModule.forRoot(),

  ],
  providers: [JwtHelperService],
  bootstrap: [AppComponent]
})
export class AppModule { }
