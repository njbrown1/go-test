package main

import "fmt"

// import "math"

func main() {

	var slice []int
	for i := 1; i <= 20; i++ { // add the numbers 1-20 to the slice
		slice = append(slice, i)
	}

	fmt.Println(slice)

}
