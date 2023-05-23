package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 6898
	b := [5]int{1, 2, 3, 4, 5}
	aString := strconv.Itoa(a) // initialize aString, which is just a string (not int) equivalent of 'a'

	array := make([]int, 0, len(aString))

	for i, ch := range aString {
		// array[i] = int(ch)
		array = append(array, int(ch))
		fmt.Printf("Index: %d, Character: %c\n", i, ch)
	} // this code was added by ChatGPT

	fmt.Println("initial array:", array)
	fmt.Println("b:", b)
	fmt.Println(aString)
}
