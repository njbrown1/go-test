package main

import "fmt"

func find_length_of_collatz_sequence(number int) int {
	number_of_terms := 1
	for number_equals_1 := false; number_equals_1 != true; number_of_terms++ {
		fmt.Println(number) // for debugging
		if number == 1 {    // if the number equals 1, break the for loop
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
	fmt.Println("hello world")
	result := find_length_of_collatz_sequence(5)
	fmt.Println("number of terms:", result)
}
