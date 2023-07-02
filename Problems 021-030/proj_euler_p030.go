package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pow(2, 5))

	k := 28935

	sum_of_fifth_powers := 0
	for k > 0 {
		digit := (k % 10)
		k -= digit
		k /= 10
		sum_of_fifth_powers += int(math.Pow(float64(digit), 5))
	}
	fmt.Println(sum_of_fifth_powers)

	for number := 2; number <= 999999; number++ {

	}
}
