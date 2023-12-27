package main

import (
	"fmt"
	"slices"
	"strings"
)

func getSpace(input []string) [][]string {
	space := make([][]string, len(input))
	for i, row := range input {
		spaceRow := strings.Split(row, "")
		space[i] = spaceRow
	}

	return space
}

func expandSpace(space [][]string) [][]string {
	var expandedSpace [][]string

	for _, row := range space {
		doesRowContainsGalaxy := false
		for _, col := range row {
			if col == "#" {
				doesRowContainsGalaxy = true
				break
			}
		}

		expandedSpace = append(expandedSpace, row)
		if doesRowContainsGalaxy == false {
			expandedSpace = append(expandedSpace, row)
		}
	}

	var colsToExpand []int
	for colIndex := 0; colIndex < len(expandedSpace[0]); colIndex += 1 {
		fmt.Println(len(expandedSpace[0]))
		doesColContainsGalaxy := false

		for rowIndex := 0; rowIndex < len(expandedSpace); rowIndex += 1 {
			if expandedSpace[rowIndex][colIndex] == "#" {
				doesColContainsGalaxy = true
				break
			}
		}

		if doesColContainsGalaxy == true {
			continue
		}

		colsToExpand = append(colsToExpand, colIndex)

	}

	for rowIndex := 0; rowIndex < len(expandedSpace); rowIndex += 1 {
		newRow := make([]string, 0)
		for i, col := range expandedSpace[rowIndex] {
			newRow = append(newRow, col)

			if slices.Contains(colsToExpand, i) {
				newRow = append(newRow, ".")
			}
		}
		expandedSpace[rowIndex] = newRow
	}

	return expandedSpace
}
