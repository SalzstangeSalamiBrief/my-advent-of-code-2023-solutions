package main

import (
	"fmt"
)

const SHOULD_RESPECT_LAST_LETTER = false

func main() {
	input := getInput("puzzleInput.txt")
	instructions := getInstructionSet(input[0])
	mapDictionary := getMapDictionary(input[2:])
	startPoints, endPoints := getStartAndEndPoints(mapDictionary, SHOULD_RESPECT_LAST_LETTER)
	visitedNodes := moveFromStartToGoal(instructions, mapDictionary, startPoints, endPoints)
	numberOfStepsTaken := getNumberOfStepsFromStartToGoal(visitedNodes)
	leastCommonMultiple := getLeastCommonMultiple(numberOfStepsTaken)
	fmt.Printf("leastCommonMultiple: %v\n", leastCommonMultiple)
}

func moveFromStartToGoal(instructions []string, mapDictionary MapDictionary, startPoints []string, endPoints []string) (visitedNodes [][]string) {
	visitedNodes = make([][]string, len(startPoints))

	for i, currentStartingPoint := range startPoints {
		fmt.Printf("currentStartingPoint: %v\n", currentStartingPoint)
		currentLocation := currentStartingPoint
		foundGoal := false
		for foundGoal == false {
			for _, instruction := range instructions {

				if instruction == "L" {
					currentLocation = mapDictionary[currentLocation].left
				}

				if instruction == "R" {
					currentLocation = mapDictionary[currentLocation].right
				}

				visitedNodes[i] = append(visitedNodes[i], currentLocation)
				foundGoal = checkIfGoalHasBeenReached(currentLocation)
				if foundGoal {
					break
				}
			}
		}
	}

	return visitedNodes
}

func checkIfGoalHasBeenReached(currentLocation string) bool {
	lastCharacter := currentLocation[len(currentLocation)-1]
	isCurrentLocationGoal := lastCharacter == 'Z'
	return isCurrentLocationGoal
}

func getNumberOfStepsFromStartToGoal(visitedNodes [][]string) []int {
	numberOfSteps := make([]int, len(visitedNodes))

	for i, visitedNode := range visitedNodes {
		numberOfSteps[i] = len(visitedNode)
	}

	return numberOfSteps
}

func getLeastCommonMultiple(numberOfStepsTake []int) uint64 {
	//lcm(a, b, c) = lcm(a, b) * c / (gcd(a, c) * gcd(b, c))
	//lcm(a, b, c) = (a * b * c) / (gcd(a, b) * gcd(a, c) * gcd(b, c))

	previous := 1
	for _, current := range numberOfStepsTake {
		greatestCommonDivisor := getGreatestCommonDivisorViaEuclidsAlgorithm(previous, current)
		previous = (previous * current) / greatestCommonDivisor
	}

	fmt.Printf("previous: %v\n", previous)
	return uint64(previous)
}

func getGreatestCommonDivisorViaEuclidsAlgorithm(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
