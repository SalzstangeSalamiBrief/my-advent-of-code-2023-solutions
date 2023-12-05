package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := getInput("puzzleInput.txt")
	// result := getSumOfPartNumbersAdjancentToCharacters(input)
	result := getGears(input)
	fmt.Printf("sum %v\n", result)
}

func getSumOfPartNumbersAdjancentToCharacters(input []string) int {
	splitTwoDimensionalArray := createSplitTwoDimensionalArray(input)
	fmt.Printf("splitTwoDimensionalArray %v\n", splitTwoDimensionalArray)
	sum := 0

	for row, line := range splitTwoDimensionalArray {

		column := 0
		for column < len(line) {
			possibleNumericalValue := getNumberFromSplitCharacters(column, line)
			numericalValue, transformToNumericalValueError := strconv.Atoi(possibleNumericalValue)
			if transformToNumericalValueError != nil {
				column += 1
				continue
			}

			endColumn := column + len(possibleNumericalValue) - 1
			isAdjancedToCharacter := checkIfNumberIsAdjancedToCharacter(row, column, endColumn, splitTwoDimensionalArray)

			if isAdjancedToCharacter {
				sum += numericalValue
			}

			column = endColumn + 1
		}
	}

	return sum
}

func createSplitTwoDimensionalArray(input []string) [][]string {
	result := make([][]string, len(input))
	for i, line := range input {
		result[i] = splitLineIntoNumbersAndCharacters(line)
	}
	return result
}

func getNumberFromSplitCharacters(currentColumnIndex int, line []string) string {
	regexForNumbers := regexp.MustCompile(`\d`)
	var result string
	for i := currentColumnIndex; i < len(line); i += 1 {
		currentCharacterOrNumber := line[i]
		isNumber := regexForNumbers.MatchString(currentCharacterOrNumber)
		if isNumber == false {
			break
		}

		result += currentCharacterOrNumber
	}

	return result
}

func splitLineIntoNumbersAndCharacters(line string) []string {
	regexForNumbers := regexp.MustCompile(`\d|\D`)
	splittedString := regexForNumbers.FindAllString(line, -1)
	return splittedString
}

func checkIfNumberIsAdjancedToCharacter(row int, startColumn int, endColumn int, splitTwoDimensionalArray [][]string) bool {
	possibleXCoordsForMatches := []int{row - 1, row, row + 1}
	possibleYCoordsForMatches := getRangeOfNumbers(startColumn-1, endColumn+1)

	for _, xCoord := range possibleXCoordsForMatches {
		isOutOfBoundariesXAxis := xCoord < 0 || xCoord >= len(splitTwoDimensionalArray)
		if isOutOfBoundariesXAxis {
			continue
		}

		isSameRow := xCoord == row
		currentRow := splitTwoDimensionalArray[xCoord]
		for _, yCoord := range possibleYCoordsForMatches {

			isOutOfBoundariesYAxis := yCoord < 0 || yCoord >= len(splitTwoDimensionalArray[xCoord])
			if isOutOfBoundariesYAxis {
				continue
			}

			isInsidelengthOfNumberInDigits := yCoord >= startColumn && yCoord <= endColumn
			if isSameRow && isInsidelengthOfNumberInDigits {
				continue
			}

			currentColumnWithCharacter := currentRow[yCoord]
			isDot := currentColumnWithCharacter == "."
			if isDot {
				continue
			}

			return true
		}
	}

	return false
}

func getRangeOfNumbers(start int, end int) []int {
	result := make([]int, end-start+1)
	for i := range result {
		result[i] = start + i
	}
	return result
}

func getGears(input []string) int {
	splitTwoDimensionalArray := createSplitTwoDimensionalArray(input)
	fmt.Printf("splitTwoDimensionalArray %v\n", splitTwoDimensionalArray)
	sum := 0

	for row, line := range splitTwoDimensionalArray {
		column := 0
		for column < len(line) {
			currentCharacter := line[column]
			if currentCharacter != "*" {
				column += 1
				continue
			}

			indicesOfAdjancedIndexesOfDigitsOfNumbers := getAdjancentIndexOfDigitOfNumber(row, column, splitTwoDimensionalArray)
			if len(indicesOfAdjancedIndexesOfDigitsOfNumbers) != 2 {
				column += 1
				continue
			}

			fmt.Printf("indicesOfAdjancedIndexesOfDigitsOfNumbers %v\n", indicesOfAdjancedIndexesOfDigitsOfNumbers)
			factors := make([]int, 2)
			for i, indices := range indicesOfAdjancedIndexesOfDigitsOfNumbers {
				col := indices[1]
				row := indices[0]
				numberAsString := splitTwoDimensionalArray[row][col]
				leftPartOfTheNumber := getLeftPartOfTheNumber(col, row, splitTwoDimensionalArray)
				numberAsString = leftPartOfTheNumber + numberAsString
				rightPartOfTheNumber := getRightPartOfTheNumber(col, row, splitTwoDimensionalArray)
				numberAsString = numberAsString + rightPartOfTheNumber

				number, transformToNumericalValueError := strconv.Atoi(numberAsString)
				if transformToNumericalValueError != nil {
					continue
				}

				factors[i] = number
			}

			product := factors[0] * factors[1]
			sum += product
			column += 1
		}
	}

	return sum
}

