import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UpdateAccommodationOfferComponent } from './update-accommodation-offer.component';

describe('UpdateAccommodationOfferComponent', () => {
  let component: UpdateAccommodationOfferComponent;
  let fixture: ComponentFixture<UpdateAccommodationOfferComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UpdateAccommodationOfferComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UpdateAccommodationOfferComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
