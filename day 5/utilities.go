package main

import (
	"log"
	"regexp"
	"strconv"
)

var regexpForNumbers = regexp.MustCompile(`\d+`)
var regexpForLabels = regexp.MustCompile(`([a-z]+-to-[a-z]+ map:)`)
var regexpToExtractSourceAndDestinationFromLabel = regexp.MustCompile(`-to-| map:`)

func getNumbersFromString(input string) []int {
	numbersAsString := regexpForNumbers.FindAllString(input, -1)
	numbers := make([]int, len(numbersAsString))

	for i, numberAsString := range numbersAsString {
		number, err := strconv.Atoi(numberAsString)
		if err != nil {
			log.Panic(err.Error())
		}

		numbers[i] = number
	}

	return numbers
}

func getRangeMap(lines []string) RangesMap {
	rangesMap := make(RangesMap)
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
			rangesMap[currentSource] = newRangeItem
			continue
		}

		numbers := getNumbersFromString(line)
		if len(numbers) != 3 {
			log.Panic("Invalid mapping line")
		}

		if currentSource == "" {
			log.Panic("No source found")
		}

		entry, ok := rangesMap[currentSource]
		if ok == false {
			log.Panic("No source found")
		}

		newRange := Range{
			destinationRangeStart: numbers[0],
			sourceRangeStart:      numbers[1],
			rangeLength:           numbers[2],
		}

		entry.ranges = append(entry.ranges, newRange)
		rangesMap[currentSource] = entry
	}

	return rangesMap
}

func getLabelPartsFromLabel(label string) (source string, destination string) {
	labelParts := regexpToExtractSourceAndDestinationFromLabel.Split(label, -1)
	return labelParts[0], labelParts[1]
}
