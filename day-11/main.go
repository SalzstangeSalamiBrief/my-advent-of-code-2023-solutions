package main

import (
	"fmt"
)

type Galaxy struct {
	row int
	col int
}

func main() {

	input := getInput("puzzleInput.txt")
	space := getSpace(input)
	expandedSpace := expandSpace(space)
	for _, row := range expandedSpace {
		fmt.Printf("%v\n", row)
	}
	fmt.Printf("rows: '%v'; cols: '%v'\n", len(expandedSpace), len(expandedSpace[0]))
	galaxies := getGalaxies(expandedSpace)
	for _, galaxy := range galaxies {
		fmt.Printf("%v\n", galaxy)
	}
	distances := getDistancesBetweenGalaxies(galaxies)
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

func getDistancesBetweenGalaxies(galaxies []Galaxy) [][]int {
	distances := make([][]int, len(galaxies)-1)

	for i := 0; i < len(galaxies)-1; i += 1 {
		distances[i] = make([]int, 0)
		for j := i + 1; j < len(galaxies); j += 1 {
			startGalaxy := galaxies[i]
			endGalaxy := galaxies[j]

			rowDistance := endGalaxy.row - startGalaxy.row
			colDistance := endGalaxy.col - startGalaxy.col
			if colDistance < 0 {
				colDistance = startGalaxy.col - endGalaxy.col
			}

			distance := rowDistance + colDistance
			fmt.Printf("i: %v; j: %v; distance %v\n", i, j, distance)
			distances[i] = append(distances[i], distance)
		}
	}

	return distances
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
