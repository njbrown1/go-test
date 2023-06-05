package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
)

func factorialOfNumber(number int64) int64 {
	k := int64(1)
	for i := int64(1); i <= number; i++ {
		k *= i
	}
	return k
}

func main() {

	// funnily enough, I came up with this problem myself! I remembered some vague answer / way to use combinatorics
	// and after some pain importing the combin package, I figured out that on an LxL board, running the combination
	// nCr (where) n = 2L and r = L gives the correct answer experimentally. there's probably a better way to do this,
	// but I cannot be bothered. I made a whole Math StackExchange post on it too.

	L := 10
	n := 2 * L
	r := L

	comb := combin.Combinations(n, r)
	fmt.Println("there are", len(comb), "paths for an LxL board, where L =", L)

	// OMG it killed the signal when it tried to calculate all of the combinations for 20 C 10. fine, I'll do the
	// calculation myself. I don't even need the combin import. I'll make my own factorial function.

	number_of_valid_paths := int64(factorialOfNumber(40))
	fmt.Println(int64(number_of_valid_paths))
}
