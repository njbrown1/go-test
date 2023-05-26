package main

import "fmt"

func generate_primes(limit int) []int { // returns a slice with prime numbers from 2 - limit.
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

func find_prime_factorization(limit int, number int) map[int]int { // returns a map with int keys and int values.

	using_prime_slice := generate_primes(limit) // grab a prime_slice from the function generate_primes.

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

func update_current_prime_factor_map(limit int, number int, current_prime_factor_map map[int]int) map[int]int { // returns an updated prime factor map.

	individual_prime_factor_map := find_prime_factorization(limit, number)

	for key, value := range individual_prime_factor_map {

		if current_prime_factor_map[key] == 0 { // if there isn't a value assigned to the key in the current_prime_factor_map
			current_prime_factor_map[key] = individual_prime_factor_map[key]

		} else if individual_prime_factor_map[key] > current_prime_factor_map[key] {
			current_prime_factor_map[key] = individual_prime_factor_map[key]
		}

		fmt.Println("Key:", key, "=>", "Element:", value) // prints

	}

	return current_prime_factor_map

}

func main() {

	fmt.Println("ucpfm output:")

	old_number := 12
	new_number := 13 // update with number x

	result := update_current_prime_factor_map(new_number, new_number, find_prime_factorization(old_number, old_number))
	fmt.Println("new map:", result)

}
