package main

import (
	day_03 "2022/golang/day_03/helper"
)

func main() {

	doA()

	doB()

}

func doA() {

	var inputData = day_03.GetInputFromFile("./input.example.txt")

	if day_03.DayThree_A(inputData) != 157 {
		panic("DayThree_A() failed, expected 157")
	}
	inputData = day_03.GetInputFromFile("./input.txt")
	println(day_03.DayThree_A(inputData))
}

func doB() {

	var inputData = day_03.GetInputFromFile("./input.example.txt")
	if day_03.DayThree_B(inputData) != 70 {
		panic("DayThree_B() failed, expected 70")
	}
	inputData = day_03.GetInputFromFile("./input.txt")
	println(day_03.DayThree_B(inputData))

}
