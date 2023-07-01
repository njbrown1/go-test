package main

import "fmt"
import "math"

func is_n_prime(n int) bool {
	square_root_of_n := int(math.Floor(math.Sqrt(float64(n)))) // finds square root of n, rounded DOWN
	for i := 1; i <= square_root_of_n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("hi")
	// a and b must both be odd

	continue_running_until_non_prime_found := true
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			number_of_primes_found_before_failure := 0
			n := 0
			for continue_running_until_non_prime_found == true {
				evaluation := (n * n) + (a * n) + b
				if find_sum_of_divisors_of_n_v027(evaluation) != 1 {
					fmt.Println(a, b, n, number_of_primes_found_before_failure)
					break
				} else {
					number_of_primes_found_before_failure++
				}
				n++
			}
		}
	}
}
