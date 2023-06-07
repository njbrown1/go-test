package main

import (
	"fmt"
	// "math"
	"strconv"
)

func add_leading_zeros(number int) string { // takes any number from 0-999, and returns a string with leading zeros added (if the original number is only one or two digits).
	var number_string string              // 'number_string' is the output that will be returned.
	var input string = fmt.Sprint(number) // 'input' is the string equivalent of 'number'. 'input' is created so string concatenation is possible.
	if number < 10 {
		number_string = "00" + input // if number only has one digit, add TWO leading zeros.
	} else if number < 100 {
		number_string = "0" + input // if number has two digits, add ONE leading zeros.
	} else if number < 1000 {
		number_string = "" + input // if number has three digits, add NO leading zeros.
	}
	return number_string
}

func find_length_of_written_number(three_digit_number string) int {
	hundreds_digit, _ := strconv.Atoi(string(three_digit_number[0]))
	tens_digit, _ := strconv.Atoi(string(three_digit_number[1]))
	ones_digit, _ := strconv.Atoi(string(three_digit_number[2]))
	fmt.Println(hundreds_digit, tens_digit, ones_digit)

	return 4
}

func main() {
	fmt.Println("hello world")
	fmt.Println(add_leading_zeros(0))
}
