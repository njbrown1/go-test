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

	digits_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // also works correctly if you decide to change the # of digits. I used this for testing with the digits 0-3.
	var full_permutation string                         // displays the full permutation
	lexographical_number := 1000000                     // one million
	remainder := (lexographical_number - 1)

	for len(digits_slice) > 0 { // while there are still digits in digits_slice:
		permutation_index := math.Floor(float64(remainder / factorial_of(len(digits_slice)-1)))
		perm_digit := digits_slice[int(permutation_index)]
		full_permutation += fmt.Sprint(perm_digit)                              // concatenate perm_digit to the end of full_permutation
		remainder -= int(permutation_index) * factorial_of(len(digits_slice)-1) // adjust the remainder
		digits_slice = removeElement(digits_slice, int(permutation_index))      // remove the digit that was used (perm_digit) from digits_slice
	}
	fmt.Println("lexicographic permutation of", lexographical_number, "–", full_permutation, "– with digits 0-9")
}
