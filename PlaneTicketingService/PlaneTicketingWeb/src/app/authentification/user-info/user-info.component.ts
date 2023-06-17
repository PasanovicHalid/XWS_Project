import { Component, OnInit } from '@angular/core';
import { AuthentificationService } from '../services/authentification.service';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-user-info',
  templateUrl: './user-info.component.html',
  styleUrls: ['./user-info.component.css']
})
export class UserInfoComponent implements OnInit { 
  userInfo : any
  apiKeyForever: boolean = false;

  constructor(
    private authService: AuthentificationService,
    private toastr: ToastrService,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.loadUserInfo();
  }

  private loadUserInfo() {
    this.authService.GetUserInfo().subscribe({
      next: (userInfo) => {
        this.userInfo = userInfo;
      },
      error: (e) => {
        this.toastr.error("Something failed with getting user info");
      }
    });
  }

  GenerateApiKey() {
    let duration = new Date(this.userInfo.identity.apiKeyDuration);
    this.authService.GenerateApiKey(duration, this.apiKeyForever).subscribe({
      next: (apiKey) => {
        this.toastr.success("Api key generated successfully")
        this.loadUserInfo();
      },
      error: (e) => {
        this.toastr.error("Something failed with generating api key")
      }
    })
  }

}
