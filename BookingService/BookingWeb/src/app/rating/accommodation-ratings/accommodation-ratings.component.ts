import { Component, OnInit } from '@angular/core';
import { RateDialogComponent, RatingType } from '../rate-dialog/rate-dialog.component';
import { MatDialog } from '@angular/material/dialog';
import { RatingService } from '../services/rating.service';
import { ToastrService } from 'ngx-toastr';
import { AuthentificationService } from 'src/app/authentification/services/authentification.service';

@Component({
  selector: 'app-accommodation-ratings',
  templateUrl: './accommodation-ratings.component.html',
  styleUrls: ['./accommodation-ratings.component.scss']
})
export class AccommodationRatingsComponent implements OnInit {
  displayedColumns: string[] = ['name', 'address', 'averageRating', 'rating'];
  dataSource: any;

  constructor(private dialog: MatDialog,
    private ratingService: RatingService,
    private toastr: ToastrService,
    private authService : AuthentificationService) { }

  ngOnInit(): void {
    this.reloadTable();
  }

  private reloadTable() {
    this.ratingService.GetAccommodationsForRating().subscribe({
      next: (response) => {
        this.dataSource = response.accommodations;
      },
      error: () => {
        this.toastr.error("Something went wrong.");
      }
    });
  }

  openDialog(accommodation: any) {
    const dialogRef = this.dialog.open(RateDialogComponent, {
      width: '500px',
      height: '500px',
      data: { data: accommodation }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result.type == RatingType.ADD) {
        this.ratingService.CreateAccommodationRating({ userId: this.authService.GetIdentityId(), rating: result.rating }, accommodation.id).subscribe({
          next: (response) => {
            if (response.code != 200) {
              this.toastr.error(response.message)
              return
            }
            this.toastr.success("Successfully added rating")
            this.reloadTable();
          },
          error: () => {
            this.toastr.error("Something went wrong.");
          }
        });
      } else if (result.type == RatingType.EDIT) {
        this.ratingService.UpdateRating({ id: accommodation.ratingId, rating: result.rating }).subscribe({
          next: (response) => {
            if (response.code != 200) {
              this.toastr.error(response.message)
              return
            }
            this.toastr.success("Successfully updated rating")
            this.reloadTable();
          },
          error: () => {
            this.toastr.error("Something went wrong.");
          }
        });
      } else if (result.type == RatingType.DELETE) {
        this.ratingService.DeleteRating(accommodation.ratingId).subscribe({
          next: (response) => {
            if (response.code != 200) {
              this.toastr.error(response.message)
              return
            }
            this.toastr.success("Successfully deleted rating")
            this.reloadTable();
          },
          error: () => {
            this.toastr.error("Something went wrong.");
          }
        });
      }
    });
  }
}
