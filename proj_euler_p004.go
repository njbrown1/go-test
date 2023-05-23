package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 123456
	b := [5]int{1, 2, 3, 4, 5}
	aString := strconv.Itoa(a) // initialize aString, which is just a string (not int) equivalent of 'a'

	array := make([]int, 0, len(aString))

	for i := 1; i <= len(aString); i++ {
		array = append(array, 1)
	}

	// for i, ch := range aString {
	// digit, _ := strconv.Atoi(string(ch))
	// array[i] = digit
	//	}

	fmt.Println(array)
	fmt.Println("b:", b)
	fmt.Println(aString)
}
