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

	for a <= 999 {
		for b <= 999 {

			n = a * b
			nString = strconv.Itoa(n) // convert n to a string, 'nString''
			b++

			if n > 99999 { // if n has 6 digits

				if (nString[0] == nString[5] && nString[1] == nString[4]) && nString[2] == nString[3] { // taking the nth element from the string and running comparisons

					if n > largest_palindrome { // if n is larger than the largest palindrome, make n the new largest palindrome
						largest_palindrome = n
						fmt.Println("new LARGEST palindrome found!")
						fmt.Printf("a: %d, b: %d, n: %d, length: %d, nString: %s\n", a, b, n, len(nString), nString) // showing a, b, n, and the new largest palindrome
					}

				}

			} else if n > 9999 { // if n has 5 digits

				if nString[0] == nString[4] && nString[1] == nString[3] {

					if n > largest_palindrome { // if n is larger than the largest palindrome, make n the new largest palindrome
						largest_palindrome = n
						fmt.Println("new LARGEST palindrome found!")
						fmt.Printf("a: %d, b: %d, n: %d, length: %d, nString: %s\n", a, b, n, len(nString), nString) // showing a, b, n, and the new largest palindrome
					}
				}

			}

		}

		b = 100
		a++

	}

	fmt.Println("The largest palindrome attainable by multiplying two 3-digit integers together is:", largest_palindrome)
	// solved. answer is: 906609
}
