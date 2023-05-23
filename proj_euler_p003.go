package main

import "fmt"
import "math"

func main() {
	var n int = 24 // where n is the composite number we are testing
	var s = math.Sqrt(float64(n))
	fmt.Println(n)
	fmt.Println(s)
}
