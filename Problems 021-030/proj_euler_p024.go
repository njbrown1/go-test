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
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

func main() {

	fmt.Println("hi")
	fmt.Println("factorial of 4:", factorial_of(1))
	// 10! = 3628800
	// "0123456789"
	// len(digits_slice) - 1
	// xth_lexico_permutation := 4
	// subtract 1 to perm to make work?
	// small_digits_slice := []int{0, 1, 2, 3}

	digits_slice := []int{0, 1, 2, 3}
	position := 12
	// remainder := 12

	permutation_index := math.Floor(float64((position - 1) / (factorial_of(len(digits_slice) - 1))))
	fmt.Println(digits_slice, permutation_index)
	removeElement(digits_slice, int(permutation_index))
	fmt.Println(digits_slice, permutation_index)

}
