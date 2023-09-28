package main

import (
	"fmt"
	"math/rand"
)

type TheRock struct {
	myScore int
}

func (m *TheRock) GainOnePoint() {
	m.myScore++
}

func addNumbers(x int, y int) int {
	return (x + y)
}

func main() {
	fmt.Println("hello world")
	fmt.Println(rand.Intn(25))
	theRockOne := TheRock{4}
	fmt.Println(theRockOne)
	theRockOne.GainOnePoint()
	fmt.Println(theRockOne)
	theRockOne.myScore = addNumbers(3, 7)
	fmt.Println(theRockOne)
}
