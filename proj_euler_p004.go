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
			fmt.Println(a, ",", b, ":", n)
			b++
		}

		b = 100
		a++

	}
}
