package main

import (
	"fmt"
	"math"
	// "sort"
)

func main() {
	results_slice := []int{}
	for a := 2; a <= 5; a++ {
		for b := 2; b <= 5; b++ {
			result := math.Pow(float64(a), float64(b))
			fmt.Println(a, b, result)

			add_result := true
			for _, entry := range results_slice {
				if entry == int(result) {
					add_result = false
				}
			}
			if add_result == true {
				results_slice = append(results_slice, int(result))
			}
		}
	}
	fmt.Println(results_slice, len(results_slice))
}
