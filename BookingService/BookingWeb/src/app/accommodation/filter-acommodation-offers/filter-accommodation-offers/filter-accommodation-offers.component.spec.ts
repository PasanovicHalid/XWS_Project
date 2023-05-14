import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FilterAccommodationOffersComponent } from './filter-accommodation-offers.component';

describe('FilterAccommodationOffersComponent', () => {
  let component: FilterAccommodationOffersComponent;
  let fixture: ComponentFixture<FilterAccommodationOffersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FilterAccommodationOffersComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FilterAccommodationOffersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
