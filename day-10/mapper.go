package main

import "strings"

type MapGrid [][]string

var Map MapGrid

func initMap(fileName string) {
	input := getInput(fileName)
	grid := make([][]string, len(input))
	for i, line := range input {
		cells := strings.Split(line, "")
		grid[i] = cells
	}

	Map = grid
}

func (m *MapGrid) getStartingNode() Node {
	for i, row := range *m {
		for j, column := range row {
			if column == "S" {
				startingNode := Node{
					coordinates:    [2]int{i, j},
					connectionType: StartingPipe,
				}

				return startingNode
			}
		}
	}

	panic("No starting node found in grid")
}
