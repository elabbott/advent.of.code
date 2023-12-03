package main

import (
	day_03 "2023/day_03/helpers"
	global "2023/global"
)

func main() {
	// do_01()
	do_02()
}

func do_01() {
	// input := global.LoadInput("/home/eric.abbott/.source/elabbott/advent.of.code/2023/day_03/tests/input.txt")
	input := global.LoadInput("/home/eric.abbott/.source/elabbott/advent.of.code/2023/day_03/input.txt")

	schematic := day_03.BuildSchematic(input)

	// day_03.PrintSchematic(schematic)

	partNumbers := day_03.FindEnginePartNumbers(schematic)

	// day_03.PrintPartNumbers(partNumbers)

	sum := day_03.SumPartNumbers(partNumbers)

	day_03.PrintSum(sum)
}

func do_02() {
	// input := global.LoadInput("/home/eric.abbott/.source/elabbott/advent.of.code/2023/day_03/tests/input.txt")
	input := global.LoadInput("/home/eric.abbott/.source/elabbott/advent.of.code/2023/day_03/input.txt")

	schematic := day_03.BuildSchematic(input)

	// day_03.PrintSchematic(schematic)

	_, gears := day_03.FindGearsNumbers(schematic)

	gears = day_03.RemoveInvalidGears(gears)

	day_03.PrintGears(gears)

	sum := day_03.SumGearRatios(gears)

	day_03.PrintTotalGearRatio(sum)
}
