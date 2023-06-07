package main

import "math/big"
import "fmt"
import "strconv"

func main() {
	base := big.NewInt(2)
	exponent := big.NewInt(1000)
	result := big.NewInt(0)

	result.Exp(base, exponent, nil)

	string_of_result := result.String()

	var current_sum int64 = 0
	for i := 0; i < len(string_of_result); i++ {
		digit, _ := strconv.Atoi(string(string_of_result[i]))
		current_sum += int64(digit)
	}

	fmt.Println(current_sum)
}
