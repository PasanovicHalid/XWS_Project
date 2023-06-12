import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllRatingsComponent } from './all-ratings.component';

describe('AllRatingsComponent', () => {
  let component: AllRatingsComponent;
  let fixture: ComponentFixture<AllRatingsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AllRatingsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AllRatingsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
