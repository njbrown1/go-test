package main

import "fmt"
import "math"

func main() {

	remainder := 1
	divisor := 7
	quotient := "."
	for i := 1; i <= 99; i++ {
		if remainder == 0 {
			break
		}
		for divisor > remainder {
			remainder *= 10
		}
		digit_or_digits := math.Floor(float64(remainder) / float64(divisor))
		quotient += fmt.Sprint(digit_or_digits)
		fmt.Println(remainder, i, divisor, quotient)
		remainder -= int(digit_or_digits) * divisor
	}

}
