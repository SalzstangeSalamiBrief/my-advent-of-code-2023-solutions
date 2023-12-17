package main

import (
	"fmt"
)

func main() {
	input := getInput("puzzleInput.txt")
	//extrapolatedSequences := extrapolateForward(input)
	extrapolatedSequences := extrapolateBackward(input)
	sumOfExtrapolatedSequences := getSumOfExtrapolatedSequences(extrapolatedSequences)
	fmt.Printf("The sum of the extrapolated sequences is: %v\n", sumOfExtrapolatedSequences)
}
