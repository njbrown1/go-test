package main

import "fmt"

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

func main() {
	fmt.Println("hi")
	// a and b must both be odd

	fmt.Println(generate_primes_v027(100))

	for a := 0; a <= 1; a++ {
		for b := 0; b <= 1; b++ {
			fmt.Println(a, b)
		}
	}
}
