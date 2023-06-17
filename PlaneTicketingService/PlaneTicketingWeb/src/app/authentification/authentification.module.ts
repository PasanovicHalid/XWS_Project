import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { RouterModule, Routes } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { SignUpComponent } from './sign-up/sign-up.component';
import { UserInfoComponent } from './user-info/user-info.component';

@NgModule({
  declarations: [
    LoginComponent,
    SignUpComponent,
    UserInfoComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
  ]
})
export class AuthentificationModule { }
