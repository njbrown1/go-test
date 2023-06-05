package main

import "fmt"

func find_length_of_collatz_sequence(number int) int { // this function takes an integer as input, and outputs the number of terms for that integer's Collatz sequence.
	number_of_terms := 1
	for number_equals_1 := false; number_equals_1 != true; number_of_terms++ {
		if number == 1 { // if the number equals 1, break the 'for' loop
			break
		} else if number%2 == 0 { // if n is even: n -> n/2
			number = (number / 2)
		} else if number%2 == 1 { // if n is odd: n -> 3n + 1
			number = (3 * number) + 1
		}
	}
	return number_of_terms
}

func main() {
	longest_number_of_terms := 0
	for i := 1; i <= 999999; i++ {
		number_of_terms_result := find_length_of_collatz_sequence(i)
		if number_of_terms_result > longest_number_of_terms {
			longest_number_of_terms = number_of_terms_result
			fmt.Println(i, "has", number_of_terms_result, "terms in its Collatz sequence")
		}
	}
}
