package main

import (
	"fmt"
	"math/big"
)

func main() {
	results_slice := []string{}
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			base := big.NewInt(int64(a))
			exponent := big.NewInt(int64(b))
			result := new(big.Int).Exp(base, exponent, nil)
			fmt.Println(a, b, result)

			add_result := true
			for _, entry := range results_slice {
				if entry == result.String() {
					add_result = false
				}
			}
			if add_result == true {
				results_slice = append(results_slice, result.String())
			}
		}
	}
	fmt.Println(results_slice, len(results_slice))
}
