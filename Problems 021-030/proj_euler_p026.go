package main

import "fmt"
import "math/big"

func main() {

	decimal_representation := new(big.Float)
	decimal_representation.SetFloat64(1.0 / 3.0)
	fmt.Println(decimal_representation.Text('g', 20))
}
