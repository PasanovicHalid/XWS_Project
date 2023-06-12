import { Pipe, PipeTransform } from "@angular/core";

@Pipe({name: 'ratingZero'})
export class RatingZeroPipe implements PipeTransform {
    transform(value: any, ...args: any[]) {
        return value != 0 ? value : "-";
    }
}
