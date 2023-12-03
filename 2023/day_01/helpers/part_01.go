package day_01

import (
	"2023/global"
	"strconv"
)

func IsOneDigitNumber(input string) bool {
	num, err := strconv.Atoi(input)
	if err == nil && num >= 0 && num <= 9 {
		return true
	}
	return false
}

func TransformStringArrayToInt(input []string) int {
	result := 0
	numArray := []string{}

	for _, char := range input {
		if IsOneDigitNumber(char) {
			numArray = append(numArray, char)
		}
	}

	length := len(numArray)

	switch length {
	case 0:
		result = -1
	case 1:
		result, _ = strconv.Atoi(numArray[0] + numArray[0])
	case 2:
		result, _ = strconv.Atoi(numArray[0] + numArray[1])
	default:
		result, _ = strconv.Atoi(numArray[0] + numArray[length-1])
	}

	return result
}

func TransformStringToInt(input string) int {
	arrayInput := global.StringToArray(input)

	return TransformStringArrayToInt(arrayInput)
}

func SumCalibration(calibrationNumbers []int) int {
	sum := 0
	for _, calibrationNumber := range calibrationNumbers {
		if calibrationNumber >= 0 {
			sum += calibrationNumber
		}
	}
	return sum
}

func DecipherCalibration(calibrationStrings []string) []int {
	var calibrationNumbers []int

	for _, calibrationString := range calibrationStrings {
		calibrationNumber := TransformStringToInt(calibrationString)
		calibrationNumbers = append(calibrationNumbers, calibrationNumber)
	}

	return calibrationNumbers
}
