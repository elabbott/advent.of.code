package day_02

// "2023/global"

type GameMinimum struct {
	Blue  int
	Red   int
	Green int
	Power int
	// Game
}

func SumGameMinimumPowers(gameMinimums []GameMinimum) int {
	sum := 0

	for _, gameMinimum := range gameMinimums {
		sum += gameMinimum.Power
	}

	return sum
}

func GetGameMinimums(games []Game) []GameMinimum {
	var gameMinimums []GameMinimum

	for _, game := range games {
		gameMinimums = append(gameMinimums, DetermineGameMinimums(game))
	}

	return gameMinimums
}

func DetermineGameMinimums(game Game) GameMinimum {
	gameMinimum := GameMinimum{
		Blue:  0,
		Red:   0,
		Green: 0,
		Power: 0,
		// Game:  game,
	}

	for _, set := range game.Sets {
		if set.Blue > gameMinimum.Blue || gameMinimum.Blue == 0 {
			gameMinimum.Blue = set.Blue
		}
		if set.Red > gameMinimum.Red || gameMinimum.Red == 0 {
			gameMinimum.Red = set.Red
		}
		if set.Green > gameMinimum.Green || gameMinimum.Green == 0 {
			gameMinimum.Green = set.Green
		}
	}

	gameMinimum.Power = GetMinimumPower(gameMinimum)

	return gameMinimum
}

func GetMinimumPower(gameMinimum GameMinimum) int {
	return gameMinimum.Blue * gameMinimum.Red * gameMinimum.Green
}
