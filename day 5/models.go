package main

type Range struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

type RangeMapItem struct {
	destinationString string
	ranges            []Range
}

type RangesMap map[string]RangeMapItem
