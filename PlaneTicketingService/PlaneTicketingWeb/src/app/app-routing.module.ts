import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './authentification/login/login.component';
import { AdminGuard, CustomerGuard } from './authentification/services/auth.guard';
import { SignUpComponent } from './authentification/sign-up/sign-up.component';
import { AdminHomepageComponent } from './home-pages/admin-homepage/admin-homepage.component';
import { CustomerHomepageComponent } from './home-pages/customer-homepage/customer-homepage.component';
import { DefaultHomepageComponent } from './home-pages/default-homepage/default-homepage.component';
import { SearchFlightsComponent } from './flight/search-flights/search-flights.component';
import { UserInfoComponent } from './authentification/user-info/user-info.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'sign-up', component: SignUpComponent },
  { path: 'search-flights', component: SearchFlightsComponent},
  { path: 'admin-homepage', component: AdminHomepageComponent, canActivate: [AdminGuard] },
  { path: 'customer-homepage', component: CustomerHomepageComponent, canActivate: [CustomerGuard] },
  { path: 'user-info', component: UserInfoComponent, canActivate: [CustomerGuard] },
  { path: '', component: DefaultHomepageComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
