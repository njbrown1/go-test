package main

import "fmt"

func main() {
	a := 100
	b := 100
	n := 0

	fmt.Println("a initial:", a, "\nb initial:", b)
	fmt.Println("a * b initial:", a*b)

	for a <= 104 {
		for b <= 104 {
			n = a * b
			fmt.Printf("a: %d, b: %d, n: %d \n", a, b, n) // ensuring that the 'for' loop works correctly
			b++

			if n%10 != 0 {

			}
		}

		b = 100
		a++

	}
}

func palindrome_check() {
	d1 := 7
	d2 := 6
	d3 := 5
	d4 := 4
	d5 := 3
	d6 := n % 10
	fmt.Println(d1, d2, d3, d4, d5, d6) // doesn't work correctly, since palindrome_check() doesn't know what n is.
}
