package main

import "fmt"

func generate_primes(limit int) []int {

	// this is a classic prime sieve that will generate primes from 2 – limit.

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

func find_prime_factorization(limit int, number int) map[int]int { // returns a map with int keys and values.

	using_prime_slice := generate_primes(limit)

	primeFactorMap := make(map[int]int) // generate a new map with BOTH keys and values as integers.

	for i := 0; i < len(using_prime_slice); i++ { // iterate over all the primes
		var prime int = using_prime_slice[i] // define the prime
		if number%prime == 0 {
			primeFactorMap[prime]++
			number = number / prime
			i-- // check the prime again
		}
	}

	return primeFactorMap

}

func update_current_prime_factor_map(limit int, number int, current_prime_factor_map map[int]int) map[int]int {

	individual_prime_factor_map := find_prime_factorization(limit, number)

	return individual_prime_factor_map

}

func main() {

	fmt.Println("gp output:")
	result := update_current_prime_factor_map(1500, 75, find_prime_factorization(1500, 75))
	fmt.Println(result)

}
