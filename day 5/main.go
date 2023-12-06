package main

import (
	"fmt"
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

// ! sourceNr that are not mapped correspond to the same destination
// 	- example: 10 => 10

// seed-to-soil-map: [seed] [soil]
// seed  soil
// 0     0
// 1     1
// ...   ...
// 48    48
// 49    49
// 50    52
// 51    53
// ...   ...
// 96    98
// 97    99
// 98    50
// 99    51

// TODO
func main() {
	input := getInput("exampleInput.txt")
	fmt.Printf("%#v", input)
}
