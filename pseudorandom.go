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

func (p *Player) ChooseNumber(numbersChosenLastRound []int) {
	p.numbersChosenLastRound = numbersChosenLastRound
	switch p.behavior {
	case "rock":
		if len(p.numbersIHaveChosenInThePast) == 0 {
			p.myCurrentlyChosenNumber = GenerateRandomNumberFromOneToNInclusive(10)
		}
	case "opportunist":
		leastChosenNumbers := findLeastCommonNumbers(p.numbersChosenLastRound)
		p.myCurrentlyChosenNumber = chooseRandomNumberFromSlice(leastChosenNumbers)
	case "realEstateAgent":
		mostChosenNumbers := findMostCommonNumbers(p.numbersChosenLastRound)
		p.myCurrentlyChosenNumber = chooseRandomNumberFromSlice(mostChosenNumbers)
	case "completelyRandom":
		p.myCurrentlyChosenNumber = GenerateRandomNumberFromOneToNInclusive(10)
	}
	p.numbersIHaveChosenInThePast = append(p.numbersIHaveChosenInThePast, p.myCurrentlyChosenNumber) // add the number chosen to the history
}

func chooseRandomNumberFromSlice(numbers []int) int {
	rand.Seed(time.Now().UnixNano()) // choose a random seed based on the time
	randomIndex := rand.Intn(len(numbers))
	return numbers[randomIndex]
}

func main() {
	numbersChosenLastRound := []int{}
	allPlayers := []Player{
		{0, []int{}, []int{}, 0, "completelyRandom"},
		{0, []int{}, []int{}, 0, "rock"},
	}
	for _, player := range allPlayers {
		player.ChooseNumber(numbersChosenLastRound)
		fmt.Println(player, "first")
	}
	numbersChosenLastRound = []int{} // clear the slice
	for _, player := range allPlayers {
		fmt.Println(player)
		numbersChosenLastRound = append(numbersChosenLastRound, player.myCurrentlyChosenNumber)
		fmt.Println(player)
	}
	fmt.Println(numbersChosenLastRound)
}
