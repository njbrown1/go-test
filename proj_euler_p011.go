package main

import (
	"fmt"
	"os"
	// "strconv"
	"strings"
)

func main() {

	grid_file, _ := os.ReadFile("20x20_grid.txt") // grab the numbers from the file
	var data_without_newlines string = strings.ReplaceAll(string(grid_file), "\n", "")
	var condensed_data string = strings.ReplaceAll(string(data_without_newlines), " ", "")

	fmt.Println("data:", string(condensed_data))                                                          // for debugging
	fmt.Println("length of data:", len(condensed_data), "| data index of 0:", string(condensed_data[41])) // for debugging

	var twoD [20][20]int
	fmt.Println("2d: ", twoD)

}

// I need four functions:
// 1. a function that finds the largest product for numbers up / down
// 2. a function that finds the largest product for numbers left / right
// 3. a function that finds the largest product for numbers diagonally, up and to the right
// 4. a function that finds the largest product for numbers diagonally, up and to the left
