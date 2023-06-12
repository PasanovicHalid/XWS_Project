import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HostRatingsComponent } from './host-ratings.component';

describe('HostRatingsComponent', () => {
  let component: HostRatingsComponent;
  let fixture: ComponentFixture<HostRatingsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HostRatingsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HostRatingsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