func getAdjancentIndexOfDigitOfNumber(row int, column int, splitTwoDimensionalArray [][]string) [][]int {
	adjancentPartNumberCoordinates := make([][]int, 0)
	hasMatchStraightAbove := checkIfCellIsNumber(row-1, column, splitTwoDimensionalArray)
	if hasMatchStraightAbove {
		adjancentPartNumberCoordinates = append(adjancentPartNumberCoordinates, []int{row - 1, column})
	}

	hasMatchStraightBelow := checkIfCellIsNumber(row+1, column, splitTwoDimensionalArray)
	if hasMatchStraightBelow {
		adjancentPartNumberCoordinates = append(adjancentPartNumberCoordinates, []int{row + 1, column})
	}

	hasMatchStraightLeft := checkIfCellIsNumber(row, column-1, splitTwoDimensionalArray)
	if hasMatchStraightLeft {
		adjancentPartNumberCoordinates = append(adjancentPartNumberCoordinates, []int{row, column - 1})
	}

	hasMatchStraightRight := checkIfCellIsNumber(row, column+1, splitTwoDimensionalArray)
	if hasMatchStraightRight {
		adjancentPartNumberCoordinates = append(adjancentPartNumberCoordinates, []int{row, column + 1})
	}

	canCheckTopCorners := hasMatchStraightAbove == false
	canCheckBottomCorners := hasMatchStraightBelow == false

	if canCheckTopCorners {
		hasMatchTopLeft := checkIfCellIsNumber(row-1, column-1, splitTwoDimensionalArray)
		if hasMatchTopLeft {
			adjancentPartNumberCoordinates = append(adjancentPartNumberCoordinates, []int{row - 1, column - 1})
		}

		hasMatchTopRight := checkIfCellIsNumber(row-1, column+1, splitTwoDimensionalArray)
		if hasMatchTopRight {
			adjancentPartNumberCoordinates = append(adjancentPartNumberCoordinates, []int{row - 1, column + 1})
		}
	}

	if canCheckBottomCorners {
		hasMatchBottomLeft := checkIfCellIsNumber(row+1, column-1, splitTwoDimensionalArray)
		if hasMatchBottomLeft {
			adjancentPartNumberCoordinates = append(adjancentPartNumberCoordinates, []int{row + 1, column - 1})
		}

		hasMatchBottomRight := checkIfCellIsNumber(row+1, column+1, splitTwoDimensionalArray)
		if hasMatchBottomRight {
			adjancentPartNumberCoordinates = append(adjancentPartNumberCoordinates, []int{row + 1, column + 1})
		}
	}

	return adjancentPartNumberCoordinates
}

func checkIfCellIsNumber(row int, column int, splitTwoDimensionalArray [][]string) bool {
	isOutsideXAxisBoundaries := row < 0 || row >= len(splitTwoDimensionalArray)
	isOutsideYAxisBoundaries := column < 0 || column >= len(splitTwoDimensionalArray[row])

	if isOutsideXAxisBoundaries || isOutsideYAxisBoundaries {
		return false
	}

	currentCharacter := splitTwoDimensionalArray[row][column]
	_, transformToNumericalValueError := strconv.Atoi(currentCharacter)

	return transformToNumericalValueError == nil
}

func getLeftPartOfTheNumber(col int, row int, splitTwoDimensionalArray [][]string) string {
	var leftNumberString string

	currentCol := col - 1 // start left
	for true {
		if currentCol < 0 {
			break
		}

		possibleNumericalValue := splitTwoDimensionalArray[row][currentCol]
		_, transformToNumericalValueError := strconv.Atoi(possibleNumericalValue)
		if transformToNumericalValueError != nil {
			break
		}

		leftNumberString = possibleNumericalValue + leftNumberString
		currentCol -= 1
	}

	return leftNumberString
}

func getRightPartOfTheNumber(col int, row int, splitTwoDimensionalArray [][]string) string {
	currentCol := col + 1 // start right
	var rightNumberString string

	for true {
		if currentCol >= len(splitTwoDimensionalArray[row]) {
			break
		}

		possibleNumericalValue := splitTwoDimensionalArray[row][currentCol]
		_, transformToNumericalValueError := strconv.Atoi(possibleNumericalValue)
		if transformToNumericalValueError != nil {
			break
		}

		rightNumberString = rightNumberString + possibleNumericalValue
		currentCol += 1
	}

	return rightNumberString
}
