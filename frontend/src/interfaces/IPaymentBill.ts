import { BookingInterface } from "./IBooking";
import { EmployeeInterface } from "./IEmployee";
import { SemesterInterface } from "./ISemester";

export interface PaymentBillInterface {
    ID?:                    number;
    Billing_Date?:          Date | null;
    Electric_Bill?:         number;
    Water_Bill?:            number;
    Payment_Balance?:       number;
    EmployeeID?:            number;
    Employee?:              EmployeeInterface;
    BookingID?:             number;
    Booking?:               BookingInterface;
    SemesterID?:            number;
    Semester?:              SemesterInterface;
}