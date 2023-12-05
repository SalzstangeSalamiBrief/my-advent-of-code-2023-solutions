package main

import (
	"fmt"
	"regexp"
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
		numberOfWinningCardsDrawn := getNumberOfWinningNumbersDrawn(card)
		pointsOfCard := transformNumberOfWinningNumbersDrawnToIn(numberOfWinningCardsDrawn)
		totalPoints += pointsOfCard
	}

	return totalPoints
}

func getNumberOfWinningNumbersDrawn(card Card) int {
	numberOfWinningCardsDrawn := 0

	for _, winningNumber := range card.winningNumbers {
		for _, drawnNumber := range card.drawnNumbers {
			isNumberDrawn := winningNumber == drawnNumber
			if isNumberDrawn {
				numberOfWinningCardsDrawn += 1
				break
			}

		}
	}

	return numberOfWinningCardsDrawn
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
