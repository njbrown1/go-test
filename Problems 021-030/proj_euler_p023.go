package main

import (
	"fmt"
	"math"
	"sort"
)

func find_sum_of_divisors_of_n(n int) int {

	factors_slice := []int{1}                                  // initialize a slice (with 1) that factors will be added to
	square_root_of_n := int(math.Floor(math.Sqrt(float64(n)))) // finds square root of n, rounded DOWN

	for i := 2; i <= square_root_of_n; i++ {
		if n%i == 0 {
			if float64(i) == math.Sqrt(float64(n)) { // if i is the square root of n...
				factors_slice = append(factors_slice, i) // append the factor only ONCE
			} else { // otherwise...
				factors_slice = append(factors_slice, i, (n / i)) // append the factor and its factor pair
			}
		}
	}
	sum_of_factors := 0
	for i := range factors_slice { // add together all of the factors in factors_slice
		sum_of_factors += factors_slice[i]
	}
	return sum_of_factors
}

func main() {
	abundant_numbers := []int{12}
	for number := 13; number <= 28123; number++ { // max: 28123
		sum_of_factors := find_sum_of_divisors_of_n(number)
		if sum_of_factors > number {
			abundant_numbers = append(abundant_numbers, number)
		}
	}
	fmt.Println(abundant_numbers)
	fmt.Println("len:", len(abundant_numbers))

	sum_of_unruly_numbers := 0 // an unruly number is a number that is NOT the sum of any two abundant numbers

	for integer := 1; integer <= 28123; integer++ {
		for _, abundant_number_1 := range abundant_numbers {
			abundant_number_2 := integer - abundant_number_1
			fmt.Println("integer:", integer, "| ab1:", abundant_number_1, "| ab2:", abundant_number_2)
			index := sort.SearchInts(abundant_numbers, abundant_number_2)

			if abundant_numbers[index] == abundant_number_2 { // confirming the SearchInts function found a correct abundant number
				fmt.Println(integer, "is the sum of", abundant_number_1, "and", abundant_number_2)
				break
			}

			if abundant_number_2 < 0 {
				fmt.Println(integer, "is not the sum of any two abundant numbers")
				sum_of_unruly_numbers += integer
				break
			}
		}
	}
	fmt.Println(sum_of_unruly_numbers) // solved! answer: 4179871
}
