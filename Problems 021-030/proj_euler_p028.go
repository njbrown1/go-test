package main

import "fmt"

func main() {

	spiral_side_length := 5
	increment := 2
	sum_of_diagonals := 0

	for n := 1; n <= (spiral_side_length * spiral_side_length); {
		for k := 1; k <= 4; k++ {
			n += increment
			sum_of_diagonals += n
			fmt.Println("after:", sum_of_diagonals, n, increment)
		}
		increment += 2
	}

	fmt.Println(sum_of_diagonals)
}
