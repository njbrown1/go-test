package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find_product(thousand_digit_number string, number_of_consecutive_digits int) int {

	var slice_of_x_numbers = make([]int, number_of_consecutive_digits)

	current_product := 1

	for i := 0; i <= 12; i++ {

		current_product = 1

		for j := 0; j <= (number_of_consecutive_digits - 1); j++ {

			character, _ := strconv.Atoi(string(thousand_digit_number[i+j]))
			slice_of_x_numbers[j] = character
			fmt.Println(i, j, slice_of_x_numbers, current_product)

		}

	}

	return current_product

}

func main() {

	data, _ := os.ReadFile("1000_digit_number.txt")                               // read the file with the 1000 digit number
	var thousand_digit_number string = strings.ReplaceAll(string(data), "\n", "") // remove newline characters from the data
	fmt.Println(thousand_digit_number)                                            // debugging purposes
	fmt.Println(find_product(thousand_digit_number, 4))

}
