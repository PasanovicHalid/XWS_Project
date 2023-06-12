import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';

export interface Data {
  data : any;
}

export interface Output {
  rating: number;
  type: RatingType;
}

export enum RatingType {
  ADD = 0,
  EDIT = 1,
  DELETE = 2
}

@Component({
  selector: 'app-rate-dialog',
  templateUrl: './rate-dialog.component.html',
  styleUrls: ['./rate-dialog.component.scss']
})
export class RateDialogComponent {
  output: Output = {rating: 0, type: RatingType.EDIT};

  constructor(
    public dialogRef: MatDialogRef<RateDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public inputData: Data,
  ) {
    if (inputData.data.ratingId == "") {
      this.output.type = RatingType.ADD;
    } else {
      this.output.type = RatingType.EDIT;
      this.output.rating = inputData.data.rating;
    }
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  onOkClick(): void {
    this.dialogRef.close(this.output);
  }

  onDeleteClick(): void {
    this.output.type = RatingType.DELETE;
    this.dialogRef.close(this.output);
  }
}
