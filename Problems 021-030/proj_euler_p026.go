package main

import "fmt"

import "math"

func perform_long_divison(divisor int, decimal_places int) (string, int) { // where the numerator HAS to be 1, and the divisor is d (as specified in p#026).

	remainder_list := []int{}
	remainder := 1
	quotient := "0."
	is_terminating := false

	for i := 1; i <= decimal_places; i++ {

		if remainder == 0 {
			is_terminating = true
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

	}

	// check to find the longest recurring sequence
	length_of_recurring_sequence := 0
	if is_terminating == false {
		last_index := len(remainder_list) - 1
		very_last_remainder := remainder_list[last_index]
		fmt.Println(very_last_remainder, last_index)
		i := last_index - 1
		fmt.Println(very_last_remainder, i, last_index)
		for very_last_remainder != remainder_list[i] {
			i--
		}
		fmt.Println(last_index, i)
		length_of_recurring_sequence = last_index - i
	}

	return quotient, length_of_recurring_sequence
}

func main() {
	d := 9
	quotient, length_of_recurring_sequence := perform_long_divison(d, 20)
	fmt.Println(quotient, length_of_recurring_sequence)
}
