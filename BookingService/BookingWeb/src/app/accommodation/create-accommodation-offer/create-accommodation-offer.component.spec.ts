import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateAccommodationOfferComponent } from './create-accommodation-offer.component';

describe('CreateAccommodationOfferComponent', () => {
  let component: CreateAccommodationOfferComponent;
  let fixture: ComponentFixture<CreateAccommodationOfferComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateAccommodationOfferComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateAccommodationOfferComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
