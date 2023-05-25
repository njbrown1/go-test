package main

import "fmt"

// import "math"

func find_multiplier(i int) {
	for a := i; a > 0; a-- {

	}

}

func main() {

	var slice []int
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)

	// current_divisible_number := 6

	for i := 4; i <= 20; i++ { // for the numbers 4-x (here, 20)

		slice = append(slice, i) // add i to the slice
	}
}
