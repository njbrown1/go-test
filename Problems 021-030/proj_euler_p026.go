package main

import "fmt"
import "math"

func perform_long_divison(divisor int, decimal_places int) string { // where the numerator HAS to be 1, and the divisor is d (as specified in p#026).
	remainder := 1
	quotient := "0."
	for i := 1; i <= decimal_places; i++ {
		if remainder == 0 {
			break
		}
		for divisor > remainder {
			remainder *= 10
		}
		digit := math.Floor(float64(remainder) / float64(divisor))
		quotient += fmt.Sprint(digit)
		fmt.Println(remainder, digit, quotient)
		remainder -= int(digit) * divisor
	}
	return quotient
}

func main() {
	d := 260
	result := perform_long_divison(d, 30)
	fmt.Println("1 divided by", d, "is approximately:", result)
}
