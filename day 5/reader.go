package main

import (
	"log"
	"os"
	"regexp"
)

func getInput(file string) []string {
	fileContentBytes, readFileErr := os.ReadFile(file)
	if readFileErr != nil {
		log.Panic(readFileErr.Error())
	}

	fileContentString := string(fileContentBytes)
	emptyLineRegex := regexp.MustCompile(`\r\n`)
	// TODO SPLIT BY EMPTY LINES
	contentSplitByEmptyLine := emptyLineRegex.Split(fileContentString, -1)
	linesWithContent := removeLinesWithoutContent(contentSplitByEmptyLine)
	return linesWithContent
}

func removeLinesWithoutContent(lines []string) []string {
	var nonEmptyLines []string
	for _, line := range lines {
		isEmptyLine := line == ""
		if isEmptyLine {
			continue
		}
		nonEmptyLines = append(nonEmptyLines, line)
	}
	return nonEmptyLines
}
