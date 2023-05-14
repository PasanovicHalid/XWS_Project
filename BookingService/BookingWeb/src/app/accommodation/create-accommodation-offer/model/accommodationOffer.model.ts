import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';

export class CreateOfferRequest {
    id: string = "";
    accommodation_id: string = "";
    start_date_time_utc: string = "";
    end_date_time_utc: string = "";
    price: number = 0;
    per_guest: boolean = false;
  }