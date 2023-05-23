package main

import "fmt"

func main() {
	a := 300
	b := 300
	n := 0
	// c := 0

	fmt.Println("a initial:", a, "\nb initial:", b)
	fmt.Println("a * b initial:", a*b)

	for a <= 302 {
		for b <= 302 {
			n = a * b
			fmt.Printf("a: %d, b: %d, n: %d \n", a, b, n) // ensuring that the 'for' loop works correctly
			b++

			if n > 99999 {
				fmt.Println("6 digits long")
			} else {
				fmt.Println("5 digits long")
			}

		}

		b = 300
		a++

	}
}
