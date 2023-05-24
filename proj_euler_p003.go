package main

import "fmt"

// foundational thoughts:
// we are asked to find the largest PRIME factor of 600851475143, a 12-digit number.
// consider the case of a composite number like 58 – it is divisible by 29 and 2, which
// are both prime. so 29 is the largest prime factor of 58. and 600851475143 is not a
// prime number – because I plugged 600851475143 into the answer box, and it was wrong.

// a composite number (ie. not a prime number) must have at least 2 composite factors.
// simple modulo division by early prime numbers (2, 3, 5, 7...)

func main() {

	var slice []int                // initialize an empty slice
	slice = append(slice, 2, 3, 5) // start with the first few prime numbers
	// current_number := 600851475143
	// largest_prime_number := 0

	for i := 2; i <= 1000; i++ { // creating an array with prime numbers from 1 - 10000 (1 is not a prime)

		for j := 0; j < len(slice); j++ {

			if i%slice[j] != 0 { // if 'i%array[j]' equals 0, that means that the number is NOT prime

			} else {
				break
			}

			if j == (len(slice) - 1) { // if i is not divisible by any of the prime numbers in the slice

				fmt.Println(i, "is prime!")
				slice = append(slice, i)

			}
		}
	}

	fmt.Println("new array:", slice)

}
