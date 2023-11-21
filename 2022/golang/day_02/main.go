package main

import (
	"day_02/helper"
	"fmt"
)

func main() {
	data := helper.LoadShard(helper.FilePath)
	score := 0

	if data == nil {
		fmt.Println("Error loading file")
		return
	}

	for _, shard := range data {
		elf := helper.GetGuideValueA(shard[0])
		me := helper.GetGuideValueA(shard[1])

		matchScores := helper.PlayGame(me, elf)
		score += helper.GetPlayerScore(matchScores, 1)
	}

	fmt.Println(score)

	score = 0
	for _, shard := range data {
		elf := shard[0]
		elfPlay := helper.GetGuideValueA(elf)
		matchResult := shard[1]
		me := helper.GetMyPlay(elf, matchResult)

		matchScores := helper.PlayGame(me, elfPlay)

		score += helper.GetPlayerScore(matchScores, 1)
	}

	fmt.Println(score)
}
