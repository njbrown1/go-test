package main

import "fmt"

import "math"

func perform_long_divison(divisor int, decimal_places int) (string, int) { // where the numerator HAS to be 1, and the divisor is d (as specified in p#026).

	remainder_list := []int{}
	remainder := 1
	quotient := "0."
	length_of_recurring_sequence := 0

	for i := 1; i <= decimal_places; i++ {
		if remainder == 0 {
			break
		}
		number_of_multiplications := 0
		for divisor > remainder {
			remainder *= 10
			remainder_list = append(remainder_list, remainder)
			number_of_multiplications++
			if number_of_multiplications > 1 {
				quotient += "0"
			}
		}
		digit := math.Floor(float64(remainder) / float64(divisor))
		quotient += fmt.Sprint(digit)

		fmt.Println(remainder, digit, quotient, remainder_list)

		remainder -= int(digit) * divisor
		for i := range remainder_list {
			fmt.Println(i, remainder_list)
			if remainder == remainder_list[i] { // if the last element in remainder_list = an entry in remainder_list
				fmt.Println("if statement does run")
				length_of_recurring_sequence = (len(remainder_list) - 1) - i
			}
		}
	}
	return quotient, length_of_recurring_sequence
}

func main() {
	d := 70
	quotient, length_of_recurring_sequence := perform_long_divison(d, 20)
	fmt.Println(quotient, length_of_recurring_sequence)
}
