package main

import (
	"fmt"
	"math"
)

type Galaxy struct {
	row int64
	col int64
}

func main() {
	input := getInput("puzzleInput.txt")
	space := getSpace(input)
	emptyRows := getEmptyRows(space)
	emptyCols := getEmptyCols(space)
	galaxies := getGalaxies(space, emptyRows, emptyCols, 1000000)
	distances := getDistancesBetweenGalaxies(galaxies)
	sumOfDistances := getSumOfAllDistances(distances)
	fmt.Printf("sum of distances: %v\n", sumOfDistances)
}

func getGalaxies(space [][]string, emptyRows []int64, emptyCols []int64, valueOfEmptyElement int64) []Galaxy {
	galaxies := make([]Galaxy, 0)

	for i, row := range space {
		for j, cell := range row {
			if cell != "#" {
				continue
			}

			numberOfEmptyRowsBetweenStartAndSpace := getNumberOfEmptyElementsBetweenStartAndEnd(0, int64(i), emptyRows)
			rowsToAdd := numberOfEmptyRowsBetweenStartAndSpace * (valueOfEmptyElement - 1) // remove one to not count the same row twice
			numberOfEmptyColsBetweenStartAndSpace := getNumberOfEmptyElementsBetweenStartAndEnd(0, int64(j), emptyCols)
			colsToAdd := numberOfEmptyColsBetweenStartAndSpace * (valueOfEmptyElement - 1) // remove one to not count the same col twice
			newGalaxy := Galaxy{
				row: int64(i) + rowsToAdd,
				col: int64(j) + colsToAdd,
			}

			galaxies = append(galaxies, newGalaxy)
		}
	}

	return galaxies
}

func getDistancesBetweenGalaxies(galaxies []Galaxy) [][]int64 {
	distances := make([][]int64, len(galaxies)-1)

	for i := 0; i < len(galaxies)-1; i += 1 {
		distances[i] = make([]int64, 0)
		for j := i + 1; j < len(galaxies); j += 1 {
			distance := getManhattenDistance(galaxies[i], galaxies[j])
			distances[i] = append(distances[i], distance)
		}
	}

	return distances
}

func getNumberOfEmptyElementsBetweenStartAndEnd(start int64, end int64, emptyElements []int64) int64 {
	startElement := start
	endElement := end
	if start > end {
		startElement = end
		endElement = start
	}

	var numberOfEmptyElements int64

	for _, emptyElement := range emptyElements {
		if emptyElement > startElement && emptyElement < endElement {
			numberOfEmptyElements += 1
		}
	}

	return numberOfEmptyElements
}

func getSumOfAllDistances(distances [][]int64) int64 {
	var sum int64

	for _, row := range distances {
		for _, distance := range row {
			sum += distance
		}
	}

	return sum
}

func getManhattenDistance(firstGalaxy Galaxy, secondGalaxy Galaxy) int64 {
	colDifference := float64(firstGalaxy.col - secondGalaxy.col)
	rowDifference := float64(firstGalaxy.row - secondGalaxy.row)
	distance := math.Abs(rowDifference) + math.Abs(colDifference)
	return int64(distance)
}
