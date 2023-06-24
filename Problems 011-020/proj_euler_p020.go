package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func factorial_of_number_020(number int64) string {
	product := big.NewInt(1)
	for i := 1; i <= int(number); i++ {
		product.Mul(product, big.NewInt(int64(i)))
	}
	return product.String()
}

func main() {
	result := factorial_of_number_020(100)
	sum_of_factorial_digits := 0
	digits_slice := []int{}
	for i := range result {
		individual_digit, _ := strconv.Atoi(string(result[i]))
		digits_slice = append(digits_slice, individual_digit)
	}
	for i := range digits_slice {
		sum_of_factorial_digits += digits_slice[i]
	}
	fmt.Println(sum_of_factorial_digits) // problem solved! answer: 648
}
