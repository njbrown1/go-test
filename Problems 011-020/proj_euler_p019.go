package main

import "fmt"

func main() {

	// a 40,000 foot view of how this code works:
	// 1. I set a number of variables and maps. I've established a convention that 0 = Sunday, 1 = Monday,
	// ..., 5 = Friday, 6 = Saturday. this is in place so I can use modulo to my advantage.
	// 2. I determine if the given year is a leap year, and set February to the correct number of days in
	// the days_in_month map.
	// 3. for every month in that year, I adjust the variable current_first_day_of_month BASED ON the
	// number of days in the PREVIOUS month. this is why you see days_in_month[month-1] used. after a bit of
	// experimenting, I found that using modulo was the easiest way to find the current_first_day_of_month.
	// 4. if current_first_day_of_month == 0, I increment the number_of_first_month_Sundays variable by 1.

	number_of_first_month_Sundays := 0
	current_first_day_of_month := 6 // December 1st, 1900 was a Saturday
	days_in_month := map[int]int{0: 31, 1: 31, 3: 31, 4: 30, 5: 31,
		6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30}
	weekday_of := map[int]string{0: "Sun", 1: "Mon", 2: "Tue",
		3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat"}
	month_of := map[int]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun",
		7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}

	for year := 1901; year <= 2000; year++ { // inclusive

		if (year%4 == 0) && (year%100 != 0) {
			days_in_month[2] = 29
		} else if year%400 == 0 {
			days_in_month[2] = 29
		} else {
			days_in_month[2] = 28
		}

		for month := 1; month <= 12; month++ {
			current_first_day_of_month = ((days_in_month[month-1] + current_first_day_of_month) % 7)
			fmt.Println(month_of[month], "1", year, "is a", weekday_of[current_first_day_of_month])
			if current_first_day_of_month == 0 {
				number_of_first_month_Sundays++
				fmt.Println("current Sunday counter:", number_of_first_month_Sundays)
			}
		}
	}
	fmt.Println(number_of_first_month_Sundays) // solved! answer: 171
}
