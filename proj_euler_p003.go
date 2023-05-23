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

	j := 0
	var array []int
	array = append(array, 2)
	fmt.Println("array:", array)

	for i := 4; i <= 10; i++ { // creating an array with prime numbers from 1 - 10000 (1, 2, and 3 are not prime)

		j = 0

		for j < len(array) {

			if i%array[j] == 0 {

			}

			j++
		}
	}
}
