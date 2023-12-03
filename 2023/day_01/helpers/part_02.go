package day_01

import (
	"strings"
	// "2023/global"
)

func DecipherStringNumber(calibrationString string, lastLetter string) string {
	switch calibrationString {
	case "zero":
		return "0"
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
		if len(lastLetter) > 0 {
			return DecipherStringNumber(lastLetter+calibrationString[0:], "")
		}
		return ""
	}
}

// if string starts with number, return length of string
func StartsWithNumber(previousLastLetter, calibrationString string) (bool, int, string) {
	// if strings.HasPrefix(calibrationString, "zero") {
	// 	return true, 4
	// } else
	if strings.HasPrefix(calibrationString, "one") {
		return true, 3, "e"
	} else if strings.HasPrefix(calibrationString, "two") {
		return true, 3, "o"
	} else if strings.HasPrefix(calibrationString, "three") {
		return true, 5, "e"
	} else if strings.HasPrefix(calibrationString, "four") {
		return true, 4, "r"
	} else if strings.HasPrefix(calibrationString, "five") {
		return true, 4, "e"
	} else if strings.HasPrefix(calibrationString, "six") {
		return true, 3, "x"
	} else if strings.HasPrefix(calibrationString, "seven") {
		return true, 5, "n"
	} else if strings.HasPrefix(calibrationString, "eight") {
		return true, 5, "t"
	} else if strings.HasPrefix(calibrationString, "nine") {
		return true, 4, "e"
	} else if len(previousLastLetter) > 0 {
		startsWithNumber, length, lastLetter := StartsWithNumber("", previousLastLetter+calibrationString[0:])
		return startsWithNumber, length - 1, lastLetter
	} else {
		return false, 0, ""
	}
}

func GetNumberStringArray(calibrationString string) []string {
	result := []string{}
	lastLetter := ""

	for len(calibrationString) > 0 {
		isOneDigitNumber := IsOneDigitNumber(string(calibrationString[0]))
		startsWithNumber, length, nextLastLetter := StartsWithNumber(lastLetter, calibrationString)

		if isOneDigitNumber {
			result = append(result, calibrationString[0:1])
			// remove first character
			calibrationString = calibrationString[1:]
			// reset lastLetter
			lastLetter = ""
		} else if startsWithNumber {
			stringNumber := calibrationString[0:length]
			number := DecipherStringNumber(stringNumber, lastLetter)
			result = append(result, number)
			// remove first characters
			calibrationString = calibrationString[length:]
			// set lastLetter
			lastLetter = nextLastLetter
		} else {
			// remove first character
			calibrationString = calibrationString[1:]
			// reset lastLetter
			lastLetter = ""
		}
	}
	return result
}

func DecipherCalibration2(calibrationStrings []string) []int {
	var arrayInput []string

	for _, calibrationString := range calibrationStrings {
		numberStringArray := GetNumberStringArray(calibrationString)

		asString := strings.Join(numberStringArray, "")
		arrayInput = append(arrayInput, asString)
	}

	return DecipherCalibration(arrayInput)
}
