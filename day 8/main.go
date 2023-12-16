package main

import (
	"fmt"
)

// left/right to navigate through the network

const START = "AAA"
const GOAL = "ZZZ"

func main() {
	input := getInput("puzzleInput.txt")
	instructions := getInstructionSet(input[0])
	mapDictionary := getMapDictionary(input[2:])
	fmt.Printf("instructions: %v\n", instructions)
	fmt.Printf("mapDictionary: %v\n", mapDictionary)
	visitedNodes := moveFromStartToGoal(instructions, mapDictionary)
	fmt.Printf("visitedNodes: %v\n", visitedNodes)
	numberOfStepsTaken := getNumberOfStepsFromStartToGoal(visitedNodes)
	fmt.Printf("numberOfStepsTaken: %v\n", numberOfStepsTaken)
}

// A => l => b => l => a => r => b => l => a => b => r =>

func moveFromStartToGoal(instructions []string, mapDictionary MapDictionary) (visitedNodes []string) {
	visitedNodes = make([]string, 0)
	foundGoal := false
	currentLocation := START

	for foundGoal == false {
		for _, instruction := range instructions {

			if instruction == "L" {
				currentLocation = mapDictionary[currentLocation].left
			}

			if instruction == "R" {
				currentLocation = mapDictionary[currentLocation].right
			}

			visitedNodes = append(visitedNodes, currentLocation)
			if currentLocation == GOAL {
				foundGoal = true
				break
			}
		}
	}

	return visitedNodes
}

func getNumberOfStepsFromStartToGoal(visitedNodes []string) int {
	return len(visitedNodes)
}
