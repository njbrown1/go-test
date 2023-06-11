package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers_file, _ := os.ReadFile("triangle_numbers.txt")                                    // grab the numbers from the file
	var data_without_newlines string = strings.ReplaceAll(string(numbers_file), "\n", "")     // remove newline characters
	var condensed_numbers string = strings.ReplaceAll(string(data_without_newlines), " ", "") // remove spaces
	fmt.Println(condensed_numbers)

	two_dimensional_slice := make([][]int, 15)

	for i := 0; i <= 14; i++ {
		two_dimensional_slice[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			tens_digit, _ := strconv.Atoi(string(condensed_numbers[2*j]))
			ones_digit, _ := strconv.Atoi(string(condensed_numbers[2*i+2*j+1]))
			two_dimensional_slice[i][j] = (tens_digit * 10) + (ones_digit) // assign the two-digit number to the two-dimensional array
		}
	}

	fmt.Println(two_dimensional_slice)
	fmt.Println(len(two_dimensional_slice[14]))
}
