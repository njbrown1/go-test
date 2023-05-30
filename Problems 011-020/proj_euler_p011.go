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
		two_dimensional_slice[i] = make([]int, 20) // make the inner slices with length 20
		for j := 0; j <= 19; j++ {
			tens_digit, _ := strconv.Atoi(string(condensed_data[(40*i)+(2*j)]))
			ones_digit, _ := strconv.Atoi(string(condensed_data[(40*i)+(2*j)+1]))
			two_dimensional_slice[i][j] = (tens_digit * 10) + (ones_digit) // assign the two-digit number to the two-dimensional array
		}
	}

	fmt.Println(find_largest_product_up_and_down(two_dimensional_slice))
	fmt.Println(find_largest_product_left_and_right(two_dimensional_slice))
	fmt.Println(find_largest_product_up_right_diagonally(two_dimensional_slice))
	fmt.Println(find_largest_product_down_right_diagonally(two_dimensional_slice))

	// I just had it print the largest product for all four functions, then identified the largest number.
	// solved! answer: 70600674
}

// NOTES ON THE FOUR FUNCTIONS:
// 1. I first wrote the function find_largest_product_up_and_down, and I wanted to make the function extremely reusable.
// The only difference between all four functions is the presence of y++, y--, x++, and x--.
// 2. Each function accesses every number in the grid as a 'starting point', and then continues to go in the direction specified
// in the function name.
// 3. There might be a better way to do this, but copy-pasting code is really easy.

func find_largest_product_up_and_down(grid [][]int) int { // input: the two_dimensional_slice (grid) of numbers. output: the largest product.

	current_largest_product := 0

	for row := 0; row <= 19; row++ {
		for column := 0; column <= 19; column++ { // access EVERY number in the grid

			x := column                                            // so that x can be adjusted without adjusting 'column'
			y := row                                               // so that y can be adjusted without adjusting 'row'
			out_of_range := (x < 0 || x > 19) || (y < 0 || y > 19) // a boolean that returns true if the requested value is out of range
			var four_number_slice []int

			for k := 1; k <= 4; k++ {
				if out_of_range == true {
					break // if x OR y is out of range, then nothing else is added to four_number_slice
				}
				four_number_slice = append(four_number_slice, grid[y][x])
				y++
				out_of_range = (x < 0 || x > 19) || (y < 0 || y > 19)
			}

			// the code below multiplies together the numbers in the four_number_slice, then updates the current_largest_product if product > current_largest_product.

			product := 1
			for j := 0; j <= len(four_number_slice)-1; j++ {
				product = product * four_number_slice[j]
			}

			if product > current_largest_product {
				current_largest_product = product
				fmt.Println(four_number_slice, product)
			}
		}
	}
	return current_largest_product
}

func find_largest_product_left_and_right(grid [][]int) int { // input: the two_dimensional_slice (grid) of numbers. output: the largest product.

	current_largest_product := 0

	for row := 0; row <= 19; row++ {
		for column := 0; column <= 19; column++ { // access EVERY number in the grid

			x := column                                            // so that x can be adjusted without adjusting 'column'
			y := row                                               // so that y can be adjusted without adjusting 'row'
			out_of_range := (x < 0 || x > 19) || (y < 0 || y > 19) // a boolean that returns true if the requested value is out of range
			var four_number_slice []int

			for k := 1; k <= 4; k++ {
				if out_of_range == true {
					break
				}
				four_number_slice = append(four_number_slice, grid[y][x])
				x++
				out_of_range = (x < 0 || x > 19) || (y < 0 || y > 19)
			}

			// the code below multiplies together the numbers in the four_number_slice, then updates the current_largest_product if product > current_largest_product.

			product := 1
			for j := 0; j <= len(four_number_slice)-1; j++ {
				product = product * four_number_slice[j]
			}

			if product > current_largest_product {
				current_largest_product = product
				fmt.Println(four_number_slice, product)
			}
		}
	}
	return current_largest_product
}

func find_largest_product_up_right_diagonally(grid [][]int) int { // input: the two_dimensional_slice (grid) of numbers. output: the largest product.

	current_largest_product := 0

	for row := 0; row <= 19; row++ {
		for column := 0; column <= 19; column++ { // access EVERY number in the grid

			x := column                                            // so that x can be adjusted without adjusting 'column'
			y := row                                               // so that y can be adjusted without adjusting 'row'
			out_of_range := (x < 0 || x > 19) || (y < 0 || y > 19) // a boolean that returns true if the requested value is out of range
			var four_number_slice []int

			for k := 1; k <= 4; k++ {
				if out_of_range == true {
					break
				}
				four_number_slice = append(four_number_slice, grid[y][x])
				x++
				y--
				out_of_range = (x < 0 || x > 19) || (y < 0 || y > 19)
			}

			// the code below multiplies together the numbers in the four_number_slice, then updates the current_largest_product if product > current_largest_product.

			product := 1
			for j := 0; j <= len(four_number_slice)-1; j++ {
				product = product * four_number_slice[j]
			}

			if product > current_largest_product {
				current_largest_product = product
				fmt.Println(four_number_slice, product)
			}
		}
	}
	return current_largest_product
}

func find_largest_product_down_right_diagonally(grid [][]int) int { // input: the two_dimensional_slice (grid) of numbers. output: the largest product.

	current_largest_product := 0

	for row := 0; row <= 19; row++ {
		for column := 0; column <= 19; column++ { // access EVERY number in the grid

			x := column                                            // so that x can be adjusted without adjusting 'column'
			y := row                                               // so that y can be adjusted without adjusting 'row'
			out_of_range := (x < 0 || x > 19) || (y < 0 || y > 19) // a boolean that returns true if the requested value is out of range
			var four_number_slice []int

			for k := 1; k <= 4; k++ {
				if out_of_range == true {
					break
				}
				four_number_slice = append(four_number_slice, grid[y][x])
				x++
				y++
				out_of_range = (x < 0 || x > 19) || (y < 0 || y > 19)
			}

			// the code below multiplies together the numbers in the four_number_slice, then updates the current_largest_product if product > current_largest_product.

			product := 1
			for j := 0; j <= len(four_number_slice)-1; j++ {
				product = product * four_number_slice[j]
			}

			if product > current_largest_product {
				current_largest_product = product
				fmt.Println(four_number_slice, product)
			}
		}
	}
	return current_largest_product
}
