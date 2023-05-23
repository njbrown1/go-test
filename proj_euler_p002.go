package main

import "fmt"

func main() {
	n := 0 // incremented each time j is even, to find the sum (and answer)
	i := 0
	j := 1
	for j < 4000000 {
		j = i + j
		i = j - i
		// fmt.Println(i,j)
		if j%2 == 0 {
			if j <= 4000000 { // ensuring j is less than or equal to 4 million
				n = n + j
			}
		}
	}
	fmt.Println(n)
	// solved. answer: 4613732
}
