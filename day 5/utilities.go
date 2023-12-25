package main

import (
	"log"
	"regexp"
	"slices"
	"strconv"
)

var regexpForNumbers = regexp.MustCompile(`\d+`)
var regexpForLabels = regexp.MustCompile(`([a-z]+-to-[a-z]+ map:)`)
var regexpToExtractSourceAndDestinationFromLabel = regexp.MustCompile(`-to-| map:`)

func getSeedRanges(lineWithSeeds string, isPartTwo bool) [][2]int {
	numbersAsString := regexpForNumbers.FindAllString(lineWithSeeds, -1)
	resultTuples := make([][2]int, 0)

	if isPartTwo == false {
		for _, numberAsString := range numbersAsString {
			number, err := strconv.Atoi(numberAsString)
			if err != nil {
				log.Panic(err.Error())
			}
			currentInterval := [2]int{number, number + 1}
			resultTuples = append(resultTuples, currentInterval)
		}
	}

	if isPartTwo == true {
		for i := 0; i < len(numbersAsString); i += 2 {
			intervalStart, err := strconv.Atoi(numbersAsString[i])
			if err != nil {
				log.Panic(err.Error())
			}

			intervalLength, err := strconv.Atoi(numbersAsString[i+1])
			if err != nil {
				log.Panic(err.Error())
			}

			intervalEnd := intervalStart + intervalLength
			currentInterval := [2]int{intervalStart, intervalEnd}
			resultTuples = append(resultTuples, currentInterval)
		}
	}

	return resultTuples
}

func getNumbersFromString(input string, shouldUseRangesOfSeeds bool) []int {
	numbersAsString := regexpForNumbers.FindAllString(input, -1)
	numbers := make([]int, len(numbersAsString))

	for i, numberAsString := range numbersAsString {
		number, err := strconv.Atoi(numberAsString)
		if err != nil {
			log.Panic(err.Error())
		}

		numbers[i] = number
	}

	if shouldUseRangesOfSeeds == true {
		if len(numbers)%2 != 0 {
			log.Panic("Invalid seed ranges")
		}

		numbersBasedOnRanges := make([]int, 0)
		for i := 0; i < len(numbers); i += 2 {
			intervalStart := numbers[i]
			intervalLength := numbers[i+1]
			intervalEnd := intervalStart + intervalLength
			for j := intervalStart; j < intervalEnd; j += 1 {
				numbersBasedOnRanges = append(numbersBasedOnRanges, j)
			}
		}

		numbers = numbersBasedOnRanges
	}

	return numbers
}

func getRangeMap(lines []string) RangesMap {
	resultMap := make(RangesMap)
	var currentSource string

	for _, line := range lines {
		label := regexpForLabels.FindString(line)
		if label != "" {
			source, destination := getLabelPartsFromLabel(label)
			currentSource = source
			newRangeItem := RangeMapItem{
				destinationString: destination,
				ranges:            make([]Range, 0),
			}
			resultMap[currentSource] = newRangeItem
			continue
		}

		numbers := getNumbersFromString(line, false)
		if len(numbers) != 3 {
			log.Panic("Invalid mapping line")
		}

		if currentSource == "" {
			log.Panic("No source found")
		}

		entry, ok := resultMap[currentSource]
		if ok == false {
			log.Panic("No source found")
		}

		newRange := Range{
			destinationRangeStart: numbers[0],
			sourceRangeStart:      numbers[1],
			rangeLength:           numbers[2],
		}

		entry.ranges = append(entry.ranges, newRange)
		resultMap[currentSource] = entry
	}

	for _, v := range resultMap {
		slices.SortFunc(v.ranges, func(a, b Range) int { return a.sourceRangeStart - b.sourceRangeStart })
	}

	return resultMap
}

func getLabelPartsFromLabel(label string) (source string, destination string) {
	labelParts := regexpToExtractSourceAndDestinationFromLabel.Split(label, -1)
	return labelParts[0], labelParts[1]
}
