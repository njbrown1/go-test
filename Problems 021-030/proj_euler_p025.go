package main

import "fmt"
import "math/big"

func main() {
	a := big.NewInt(1)
	b := big.NewInt(1)
	storage := big.NewInt(1)

	for i := 1; len(b.String()) < 1000; i++ {
		storage.Add(a, b)
		a.Set(b)
		b.Set(storage)
		fmt.Println("a:", a, "| b:", b, "| i:", i, "| len of b:", len(b.String())) // final answer is 4782 = i + 2 = 4780 + 2. it's just an adjustment for the indexing.
	}
}
