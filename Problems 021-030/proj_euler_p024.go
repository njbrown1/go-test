package main

import (
	"fmt"
	"math"
)

func factorial_of(number int) int {
	multiplication_total := 1
	for i := 2; i <= number; i++ {
		multiplication_total *= i
	}
	return multiplication_total
}

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {

	fmt.Println("hi")
	fmt.Println("factorial of 4:", factorial_of(1))
	// 10! = 3628800
	// "0123456789"
	// xth_lexico_permutation := 4
	// subtract 1 to perm to make work?
	// small_digits_slice := []int{0, 1, 2, 3}

	digits_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := removeElement(digits_slice, 4)
	fmt.Println(result)

	position := 1
	first_digit := math.Floor(float64(position-1) / float64(factorial_of(3)))
	fmt.Println(first_digit)

}
