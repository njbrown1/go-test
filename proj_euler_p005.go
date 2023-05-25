package main

import "fmt"

// import "math"

func main() {

	current_divisible_number := 12
	var slice []int

	for i := 1; i <= 20; i++ { // add the numbers 1-20 to the slice
		slice = append(slice, i)
	}

	fmt.Println(slice)
	for j := 0; j <= 19; j++ {

		current_divisor := slice[j]

		if current_divisible_number%slice[j] == 0 {
			current_divisor = (current_divisible_number / slice[j])
			fmt.Println("current_divisor:", current_divisor, "current_divisible_number:", current_divisible_number)
		}

	}

}
