export class Accommodation {
    id: string = "";
    name: string = "";
    ownerId: string = "";
    location: string = "";
    wifi: boolean = false;
    kitchen: boolean = false;
    air_conditioner: boolean = false;
    parking: boolean = false;
    min_number_of_guests: number = 0;
    max_number_of_guests: number = 0;
    images : string[] = []

    public constructor(obj?: any) {
        if (obj) { 
        }
    }
  }