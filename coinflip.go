package main

import (
	"fmt"
	"math/rand"
)

func main() {
	total := 0
	for k := 0; k <= 100; k++ {
		flip := rand.Intn(2) // will psuedorandomly generate either a 0 or a 1. 0 = tails, and 1 = heads
		adjust_total := -1   // set adjust_total_by as a default -1 (tails)
		if flip == 1 {
			adjust_total = 1
		}
		fmt.Print(flip)
		total += adjust_total
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println(total)
}
