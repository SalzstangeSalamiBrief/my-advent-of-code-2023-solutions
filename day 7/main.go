package main

import (
	"fmt"
	"log"
	"slices"
)

type TypeOfHandEnum int

const (
	FiveOfAKind  TypeOfHandEnum = 0
	FourOfAKind  TypeOfHandEnum = 1
	FullHouse    TypeOfHandEnum = 2
	ThreeOfAKind TypeOfHandEnum = 3
	TwoPair      TypeOfHandEnum = 4
	OnePair      TypeOfHandEnum = 5
	HighCard     TypeOfHandEnum = 6
)

// five cards per hand
// [highest] A K Q J T 9 8 7 6 5 4 3 2 [lowest]

// types of hand [highest to lowest?]:
// 5 of a kind 5x
// 4 of a kind 4x
// full house 3x + 2x
// 3 of a kind 3x
// 2 pair 2x + 2x + 1x
// 1 pair 2x + 1x + 1x + 1x
// high card all labels are distinct

// two hand have the same type:
// compare each card from start to finish of each hand and the higher card wins => same card move one forward

// winning card wins: bid * rank
// rank is index in all number of hands

func main() {
	input := getInput("puzzleInput.txt")
	handsGroupedByType := getHandsGroupedByType(input)
	rankedHands := rankHands(handsGroupedByType)
	totalsOfWinnings := getTotalOfWinnings(rankedHands, len(input))
	fmt.Printf("%v", totalsOfWinnings)
}

func rankHands(handsGroupedByType [][]Hand) [][]Hand {
	rankedHands := getDefaultHandGroups()
	for k, handGroup := range handsGroupedByType {
		sortedHands := sortHandsOfTypeViaBubbleSort(handGroup)
		rankedHands[k] = sortedHands
	}

	return rankedHands
}

func sortHandsOfTypeViaBubbleSort(hands []Hand) []Hand {
	sortedHands := hands
	slices.SortFunc(sortedHands, func(firstHand, secondHand Hand) int {
		result := 0
		for k := 0; k < len(firstHand.cardsAsStrength); k += 1 {
			firstHandCard := firstHand.cardsAsStrength[k]
			secondHandCard := secondHand.cardsAsStrength[k]
			if firstHandCard != secondHandCard {
				result = firstHandCard - secondHandCard
				break
			}

		}
		return result
	})

	return sortedHands
}

func getHigherHand(firstHand Hand, secondHand Hand) int {
	for i, firstHandCard := range firstHand.cardsAsStrength {
		secondHandCard := secondHand.cardsAsStrength[i]
		if firstHandCard == secondHandCard {
			continue
		}

		compareResult := firstHandCard - secondHandCard

		if compareResult < 0 {
			return -1
		}

		if compareResult > 0 {
			return 1
		}

	}

	return 0
}

func getTotalOfWinnings(handsGrouped [][]Hand, totalNumberOfHands int) int {

	total := 0
	currentRank := totalNumberOfHands

	for _, handGroup := range handsGrouped {
		for _, hand := range handGroup {
			product := hand.bid * currentRank
			fmt.Printf("%v * %v = %v\n", hand.bid, currentRank, product)
			total += product
			currentRank -= 1
		}
	}

	if currentRank != 0 {
		log.Panicf("currentRank '%v' is not 0", currentRank)
	}

	return total
}
