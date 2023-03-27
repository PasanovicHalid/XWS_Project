import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DefaultHomepageComponent } from './default-homepage.component';

describe('DefaultHomepageComponent', () => {
  let component: DefaultHomepageComponent;
  let fixture: ComponentFixture<DefaultHomepageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DefaultHomepageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DefaultHomepageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
