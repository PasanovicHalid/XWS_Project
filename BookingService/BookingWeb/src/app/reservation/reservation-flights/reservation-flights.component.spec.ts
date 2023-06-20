import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReservationFlightsComponent } from './reservation-flights.component';

describe('ReservationFlightsComponent', () => {
  let component: ReservationFlightsComponent;
  let fixture: ComponentFixture<ReservationFlightsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReservationFlightsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReservationFlightsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
