import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { EmptyPipe } from './pipes/empty-pipe';
import { FirstLetterLowercasePipe } from './pipes/first-letter-lowercase-pipe';
import { RatingZeroPipe } from './pipes/rating-zero';



@NgModule({
  declarations: [
    EmptyPipe,
    FirstLetterLowercasePipe,
    RatingZeroPipe
  ],
  imports: [
    CommonModule
  ],
  exports: [
    EmptyPipe,
    FirstLetterLowercasePipe,
    RatingZeroPipe
  ]
})
export class CommonCodeModule { }
