import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ToastrService } from 'ngx-toastr';
import { AuthentificationService } from 'src/app/authentification/services/authentification.service';
import { RateDialogComponent, RatingType } from '../rate-dialog/rate-dialog.component';
import { RatingService } from '../services/rating.service';

@Component({
  selector: 'app-host-ratings',
  templateUrl: './host-ratings.component.html',
  styleUrls: ['./host-ratings.component.scss']
})
export class HostRatingsComponent {
  displayedColumns: string[] = ['name', 'averageRating', 'rating'];
  dataSource: any;

  constructor(private dialog: MatDialog,
    private ratingService: RatingService,
    private toastr: ToastrService,
    private authService : AuthentificationService) { }

  ngOnInit(): void {
    this.reloadTable();
  }

  private reloadTable() {
    this.ratingService.GetHostsForRating().subscribe({
      next: (response) => {
        this.dataSource = response.hosts;
      },
      error: () => {
        this.toastr.error("Something went wrong.");
      }
    });
  }

  openDialog(host: any) {
    const dialogRef = this.dialog.open(RateDialogComponent, {
      width: '500px',
      height: '500px',
      data: { data: host }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result.type == RatingType.ADD) {
        this.ratingService.CreateHostRating({ userId: this.authService.GetIdentityId(), rating: result.rating }, host.id).subscribe({
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
        this.ratingService.UpdateRating({ id: host.ratingId, rating: result.rating }).subscribe({
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
        this.ratingService.DeleteRating(host.ratingId).subscribe({
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
