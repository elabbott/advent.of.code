package day_01

func TransformStringToInt(input string) int {
	return 0
}

func SumCalibration(calibrationNumbers []int) int {
	sum := 0
	for _, calibrationNumber := range calibrationNumbers {
		sum += calibrationNumber
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
