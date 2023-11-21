package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := loadShard(filePath)
	score := 0

	if data == nil {
		fmt.Println("Error loading file")
		return
	}

	for _, shard := range data {
		elf := getGuideValueA(shard[0])
		me := getGuideValueA(shard[1])
		score += playGame(me, elf)
	}

	fmt.Println(score)

	score = 0
	for _, shard := range data {
		elf := shard[0]
		elfPlay := getGuideValueA(elf)
		matchResult := shard[1]
		me := getMyPlay(elf, matchResult)

		score += playGame(me, elfPlay)
	}

	fmt.Println(score)
}

const (
	rock     = 1
	paper    = 2
	scissors = 3
	win      = 6
	lose     = 0
	draw     = 3
	filePath = "./input.txt"
)

func getMyPlay(elf string, matchResult string) int {

	switch matchResult {
	case "X": // Lose
		return getLosingPlay(elf)
	case "Y": // Draw
		return getDrawPlay(elf)
	case "Z": // Win
		return getWinningPlay(elf)
	default:
		return 0
	}
}

func getWinningPlay(elf string) int {

	switch elf {
	case "A": // Rock
		return paper
	case "B": // Paper
		return scissors
	case "C": // Scissors
		return rock
	default:
		return 0
	}
}

func getLosingPlay(elf string) int {

	switch elf {
	case "A": // Rock
		return scissors
	case "B": // Paper
		return rock
	case "C": // Scissors
		return paper
	default:
		return 0
	}
}

func getDrawPlay(elf string) int {

	switch elf {
	case "A": // Rock
		return rock
	case "B": // Paper
		return paper
	case "C": // Scissors
		return scissors
	default:
		return 0
	}
}

func getGuideValueA(value string) int {
	switch value {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	case "X":
		return rock
	case "Y":
		return paper
	case "Z":
		return scissors
	default:
		return 0
	}
}

func playGame(player1, player2 int) int {
	if player1 == player2 {
		return draw + player1
	}

	if player1 == rock && player2 == scissors {
		return win + player1
	}

	if player1 == paper && player2 == rock {
		return win + player1
	}

	if player1 == scissors && player2 == paper {
		return win + player1
	}

	return lose + player1
}

func loadShard(filePath string) [][2]string {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var guide [][2]string
	var game [2]string

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		match := strings.Split(line, " ")
		game[0] = match[0]
		game[1] = match[1]

		guide = append(guide, game)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return guide
}
