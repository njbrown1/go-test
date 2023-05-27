package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, _ := os.ReadFile("1000_digit_number.txt")                               // read the file with the 1000 digit number
	var thousand_digit_number string = strings.ReplaceAll(string(data), "\n", "") // remove newline characters from the data

	x := 4
	var slice_of_x_numbers = make([]int, x)
	// largest_product := 0
	fmt.Println(thousand_digit_number)

	for i := 0; i <= 4; i++ { // this

		for j := 0; j <= 3; j++ {
			character, _ := strconv.Atoi(string(thousand_digit_number[i+j]))
			slice_of_x_numbers[j] = character
			fmt.Println(i, j, slice_of_x_numbers)
		}
	}
}
