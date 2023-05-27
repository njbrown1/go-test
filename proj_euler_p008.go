package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("1000_digit_number.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	number := string(content)
	fmt.Println("Number:", number)
}
