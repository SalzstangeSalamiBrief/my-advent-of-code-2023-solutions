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

const numberOfElementsPerBatch = 4

func main() {
	input := getInput("puzzleInput.txt")
	isPartTwo := true
	seedRanges := getSeedRanges(input[0], isPartTwo)
	rangesMap = getRangeMap(input[1:])
	lowestLocations := mapSeedsToLocation(seedRanges)
	fmt.Printf("--lowestLocations: %#v\n", lowestLocations)
}

func mapSeedsToLocation(seedRanges [][2]int) int {
	var mutex = &sync.Mutex{}
	result := 0
	semaphore := make(chan struct{}, numberOfElementsPerBatch)
	var wg sync.WaitGroup

	for _, currentRange := range seedRanges {
		for i := currentRange[0]; i < currentRange[1]; i += 1 {
			wg.Add(1)
			go func(currentIndex int) {
				defer wg.Done()
				semaphore <- struct{}{}
				currentLocation := getLocationForSeed(currentIndex)
				mutex.Lock()
				if result == 0 || result > currentLocation {
					result = currentLocation
				}
				mutex.Unlock()
				<-semaphore
			}(i)
		}
	}

	wg.Wait()
	return result
}

func getLocationForSeed(seed int) int {
	location := mapSeedToSoil(seed, startingSource)
	if location == -1 {
		location = seed
	}
	return location
}

func mapSeedToSoil(seed int, sourceString string) int {
	currentMapName := sourceString
	resultSource := seed
	for currentMapName != "" && currentMapName != endingDestination {
		c, m := getDestinationOfSeed(resultSource, rangesMap[currentMapName])
		if m == "" {
			break
		}
		currentMapName = m
		resultSource = c
	}
	return resultSource
}

func getDestinationOfSeed(source int, rangeItem RangeMapItem) (newSource int, newMapName string) {
	if source < rangeItem.ranges[0].sourceRangeStart { // requires that the ranges are sorted
		return source, rangeItem.destinationString
	}

	for _, r := range rangeItem.ranges {
		if source < r.sourceRangeStart {
			continue
		}

		intervalEnd := r.sourceRangeStart + r.rangeLength - 1
		if source > intervalEnd {
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
