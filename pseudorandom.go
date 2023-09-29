package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	myScore                     int
	numbersIHaveChosenInThePast []int
	numbersChosenLastRound      []int
	myCurrentlyChosenNumber     int
	behavior                    string
}

func (p *Player) SimulateRound(numbersChosenLastRound []int) {
	p.numbersChosenLastRound = numbersChosenLastRound
	p.ChooseNumber()
	fmt.Println(p)
}

func (p *Player) ChooseNumber() {
	switch p.behavior {
	case "rock":
		if p.myCurrentlyChosenNumber == 0 {
			fmt.Println("WHOOPS rock forgot to choose an initial number with ChooseInitialNumberRandomly")
		}
		// else, do nothing. rock just keeps his number for every single round.
	case "opportunist":
		// do stuff
	case "cherry":
		// do stuff
	case "completelyRandom":
		p.myCurrentlyChosenNumber = rand.Intn(10) + 1 // generates an integer from 1-10
	}
	p.numbersIHaveChosenInThePast = append(p.numbersIHaveChosenInThePast, p.myCurrentlyChosenNumber)
}

func chooseRandomNumberFromSlice(numbers []int) int {
	rand.Seed(time.Now().UnixNano()) // choose a random seed based on the time
	randomIndex := rand.Intn(len(numbers))
	return numbers[randomIndex]
}

func main() {
	numbersChosenLastRound := []int{0, 7}
	allPlayers := []Player{
		{0, []int{}, []int{}, 0, "completelyRandom"},
		{0, []int{}, []int{}, 4, "rock"},
	}
	for player := range allPlayers {
		allPlayers[player].SimulateRound(numbersChosenLastRound)
	}
}
