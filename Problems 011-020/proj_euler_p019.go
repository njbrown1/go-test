package main

import "fmt"

func main() {
	fmt.Println("great")

	year := 1901
	month := "Jan"
	number_of_days := map[string]int{"Jan": 31, "Mar": 31, "Apr": 30, "May": 31,
		"Jun": 30, "Jul": 31, "Aug": 31, "Ssep": 30, "Oct": 31, "Nov": 30, "Dec": 30}

	fmt.Println(year, month, number_of_days)
	number_of_days["Feb"] = 28
	fmt.Println(year, month, number_of_days)
	number_of_days["Feb"] = 29
	fmt.Println(year, month, number_of_days)
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
