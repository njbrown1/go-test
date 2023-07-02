package main

import (
	"fmt"
	"math"
)

func is_k_prime(k int) bool {
	if k < 2 {
		return false
	}
	square_root_of_k := math.Ceil(math.Sqrt(float64(k)))
	for possible_divisor := 2; possible_divisor <= int(square_root_of_k); possible_divisor++ {
		if k%possible_divisor == 0 {
			return false
		}
	}
	return true
}

func main() {

	result := is_k_prime(1009)
	fmt.Println(result)

	for a := 0; a <= 1; a++ {
		for b := 0; b <= 1; b++ {
			fmt.Println(a, b)
		}
	}
}
