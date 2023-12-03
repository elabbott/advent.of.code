package day_03

import (
	"fmt"
	"strconv"
)

type Gear struct {
	RowIndex    int
	ColumnIndex int
	First       int
	Second      int
	Ratio       uint64
}

type PartNumber struct {
	RowIndex         int
	ColumnIndexStart int
	ColumnIndexEnd   int
	Number           int
	IsGearPart       bool
}

func FindGearsNumbers(schematic [][]string) ([]int, []Gear) {
	result := []int{}
	// partNumbers := []PartNumber{}
	gears := []Gear{}
	numberBuilder := ""

	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			if !IsNumeral(schematic[i][j]) {
				continue
			}

			if IsNumeral(schematic[i][j]) {
				numberBuilder = BuildNumber(schematic, i, j)
				partNumber, err := strconv.Atoi(numberBuilder)
				if err != nil {
					panic(err)
				}
				isPartNumber := false

				for index := 0; index < len(numberBuilder); index++ {
					isPartNumber, gearRowIndex, gearColumnIndex := IsPartNumber(schematic, i, j+index)
					isGearPart := false
					if !(gearColumnIndex == -1 || gearRowIndex == -1) {
						isGearPart = IsStar(schematic[gearRowIndex][gearColumnIndex])
					}
					if isPartNumber {
						// part := PartNumber{
						// 	RowIndex:         i,
						// 	ColumnIndexStart: j,
						// 	ColumnIndexEnd:   j + index,
						// 	Number:           partNumber,
						// 	IsGearPart:       isGearPart,
						// }
						if isGearPart {
							if len(gears) == 0 {
								gears = append(gears, Gear{
									RowIndex:    gearRowIndex,
									ColumnIndex: gearColumnIndex,
									First:       partNumber,
									Second:      -1,
								})
							} else {
								existingGearIndex := CheckIfGearExists(gears, gearRowIndex, gearColumnIndex)
								if existingGearIndex == -1 {
									gears = append(gears, Gear{
										RowIndex:    gearRowIndex,
										ColumnIndex: gearColumnIndex,
										First:       partNumber,
										Second:      -1,
									})
								} else {
									gears[existingGearIndex].Second = partNumber
									gears[existingGearIndex].Ratio = uint64(gears[existingGearIndex].First) * uint64(gears[existingGearIndex].Second)
								}
							}
						}
						// partNumbers = append(partNumbers, part)
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
	}

	return result, gears
}

func IsStar(char string) bool {
	return char == "*"
}

func CheckIfGearExists(gears []Gear, gearRowIndex int, gearColumnIndex int) int {
	for i := 0; i < len(gears); i++ {
		if gears[i].RowIndex == gearRowIndex && gears[i].ColumnIndex == gearColumnIndex {
			return i
		}
	}

	return -1
}

func PrintGears(gears []Gear) {
	for i := 0; i < len(gears); i++ {
		PrintGear(gears[i])
	}
}

func SumGearRatios(gears []Gear) uint64 {
	result := uint64(0)
	for i := 0; i < len(gears); i++ {
		if gears[i].Ratio > 0 {
			result += gears[i].Ratio
		}
	}
	return result
}

func PrintGear(gear Gear) {
	gearString := fmt.Sprintf("Gear Location: [%d, %d], First Number: %d, Second Number: %d, Ratio: %d", gear.RowIndex, gear.ColumnIndex, gear.First, gear.Second, gear.Ratio)
	fmt.Println(gearString)
}

func RemoveInvalidGears(gears []Gear) []Gear {
	result := []Gear{}

	for i := 0; i < len(gears); i++ {
		if gears[i].Second != -1 {
			result = append(result, gears[i])
		}
	}

	return result
}

func PrintTotalGearRatio(totalRatio uint64) {
	fmt.Printf("Total Gear Ratio: %d\n", totalRatio)
}
