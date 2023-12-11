package main

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cardsAsStrength [5]int
	bid             int
	rank            int
}

// part 1 (usesJoker = true)
//var possibleCards = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

var usesJoker = false

// part 2 (usesJoker = false)
var possibleCards = []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

func getHandsGroupedByType(input []string) [][]Hand {
	hands := getDefaultHandGroups()
	for _, line := range input {
		splitLine := strings.Split(line, " ")
		inputHand, bid := splitLine[0], splitLine[1]
		cards := getCards(inputHand)
		typeOfHand := getTypeOfHand(cards)
		cardsAsStrength := getCardsAsStrength(cards)

		//index := len(hands) - int(typeOfHand)
		hand := Hand{
			cardsAsStrength: cardsAsStrength,
			bid:             getBidStringAsNumber(bid),
			rank:            0,
		}

		hands[int(typeOfHand)] = append(hands[int(typeOfHand)], hand)
	}

	return hands
}

func getDefaultHandGroups() [][]Hand {
	return [][]Hand{
		[]Hand{},
		[]Hand{},
		[]Hand{},
		[]Hand{},
		[]Hand{},
		[]Hand{},
		[]Hand{},
	}
}

func getCards(hand string) [5]string {
	splitHand := strings.Split(hand, "")
	cards := [5]string{splitHand[0], splitHand[1], splitHand[2], splitHand[3], splitHand[4]}
	return cards
}

func getCardsAsStrength(cards [5]string) [5]int {
	cardsAsStrength := [5]int{}

	for i, card := range cards {
		indexOfCard := slices.IndexFunc(possibleCards, func(c string) bool {
			return c == card
		})

		cardsAsStrength[i] = indexOfCard
	}
	return cardsAsStrength
}

func getBidStringAsNumber(bid string) int {
	number, err := strconv.Atoi(bid)

	if err != nil {
		log.Panic(err.Error())
	}

	return number
}

func getTypeOfHand(cards [5]string) TypeOfHandEnum {
	groupedCards := make(map[string]int)

	for _, card := range cards {
		_, ok := groupedCards[card]
		if ok == false {
			groupedCards[card] = 0
		}

		groupedCards[card] += 1
	}

	if checkForFiveOfAKind(groupedCards) {
		return FiveOfAKind
	}

	if checkForFourOfAKind(groupedCards) {
		return FourOfAKind
	}

	if checkForFullHouse(groupedCards) {
		return FullHouse
	}

	if checkForThreeOfAKind(groupedCards) {
		return ThreeOfAKind
	}

	if checkForPairs(groupedCards, 2) {
		return TwoPair
	}

	if checkForPairs(groupedCards, 1) {
		return OnePair
	}

	return HighCard
}

func checkForFiveOfAKind(groupedCards map[string]int) bool {
	return len(groupedCards) == 1
}

func checkForFourOfAKind(groupedCards map[string]int) bool {
	if len(groupedCards) != 2 {
		return false
	}

	for _, value := range groupedCards {
		if value == 4 {
			return true
		}
	}

	return false
}

func checkForFullHouse(groupedCards map[string]int) bool {
	if len(groupedCards) != 2 {
		return false
	}

	hasThree := false
	hasTwo := false
	for _, value := range groupedCards {
		if value == 3 {
			hasThree = true
			continue
		}

		if value == 2 {
			hasTwo = true
		}
	}

	return hasThree && hasTwo
}

func checkForThreeOfAKind(groupedCards map[string]int) bool {
	numberOfThrees := 0
	for _, value := range groupedCards {
		if value == 3 {
			numberOfThrees += 1
		}
	}

	return numberOfThrees == 1
}

func checkForPairs(groupedCards map[string]int, expectedNumberOfPairs int) bool {
	numberOfPairs := 0
	for _, value := range groupedCards {
		if value == 2 {
			numberOfPairs += 1
		}
	}

	return numberOfPairs == expectedNumberOfPairs
}
