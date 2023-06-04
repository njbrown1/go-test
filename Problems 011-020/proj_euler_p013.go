package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello world")

	numbers_file, _ := os.ReadFile("100_fifty_digit_numbers.txt")                         // grab the numbers from the file
	var data_without_newlines string = strings.ReplaceAll(string(numbers_file), "\n", "") // remove newline characters
	var condensed_numbers string = strings.ReplaceAll(string(data_without_newlines), " ", "")

	fmt.Println(condensed_numbers)
	fmt.Println(numbers_file)
}
