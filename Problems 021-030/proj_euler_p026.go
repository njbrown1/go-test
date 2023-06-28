package main

import "fmt"
import "math"

func perform_long_divison(divisor int, decimal_places int) (int, string) { // where the numerator HAS to be 1, and the divisor is d (as specified in p#026).
	recurring_cycle_length := 0
	remainder_slice := []int{}
	remainder := 1
	quotient := "0."
	for i := 1; i <= decimal_places; i++ { // for the specified number of decimal places
		if remainder == 0 { // if the remainder is 0, the fraction is terminating, and so Go should break the outer 'for' loop.
			break
		}
		number_of_multiplications := 0
		for divisor > remainder {
			remainder *= 10
			number_of_multiplications++
			if number_of_multiplications > 1 { // if there's more than one multiplication of 10, for each of them, there needs to be an extra 0 in the quotient.
				quotient += "0"
			}
		}
		digit := math.Floor(float64(remainder) / float64(divisor)) // calculate the digit
		quotient += fmt.Sprint(digit)                              // add the digit to the quotient

		for _, entry := range remainder_slice {
			if remainder == entry {
				recurring_cycle_length = len(remainder_slice)
				break
			}
		}
		remainder_slice = append(remainder_slice, remainder)
		if recurring_cycle_length != 0 {
			break
		}
		fmt.Println(remainder, digit, quotient, remainder_slice)
		remainder -= int(digit) * divisor // adjust the remainder
	}
	return recurring_cycle_length, quotient
}

func main() {
	d := 51
	length_of_recurring_sequence, quotient := perform_long_divison(d, 30)
	fmt.Println("1 divided by", d, "is approximately:", quotient)
	fmt.Println("length_of_recurring_sequence:", length_of_recurring_sequence)
}
