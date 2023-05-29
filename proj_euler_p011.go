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

	var testdata string = "01020304050607080910111213141516"

	var twoD [4][4]int
	fmt.Println("testdata:", twoD)
	fmt.Println("length of testdata:", len(testdata))

	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			tens_digit, _ := strconv.Atoi(string(testdata[(8*i)+(2*j)]))   //
			ones_digit, _ := strconv.Atoi(string(testdata[(8*i)+(2*j)+1])) // access the next index
			twoD[i][j] = (tens_digit * 10) + (ones_digit)

			fmt.Println(i, j, "|", tens_digit, ones_digit)
			fmt.Println(twoD)
		}
	}

}

// I need four functions:
// 1. a function that finds the largest product for numbers up / down
// 2. a function that finds the largest product for numbers left / right
// 3. a function that finds the largest product for numbers diagonally, up and to the right
// 4. a function that finds the largest product for numbers diagonally, up and to the left
