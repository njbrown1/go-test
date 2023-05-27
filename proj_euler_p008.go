package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data, _ := os.ReadFile("1000_digit_number.txt")                               // read the file with the 1000 digit number
	var thousand_digit_number string = strings.ReplaceAll(string(data), "\n", "") // remove newline characters from the data

	fmt.Println(thousand_digit_number)

	for i := 0; i <= 4; i++ { // start with the index i

		for j := 0; j <= 3; j++ {
			character := thousand_digit_number[i+j]
			fmt.Println(i, j, string(character))
		}
	}
}
