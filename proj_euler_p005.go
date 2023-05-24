package main

import "fmt"

func main() {

	var array []int
	array = append(array, 1, 7, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	number := 4
	counter := 0

	for counter < 12 { // until

		counter = 0

		for i := 0; i <= 11; i++ {

			if number%array[i] == 0 {
				counter++
			}

		}

		if counter < 10 {
			number++
			fmt.Println(number)
		}

	}

	fmt.Println(number)

}
