import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PendingGuestReservationsComponent } from './pending-guest-reservations.component';

describe('PendingGuestReservationsComponent', () => {
  let component: PendingGuestReservationsComponent;
  let fixture: ComponentFixture<PendingGuestReservationsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PendingGuestReservationsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PendingGuestReservationsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
