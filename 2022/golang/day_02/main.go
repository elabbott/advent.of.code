package main

import (
	"2022/golang/day_02/helper"
	"fmt"
)

func main() {
	data := helper.LoadSlice(helper.FilePath)
	score := 0

	if data == nil {
		fmt.Println("Error loading file")
		return
	}

	for _, slice := range data {
		elf := helper.GetGuideValueA(slice[0])
		me := helper.GetGuideValueA(slice[1])

		matchScores := helper.PlayGame(me, elf)
		score += helper.GetPlayerScore(matchScores, 1)
	}

	fmt.Println(score)

	score = 0
	for _, slice := range data {
		elf := slice[0]
		elfPlay := helper.GetGuideValueA(elf)
		matchResult := slice[1]
		me := helper.GetMyPlay(elf, matchResult)

		matchScores := helper.PlayGame(me, elfPlay)

		score += helper.GetPlayerScore(matchScores, 1)
	}

	fmt.Println(score)
}
