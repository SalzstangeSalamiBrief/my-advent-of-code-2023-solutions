package main

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	// TODO DEBUG
	input := getInput("puzzleInput.txt")
	extrapolatedSequences := getExtrapolateSequenceForEachInput(input)
	for _, extrapolatedSequence := range extrapolatedSequences {
		fmt.Printf("The extrapolated values are: %v\n", extrapolatedSequence)
	}
	sumOfExtrapolatedSequences := getSumOfExtrapolatedSequences(extrapolatedSequences)
	fmt.Printf("The sum of the extrapolated sequences is: %v\n", sumOfExtrapolatedSequences)
}

func getExtrapolateSequenceForEachInput(input []string) [][]int {
	extrapolatedSequences := make([][]int, len(input))

	for i, line := range input {
		currentNumbers := transformLineToIntArray(line)
		sequences := generateSequencesBasedOnStartingSequence(currentNumbers)
		fmt.Printf("The sequences are: %v\n", sequences)
		extrapolatedValues := extrapolateSequence(sequences)
		//fmt.Printf("The extrapolated values are: %v\n", extrapolatedValues)
		extrapolatedSequences[i] = extrapolatedValues
	}

	return extrapolatedSequences
}

func transformLineToIntArray(line string) []int {
	numberRegexp := regexp.MustCompile(`\d+`)
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

func generateSequencesBasedOnStartingSequence(startingSequence []int) [][]int {
	sequences := [][]int{startingSequence}

	for i := 0; i < len(sequences); i += 1 {
		newSubsequence := getSubsequence(sequences[i])
		sequences = append(sequences, newSubsequence)

		doesNewSubsequenceContainsOnlyZeroes := doesSequenceContainsOnlyZeroes(newSubsequence)
		if doesNewSubsequenceContainsOnlyZeroes {
			break
		}
	}
	return sequences
}

func getSubsequence(parentSequence []int) []int {
	newSubsequence := make([]int, 0)
	if len(parentSequence) <= 1 {
		return newSubsequence
	}

	lengthOfParentSequence := len(parentSequence)
	for i := 0; i < lengthOfParentSequence-1; i += 1 {
		nextIndex := i + 1
		if nextIndex == lengthOfParentSequence {
			newSubsequence = append(newSubsequence, 0)
			break
		}

		currentNumber := parentSequence[i]
		nextNumber := parentSequence[nextIndex]
		difference := nextNumber - currentNumber
		newSubsequence = append(newSubsequence, difference)
	}
	// This function returns the next subsequence
	// This function returns the next subsequence
	// This function returns the next subsequence
	return newSubsequence
}

func doesSequenceContainsOnlyZeroes(sequence []int) bool {
	for _, number := range sequence {
		if number != 0 {
			return false
		}
	}

	return true
}

func extrapolateSequence(sequences [][]int) []int {
	extrapolatedValues := make([]int, len(sequences))
	clonedSequences := slices.Clone(sequences)
	slices.Reverse(clonedSequences)

	for i, currentSequence := range clonedSequences {
		if i == 0 {
			extrapolatedValues[i] = 0
			continue
		}

		previousExtrapolatedValue := extrapolatedValues[i-1]
		lastElementOfCurrentSequence := currentSequence[len(currentSequence)-1]
		extrapolatedValues[i] = previousExtrapolatedValue + lastElementOfCurrentSequence
	}

	return extrapolatedValues
}

func getSumOfExtrapolatedSequences(extrapolatedValues [][]int) int {
	sum := 0

	for _, extrapolatedSequence := range extrapolatedValues {
		lastElementInSequence := extrapolatedSequence[len(extrapolatedSequence)-1]
		sum += lastElementInSequence
	}

	return sum
}
