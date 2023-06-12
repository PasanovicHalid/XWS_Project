import { Pipe, PipeTransform } from "@angular/core";

@Pipe({name: 'first_letter_lowercase'})
export class FirstLetterLowercasePipe implements PipeTransform {
    transform(value: any, ...args: any[]) {
        if (typeof value !== 'string') {
            return value;
        }
        return value.charAt(0).toLowerCase() + value.slice(1);
    }
}
