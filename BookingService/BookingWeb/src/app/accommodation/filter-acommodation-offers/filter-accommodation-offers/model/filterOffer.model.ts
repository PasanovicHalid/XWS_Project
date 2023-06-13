export class AccommodationFilterOffer {
    location: string = "";
    from: string = "";
    to: string = "";
    guestNumber: number = 0;
    minPrice: number = 0;
    maxPrice: number = 0;
    filterByRating: boolean = false;
    ratingBottom: number = 0;
    ratingTop: number = 0;
    hostIsDistinguished : boolean = false;
    wifi: boolean = false;
    kitchen: boolean = false;
    airConditioner: boolean = false;
    parking: boolean = false;

    public constructor(obj?: any) {
        if (obj) { 
        }
    }
  }