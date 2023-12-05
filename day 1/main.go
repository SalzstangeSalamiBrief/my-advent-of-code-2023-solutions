package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := getInput("input2.txt")
	calibration := getCalibrationSumOfLines(input)
	fmt.Println(calibration)
}

func getCalibrationSumOfLines(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += getCalibrationOfLine(line)
	}
	return sum
}

func getCalibrationOfLine(line string) int {
	stringWithoutDigitWords := transformDigitStringsInStringToNumber(line)
	numberRegexp := regexp.MustCompile("[0-9]{1}")
	arrayOfStringifiedDigits := numberRegexp.FindAllString(stringWithoutDigitWords, -1)
	concatedDigitsString := "0"

	numberOfDigits := len(arrayOfStringifiedDigits)

	if numberOfDigits == 1 {
		currentNumberString := transformNumberWordToNumber(arrayOfStringifiedDigits[0])
		concatedDigitsString = fmt.Sprintf("%v%v", currentNumberString, currentNumberString)
	}

	if numberOfDigits > 1 {
		firstNumberString := transformNumberWordToNumber(arrayOfStringifiedDigits[0])
		lastNumberString := transformNumberWordToNumber(arrayOfStringifiedDigits[numberOfDigits-1])
		concatedDigitsString = fmt.Sprintf("%v%v", firstNumberString, lastNumberString)
	}

	number, err := strconv.Atoi(concatedDigitsString)
	if err != nil {
		log.Fatal(err)
	}

	return number
}

func transformDigitStringsInStringToNumber(line string) string {
	withoutOne := strings.ReplaceAll(line, "one", "o1e")
	withoutTwo := strings.ReplaceAll(withoutOne, "two", "t2o")
	withoutThree := strings.ReplaceAll(withoutTwo, "three", "t3e")
	withoutFour := strings.ReplaceAll(withoutThree, "four", "f4r")
	withoutFive := strings.ReplaceAll(withoutFour, "five", "f5e")
	withoutSix := strings.ReplaceAll(withoutFive, "six", "s6x")
	withoutSeven := strings.ReplaceAll(withoutSix, "seven", "s7n")
	withoutEight := strings.ReplaceAll(withoutSeven, "eight", "e8t")
	withoutNine := strings.ReplaceAll(withoutEight, "nine", "n9e")
	return withoutNine
}

func transformNumberWordToNumber(numberWord string) string {
	switch numberWord {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return numberWord
	}
}
