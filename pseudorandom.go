package main

import (
	"fmt"
	"math/rand"
)

func chooseRandomNumberFromSlice(numbers []int) int {
	return 4
}

// type 1. always chooses the same number every time

type TheRock struct {
	myScore          int
	lastNumberIChose int
}

func (r *TheRock) GainOnePoint() {
	r.myScore++
}

func (r *TheRock) ChooseNumber() int {
	return 4
}

func addNumbers(x int, y int) int {
	return (x + y)
}

func main() {
	fmt.Println("hello world")
	fmt.Println(rand.Intn(25))
	theRockOne := TheRock{4, 0}
	fmt.Println(theRockOne)
	theRockOne.GainOnePoint()
	fmt.Println(theRockOne)
	theRockOne.myScore = addNumbers(3, 7)
	fmt.Println(theRockOne)
}
