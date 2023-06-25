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

	digits_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var full_permutation string
	lexographical_number := 1000000
	remainder := (lexographical_number - 1)

	for len(digits_slice) > 0 {
		permutation_index := math.Floor(float64(remainder / factorial_of(len(digits_slice)-1)))
		perm_digit := digits_slice[int(permutation_index)]
		full_permutation += fmt.Sprint(perm_digit)
		remainder -= int(permutation_index) * factorial_of(len(digits_slice)-1)
		digits_slice = removeElement(digits_slice, int(permutation_index))
	}

	fmt.Println("lexicographic permutation of", lexographical_number, "–", full_permutation)

}
