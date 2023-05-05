import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SideNavComponent } from './side-nav/side-nav.component';
import { NavComponent } from './nav/nav.component';
import { AngularMaterialModule } from '../library-modules/angular-material.module';
import { RouterModule } from '@angular/router';



@NgModule({
  declarations: [
    SideNavComponent,
    NavComponent,
  ],
  imports: [
    CommonModule,
    AngularMaterialModule, 
    RouterModule
  ],
  exports: [
    SideNavComponent,
    NavComponent,
  ]
})
export class NavsModule { }
