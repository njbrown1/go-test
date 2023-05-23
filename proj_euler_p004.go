package main

import "fmt"
import "strconv"

func main() {
	i := 100
	j := 100

	for i <= 102 {
		i++
		j++
		fmt.Println(i*j)
	}

	n := 16889896
	var nString string = strconv.Itoa(n)
	a := len(nString)
	fmt.Println(a)
	fmt.Println(n)

}
