package main

import "slices"

func extrapolateBackward(input []string) [][]int {
	extrapolatedSequences := make([][]int, len(input))

	for i, line := range input {
		currentNumbers := transformLineToIntArray(line)
		sequences := generateBackwardExtrapolatedSequences(currentNumbers)
		extrapolatedValues := backwardExtrapolateSequence(sequences)
		extrapolatedSequences[i] = extrapolatedValues
	}

	return extrapolatedSequences
}

func generateBackwardExtrapolatedSequences(startingSequence []int) [][]int {
	sequences := [][]int{startingSequence}

	for i := 0; i < len(sequences); i += 1 {
		newSubsequence := getBackwardSequence(sequences[i])
		sequences = append(sequences, newSubsequence)

		doesNewSubsequenceContainsOnlyZeroes := doesSequenceContainsOnlyZeroes(newSubsequence)
		if doesNewSubsequenceContainsOnlyZeroes {
			break
		}
	}

	return sequences
}

func getBackwardSequence(parentSequence []int) []int {
	lengthOfParentSequence := len(parentSequence)
	newSubsequence := make([]int, 0)
	if lengthOfParentSequence <= 1 {
		return newSubsequence
	}

	for i := lengthOfParentSequence - 1; i >= 0; i -= 1 {
		if i == 0 {
			//newSubsequence = append(newSubsequence, 0)
			break
		}

		previousIndex := i - 1
		currentNumber := parentSequence[i]
		previousNumber := parentSequence[previousIndex]
		difference := currentNumber - previousNumber
		newSubsequence = append([]int{difference}, newSubsequence...)
	}

	return newSubsequence
}

func backwardExtrapolateSequence(sequences [][]int) []int {
	extrapolatedValues := make([]int, len(sequences))
	clonedSequences := slices.Clone(sequences)
	slices.Reverse(clonedSequences)

	for i, currentSequence := range clonedSequences {
		if i == 0 {
			extrapolatedValues[i] = 0
			continue
		}

		previousExtrapolatedValue := extrapolatedValues[i-1]
		firstElementInSequence := currentSequence[0]
		extrapolatedValues[i] = firstElementInSequence - previousExtrapolatedValue
	}

	return extrapolatedValues
}
