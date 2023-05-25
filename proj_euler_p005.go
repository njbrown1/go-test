package main

import "fmt"

// import "math"

func main() {

	var slice []int
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)

	// current_divisible_number := 6

	for i := 4; i <= 20; i++ { // for the numbers 4-x (here, 20)

		slice = append(slice, i) // add i to the slice
		current_multiplier := i

		fmt.Println(i)

		for a := 0; a < len(slice); a++ {

			if current_multiplier%slice[a] == 0 {
				current_multiplier = (current_multiplier / slice[a])
				fmt.Println("i:", i, "current_multiplier:", current_multiplier)
			}
		}
	}
}
