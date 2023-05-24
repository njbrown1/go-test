package main

import "fmt"

import "math"

func main() {

	sum_of_squares := 0 // sum of 1^2 + 2^2 + 3^2 ... + 100^2
	sum := 0            // sum of (1 + 2 + 3 + ... + 100)
	square_of_sum := 0  // sum of (1 + 2 + 3 + ... + 100)^2
	fmt.Println("hello world")

	for i := 1; i <= 100; i++ {
		sum_of_squares = sum_of_squares + (i ^ 2) // NEED TO FIX
		sum = sum + (i)
		fmt.Println("i:", i, "sum of squares:", sum_of_squares, "sum:", sum)
	}

	square_of_sum = (sum ^ 2)

	fmt.Println(square_of_sum)
	fmt.Println(sum_of_squares)
	fmt.Println(6 ^ 2)
	fmt.Println("The difference between the sum of the squares and the square of the sum is:", (sum_of_squares - square_of_sum))
}
