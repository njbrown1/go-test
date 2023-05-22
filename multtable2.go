package main

import "fmt"

func main() {

	a := 1
	b := 1

	for a <= 10 {
		fmt.Println(a*b, a * (b + 1), a*(b+2), a*(b+3), a*(b+4), a*(b+5), a*(b+6), a*(b+7), a*(b+8), a*(b+9))
		a++
	}
}
