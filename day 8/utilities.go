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
	keyRegexp := regexp.MustCompile("^[0-9A-Z]{3}")
	key = keyRegexp.FindString(line)
	connectedNodesRegexp := regexp.MustCompile("[0-9A-Z]{3}, [0-9A-Z]{3}")
	connectedNodesString := connectedNodesRegexp.FindString(line)
	connectNodes := strings.Split(connectedNodesString, ", ")
	node = Node{left: connectNodes[0], right: connectNodes[1]}
	return key, node
}

func getStartAndEndPoints(mapDictionary MapDictionary, shouldRespectLastLetter bool) (startgPoints []string, endPoints []string) {
	if shouldRespectLastLetter == false {
		startgPoints = []string{"AAA"}
		endPoints = []string{"ZZZ"}
		return startgPoints, endPoints
	}

	startgPoints = make([]string, 0)
	endPoints = make([]string, 0)

	for key, _ := range mapDictionary {
		lastCharacter := key[len(key)-1]

		if lastCharacter == 'A' {
			startgPoints = append(startgPoints, key)
		}

		if lastCharacter == 'Z' {
			endPoints = append(endPoints, key)
		}
	}

	return startgPoints, endPoints
}
