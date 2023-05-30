package main

import "fmt"

func main() {
	sum_of_primes := 0
	prime_slice := generate_primes(2000000)

	for i := 0; i < len(prime_slice); i++ {
		sum_of_primes = sum_of_primes + prime_slice[i]
		fmt.Println("sum_of_primes:", sum_of_primes, "prime_slice[i]", prime_slice[i])
	}
	fmt.Println(sum_of_primes) // took about 2 min, but got the right answer of 142913828922. wonder if there's a faster way to do this
}
