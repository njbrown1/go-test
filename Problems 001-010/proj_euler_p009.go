package main

import "fmt"
import "math"

func main() {

	for a := 0; a <= 500; a++ {
		for b := 0; b <= 500; b++ {
			c_float := math.Sqrt(float64(a*a + b*b))
			c := int(math.Round(c_float))

			sum_is_1k := a+b+c == 1000
			is_pythag := a*a+b*b == c*c
			in_order := a < b && b < c
			if sum_is_1k && is_pythag && in_order {
				successful_slice := []int{a, b, c}
				fmt.Println(successful_slice, (a * b * c))
			}
		}
	}
}
