package helper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
	Win      = 6
	Lose     = 0
	Draw     = 3

	RockString     = "A"
	PaperString    = "B"
	ScissorsString = "C"

	RockString2     = "X"
	PaperString2    = "Y"
	ScissorsString2 = "Z"

	LoseString = "X"
	DrawString = "Y"
	WinString  = "Z"

	FilePath = "./input.txt"
)

func GetMyPlay(elf string, matchResult string) int {

	switch matchResult {
	case LoseString: // Lose
		return getLosingPlay(elf)
	case DrawString: // Draw
		return getDrawPlay(elf)
	case WinString: // Win
		return getWinningPlay(elf)
	default:
		return 0
	}
}

func getWinningPlay(elf string) int {

	switch elf {
	case RockString: // Rock
		return Paper
	case PaperString: // Paper
		return Scissors
	case ScissorsString: // Scissors
		return Rock
	default:
		return 0
	}
}

func getLosingPlay(elf string) int {

	switch elf {
	case RockString: // Rock
		return Scissors
	case PaperString: // Paper
		return Rock
	case ScissorsString: // Scissors
		return Paper
	default:
		return 0
	}
}

func getDrawPlay(elf string) int {

	switch elf {
	case RockString: // Rock
		return Rock
	case PaperString: // Paper
		return Paper
	case ScissorsString: // Scissors
		return Scissors
	default:
		return 0
	}
}

func GetGuideValueA(value string) int {
	switch value {
	case RockString:
		return Rock
	case PaperString:
		return Paper
	case ScissorsString:
		return Scissors
	case RockString2:
		return Rock
	case PaperString2:
		return Paper
	case ScissorsString2:
		return Scissors
	default:
		return 0
	}
}

func PlayGame(player1, player2 int) [2]int {
	if player1 == player2 {
		return [2]int{Draw + player1, Draw + player2}
	}

	if player1 == Rock && player2 == Scissors {
		return [2]int{Win + player1, Lose + player2}
	}

	if player1 == Paper && player2 == Rock {
		return [2]int{Win + player1, Lose + player2}
	}

	if player1 == Scissors && player2 == Paper {
		return [2]int{Win + player1, Lose + player2}
	}

	return [2]int{Lose + player1, Win + player2}
}

func GetPlayerScore(scores [2]int, player int) int {
	return scores[player-1]
}

func LoadSlice(filePath string) [][2]string {
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
