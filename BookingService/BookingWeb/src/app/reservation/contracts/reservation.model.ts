export class Reservation {
    id: string = '';
    accommodationOfferId: string = '';
    customerId: string = '';
    hostId: string = '';
    reservationStatus: ReservationStatus = ReservationStatus.PENDING;
    numberOfGuests: number = 0;
    startDateTimeUtc: string = '';
    endDateTimeUtc: string = '';
  }
  
  export enum ReservationStatus {
    PENDING = 0,
    ACCEPTED = 1,
    REJECTED = 2,
  }