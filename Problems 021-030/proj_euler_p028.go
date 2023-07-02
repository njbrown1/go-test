package main

import "fmt"

func main() {
	spiral_side_length := 1001
	increment := 2
	sum_of_diagonals := 1
	for n := 1; n < (spiral_side_length * spiral_side_length); {
		for k := 1; k <= 4; k++ { // k has no significance here, I just want it to repeat 4 times
			n += increment
			sum_of_diagonals += n
		}
		increment += 2
	}
	fmt.Println(sum_of_diagonals)
}
