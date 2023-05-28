package main

import "fmt"
import "math"

func main() {

	for a := 0; a <= 500; a++ {

		for b := 0; b <= 500; b++ {

			c_float := math.Sqrt(float64(math.Pow(float64(a), 2) + math.Pow(float64(b), 2)))
			c_integer := int(math.Round(c_float))

			if (a+b+int(c_integer)) == 1000 && ((a*a)+(b*b) == (c_integer)*(c_integer)) { // if a + b + c = 1000 and a ^ 2 + b ^ 2 = c ^2

				if (a < b) && (b < int(c_integer)) { // if a < b < c
					successful_slice := []int{a, b, int(c_integer)}
					fmt.Println(successful_slice, (a * b * int(c_integer)))

				}

			}
		}

	}

}
