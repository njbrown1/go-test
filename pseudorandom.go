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
	/*
		numbersChosen := []int{0, 7}
		allPlayers := []Player{
			{0, []int{}, []int{}, 0, "completelyRandom"},
			{0, []int{}, []int{}, 4, "rock"},
		}
	*/
	fmt.Println(findMostCommonNumbers([]int{4, 0, 4, 3, 1, 3, 3, 0, 0, 6, 1, 2, 6}))
	fmt.Println(findLeastCommonNumbers([]int{4, 0, 4, 3, 1, 3, 3, 0, 0, 6, 1, 2, 6}))
	fmt.Println(findLeastCommonNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 4, 5, 5, 5}))
	fmt.Println(findMostCommonNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 4, 5, 5, 5}))
}
