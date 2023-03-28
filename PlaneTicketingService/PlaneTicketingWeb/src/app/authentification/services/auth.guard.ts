import { inject, Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthentificationService } from './authentification.service';

export const AuthGuard = () =>  {
  const router = inject(Router);
  const authService = inject(AuthentificationService)
  let allowed = authService.IsLoggedIn();
  if(!allowed){
    router.navigate(['/login'])
    return false;
  }
  return true
}

export const AdminGuard = () =>  {
  const router = inject(Router);
  const authService = inject(AuthentificationService)
  let allowed = authService.IsLoggedIn() && authService.IsAdmin();
  if(!allowed){
    router.navigate(['/login'])
    return false;
  }
  return true
}

export const CustomerGuard = () =>  {
  const router = inject(Router);
  const authService = inject(AuthentificationService)
  let allowed = authService.IsLoggedIn() && authService.IsCustomer();
  if(!allowed){
    router.navigate(['/login'])
    return false;
  }
  return true 
}
