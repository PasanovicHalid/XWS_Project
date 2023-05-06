import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { AngularMaterialModule } from '../library-modules/angular-material.module';
import { RegisterComponent } from './register/register.component';
import { InitialUserInfoEntryComponent } from './initial-user-info-entry/initial-user-info-entry.component';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { UserPanelComponent } from './user-panel/user-panel.component';
import { UserInfoComponent } from './user-info/user-info.component';
import { ChangeUsernameComponent } from './change-username/change-username.component';
import { ChangePasswordComponent } from './change-password/change-password.component';



@NgModule({
  declarations: [
    LoginComponent,
    RegisterComponent,
    InitialUserInfoEntryComponent,
    UserPanelComponent,
    UserInfoComponent,
    ChangeUsernameComponent,
    ChangePasswordComponent,
  ],
  imports: [
    CommonModule,
    AngularMaterialModule,
    FormsModule,
    HttpClientModule,
  ],
})
export class AuthentificationModule { }
