package main

import (
	"log"
	"regexp"
	"strconv"
)

func transformLineToIntArray(line string) []int {
	numberRegexp := regexp.MustCompile(`-?\d+`)
	numbersAsStrings := numberRegexp.FindAllString(line, -1)
	parsedNumbers := make([]int, len(numbersAsStrings))

	for i, numberAsString := range numbersAsStrings {
		parsedNumber, err := strconv.Atoi(numberAsString)
		if err != nil {
			log.Panic(err.Error())
		}

		parsedNumbers[i] = parsedNumber
	}

	return parsedNumbers
}

func doesSequenceContainsOnlyZeroes(sequence []int) bool {
	for _, number := range sequence {
		if number != 0 {
			return false
		}
	}

	return true
}

func getSumOfExtrapolatedSequences(extrapolatedValues [][]int) int {
	sum := 0

	for _, extrapolatedSequence := range extrapolatedValues {
		lastElementInSequence := extrapolatedSequence[len(extrapolatedSequence)-1]
		sum += lastElementInSequence
	}

	return sum
}
