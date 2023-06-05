package main

import (
	"fmt"
)

func factorialOfNumber(number int64) string {
	k := int64(1)
	for i := int64(1); i <= number; i++ {
		k *= i
	}
	return fmt.Sprint(k)
}

func main() {
	result := factorialOfNumber(40)
	fmt.Println(result)
}
