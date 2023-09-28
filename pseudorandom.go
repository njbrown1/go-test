package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	myScore                 int
	numbersIHaveChosen      []int
	myCurrentlyChosenNumber int
	behavior                string
}

func (p *Person) GainOnePoint() string {
	p.myScore++
	return ("score successfully incremented by 1")
}

func (p *Person) GenerateInitialNumber() int {
	return (rand.Intn(10) + 1)
}

func (p *Person) ChooseNumber() int {
	switch p.behavior {
	case "rock":
		fmt.Println("I'm a rock")
	case "opportunist":
		// do stuff
	case "cherry":
		// do stuff
	case "completelyRandom":
		// do stuff
	}
	return 4
}

func chooseRandomNumberFromSlice(numbers []int) int {
	rand.Seed(time.Now().UnixNano()) // choose a random seed based on the time
	randomIndex := rand.Intn(len(numbers))
	return numbers[randomIndex]
}

func main() {

	// global variables: numOfRoundsPerSimulation, numOfSimulations,
	numbers := []int{0, 5, 3, 2, 1, 7, 8}
	for i := 1; i <= 10; i++ {
		fmt.Println(chooseRandomNumberFromSlice(numbers))
	}
}
