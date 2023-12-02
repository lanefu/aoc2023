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

//go:embed test2.txt
var testBytes []byte

func main() {
	testNumLines := countLines(testBytes)
	inputNumLines := countLines(inputBytes)

	fmt.Printf("Test file has %d lines\n", testNumLines)
	fmt.Printf("Input file has %d lines\n", inputNumLines)

	//	processLines(string(testBytes))

	calibrationTestFinalValue := calibrate(string(testBytes))
	fmt.Printf("Calibration value for test: %d\n", calibrationTestFinalValue)

	// calibrationFinalValue := calibrate(string(inputBytes))
	// fmt.Printf("Calibration value: %d\n", calibrationFinalValue)
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
		calibrationDigits := getCalibrationDigitsFromString(digitPlunger(line))
		calibrationValue += calibrationDigits
		fmt.Printf("source: %s calibration digit: %d calibration Value: %d\n", line, calibrationDigits, calibrationValue)
	}

	return calibrationValue

}

func digitPlunger(plungeString string) string {

	var digits [10]string
	digits[0] = "zero"
	digits[1] = "one"
	digits[2] = "two"
	digits[3] = "three"
	digits[4] = "four"
	digits[5] = "five"
	digits[6] = "six"
	digits[7] = "seven"
	digits[8] = "eight"
	digits[9] = "nine"

	var cursor, cursorTrimStart, cursorTrimEnd int
	var digitLength int
	var trimmedString string

	oldString := plungeString
	cursor = 0

	fmt.Printf("\n%s:\n", oldString)

	// has to be at least 1 characters left to match a string
	for cursor <= (len(plungeString)) {

		for i := 0; i < 10; i++ {
			digit := digits[i]
			digitLength = len(digit)
			// block probably notnneeded now
			if cursor > 0 {
				cursorTrimStart = cursor
				cursorTrimEnd = cursorTrimStart + digitLength
			} else {
				cursorTrimEnd = digitLength
			}
			// bail if word length wont fit in substring
			if cursorTrimEnd <= len(plungeString) {

				if cursor == 0 {
					trimmedString = plungeString[:digitLength]
				} else {
					trimmedString = plungeString[cursorTrimStart:cursorTrimEnd]
				}
				fmt.Printf(" Trying %s in %s[%d]\n", digit, trimmedString, digitLength)
				if strings.Contains(trimmedString, digit) {
					plungeString = strings.Replace(plungeString, digit, strconv.Itoa(i), 1)
					fmt.Printf("  Matched %s, new value %s\n", digit, plungeString)
				}
			}
		}
		cursor++
	}
	fmt.Printf("   COMPLETE: before: %s after: %s\n", oldString, plungeString)
	return plungeString
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
