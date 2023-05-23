package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 100
	j := 100

	fmt.Println(i, j)

}

func palindromecheck() {
	n := 16889896
	var nString string = strconv.Itoa(n)
	a := len(nString)

	arr := make([]string, len(nString))
	if a == 5 {
		// check to see if e1 = e5, e2 = e4 (e4 means "element 4" within the string)
	} else if a == 6 {
		// check to see if e1 = e6, e2 = e5, and e3 = e4
	}
}
