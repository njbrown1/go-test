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

func GenerateRandomNumberFromOneToNInclusive(n int) int {
	return (rand.Intn(n) + 1)
}

func (p *Player) ChooseNumber(numbersChosenLastRound []int) {
	p.numbersChosenLastRound = numbersChosenLastRound
	switch p.behavior {
	case "rock":
		if len(p.numbersIHaveChosenInThePast) == 0 {
			p.myCurrentlyChosenNumber = GenerateRandomNumberFromOneToNInclusive(10)
		} else {
			// do nothing. rock just keeps his number the same for every single round.
		}
	case "opportunist":
		// do stuff
	case "cherry":
		// do stuff
	case "completelyRandom":
		p.myCurrentlyChosenNumber = GenerateRandomNumberFromOneToNInclusive(10)
	}
	p.numbersIHaveChosenInThePast = append(p.numbersIHaveChosenInThePast, p.myCurrentlyChosenNumber)
}

func chooseRandomNumberFromSlice(numbers []int) int {
	rand.Seed(time.Now().UnixNano()) // choose a random seed based on the time
	randomIndex := rand.Intn(len(numbers))
	return numbers[randomIndex]
}

func main() {
	numbersChosen := []int{0, 7}
	allPlayers := []Player{
		{0, []int{}, []int{}, 0, "completelyRandom"},
		{0, []int{}, []int{}, 4, "rock"},
	}
	for i := 1; i <= 10; i++ {
		for player := range allPlayers {
			allPlayers[player].ChooseNumber(numbersChosen)
			fmt.Println(allPlayers[player])
		}
		numbersChosen = []int{}                   // clear the slice
		fmt.Println(numbersChosen, "after clear") // after clear
		for player := range allPlayers {
			numbersChosen = append(numbersChosen, allPlayers[player].myCurrentlyChosenNumber)
			fmt.Println(numbersChosen)
		}

		for player := range allPlayers {
			counter := 0
			for number := range numbersChosen {
				if allPlayers[player].myCurrentlyChosenNumber == numbersChosen[number] {
					counter++
				}
			}
			if counter == 1 {
				allPlayers[player].myScore++
			}
		}
	}
	fmt.Println("end")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
}
