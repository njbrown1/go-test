package main

import "fmt"

func generate_primes(limit int) []int {

	// this is a classic prime sieve.

	var prime_slice []int
	prime_slice = append(prime_slice, 2, 3, 5) // add the first few primes.

	for k := 4; k <= limit; k++ { // for finding prime numbers up to 'limit'.

		is_k_prime := true

		for l := 0; l < len(prime_slice); l++ { // iterate over all of the items in prime_slice.

			if k%prime_slice[l] == 0 {
				is_k_prime = false // if k is divisible by one of the prime numbers in the slice, k is NOT prime.
			}
		}

		if is_k_prime == true {
			prime_slice = append(prime_slice, k) // if k IS prime, add it to prime_slice.

		}
	}

	return prime_slice

}

func find_prime_factorization() {
	fmt.Println(true)
}

func main() {
	fmt.Println("hello world")
	result := generate_primes(1000)
	fmt.Println(result)
}
