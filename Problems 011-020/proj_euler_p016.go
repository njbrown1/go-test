package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	// set basic variables, all in BigInt to eliminate conversion issues.
	base := big.NewInt(2)
	exponent := big.NewInt(1000)
	result := big.NewInt(0)

	result.Exp(base, exponent, nil)     // finds the result. very fast too: took less than 2 seconds.
	string_of_result := result.String() // turns the result into a string.

	var current_sum int64 = 0
	for i := 0; i < len(string_of_result); i++ { // add each digit of the string to current_sum.
		digit, _ := strconv.Atoi(string(string_of_result[i]))
		current_sum += int64(digit)
	}
	fmt.Println(current_sum) // solved! answer: 1366
}
