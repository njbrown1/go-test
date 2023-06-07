package main

import (
	"fmt"
	// "math"
	"strconv"
)

func add_leading_zeros(three_digit_number int) string {
	var new_number string
	var number string = fmt.Sprint(three_digit_number)
	if three_digit_number < 10 {
		new_number = "00" + number
	} else if three_digit_number < 100 {
		new_number = "0" + number
	} else if three_digit_number < 1000 {
		new_number = number
	}
	return new_number
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
	fmt.Println(add_leading_zeros(2))
}
