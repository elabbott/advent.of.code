package day_01

import (
	"strconv"
)

func TransformStringToInt(input string) int {
	result := 0

	arrayInput := stringToArray(input)
	numArray := []string{}

	for _, char := range arrayInput {

		num, err := strconv.Atoi(char)
		if err == nil && num >= 0 && num <= 9 {
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

func stringToArray(input string) []string {
	arrayResult := []string{}
	for _, char := range input {
		arrayResult = append(arrayResult, string(char))
	}
	return arrayResult
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
