package main

import (
	"fmt"
	"math"
)

func main() {
	successful_numbers := []int{} // all numbers that satisfy the condition set in p#030
	for number := 2; number <= 999999; number++ {
		sum_of_fifth_powers := 0
		k := number // so we don't mess with the wider 'for' loop
		for k > 0 {
			digit := (k % 10)
			k -= digit
			k /= 10
			sum_of_fifth_powers += int(math.Pow(float64(digit), 5))
		}
		if number == sum_of_fifth_powers {
			successful_numbers = append(successful_numbers, number)
			fmt.Println(successful_numbers)
		}
	}
	sum_of_successful_numbers := 0
	for _, successful_number := range successful_numbers {
		sum_of_successful_numbers += successful_number
	}
	fmt.Println(successful_numbers, "sum:", sum_of_successful_numbers)
}
