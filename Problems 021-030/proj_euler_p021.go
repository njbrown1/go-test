package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hi")
}

func find_sum_of_divisors_of_n(n int) int {

	var factors_slice []int                                    // initialize an empty slice that factors will be added to
	square_root_of_n := int(math.Floor(math.Sqrt(float64(n)))) // finds square root of n, rounded DOWN

	for i := 1; i <= square_root_of_n; i++ {
		if n%i == 0 {
			if float64(i) == math.Sqrt(float64(n)) { // if i is the square root of n...
				factors_slice = append(factors_slice, i) // append the factor only ONCE
			} else { // otherwise...
				factors_slice = append(factors_slice, i, (n / i)) // append the factor and its factor pair
			}
		}
	}
	sum_of_factors := 0
	for i := range factors_slice {
		sum_of_factors += factors_slice[i]
	}
	return sum_of_factors
}

// evaluate the sum of all the amicable numbers under 10000
