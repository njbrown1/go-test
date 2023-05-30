package main

import "fmt"
import "math"

func main() {

	fmt.Println(find_x_divisors_of_n(150))

	// over_500_divisors := false
	// for number := 1; over_500_divisors == true; number++ {
	// 	num_of_factors := find_x_divisors_of_n(number)
	//	fmt.Println("there are", num_of_factors, "factors of the number", number)
	//
	//	if num_of_factors > 50 { // if the number of factors exceeds 50, end the loop
	//		over_500_divisors = true
	//		fmt.Println(number)
	//	}
	// }
}

func find_x_divisors_of_n(n int) int {

	var factors_slice []int
	square_root_of_n := int(math.Floor(math.Sqrt(float64(n)))) // finds square root of n, rounded up

	for i := 1; i <= square_root_of_n; i++ {

		if n%i == 0 {

			is_i_the_square_root_of_n := float64(i) == math.Sqrt(float64(n))

			if is_i_the_square_root_of_n == true {
				factors_slice = append(factors_slice, i) // append the factor only ONCE
				fmt.Println(factors_slice)               // for debugging
			} else {
				factors_slice = append(factors_slice, i, (n / i)) // append the factor and its factor pair
				fmt.Println(factors_slice)                        // for debugging
			}
		}
	}
	return len(factors_slice)
}
