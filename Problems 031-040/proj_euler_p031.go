package main

import "fmt"

func main() {
	fmt.Println("hi")
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	last_index := coins[len(coins)-1]
	fmt.Println(last_index)
}
