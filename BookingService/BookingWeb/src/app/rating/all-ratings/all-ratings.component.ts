import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ToastrService } from 'ngx-toastr';
import { AuthentificationService } from 'src/app/authentification/services/authentification.service';
import { RateDialogComponent, RatingType } from '../rate-dialog/rate-dialog.component';
import { RatingService } from '../services/rating.service';

@Component({
  selector: 'app-all-ratings',
  templateUrl: './all-ratings.component.html',
  styleUrls: ['./all-ratings.component.scss']
})
export class AllRatingsComponent {
  displayedColumns: string[] = ['userId', 'hostId', 'accommodationId', 'rating', 'timeIssued'];
  dataSource: any;

  constructor(private dialog: MatDialog,
    private ratingService: RatingService,
    private toastr: ToastrService,
    private authService : AuthentificationService) { }

  ngOnInit(): void {
    this.reloadTable();
  }

  private reloadTable() {
    this.ratingService.GetCustomerRatings(this.authService.GetIdentityId()).subscribe({
      next: (response) => {
        this.dataSource = response.ratings;
      },
      error: () => {
        this.toastr.error("Something went wrong.");
      }
    });
  }

  openDialog(rating: any) {
    const dialogRef = this.dialog.open(RateDialogComponent, {
      width: '500px',
      height: '500px',
      data: { data: {ratingId : rating.id, rating: rating.rating} }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result.type == RatingType.EDIT) {
        this.ratingService.UpdateRating({ id: rating.id, rating: result.rating }).subscribe({
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
        this.ratingService.DeleteRating(rating.id).subscribe({
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
