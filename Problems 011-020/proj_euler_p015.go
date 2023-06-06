package main

import (
	"fmt"
	"math/big"
)

func factorialOfNumber(number int64) string {

	big_int_number := big.NewInt(number)
	current_product := big.NewInt(1)
	i := big.NewInt(1)

	for i <= big_int_number; i++ {
		current_product *= i
	}
	return fmt.Sprint(k)
}

func main() {
	result := factorialOfNumber(40)
	fmt.Println(result)
}
