package main

import "fmt"
import "math"

func main() {
	var n int = 24 // where n is the composite number we are testing
	var s = math.Round(math.Sqrt(float64(n))) // where s is the square root of n, rounded up
	fmt.Println(n)
	fmt.Println(s)
}
