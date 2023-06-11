package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1. p#018 states outright that there are 16384 routes, which struck me as interesting, since
// 16384 is a power of 2. then I noticed that you can represent every route as a distinct 14-digit binary number.

// 2. at each number in the triangle, you have two possible options: go down and go LEFT <OR> go down and go RIGHT.
// this can easily be represented by 0s and 1s so that:
// 0 means go down and go LEFT
// 1 means go down and go RIGHT

// 3. I define triangle_slice[x][y] as the (y + 1)th number in the (x + 1)th row. examples:
// triangle_slice[0][0] = 75 | triangle_slice[2][2] = 82 | triangle_slice[4][0] = 20

// 4. at each number (signified by a unique triangle_slice[x][y]), you must go to the next number, as specified in step 2. so:
// 0 means go down and go LEFT | triangle_slice[x][y] -> triangle_slice[x + 1][y]
// 1 means go down and go RIGHT | triangle_slice[x][y] -> triangle_slice[x + 1][y + 1]

// 5. so this problem isn't actually that hard. we just have to generate 14-digit binary numbers for the numbers 0 - 16383,
// and use those binary numbers as inputs to the calculate_sum_of_binary_path function.

func retrieve_element_from_triangle(x int, y int) int {

	numbers_file, _ := os.ReadFile("triangle_numbers.txt")                                    // grab the numbers from the file
	var data_without_newlines string = strings.ReplaceAll(string(numbers_file), "\n", "")     // remove newline characters
	var condensed_numbers string = strings.ReplaceAll(string(data_without_newlines), " ", "") // remove spaces

	triangle_slice := make([][]int, 15)

	for i := 0; i <= 14; i++ {
		triangle_slice[i] = make([]int, i+1)

		starting_retrieval_index := 0
		for k := 0; k <= i; k++ {
			starting_retrieval_index += k
		}
		starting_retrieval_index *= 2

		for j := 0; j <= i; j++ {
			tens_digit, _ := strconv.Atoi(string(condensed_numbers[starting_retrieval_index+2*j]))
			ones_digit, _ := strconv.Atoi(string(condensed_numbers[starting_retrieval_index+2*j+1]))
			triangle_slice[i][j] = (tens_digit * 10) + (ones_digit) // assign the two-digit number to the two-dimensional array
		}
	}
	return triangle_slice[x][y]
}

func return_14_digit_binary_number(base_10_number int64) string {
	binary_number := strconv.FormatInt(base_10_number, 2)     // convert the base_10_number to binary
	number_of_leading_zeros_needed := 14 - len(binary_number) // calculate the number of leading zeros missing (to make the binary number 14 digits long)
	for i := 1; i <= number_of_leading_zeros_needed; i++ {    // add the necessary number of missing zeros
		binary_number = "0" + binary_number
	}
	return binary_number
}

func calculate_sum_of_binary_path(final_binary_number string) int {
	current_row_INDEX := 0
	current_number_INDEX := 0
	total_sum := retrieve_element_from_triangle(current_row_INDEX, current_number_INDEX)

	for i := 1; i <= 14; i++ {
		current_row_INDEX++
		ith_digit_of_final_binary_number, _ := strconv.Atoi(string(final_binary_number[i-1]))
		if ith_digit_of_final_binary_number == 1 {
			current_number_INDEX++
		} else {
		}
		total_sum += retrieve_element_from_triangle(current_row_INDEX, current_number_INDEX)
		fmt.Println(
			"crI:", current_row_INDEX,
			"| cnI:", current_number_INDEX,
			"| ith digit:", ith_digit_of_final_binary_number,
			"| element_from_triangle:", retrieve_element_from_triangle(current_row_INDEX, current_number_INDEX),
		)
	}
	return total_sum
}

func main() {
	final_binary_number := return_14_digit_binary_number(127)
	total_sum := calculate_sum_of_binary_path(final_binary_number)
	fmt.Println("final_binary_number:", final_binary_number, "total_sum:", total_sum)
}
