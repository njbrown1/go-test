package main

import "fmt"
import "math"

func factorial_of(number int) int {
	multiplication_total := 1
	for i := 2; i <= number; i++ {
		multiplication_total *= i
	}
	return multiplication_total
}

func main() {
	fmt.Println("hi")
	// 10! = 3628800
	// "0123456789"

	// xth_lexico_permutation := 4

	whats_left := 1000000

	first_digit := math.Floor(float64(whats_left / factorial_of(9)))
	whats_left -= int(first_digit) * factorial_of(9)
	fmt.Println(whats_left, first_digit, factorial_of(9))

	second_digit := math.Floor(float64(whats_left / factorial_of(8)))
	whats_left -= int(second_digit) * factorial_of(8)
	fmt.Println(whats_left, second_digit, factorial_of(8))
}
