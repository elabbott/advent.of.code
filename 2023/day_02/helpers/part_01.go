package day_02

import (
	"strconv"
	"strings"
)

const (
	RED   = 12
	GREEN = 13
	BLUE  = 14
)

type Game struct {
	Id      int
	Sets    []Set
	IsValid bool
}

type Set struct {
	Blue    int
	Red     int
	Green   int
	IsValid bool
}

func SumValidGameIds(games []Game) int {
	sum := 0

	for _, game := range games {
		if game.IsValid {
			sum += game.Id
		}
	}

	return sum
}

func GetGamesFromStringArray(sa []string) []Game {
	var games []Game

	for _, s := range sa {
		games = append(games, GetGameFromString(s))
	}

	return games
}

func GetGameFromString(s string) Game {
	var game Game

	index := strings.Index(s, ":")
	if index != -1 {
		substring := s[:index]
		split := strings.Split(strings.TrimSpace(substring), " ")
		if len(split) != 2 {
			panic("Invalid game string")
		}
		game.Id, _ = strconv.Atoi(split[1])

		game.Sets = GetSetsFromString(s[index+1:])

		game.IsValid = true

		for _, set := range game.Sets {
			if !set.IsValid {
				game.IsValid = false
				break
			}
		}
	}
	return game
}

func GetSetsFromString(s string) []Set {
	var sets []Set

	setsAsStrings := strings.Split(s, ";")

	for _, set := range setsAsStrings {
		sets = append(sets, GetSetFromString(set))
	}

	return sets
}

func GetSetFromString(s string) Set {
	var set Set
	set.Blue = 0
	set.Red = 0
	set.Green = 0

	cubeArray := strings.Split(s, ",")

	for _, cube := range cubeArray {
		if strings.Contains(cube, "red") {
			cubeObj := strings.Split(strings.TrimSpace(cube), " ")

			numberOfCubes, err := strconv.Atoi(cubeObj[0])
			if err != nil {
				panic(err)
			} else {
				set.Red = numberOfCubes
			}
		} else if strings.Contains(cube, "green") {
			cubeObj := strings.Split(strings.TrimSpace(cube), " ")

			numberOfCubes, err := strconv.Atoi(cubeObj[0])
			if err != nil {
				panic(err)
			} else {
				set.Green = numberOfCubes
			}
		} else if strings.Contains(cube, "blue") {
			cubeObj := strings.Split(strings.TrimSpace(cube), " ")

			numberOfCubes, err := strconv.Atoi(cubeObj[0])
			if err != nil {
				panic(err)
			} else {
				set.Blue = numberOfCubes
			}
		} else {
			continue
		}
	}

	if set.Blue > BLUE || set.Red > RED || set.Green > GREEN {
		set.IsValid = false
	} else {
		set.IsValid = true
	}

	return set
}
