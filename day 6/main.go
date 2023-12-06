package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := getInput("puzzleInput.txt")
	// partOne(input)
	partTwo(input)
}

func partOne(input []string) {
	if len(input) != 2 {
		log.Fatal("Input should have 2 lines")
	}

	timesInMs := getNumbersOfLine(input[0])
	distancesInMm := getNumbersOfLine(input[1])

	if len(timesInMs) != len(distancesInMm) {
		log.Fatal("Times and distances should have the same length")
	}
	allPossibleSolution := getAllPossibleSolutions(timesInMs, distancesInMm)
	fmt.Printf("Possible solutions: %v\n", allPossibleSolution)
	marginOfError := getMarginOfError(allPossibleSolution)
	fmt.Printf("Margin of error: %v\n", marginOfError)
}

func getAllPossibleSolutions(timesInMs []int, distancesInMm []int) [][]int {
	allPossibleSolution := make([][]int, len(timesInMs))

	for i, timeInMs := range timesInMs {
		distanceInMn := distancesInMm[i]
		possibleSolution := getPossibleSolutionsToReachTheGoalOfRace(timeInMs, distanceInMn)
		allPossibleSolution[i] = possibleSolution
	}

	return allPossibleSolution
}

func getNumbersOfLine(line string) []int {
	regexForNumbers := regexp.MustCompile(`\d+`)
	numbersAsString := regexForNumbers.FindAllString(line, -1)
	numbers := make([]int, len(numbersAsString))
	for i, numberAsString := range numbersAsString {
		number, err := strconv.Atoi(numberAsString)
		if err != nil {
			log.Fatal(err)
		}

		numbers[i] = number
	}

	return numbers
}

func getPossibleSolutionsToReachTheGoalOfRace(timeInMs int, distanceInMm int) []int {
	possibleSolutionsToReachTheGoal := make([]int, 0)

	// start at 1 because with zero the boat could not move a single millimeter
	// end at timeInMs - 1 because with a time equal to timeInMs the boat would start moving at the same time the race finishes
	for msToPressTheButton := 1; msToPressTheButton < timeInMs-1; msToPressTheButton += 1 {
		speedOfBoatInMs := msToPressTheButton * 1
		remainingTimeToTravelInMs := timeInMs - msToPressTheButton
		exceptedDistanceToTravelInMm := remainingTimeToTravelInMs * speedOfBoatInMs

		if exceptedDistanceToTravelInMm > distanceInMm {
			possibleSolutionsToReachTheGoal = append(possibleSolutionsToReachTheGoal, msToPressTheButton)
		}
	}

	return possibleSolutionsToReachTheGoal
}

func getMarginOfError(possibleSolutions [][]int) int {
	if len(possibleSolutions) == 0 {
		return 0
	}

	product := 1 // start by one for the first iteration to prevent a zero value as product

	for _, possibleSolution := range possibleSolutions {
		numberOfSolutions := len(possibleSolution)
		product *= numberOfSolutions
	}
	return product
}

func partTwo(input []string) {
	if len(input) != 2 {
		log.Fatal("Input should have 2 lines")
	}

	timesInMsAsString := input[0]
	timeInMs := concatPartialNumbersAndConvertToNumber(timesInMsAsString)
	timesInMs := []int{timeInMs}
	distancesInMmAsString := input[1]
	distanceInMs := concatPartialNumbersAndConvertToNumber(distancesInMmAsString)
	distancesInMm := []int{distanceInMs}
	allPossibleSolution := getAllPossibleSolutions(timesInMs, distancesInMm)
	totalNumberOfWaysToWin := getNumberOfTotalWaysToWinTheRace(allPossibleSolution)
	fmt.Printf("totalNumberOfWaysToWin: %v\n", totalNumberOfWaysToWin)
}

func concatPartialNumbersAndConvertToNumber(line string) int {
	regexForNumbers := regexp.MustCompile(`\d+`)
	partialNumbersAsString := regexForNumbers.FindAllString(line, -1)
	numberAsString := strings.Join(partialNumbersAsString, "")

	number, err := strconv.Atoi(numberAsString)
	if err != nil {
		log.Fatal(err)
	}

	return number
}

func getNumberOfTotalWaysToWinTheRace(possibleSolutions [][]int) int {
	result := 0

	for _, possibleSolution := range possibleSolutions {
		numberOfSolutions := len(possibleSolution)
		result += numberOfSolutions
	}

	return result
}
