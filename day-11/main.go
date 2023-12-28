package main

import (
	"fmt"
)

type Galaxy struct {
	row int
	col int
}

func main() {

	input := getInput("exampleInput.txt")
	space := getSpace(input)
	emptyRows := getEmptyRows(space)
	emptyCols := getEmptyCols(space)
	galaxies := getGalaxies(space)
	for _, galaxy := range galaxies {
		fmt.Printf("%v\n", galaxy)
	}
	distances := getDistancesBetweenGalaxies(galaxies, emptyRows, emptyCols, 1)
	for _, row := range distances {
		fmt.Printf("%v\n", row)
	}
	sumOfDistances := getSumOfAllDistances(distances)
	fmt.Printf("sum of distances: %v\n", sumOfDistances)
}

func getGalaxies(space [][]string) []Galaxy {
	galaxies := make([]Galaxy, 0)

	for i, row := range space {
		for j, cell := range row {
			if cell != "#" {
				continue
			}

			newGalaxy := Galaxy{
				row: i,
				col: j,
			}

			galaxies = append(galaxies, newGalaxy)
		}
	}

	return galaxies
}

func getDistancesBetweenGalaxies(galaxies []Galaxy, emptyRows []int, emptyCols []int, valueOfEmptyElement int) [][]int {
	distances := make([][]int, len(galaxies)-1)

	for i := 0; i < len(galaxies)-1; i += 1 {
		distances[i] = make([]int, 0)
		for j := i + 1; j < len(galaxies); j += 1 {
			// TODO CALUCLATIOn
			// TODO BEFORE CALCULATING THE DIFFERENCE ADD NUMBER OF EMPTY ROWS/COLS TO TARGET
			startGalaxy := galaxies[i]
			endGalaxy := galaxies[j]

			isColOfEndBeforeColOfStart := startGalaxy.col > endGalaxy.col
			if isColOfEndBeforeColOfStart {
				startGalaxy = galaxies[j]
				endGalaxy = galaxies[i]
			}

			numberOfEmptyColsBetween := getNumberOfEmptyElementsBetweenStartAndEnd(startGalaxy.col, endGalaxy.col, emptyCols)
			numberOfExpandedEmptyCols := numberOfEmptyColsBetween * valueOfEmptyElement
			numberOfEmptyRowsBetween := getNumberOfEmptyElementsBetweenStartAndEnd(startGalaxy.row, endGalaxy.row, emptyRows)
			numberOfExpandedEmptyRows := numberOfEmptyRowsBetween * valueOfEmptyElement

			rowDistance := getDifferenceBetweenTwoNumbers(startGalaxy.row, endGalaxy.row+numberOfExpandedEmptyRows)
			colDistance := getDifferenceBetweenTwoNumbers(startGalaxy.col, endGalaxy.col+numberOfExpandedEmptyCols)

			distance := rowDistance + colDistance
			fmt.Printf("i: %v; j: %v; distance %v\n", i, j, distance)
			distances[i] = append(distances[i], distance)
		}
	}

	return distances
}

func getNumberOfEmptyElementsBetweenStartAndEnd(start int, end int, emptyElements []int) int {
	startElement := start
	endElement := end
	if start > end {
		startElement = end
		endElement = start
	}

	numberOfEmptyElements := 0

	for _, emptyElement := range emptyElements {
		if emptyElement > startElement && emptyElement < endElement {
			numberOfEmptyElements += 1
		}
	}

	return numberOfEmptyElements
}

func getDifferenceBetweenTwoNumbers(firstNumber int, secondNumber int) int {
	if firstNumber > secondNumber {
		return firstNumber - secondNumber
	}

	return secondNumber - firstNumber
}

func getSumOfAllDistances(distances [][]int) int {
	sum := 0

	for _, row := range distances {
		for _, distance := range row {
			sum += distance
		}
	}

	return sum
}
