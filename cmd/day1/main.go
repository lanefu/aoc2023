package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input1.txt
var inputBytes []byte

//go:embed test1.txt
var testBytes []byte

func main() {
	testNumLines := countLines(testBytes)
	inputNumLines := countLines(inputBytes)

	fmt.Printf("Test file has %d lines\n", testNumLines)
	fmt.Printf("Input file has %d lines\n", inputNumLines)

	processLines(string(testBytes))

	calibrationTestFinalValue := calibrate(string(testBytes))
	fmt.Printf("Calibration value for test: %d\n", calibrationTestFinalValue)

	calibrationFinalValue := calibrate(string(inputBytes))
	fmt.Printf("Calibration value: %d\n", calibrationFinalValue)
}

func calibrate(calibrationData string) int {

	var calibrationValue int
	calibrationValue = 0

	lines := strings.Split(string(calibrationData), "\n")

	// Process each line
	for i, line := range lines {
		if i == len(lines)-1 && line == "" {
			break
		}
		// unsafe assumes at least one digit... which is true for this exercise
		calibrationDigits := getCalibrationDigitsFromString(line)
		calibrationValue += calibrationDigits
		// fmt.Printf("calibration digit: %d calibratin Value: %d\n", calibrationDigits, calibrationValue)
	}

	return calibrationValue

}

func countLines(content []byte) int {
	return strings.Count(string(content), "\n")
}

func processLines(content string) {
	lines := strings.Split(content, "\n")

	// Process each line
	for i, line := range lines {
		if i == len(lines)-1 && line == "" {
			break
		}
		intValue := getDigitsFromString(line)
		fmt.Printf("Line %d: %s AKA %d\n", i+1, line, intValue)
		// Add your processing logic for each line here
	}
}

func getDigitsFromString(inputString string) int {
	runeSlice := []rune(inputString)
	var digitSlice []rune

	for _, r := range runeSlice {
		if unicode.IsDigit(r) {
			digitSlice = append(digitSlice, r)
		}
	}

	// Convert the rune slice to a string
	digitString := string(digitSlice)

	// Convert the string of digits to an integer
	intValue, err := strconv.Atoi(digitString)
	if err != nil {
		fmt.Println("Error converting to integer:", err)
		return 0
	}

	return intValue
}

func getCalibrationDigitsFromString(inputString string) int {
	runeSlice := []rune(inputString)
	var digitSlice []rune
	var calibrationDigits []rune

	for _, r := range runeSlice {
		if unicode.IsDigit(r) {
			digitSlice = append(digitSlice, r)
		}
	}

	// Convert the rune slice to a string
	digitString := string(digitSlice)
	calibrationDigits = append(calibrationDigits, rune(digitString[0]))
	calibrationDigits = append(calibrationDigits, rune(digitString[len(digitString)-1]))

	// Convert the string of digits to an integer
	intValue, err := strconv.Atoi(string(calibrationDigits))
	if err != nil {
		fmt.Println("Error converting to integer:", err)
		return 0
	}

	return intValue
}
