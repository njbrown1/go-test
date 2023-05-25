package main

import "fmt"

func find_multiplier(i int) {

	current_multiplier := i

	for a := (i - 1); a > 0; a-- {

		if current_multiplier%a == 0 {
			current_multiplier = (current_multiplier / a) // if current_multiplier is divisible by a, divide the current multiplier by a
		}

	}

	fmt.Println("current_multiplier:", current_multiplier, "| for i:", i)

}

func main() {

	for j := 2; j <= 20; j++ {
		find_multiplier(j)
	}
}
