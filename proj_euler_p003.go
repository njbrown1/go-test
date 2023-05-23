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
	fmt.Println("hello world")
	fmt.Println(600851475143 % 2)
	fmt.Println(600851475143 % 3)
	fmt.Println(600851475143 % 5)
	fmt.Println(600851475143 % 7)
	fmt.Println(600851475143 % 11)
	fmt.Println(600851475143 % 13)
	fmt.Println(600851475143 % 17)
	fmt.Println(600851475143 % 19)
	fmt.Println(600851475143 % 23)
	fmt.Println(600851475143 % 29)
	fmt.Println(600851475143 % 33)
	fmt.Println(600851475143 % 37)
}
