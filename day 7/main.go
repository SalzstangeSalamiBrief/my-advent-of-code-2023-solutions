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

var debugCategories = [7]string{
	"FiveOfAKind",
	"FourOfAKind",
	"FullHouse",
	"ThreeOfAKind",
	"TwoPair",
	"OnePair",
	"HighCard",
}
var shouldUseDefaultJoker = false

func main() {
	input := getInput("puzzleInput.txt")
	handsGroupedByType := getHandsGroupedByType(input)
	rankedHands := rankHands(handsGroupedByType)
	totalsOfWinnings := getTotalOfWinnings(rankedHands, len(input))
	fmt.Printf("Total winnings: %v", totalsOfWinnings)
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
		for k := 0; k < len(firstHand.cardsAsStrength); k += 1 {
			firstHandCard := firstHand.cardsAsStrength[k]
			secondHandCard := secondHand.cardsAsStrength[k]

			if firstHandCard == secondHandCard {
				continue
			}

			if firstHandCard > secondHandCard {
				return -1
			}

			if firstHandCard < secondHandCard {
				return 1
			}

		}
		return 0
	})

	return sortedHands
}

func getTotalOfWinnings(handsGrouped [][]Hand, totalNumberOfHands int) int {

	total := 0
	currentRank := totalNumberOfHands

	for i, handGroup := range handsGrouped {
		fmt.Printf("handGroup %v\n", debugCategories[i])
		for _, hand := range handGroup {
			product := hand.bid * currentRank
			fmt.Printf("Rank: '%v' cards: '%v', bid: '%v'\n", currentRank, hand.cardsAsStrength, hand.bid)
			total += product
			currentRank -= 1
		}
	}

	if currentRank != 0 {
		log.Panicf("currentRank '%v' is not 0", currentRank)
	}

	return total
}
