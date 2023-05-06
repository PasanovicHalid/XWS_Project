import { inject } from "@angular/core";
import { Router } from "@angular/router";
import { AuthentificationService } from "../authentification/services/authentification.service";

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