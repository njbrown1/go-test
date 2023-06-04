package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	numbers_file, _ := os.ReadFile("100_fifty_digit_numbers.txt")                             // grab the numbers from the file
	var data_without_newlines string = strings.ReplaceAll(string(numbers_file), "\n", "")     // remove newline characters
	var condensed_numbers string = strings.ReplaceAll(string(data_without_newlines), " ", "") // remove spaces

	var fifty_digit_number_slice []int
	sum_of_numbers := 0

	for i := 0; i <= 4999; i = i + 50 {
		var fifty_digit_string string
		for j := 0; j <= 12; j++ {
			fifty_digit_string += string(condensed_numbers[i+j])
		}
		fmt.Println(fifty_digit_string)
		number, _ := strconv.Atoi(fifty_digit_string)
		fifty_digit_number_slice = append(fifty_digit_number_slice, number)
	}

	for i := 0; i <= 99; i++ {
		sum_of_numbers = sum_of_numbers + fifty_digit_number_slice[i]
	}

	fmt.Println("the sum of the numbers is:", sum_of_numbers)
}
