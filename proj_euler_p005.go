package main

import "fmt"

func find_prime_factorization(i int) {

	var prime_slice []int // initialize an empty slice
	prime_slice = append(prime_slice, 2, 3, 5)
	fmt.Println(prime_slice) // testing

	for k := 4; k <= 100; k++ { // for finding prime numbers up to 100

		var is_k_prime string = "yes"

		for l := 0; l < len(prime_slice); l++ {

			if k%prime_slice[l] == 0 {
				is_k_prime = "no" // if k is divisible by one of the prime numbers in the slice, k is not prime.
			}
		}

		if is_k_prime == "yes" {
			prime_slice = append(prime_slice, k) // if k is prime, add it to prime_slice.
		}

	}

	fmt.Println(prime_slice)

}

func main() {
	fmt.Println("hello world")
	find_prime_factorization(4)
}
