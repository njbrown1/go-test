package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	grid_file, _ := os.ReadFile("20x20_grid.txt")                                          // grab the numbers from the file
	var data_without_newlines string = strings.ReplaceAll(string(grid_file), "\n", "")     // remove newline characters
	var condensed_data string = strings.ReplaceAll(string(data_without_newlines), " ", "") // remove spaces

	two_dimensional_slice := make([][]int, 20)

	for i := 0; i <= 19; i++ {

		two_dimensional_slice[i] = make([]int, 20)

		for j := 0; j <= 19; j++ {
			tens_digit, _ := strconv.Atoi(string(condensed_data[(40*i)+(2*j)]))
			ones_digit, _ := strconv.Atoi(string(condensed_data[(40*i)+(2*j)+1])) // access the next index
			two_dimensional_slice[i][j] = (tens_digit * 10) + (ones_digit)        // assign the two-digit number to the two-dimensional array
		}
	}
	fmt.Println(find_largest_product_up_and_down(two_dimensional_slice)) // necessary
}

func find_largest_product_up_and_down(grid [][]int) int { // input: the two_dimensional_slice (grid) of numbers. output: the largest product.

	for row := 0; row <= 19; row++ {
		for column := 0; column <= 19; column++ {

			x := column
			y := row
			out_of_range := (x < 0 || x > 19) || (y < 0 || y > 19)
			var four_number_slice []int

			for k := 1; k <= 4; k++ {
				if out_of_range == true {
					break
				}
				four_number_slice = append(four_number_slice, grid[y][x])
				x++
				out_of_range = (x < 0 || x > 19) || (y < 0 || y > 19)
			}

			fmt.Println(four_number_slice)
		}
	}
	return 1024
}

// I need four functions:
// 1. a function that finds the largest product for numbers up / down
// 2. a function that finds the largest product for numbers left / right
// 3. a function that finds the largest product for numbers diagonally, up and to the right
// 4. a function that finds the largest product for numbers diagonally, up and to the left
