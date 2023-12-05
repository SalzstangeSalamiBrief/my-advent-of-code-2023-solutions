package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

type Card struct {
	id             int
	winningNumbers []int
	drawnNumbers   []int
}

func main() {
	input := getInput("puzzleInput.txt")
	cards := getCardsOfInput(input)
	totalPoints := getTotalPointsOfWinningNumbersDrawn(cards)
	fmt.Printf("totalPoints: %v\n", totalPoints)
	numberOfScratchCardsWon := getNumberOfTotalScratchCardsWon(cards, cards)
	fmt.Printf("numberOfScratchCardsWon: %v\n", numberOfScratchCardsWon)
}

func getCardsOfInput(input []string) []Card {
	splitCardIntoCardsRegExp := regexp.MustCompile(`:|\|`)
	cards := make([]Card, len(input))
	for cardIndex, cardInput := range input {
		cardParts := splitCardIntoCardsRegExp.Split(cardInput, -1)

		if len(cardParts) != 3 {
			panic("Invalid card input")
		}

		var card Card
		card.id = getCardId(cardParts[0])
		card.winningNumbers = getNumbersOfCardSection(cardParts[1])
		card.drawnNumbers = getNumbersOfCardSection(cardParts[2])

		cards[cardIndex] = card
	}

	return cards
}

func getCardId(cardPartWithId string) int {
	cardIdRegExp := regexp.MustCompile(`\d+`)
	cardIdAsString := cardIdRegExp.FindString(cardPartWithId)
	cardIdAsInt, err := strconv.Atoi(cardIdAsString)

	if err != nil {
		panic("Invalid card id")
	}

	return cardIdAsInt
}

func getNumbersOfCardSection(cardSection string) []int {
	numberRegExp := regexp.MustCompile(`\d+`)
	numbersAsString := numberRegExp.FindAllString(cardSection, -1)
	numbers := make([]int, len(numbersAsString))

	for numberIndex, numberAsString := range numbersAsString {
		numberAsInt, err := strconv.Atoi(numberAsString)
		if err != nil {
			panic("Invalid number")
		}

		numbers[numberIndex] = numberAsInt
	}

	return numbers
}

func getTotalPointsOfWinningNumbersDrawn(cards []Card) int {
	totalPoints := 0

	for _, card := range cards {
		winningCardsDrawn := getIdsOfWinningCardsDrawn(card)
		numberOfWinningCardsDrawn := len(winningCardsDrawn)
		winningNumbersDrawn := transformNumberOfWinningNumbersDrawnToIn(numberOfWinningCardsDrawn)
		totalPoints += winningNumbersDrawn
	}

	return totalPoints
}

func getNumberOfTotalScratchCardsWon(cards []Card, initialCards []Card) int {
	numberOfScratchCardsWon := 0

	for _, card := range cards {
		numberOfScratchCardsWon += 1 // the current card will be added to the total number of scratch cards won

		idsOfWinningCardsDrawn := getIdsOfWinningCardsDrawn(card)
		numberOfWinningCardsDrawn := len(idsOfWinningCardsDrawn)
		if numberOfWinningCardsDrawn == 0 {
			continue
		}

		idsOfReceivedCards := getIdsOfNextCards(card.id, numberOfWinningCardsDrawn)
		winningCardsDrawn := getCardsById(idsOfReceivedCards, initialCards)
		totalNumberOfScratches := getNumberOfTotalScratchCardsWon(winningCardsDrawn, initialCards)
		numberOfScratchCardsWon += totalNumberOfScratches

	}

	return numberOfScratchCardsWon
}

func getIdsOfWinningCardsDrawn(card Card) []int {
	var winningNumbers []int

	for _, winningNumber := range card.winningNumbers {
		for _, drawnNumber := range card.drawnNumbers {
			isNumberDrawn := winningNumber == drawnNumber
			if isNumberDrawn {
				winningNumbers = append(winningNumbers, winningNumber)
			}

		}
	}

	return winningNumbers
}

func transformNumberOfWinningNumbersDrawnToIn(numberOfWinningCardsDrawn int) int {
	sum := 0

	for i := 1; i <= numberOfWinningCardsDrawn; i += 1 {
		if sum == 0 {
			sum = 1
			continue
		}

		sum *= 2
	}

	return sum
}

func getCardsById(cardIds []int, cards []Card) []Card {
	drawnCards := make([]Card, len(cardIds))

	for i, cardId := range cardIds {
		indexOfCard := slices.IndexFunc[[]Card](cards, func(card Card) bool { return card.id == cardId })
		if indexOfCard == -1 {
			continue
		}

		drawnCards[i] = cards[indexOfCard]
	}

	return drawnCards
}

func getIdsOfNextCards(initialCardId int, numberOfCards int) []int {
	result := make([]int, numberOfCards)
	for i := range result {
		result[i] = initialCardId + i + 1
	}

	return result
}
