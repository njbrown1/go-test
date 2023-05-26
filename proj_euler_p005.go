package main

import "fmt"

func find_prime_factorization(i int) {

	// BEGINNING of prime sieve

	var prime_slice []int                      // initialize an empty slice
	prime_slice = append(prime_slice, 2, 3, 5) // add the first few primes

	for k := 4; k <= 100; k++ { // for finding prime numbers up to 100

		var is_k_prime string = "yes"

		for l := 0; l < len(prime_slice); l++ { // iterate over all of the items in prime_slice.

			if k%prime_slice[l] == 0 {
				is_k_prime = "no" // if k is divisible by one of the prime numbers in the slice, k is NOT prime.
			}
		}

		if is_k_prime == "yes" {
			prime_slice = append(prime_slice, k) // if k IS prime, add it to prime_slice.
		}
	}

	fmt.Println(prime_slice)

	// END of prime sieve

	storage := make(map[int]int)
	storage[1] = 2
	storage[2] = 3

	fmt.Println(storage[2])

}

func main() {
	fmt.Println("hello world")
	find_prime_factorization(4)
}
