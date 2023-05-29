package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	grid_file, _ := os.ReadFile("20x20_grid.txt")                      // grab the numbers from the file
	var data string = strings.ReplaceAll(string(grid_file), "\n", " ") // convert newline characters to spaces
	data = strings.ReplaceAll(string(grid_file), " ", "")              // remove spaces
	// var initial_slice []int

	fmt.Println(string(data)) // for debugging
	fmt.Println(len(data))    // for debugging

	// for i := 0; i <= 57; {

	// 	first_digit, _ := strconv.Atoi(string(data[i]))
	// 	second_digit, _ := strconv.Atoi(string(data[i+1]))
	// 	initial_slice = append(initial_slice, first_digit*10+second_digit)
	//	i = i + 3
	//	fmt.Printf("%02d\n", initial_slice) // for debugging
	// }

	var twoD [20][20]int
	for i := 0; i <= 20; {
		for j := 0; j <= 20; {

			tens_digit, _ := strconv.Atoi(string(data[i]))
			ones_digit, _ := strconv.Atoi(string(data[i+1]))
			two_digit_number := tens_digit*10 + ones_digit

			twoD[i][j] = two_digit_number

			j = j + 3
		}
		i = i + 60
	}
	fmt.Println("2d: ", twoD)
	fmt.Println(data)
}

// I need four functions:
// 1. a function that finds the largest product for numbers up / down
// 2. a function that finds the largest product for numbers left / right
// 3. a function that finds the largest product for numbers diagonally, up and to the right
// 4. a function that finds the largest product for numbers diagonally, up and to the left
