package main

import "fmt"

func generate_prime_slice(limit int) []int { // returns a slice with prime numbers from 2 - limit. function copied from problem 005.

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

func main() {

	prime_slice := generate_primes(120000)         // quick and dirty way. I called generate_prime_slice(100000) and error was
	fmt.Println(prime_slice[10001-1])              // 'panic: runtime error: index out of range [10000] with length 9592'. I figured
	fmt.Println("guess it worked")                 // letting the function generate 20,000 more numbers would result in at least 409
	fmt.Println("the answer:", prime_slice[10000]) // more primes, so I called generate_prime_slice(120000) and it worked!
}
