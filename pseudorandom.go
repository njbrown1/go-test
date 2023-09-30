package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	myScore                     int
	numbersIHaveChosenInThePast []int
	myCurrentlyChosenNumber     int
	behavior                    string
}

func generateRandomNumberFromOneToNInclusive(n int) int {
	return (rand.Intn(n) + 1)
}

func chooseRandomNumberFromSlice(numbers []int) int {
	rand.Seed(time.Now().UnixNano()) // choose a random seed based on the time
	randomIndex := rand.Intn(len(numbers))
	return numbers[randomIndex]
}

func findMostCommonNumbers(nums []int) []int {
	countMap := make(map[int]int)
	mostCommon := []int{}
	maxCount := 0

	for _, num := range nums {
		countMap[num]++
		if countMap[num] > maxCount {
			maxCount = countMap[num]
		}
	}

	for num, count := range countMap {
		if count == maxCount {
			mostCommon = append(mostCommon, num)
		}
	}
	return mostCommon
}

func findLeastCommonNumbers(nums []int) []int {
	countMap := make(map[int]int)
	leastCommonNumbers := []int{}
	minCount := len(nums) + 1

	for _, num := range nums {
		countMap[num]++
	}

	for num := 1; num <= 10; num++ {
		if countMap[num] < minCount {
			minCount = countMap[num]
		}
	}

	for num := 1; num <= 10; num++ {
		if countMap[num] == minCount {
			leastCommonNumbers = append(leastCommonNumbers, num)
		}
	}
	return leastCommonNumbers
}

func (p *Player) chooseNumber(numbersChosenLastRound []int) {
	switch p.behavior {
	case "rock":
		if len(p.numbersIHaveChosenInThePast) == 0 {
			p.myCurrentlyChosenNumber = generateRandomNumberFromOneToNInclusive(10)
		}
	case "opportunist":
		leastChosenNumbers := findLeastCommonNumbers(numbersChosenLastRound)
		p.myCurrentlyChosenNumber = chooseRandomNumberFromSlice(leastChosenNumbers)
	case "realEstateAgent":
		mostChosenNumbers := findMostCommonNumbers(numbersChosenLastRound)
		p.myCurrentlyChosenNumber = chooseRandomNumberFromSlice(mostChosenNumbers)
	case "completelyRandom":
		p.myCurrentlyChosenNumber = generateRandomNumberFromOneToNInclusive(10)
	}
	p.numbersIHaveChosenInThePast = append(p.numbersIHaveChosenInThePast, p.myCurrentlyChosenNumber) // add the number chosen to the history
}

func (p *Player) UpdateScore(numbersChosenThisRound []int) {
	counter := 0
	fmt.Println(p.myCurrentlyChosenNumber)
	for number := range numbersChosenThisRound {
		if p.myCurrentlyChosenNumber == number {
			counter++
			fmt.Println(number, counter)
		}
	}
	if counter == 1 { // if the number only occurs once in the numbersChosenThisRound (ie. the player chose a unique number)
		p.myScore++
	}
}

func main() {
	numbersChosenLastRound := []int{}
	numbersChosenThisRound := []int{}
	allPlayers := []Player{
		{0, []int{}, 0, "completelyRandom"},
		{0, []int{}, 0, "rock"},
	}
	for _, player := range allPlayers {
		player.chooseNumber(numbersChosenLastRound)
		fmt.Println(player, "first")
		numbersChosenThisRound = append(numbersChosenThisRound, player.myCurrentlyChosenNumber)
	}
	fmt.Println(numbersChosenLastRound)
	fmt.Println(numbersChosenThisRound)
	fmt.Println(allPlayers, "after numbersChosenThisRound")
	for _, player := range allPlayers {
		player.UpdateScore(numbersChosenThisRound)
	}
	fmt.Println(allPlayers)
}
