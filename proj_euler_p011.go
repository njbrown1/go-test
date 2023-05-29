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

	fmt.Println("data:", string(condensed_data))                                                          // for debugging
	fmt.Println("length of data:", len(condensed_data), "| data index of 0:", string(condensed_data[41])) // for debugging

	two_dimensional_slice := make([][]int, 20)
	for i := 0; i <= 19; i++ {

		two_dimensional_slice[i] = make([]int, 20)

		for j := 0; j <= 19; j++ {
			tens_digit, _ := strconv.Atoi(string(condensed_data[(40*i)+(2*j)]))
			ones_digit, _ := strconv.Atoi(string(condensed_data[(40*i)+(2*j)+1])) // access the next index
			two_dimensional_slice[i][j] = (tens_digit * 10) + (ones_digit)        // assign the two-digit number to the two-dimensional array
		}
		fmt.Println(two_dimensional_slice) // debugging purposes
	}

	fmt.Println(find_largest_product_up_and_down(two_dimensional_slice)) // debugging purposes
}

func find_largest_product_up_and_down(grid [][]int) int { // input: the two_dimensional_slice (grid) of numbers. output: the largest product.

	// this function starts

	for row := 1; row <= 20; row++ { // indexing adjusted
		for column := 1; column <= 20; column++ {

			number_1 := grid[row-1+0][column-1]
			number_2 := grid[row-1+1][column-1]
			number_3 := grid[row-1+2][column-1]
			number_4 := grid[row-1+3][column-1]
			fmt.Println(number_1, number_2, number_3, number_4)
		}
	}
	return 1024
}

// I need four functions:
// 1. a function that finds the largest product for numbers up / down
// 2. a function that finds the largest product for numbers left / right
// 3. a function that finds the largest product for numbers diagonally, up and to the right
// 4. a function that finds the largest product for numbers diagonally, up and to the left
