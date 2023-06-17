package main

import "fmt"

func main() {

	number_of_first_month_Sundays := 0
	current_first_day_of_month := 6 // December 1st, 1900 was a Saturday
	days_in_month := map[int]int{0: 31, 1: 31, 3: 31, 4: 30, 5: 31,
		6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30} // don't assign a number of days to February – yet
	weekday_of := map[int]string{0: "Sun", 1: "Mon", 2: "Tue",
		3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat"}

	for year := 1901; year <= 2000; year++ { // probably inclusive

		// determine whether a year is a leap year, and set February to that many days
		if (year%4 == 0) && (year%100 != 0) {
			days_in_month[2] = 29
		} else if year%400 == 0 {
			days_in_month[2] = 29
		} else {
			days_in_month[2] = 28
		}

		for month := 1; month <= 12; month++ {
			current_first_day_of_month = ((days_in_month[month-1] + current_first_day_of_month) % 7)

			fmt.Println("The first day of the", month, "th month of",
				year, "is", weekday_of[current_first_day_of_month], "|",
				"number of days:", days_in_month[month],
				"| number equivalent of day:", current_first_day_of_month) // debugging output

			if current_first_day_of_month == 0 {
				number_of_first_month_Sundays++
				fmt.Println("current Sunday counter:", number_of_first_month_Sundays)
			}
		}
	}
	fmt.Println(number_of_first_month_Sundays)
}

// You are given the following information, but you may prefer to do some research for yourself.

// 1 Jan 1900 was a Monday.
// Thirty days has September,
// April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine.
// And on leap years, twenty-nine.
// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
