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

	for a <= 330 {
		for b <= 330 {

			n = a * b
			nString = strconv.Itoa(n)
			b++

			if n > 99999 {
				if (nString[0] == nString[5] && nString[1] == nString[4]) && nString[2] == nString[3] {
					fmt.Println("palindrome found!")
					fmt.Printf("a: %d, b: %d, n: %d, length: %d, nString: %s\n", a, b, n, len(nString), nString) // ensuring that the 'for' loop works correctly
				}
			} else {
				if nString[0] == nString[4] && nString[1] == nString[3] {
					fmt.Println("palindrome found!")
					fmt.Printf("a: %d, b: %d, n: %d, length: %d, nString: %s\n", a, b, n, len(nString), nString) // ensuring that the 'for' loop works correctly
				}
			}

		}

		b = 300
		a++

	}
}
