import { trigger, transition, query, style, animate } from '@angular/animations';
import { Component } from '@angular/core';
import { fadeAnimation } from 'src/app/animations/fade';

@Component({
  selector: 'app-private-layout',
  templateUrl: './private-layout.component.html',
  styleUrls: ['./private-layout.component.scss'],
  animations: [
    fadeAnimation
  ],
})
export class PrivateLayoutComponent {

}
