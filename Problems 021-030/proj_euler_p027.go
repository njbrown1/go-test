package main

import (
	"fmt"
	"math"
)

func generate_primes_v027(limit int) []int { // returns a slice with prime numbers from 2 - limit.
	prime_slice := []int{2, 3, 5}
	for i := 4; i <= limit; i++ { // for finding prime numbers up to 'limit'.
		is_i_prime := true
		for _, prime := range prime_slice {
			if i%prime == 0 {
				is_i_prime = false
			}
		}
		if is_i_prime == true {
			prime_slice = append(prime_slice, i) // if k IS prime, add it to prime_slice.
		}
	}
	return prime_slice
}

func is_n_prime(n int) bool {
	square_root_of_n := math.Ceil(math.Sqrt(float64(n)))
	prime_slice := generate_primes_v027(int(square_root_of_n))
	for prime := range prime_slice {
		if n%prime == 0 {
			return false
		}
	}
	return true
}

func main() {

	evaluation := 42
	fmt.Println(generate_primes_v027(100))

	for a := 0; a <= 1; a++ {
		for b := 0; b <= 1; b++ {
			fmt.Println(a, b)
		}
	}
}
