package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// this problem was very easy to conceptualize. here are a few things you should know:

// 1. I define triangle_slice[x][y] as the (y + 1)th number in the (x + 1)th row. examples:
// | triangle_slice[0][0] = 75 | triangle_slice[2][2] = 82 | triangle_slice[4][0] = 20 |

// 2. the whole first section of the code is devoted to extracting the raw file of numbers
// from triangle_numbers.txt, turning it into condensed_numbers (ie. no newlines or spaces)
// and putting the numbers into the two-dimensional slice (triangle_slice).

func find_starting_index_of_slice_n(n int) int {
	starting_index := 0
	for i := 0; i <= n; i++ {
		starting_index += i
	}
	starting_index *= 2
	return starting_index
}

func main() {
	numbers_file, _ := os.ReadFile("triangle_numbers.txt")                                    // grab the numbers from the file
	var data_without_newlines string = strings.ReplaceAll(string(numbers_file), "\n", "")     // remove newline characters
	var condensed_numbers string = strings.ReplaceAll(string(data_without_newlines), " ", "") // remove spaces
	fmt.Println(condensed_numbers)

	triangle_slice := make([][]int, 15)

	for i := 0; i <= 14; i++ {
		triangle_slice[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			tens_digit, _ := strconv.Atoi(string(condensed_numbers[2*j]))
			ones_digit, _ := strconv.Atoi(string(condensed_numbers[2*j+1]))
			triangle_slice[i][j] = (tens_digit * 10) + (ones_digit) // assign the two-digit number to the two-dimensional array
		}
	}

	fmt.Println(triangle_slice)
	fmt.Println(len(triangle_slice[14]))
	fmt.Println(find_starting_index_of_slice_n(4))
}
