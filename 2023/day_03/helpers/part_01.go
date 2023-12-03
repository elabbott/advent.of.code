package day_03

import (
	"strconv"
)

func BuildSchematic(input []string) [][]string {
	result := [][]string{}

	for i := 0; i < len(input); i++ {
		result = append(result, []string{})
		for j := 0; j < len(input[i]); j++ {
			result[i] = append(result[i], string(input[i][j]))
		}
	}

	return result
}

func SumPartNumbers(partNumbers []int) int {
	result := 0

	for i := 0; i < len(partNumbers); i++ {
		result += partNumbers[i]
	}

	return result
}

func FindEnginePartNumbers(schematic [][]string) []int {
	result := []int{}
	numberBuilder := ""

	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			if !IsNumeral(schematic[i][j]) {
				continue
			}

			numberBuilder = BuildNumber(schematic, i, j)
			partNumber, err := strconv.Atoi(numberBuilder)
			if err != nil {
				panic(err)
			}
			isPartNumber := false

			for index := 0; index < len(numberBuilder); index++ {
				isPartNumber, _, _ = IsPartNumber(schematic, i, j+index)
				if isPartNumber {
					break
				}
			}

			if isPartNumber {
				result = append(result, partNumber)
			}

			j += len(numberBuilder) - 1
			numberBuilder = ""
		}
	}

	return result
}

func BuildNumber(schematic [][]string, row int, column int) string {
	number := schematic[row][column]

	for IsSafeToCheckRight(schematic, row, column) {
		if IsNumeral(schematic[row][column+1]) {
			number += schematic[row][column+1]
			column++
		} else {
			break
		}
	}

	return number
}

func IsPartNumber(schematic [][]string, row int, column int) (bool, int, int) {
	if IsSafeToCheckLeft(schematic, row, column) {
		if IsSpecialCharacter(schematic[row][column-1]) {
			return true, row, column - 1
		}
	}
	if IsSafeToCheckRight(schematic, row, column) {
		if IsSpecialCharacter(schematic[row][column+1]) {
			return true, row, column + 1
		}
	}
	if IsSafeToCheckUp(schematic, row, column) {
		if IsSpecialCharacter(schematic[row-1][column]) {
			return true, row - 1, column
		}
	}
	if IsSafeToCheckDown(schematic, row, column) {
		if IsSpecialCharacter(schematic[row+1][column]) {
			return true, row + 1, column
		}
	}
	if IsSafeToCheckUpLeft(schematic, row, column) {
		if IsSpecialCharacter(schematic[row-1][column-1]) {
			return true, row - 1, column - 1
		}
	}
	if IsSafeToCheckUpRight(schematic, row, column) {
		if IsSpecialCharacter(schematic[row-1][column+1]) {
			return true, row - 1, column + 1
		}
	}
	if IsSafeToCheckDownLeft(schematic, row, column) {
		if IsSpecialCharacter(schematic[row+1][column-1]) {
			return true, row + 1, column - 1
		}
	}
	if IsSafeToCheckDownRight(schematic, row, column) {
		if IsSpecialCharacter(schematic[row+1][column+1]) {
			return true, row + 1, column + 1
		}
	}

	return false, -1, -1
}

func IsSafeToCheckLeft(schematic [][]string, rowIndex int, columnIndex int) bool {
	return columnIndex > 0
}

func IsSafeToCheckRight(schematic [][]string, rowIndex int, columnIndex int) bool {
	return columnIndex < len(schematic[rowIndex])-1
}

func IsSafeToCheckUp(schematic [][]string, rowIndex int, columnIndex int) bool {
	return rowIndex > 0
}

func IsSafeToCheckDown(schematic [][]string, rowIndex int, columnIndex int) bool {
	return rowIndex < len(schematic)-1
}

func IsSafeToCheckUpLeft(schematic [][]string, rowIndex int, columnIndex int) bool {
	return rowIndex > 0 && columnIndex > 0
}

func IsSafeToCheckUpRight(schematic [][]string, rowIndex int, columnIndex int) bool {
	return rowIndex > 0 && columnIndex < len(schematic[rowIndex])-1
}

func IsSafeToCheckDownLeft(schematic [][]string, rowIndex int, columnIndex int) bool {
	return rowIndex < len(schematic)-1 && columnIndex > 0
}

func IsSafeToCheckDownRight(schematic [][]string, rowIndex int, columnIndex int) bool {
	return rowIndex < len(schematic)-1 && columnIndex < len(schematic[rowIndex])-1
}

func IsNumeral(char string) bool {
	return char == "0" || char == "1" || char == "2" || char == "3" || char == "4" || char == "5" || char == "6" || char == "7" || char == "8" || char == "9"
}

func IsPeriod(char string) bool {
	return char == "."
}

func IsSpecialCharacter(char string) bool {
	return !(IsNumeral(char) || IsPeriod(char))
}

func PrintSchematic(schematic [][]string) {
	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			print(schematic[i][j])
		}
		println()
	}
}

func PrintPartNumbers(partNumbers []int) {
	for i := 0; i < len(partNumbers); i++ {
		println(partNumbers[i])
	}
}

func PrintSum(sum int) {
	println(sum)
}
