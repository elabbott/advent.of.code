package main

import (
	day_03 "day_03/helper"
)

func main() {

	var inputData = day_03.GetInputFromFile("./input.example.txt")

	if day_03.DayThree_A(inputData) != 157 {
		panic("DayThree_A() failed, expected 157")
	}

	inputData = day_03.GetInputFromFile("./input.txt")
	println(day_03.DayThree_A(inputData))

}
