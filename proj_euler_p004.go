package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 300
	b := 300
	n := 0
	var nString string // defaults to an empty string
	// c := 0

	fmt.Println("a initial:", a, "\nb initial:", b)
	fmt.Println("a * b initial:", a*b)

	for a <= 303 {
		for b <= 303 {
			n = a * b

			nString = strconv.Itoa(n)
			fmt.Println(nString)
			fmt.Printf("a: %d, b: %d, n: %d \n", a, b, n) // ensuring that the 'for' loop works correctly

			b++

			if n > 99999 {
				fmt.Println("6 digits long")
				if nString[1] == nString[6] {
					fmt.Println("kind of a palindrome")
				}
			} else {
				fmt.Println("5 digits long")
				if nString[0] == nString[4] {
					fmt.Println("kind of a palindrome")
				}
			}

		}

		b = 300
		a++

	}
}
