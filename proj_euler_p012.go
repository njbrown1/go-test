package main

import "fmt"
import "math"

func main() {

	// over_500_divisors := false
	for number := 1; number <= 100000000; number++ { // 100000000 is just an obscenely large number
		num_of_factors := find_x_divisors_of_n(number)
		fmt.Println("there are", num_of_factors, "factors of the number", number)

		if num_of_factors > 100 { // if the number of factors exceeds 50, end the loop
			fmt.Println("The first regular number with over 500 factors is:", number)
			break
		}
	}
}

func find_x_divisors_of_n(n int) int {

	var factors_slice []int
	square_root_of_n := int(math.Floor(math.Sqrt(float64(n)))) // finds square root of n, rounded up

	for i := 1; i <= square_root_of_n; i++ {

		if n%i == 0 {

			is_i_the_square_root_of_n := float64(i) == math.Sqrt(float64(n))

			if is_i_the_square_root_of_n == true {
				factors_slice = append(factors_slice, i) // append the factor only ONCE
			} else {
				factors_slice = append(factors_slice, i, (n / i)) // append the factor and its factor pair
			}
		}
	}
	return len(factors_slice)
}
