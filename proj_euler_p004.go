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

			if ((n) / (10 ^ 4)) >= 10 {
				fmt.Println("5 digits long")
			} else if ((n) / (10 ^ 4)) >= 1 {
				fmt.Println("6 digits long")
			}

		}

		b = 100
		a++

	}
}
