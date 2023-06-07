package main

import (
	"fmt"
	"math/big"
)

func factorial_of_number(number int64) *big.Int {

	// BigInts are required – instead of int64 – because the factorial of
	// numbers > 30 exceeds the integer limit for int64.

	product := big.NewInt(1)
	for i := 1; i <= int(number); i++ {
		product.Mul(product, big.NewInt(int64(i)))
	}

	return product
}

func main() {
	result := factorial_of_number(40)
	fmt.Println(result)
}
