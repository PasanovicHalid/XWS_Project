export class AccommodationTemp {
    id: string = "";
    name: string = "";
    ownerId: string = "";
    location: string = "";
    wifi: boolean = false;
    kitchen: boolean = false;
    airConditioner: boolean = false;
    parking: boolean = false;
    minNumberOfGuests: number = 0;
    maxNumberOfGuests: number = 0;
    images : string[] = []

    public constructor(obj?: any) {
        if (obj) { 
        }
    }
  }