package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 100
	b := 100
	n := 0
	largest_palindrome := 0
	var nString string // defaults to an empty string

	for a <= 300 {
		for b <= 300 {

			n = a * b
			nString = strconv.Itoa(n)
			b++

			if n > 99999 { // if n has 6 digits
				if (nString[0] == nString[5] && nString[1] == nString[4]) && nString[2] == nString[3] { // taking the nth element from the string and comparing it
					fmt.Println("palindrome found!")
					fmt.Printf("a: %d, b: %d, n: %d, length: %d, nString: %s\n", a, b, n, len(nString), nString) // ensuring that the 'for' loop works correctly

					if n > largest_palindrome {
						largest_palindrome = n
						fmt.Println("new LARGEST palindrome found!")
					}
				}

			} else {
				if nString[0] == nString[4] && nString[1] == nString[3] { // if n has 5 digits
					fmt.Println("palindrome found!")
					fmt.Printf("a: %d, b: %d, n: %d, length: %d, nString: %s\n", a, b, n, len(nString), nString) // ensuring that the 'for' loop works correctly

					if n > largest_palindrome {
						largest_palindrome = n
						fmt.Println("new LARGEST palindrome found!")
					}
				}

			}

		}

		b = 100
		a++

	}
}
