package main

import (
	"fmt"
	"math"
)

func is_number_contained_in_amicable_numbers(amicable_numbers []int, number int) bool {
	for _, i := range amicable_numbers {
		if i == number {
			return true
		}
	}
	return false
}

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
	var amicable_numbers []int
	for number := 1; number <= 10000; number++ {
		if is_number_contained_in_amicable_numbers(amicable_numbers, number) == false { // only check the number if it's not in amicable_numbers already
			a := find_sum_of_divisors_of_n(number)
			if number == find_sum_of_divisors_of_n(a) && number != a {
				amicable_numbers = append(amicable_numbers, a, number)
			}
		}
	}
	sum_of_amicable_numbers := 0
	for i := range amicable_numbers {
		sum_of_amicable_numbers += amicable_numbers[i]
	}
	fmt.Println(amicable_numbers)
	fmt.Println(sum_of_amicable_numbers) // solved! answer: 31626
}

// evaluate the sum of all the amicable numbers under 10000.
// amicable – where d(a) = b, d(b) = a, and a != b
// d(x) – sum of proper divisors of x
