package main

import "fmt"

func main() {

	start_year := 1901
	current_first_weekday_of_month := 6 // December 1st, 1900 was a Saturday
	days_in_month := map[int]int{1: 31, 3: 31, 4: 30, 5: 31,
		6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 30}
	weekday_of := map[int]string{0: "Sun", 1: "Mon", 2: "Tue",
		3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat"}
	fmt.Println(start_year, days_in_month, weekday_of)

	for year := 1901; year <= 2000; year++ { // probably inclusive, but idk

		// determine whether a year is a

	}
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
