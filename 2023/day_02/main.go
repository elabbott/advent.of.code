package main

import (
	day_02 "2023/day_02/helpers"
	global "2023/global"
)

func main() {
	do_01()
	do_02()
}

func do_01() {
	input := global.LoadInput("/home/eric.abbott/.source/elabbott/advent.of.code/2023/day_02/input.txt")

	games := day_02.GetGamesFromStringArray(input)

	println(day_02.SumValidGameIds(games))
}

func do_02() {
	// input := global.LoadInput("/home/eric.abbott/.source/elabbott/advent.of.code/2023/day_02/tests/input.txt")
	// input := global.LoadInput("/home/eric.abbott/.source/elabbott/advent.of.code/2023/day_02/input.txt")

	println("Not implemented yet.")
}
