package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// this problem was very easy to conceptualize. here are a few things you should know:

// retrieve_element_from_triangle is devoted to extracting the raw file of numbers
// from triangle_numbers.txt, turning it into condensed_numbers (ie. no newlines or spaces)
// and putting the numbers into triangle_slice (a two-dimensional slice).

// 1. I define triangle_slice[x][y] as the (y + 1)th number in the (x + 1)th row. examples:
// | triangle_slice[0][0] = 75 | triangle_slice[2][2] = 82 | triangle_slice[4][0] = 20 |

// 2. at each number (signified by a unique triangle_slice[x][y]), you must move EITHER down
// and left OR down and right. this means that triangle_slice[x][y] -> triangle_slice[x + 1][y]
// OR triangle_slice[x][y] -> triangle_slice[x + 1][y + 1].

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
	for i := 1; i < number_of_leading_zeros_needed; i++ {     // add the necessary number of missing zeros
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
		if final_binary_number[i-1] == 0 {
			current_number_INDEX++
		} else {
		}
		total_sum += retrieve_element_from_triangle(current_row_INDEX, current_number_INDEX)
	}
	return total_sum
}

func main() {
	fmt.Println(return_14_digit_binary_number(207))
	fmt.Println(calculate_sum_of_binary_path("00000000000001"))
}
