import { Component } from '@angular/core';

@Component({
  selector: 'app-admin-homepage',
  templateUrl: './admin-homepage.component.html',
  styleUrls: ['./admin-homepage.component.css']
})
export class AdminHomepageComponent {

  logout(): void {
    localStorage.removeItem("jwt")
      localStorage.removeItem("userId")
      localStorage.removeItem("userRole")
      localStorage.removeItem("user")
      localStorage.removeItem("userFirstName")
      localStorage.removeItem("userLastName")
    // add any other local storage keys you want to remove here
  }
}
