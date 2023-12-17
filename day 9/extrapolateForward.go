package main

import (
	"fmt"
	"slices"
)

func extrapolateForward(input []string) [][]int {
	extrapolatedSequences := make([][]int, len(input))

	for i, line := range input {
		currentNumbers := transformLineToIntArray(line)
		sequences := generateSequencesBasedOnStartingSequence(currentNumbers)
		fmt.Printf("The sequences are: %v\n", sequences)
		extrapolatedValues := forwardExtrapolateSequence(sequences)
		fmt.Printf("The extrapolated values are: %v\n", extrapolatedValues)
		extrapolatedSequences[i] = extrapolatedValues
	}

	return extrapolatedSequences
}

func generateSequencesBasedOnStartingSequence(startingSequence []int) [][]int {
	sequences := [][]int{startingSequence}

	for i := 0; i < len(sequences); i += 1 {
		newSubsequence := getForwardSubsequence(sequences[i])
		sequences = append(sequences, newSubsequence)

		doesNewSubsequenceContainsOnlyZeroes := doesSequenceContainsOnlyZeroes(newSubsequence)
		if doesNewSubsequenceContainsOnlyZeroes {
			break
		}
	}
	return sequences
}

func getForwardSubsequence(parentSequence []int) []int {
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

	return newSubsequence
}

func forwardExtrapolateSequence(sequences [][]int) []int {
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
