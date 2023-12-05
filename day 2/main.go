package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type SubSet struct {
	red   int
	blue  int
	green int
}

type GameInformation struct {
	id      int
	subSets []SubSet
}

func main() {
	bagContents := SubSet{
		red:   12,
		green: 13,
		blue:  14,
	}

	fileContent := getInput("input.txt")
	sumOfIds := 0
	sumOfMinimalNumberOfCubes := 0
	for _, line := range fileContent {
		gameInformation := transformLineToGameInformation(line)
		isPossible := checkIfGameIsPossible(gameInformation, bagContents)
		fmt.Println(isPossible)
		if isPossible {
			sumOfIds += gameInformation.id
		}

		minimalNumberOfCubes := getMinimalNumberOfCubes(gameInformation)
		sumOfMinimalNumberOfCubes += minimalNumberOfCubes.red * minimalNumberOfCubes.blue * minimalNumberOfCubes.green
	}

	fmt.Printf("sumOfIds: %v\n", sumOfIds)
	fmt.Printf("sumOfMinimalNumberOfCubes: %v\n", sumOfMinimalNumberOfCubes)
}

func transformLineToGameInformation(line string) GameInformation {
	stringSeparatedInGameAndSubsets := strings.Split(line, ": ")
	gameString, subsetsString := stringSeparatedInGameAndSubsets[0], stringSeparatedInGameAndSubsets[1]

	gameId := getGameId(gameString)
	subsets := transformSubsetsIntoSubset(subsetsString)
	return GameInformation{
		id:      gameId,
		subSets: subsets,
	}
}

func getGameId(line string) int {
	stringifiedId := strings.Split(line, " ")[1]
	id, err := strconv.Atoi(stringifiedId)
	if err != nil {
		log.Panic(err.Error())
	}

	return id
}

func transformSubsetsIntoSubset(subsetsString string) []SubSet {
	subSetStringArray := strings.Split(subsetsString, "; ")
	var subsets []SubSet

	for _, subsetString := range subSetStringArray {
		var subset SubSet
		parts := strings.Split(subsetString, ", ")
		for _, part := range parts {
			p := strings.Split(part, " ")
			numberOfCubes, err := strconv.Atoi(p[0])
			if err != nil {
				log.Panic(err.Error())
			}

			if p[1] == "red" {
				subset.red = numberOfCubes
			}

			if p[1] == "blue" {
				subset.blue = numberOfCubes
			}

			if p[1] == "green" {
				subset.green = numberOfCubes
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

func checkIfGameIsPossible(gameInformation GameInformation, bagContents SubSet) bool {
	isGamePossible := true

	for _, subset := range gameInformation.subSets {
		isRedPossible := subset.red <= bagContents.red
		isBluePossible := subset.blue <= bagContents.blue
		isGreenPossible := subset.green <= bagContents.green

		isSubSetPossible := isRedPossible && isBluePossible && isGreenPossible
		isGamePossible = isGamePossible && isSubSetPossible

	}

	return isGamePossible
}

func getMinimalNumberOfCubes(gameInformation GameInformation) SubSet {
	var result SubSet

	for _, subset := range gameInformation.subSets {

		if subset.blue > result.blue {
			result.blue = subset.blue
		}

		if subset.green > result.green {
			result.green = subset.green
		}

		if subset.red > result.red {
			result.red = subset.red
		}
	}

	return result
}
