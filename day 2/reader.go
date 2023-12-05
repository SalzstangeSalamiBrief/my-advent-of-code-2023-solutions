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
	contentSplitByEmptyLine := emptyLineRegex.Split(fileContentString, -1)
	return contentSplitByEmptyLine
}
