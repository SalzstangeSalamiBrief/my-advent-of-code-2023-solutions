package main

import (
	"fmt"
	"sync"
)

// [source]-to-[destination]
// [destinationRangeStart] [sourceRangeStart] [rangeLength]

// example:
// 50 98 2
// 	- sourceRange [98, 99] (sourceRangeStart + rangeLength - 1)
// 	- destinationRange [50, 51] (destinationRange + rangeLength - 1)
// 	- seedNr 98 => soilNumber 50
// 	- seedNr 99 => soilNumber 51
// 52 50 48
// 	- sourceRange [50, ...97] (sourceRangeStart + rangeLength - 1)
// 	- destinationRange [52, ...99] (destinationRange + rangeLength - 1)
// 	- seedNr 53 => soilNumber 55

var startingSource = "seed"
var endingDestination = "location"
var rangesMap = RangesMap{}

func main() {
	input := getInput("puzzleInput.txt")
	seeds := getNumbersFromString(input[0])
	rangesMap = getRangeMap(input[1:])
	locations := mapSeedsToLocation(seeds)
	fmt.Printf("locations: %#v\n", locations)
	lowestLocations := getLowestLocationFromLocations(locations)
	fmt.Printf("--lowestLocations: %#v\n", lowestLocations)
}

func mapSeedsToLocation(seeds []int) []int {
	mappings := make([]int, 0)

	var waitGroup sync.WaitGroup
	resultChannel := make(chan int)
	for _, seed := range seeds {
		waitGroup.Add(1)
		go func(s int) {
			defer waitGroup.Done()
			getLocationForSeed(s, resultChannel)
		}(seed)
	}

	go func() {
		waitGroup.Wait()
		close(resultChannel)
	}()

	for result := range resultChannel {
		mappings = append(mappings, result)
	}

	return mappings
}

func getLocationForSeed(seed int, resultChannel chan int) {
	location := mapSeedToSoil(seed, startingSource)
	if location == -1 {
		location = seed
	}
	resultChannel <- location
}

func mapSeedToSoil(seed int, sourceString string) int {

	currentMapName := sourceString
	resultSource := seed
	for currentMapName != "" && currentMapName != endingDestination {
		c, m := getDestinationOfSeed(resultSource, rangesMap[currentMapName])
		//fmt.Printf("c: %d, m: %s\n", c, m)
		if m == "" {
			break
		}
		currentMapName = m
		resultSource = c
	}
	return resultSource
}

func getDestinationOfSeed(source int, rangeItem RangeMapItem) (newSource int, newMapName string) {
	for _, r := range rangeItem.ranges {
		isInInterval := isValueInInterval(source, r.sourceRangeStart, r.rangeLength)
		if isInInterval == false {
			continue
		}

		indexOfValueInInterval := source - r.sourceRangeStart
		newSource = r.destinationRangeStart + indexOfValueInInterval
		return newSource, rangeItem.destinationString
	}

	return source, rangeItem.destinationString
}

func getLowestLocationFromLocations(locations []int) int {
	lowestLocation := locations[0]
	for i := 1; i < len(locations); i += 1 {
		if locations[i] < lowestLocation {
			lowestLocation = locations[i]
		}
	}

	return lowestLocation
}

func isValueInInterval(valueToFind int, intervalStart int, intervalLength int) bool {
	intervalEnd := intervalStart + intervalLength - 1
	return valueToFind >= intervalStart && valueToFind <= intervalEnd
}
