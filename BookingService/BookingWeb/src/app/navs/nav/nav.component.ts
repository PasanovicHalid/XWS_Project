import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthentificationService } from 'src/app/authentification/services/authentification.service';

@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.scss']
})
export class NavComponent {

  constructor(private authService : AuthentificationService,
              private router: Router) { }

  public Logout() {
    this.authService.Logout()
    this.router.navigate(['/']);
  }
}
