import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { SignUpComponent } from './sign-up/sign-up.component';
import { AngularMaterialModule } from '../library-modules/angular-material.module';
import { RegisterComponent } from './register/register.component';
import { InitialUserInfoEntryComponent } from './initial-user-info-entry/initial-user-info-entry.component';
import { AuthentificationService } from './services/authentification.service';
import { UserService } from './services/user.service';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';



@NgModule({
  declarations: [
    LoginComponent,
    SignUpComponent,
    RegisterComponent,
    InitialUserInfoEntryComponent
  ],
  imports: [
    CommonModule,
    AngularMaterialModule,
    FormsModule,
    HttpClientModule,
  ],
})
export class AuthentificationModule { }
