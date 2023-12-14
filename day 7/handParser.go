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
var possibleCardsWithJoker = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

// part 2 (usesJoker = false)
var possibleCardsWithoutJoker = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

func getHandsGroupedByType(input []string) [][]Hand {
	hands := getDefaultHandGroups()
	for _, line := range input {
		splitLine := strings.Split(line, " ")
		inputHand, bid := splitLine[0], splitLine[1]
		cards := getCards(inputHand)
		cardsAsStrength := getCardsAsStrength(cards)
		typeOfHand := getTypeOfHandWithJoker(cardsAsStrength)
		bidNumber := getBidStringAsNumber(bid)

		hand := Hand{
			cardsAsStrength: cardsAsStrength,
			bid:             bidNumber,
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
	possibleCards := possibleCardsWithJoker
	if shouldUseDefaultJoker == false {
		possibleCards = possibleCardsWithoutJoker
	}

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

func getTypeOfHandWithJoker(cards [5]int) TypeOfHandEnum {
	numberOfJokers := 0
	if shouldUseDefaultJoker == false {
		for _, card := range cards {
			if card == 0 {
				numberOfJokers += 1
			}
		}
	}

	numberOfEachCard := [13]int{}
	for _, card := range cards {
		numberOfEachCard[card] += 1
	}

	for i, _ := range numberOfEachCard {
		if shouldUseDefaultJoker == false && i == 0 {
			continue
		}

		numberOfEachCard[i] = numberOfEachCard[i] + numberOfJokers
	}

	numberOfFiveOfAKind := getNumberOfEachCard(numberOfEachCard, 5)
	numberOfFOurOfAKind := getNumberOfEachCard(numberOfEachCard, 4)
	numberOfTripples := getNumberOfEachCard(numberOfEachCard, 3)
	numberOfPairs := getNumberOfEachCard(numberOfEachCard, 2)

	if numberOfFiveOfAKind >= 1 {
		return FiveOfAKind
	}

	if numberOfFOurOfAKind >= 1 {
		return FourOfAKind
	}

	if numberOfTripples >= 1 && numberOfPairs >= 1 {
		return FullHouse
	}

	if numberOfTripples >= 1 {
		return ThreeOfAKind
	}

	if numberOfPairs >= 2 {
		return TwoPair
	}

	if numberOfPairs >= 1 {
		return OnePair
	}

	return HighCard
}

func getNumberOfEachCard(cards [13]int, expectedNumber int) int {
	var result int
	for _, card := range cards {
		if card == expectedNumber {
			result += 1

		}
	}
	return result
}
