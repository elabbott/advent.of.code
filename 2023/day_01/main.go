package main

import (
	day_01 "2023/day_01/helpers"
	global "2023/global"
)

func main() {
	do_01()
	do_02()
}

func do_01() {
	input := global.LoadInput("/home/eric.abbott/.source/elabbott/advent.of.code/2023/day_01/input.txt")

	deciphered := day_01.DecipherCalibration(input)
	sum := day_01.SumCalibration(deciphered)

	println("Sum of calibration numbers:", sum)
}

func do_02() {
	println("Not implemented")
}
