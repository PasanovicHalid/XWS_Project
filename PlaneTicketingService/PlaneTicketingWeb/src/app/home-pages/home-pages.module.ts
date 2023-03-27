import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AdminHomepageComponent } from './admin-homepage/admin-homepage.component';
import { CustomerHomepageComponent } from './customer-homepage/customer-homepage.component';
import { RouterModule, Routes } from '@angular/router';
import { DefaultHomepageComponent } from './default-homepage/default-homepage.component';

const routes: Routes = [
  { path: 'admin-homepage', component: AdminHomepageComponent },
  { path: 'customer-homepage', component: CustomerHomepageComponent },
  { path: '', component: DefaultHomepageComponent },
];

@NgModule({
  declarations: [
    AdminHomepageComponent,
    CustomerHomepageComponent,
    DefaultHomepageComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
  ]
})
export class HomePagesModule { }
