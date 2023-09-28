package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	myScore                     int
	numbersIHaveChosenInThePast []int
	numbersFromLastRound        []int
	myCurrentlyChosenNumber     int
	behavior                    string
}

func (p *Person) GainOnePoint() {
	p.myScore++
}

func (p *Person) ChooseNumberRandomly() {
	randomNumber := rand.Intn(10) + 1 // generates an integer from 1-10
	p.myCurrentlyChosenNumber = randomNumber
}

func (p *Person) ChooseNumber() {
	switch p.behavior {
	case "rock":
		if p.myCurrentlyChosenNumber == 0 {
			fmt.Println("WHOOPS rock forgot to choose an initial number with ChooseInitialNumberRandomly")
		} else {
			p.numbersIHaveChosenInThePast = append(p.numbersIHaveChosenInThePast, p.myCurrentlyChosenNumber)
		}
		// else, do nothing. rock just keeps his number the whole time.
	case "opportunist":
		// do stuff
	case "cherry":
		// do stuff
	case "completelyRandom":
		p.ChooseNumberRandomly()
		p.numbersIHaveChosenInThePast = append(p.numbersIHaveChosenInThePast, p.myCurrentlyChosenNumber)
	}
}

func chooseRandomNumberFromSlice(numbers []int) int {
	rand.Seed(time.Now().UnixNano()) // choose a random seed based on the time
	randomIndex := rand.Intn(len(numbers))
	return numbers[randomIndex]
}

func main() {
	Rock := Person{0, []int{}, []int{}, 0, "rock"}
	fmt.Println(Rock)
	Rock.ChooseNumberRandomly()
	fmt.Println(Rock)
	Rock.GainOnePoint()
	Rock.ChooseNumber()
	Rock.ChooseNumber()
	Rock.ChooseNumber()
	fmt.Println(Rock)
	Cherry := Person{0, []int{}, []int{}, 0, "completelyRandom"}
	fmt.Println(Cherry)
	Cherry.ChooseNumber()
	Cherry.ChooseNumber()
	Cherry.ChooseNumber()
	Cherry.ChooseNumber()
	Cherry.ChooseNumber()
	Cherry.ChooseNumber()
	fmt.Println(Cherry)
}
