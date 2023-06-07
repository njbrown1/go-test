package main

import (
	"fmt"
	"math/big"
)

func factorialOfNumber(number int64) string {

	product := big.NewInt(1)
	for i := 1; i <= int(number); i++ {
		product.Mul(product, big.NewInt(int64(i)))
	}
	return_string := product.String()
	return return_string
}

func main() {
	result := factorialOfNumber(40)
	fmt.Println(result)
}
