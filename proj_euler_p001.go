package main

import "fmt"

func main() {
	var t = 0
	var n = 0
		for t <= 999 {
			if t%3 == 0 {
				n = n + t
				t++
			} else if t%5 == 0 {
				n = n + t
				t++
			} else {
				t++
			}
		}

	fmt.Println(n)
	// solved. answer: 233168
}
