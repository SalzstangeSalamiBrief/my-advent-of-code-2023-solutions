package main

import (
	"regexp"
	"strings"
)

func getInstructionSet(line string) []string {
	return strings.Split(line, "")
}

func getMapDictionary(lines []string) MapDictionary {
	mapDictionary := make(MapDictionary, len(lines))
	for _, line := range lines {
		key, node := getMapDictionaryEntry(line)
		mapDictionary[key] = node
	}

	return mapDictionary
}

func getMapDictionaryEntry(line string) (key string, node Node) {
	keyRegexp := regexp.MustCompile("^[A-Z]{3}")
	key = keyRegexp.FindString(line)
	connectedNodesRegexp := regexp.MustCompile("[A-Z]{3}, [A-Z]{3}")
	connectedNodesString := connectedNodesRegexp.FindString(line)
	connectNodes := strings.Split(connectedNodesString, ", ")
	node = Node{left: connectNodes[0], right: connectNodes[1]}
	return key, node
}
