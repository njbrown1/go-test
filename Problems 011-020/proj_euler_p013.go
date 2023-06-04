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

	fmt.Println("hello world")
	fmt.Println(condensed_numbers)

	var fifty_digit_number_slice []int

	for i := 0; i <= 49; i++ {
		one_number_only, _ := strconv.Atoi(string(condensed_numbers[i]))
		fifty_digit_number_slice = append(fifty_digit_number_slice, one_number_only)

		// strconv.Atoi(string(condensed_data[(40*i)+(2*j)]))
	}

	fmt.Println(fifty_digit_number_slice)
}
