package main

import (
	"log"
	"regexp"
	"strconv"
)

// TODO
type SeedToSoilMap struct {
	label string
	seed  int
	soil  int
}

func transformInputToSeedsAndMap(input []string) ([]int, []SeedToSoilMap) {
	seeds := getInput(input[0])
}

func getNumbersFromString(line string) []int {
	regexForNumbers := regexp.MustCompile(`\d+`)
	seedStrings := regexForNumbers.FindAllString(line, -1)
	seeds := make([]int, len(seedStrings))

	for i, seedString := range seedStrings {
		seedAsNumber, err := strconv.Atoi(seedString)
		if err != nil {
			log.Panic(err.Error())
		}

		seeds[i] = seedAsNumber
	}

	return seeds
}

func getSeedToSoilMap(lines []string) {
	if len(lines)%3 != 0 {
		log.Panic("Input is not dividable by 3")
	}

	currentIndex := 0
	for currentIndex < len(lines) {
		var seedToSoilMap SeedToSoilMap
		seedToSoilMap.label = lines[currentIndex]
		seedDestinationRangeStart, seedSourceRangeStart, seedRangeLength := getMappingParts(lines[currentIndex+1])
		soilDestinationRangeStart, soilSourceRangeStart, soilRangeLength := getMappingParts(lines[currentIndex+2])

		currentIndex += 3
	}
}

func getMappingParts(line string) (int, int, int) {
	numbers := getNumbersFromString(line)
	if len(numbers) != 3 {
		log.Panic("Invalid mapping line")
	}

	return numbers[0], numbers[1], numbers[2]
}
