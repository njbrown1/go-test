package main

import "fmt"

func main() {

	var array []int
	array = append(array, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	array = append(array, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	number := 11
	counter := 0

	counter = 0
	for i := 0; i <= 19; i++ {
		if number%array[i] == 0 {
			counter++
		}
	}
	fmt.Println(counter)
}
