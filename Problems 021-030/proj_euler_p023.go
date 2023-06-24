package main

import "math"
import "fmt"

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
	result := find_sum_of_divisors_of_n(12)
	fmt.Println(result)
}
