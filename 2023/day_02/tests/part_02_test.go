package day_02_test

import (
	day_02 "2023/day_02/helpers"
	"testing"
)

func TestDetermineGameMinimums(t *testing.T) {
	type args struct {
		game day_02.Game
	}
	tests := []struct {
		name string
		args args
		want day_02.GameMinimum
	}{
		{"DetermineGameMinimums_1", args{day_02.Game{Id: 1, Sets: []day_02.Set{{Blue: 3, Red: 4, Green: 0, IsValid: true}, {Blue: 6, Red: 1, Green: 2, IsValid: true}, {Blue: 0, Red: 0, Green: 2, IsValid: true}}, IsValid: true}}, day_02.GameMinimum{Blue: 6, Red: 4, Green: 2, Power: 48}},
		{"DetermineGameMinimums_2", args{day_02.Game{Id: 2, Sets: []day_02.Set{{Blue: 1, Red: 0, Green: 2, IsValid: true}, {Blue: 4, Red: 1, Green: 3, IsValid: true}, {Blue: 1, Red: 0, Green: 1, IsValid: true}}, IsValid: true}}, day_02.GameMinimum{Blue: 4, Red: 1, Green: 3, Power: 12}},
		{"DetermineGameMinimums_3", args{day_02.Game{Id: 3, Sets: []day_02.Set{{Blue: 6, Red: 20, Green: 8, IsValid: false}, {Blue: 5, Red: 4, Green: 13, IsValid: true}, {Blue: 0, Red: 1, Green: 5, IsValid: true}}, IsValid: false}}, day_02.GameMinimum{Blue: 6, Red: 20, Green: 13, Power: 1560}},
		{"DetermineGameMinimums_4", args{day_02.Game{Id: 4, Sets: []day_02.Set{{Blue: 6, Red: 3, Green: 1, IsValid: true}, {Blue: 0, Red: 6, Green: 3, IsValid: true}, {Blue: 15, Red: 14, Green: 3, IsValid: false}}, IsValid: false}}, day_02.GameMinimum{Blue: 15, Red: 14, Green: 3, Power: 630}},
		{"DetermineGameMinimums_5", args{day_02.Game{Id: 5, Sets: []day_02.Set{{Blue: 1, Red: 6, Green: 3, IsValid: true}, {Blue: 2, Red: 1, Green: 2, IsValid: true}}, IsValid: true}}, day_02.GameMinimum{Blue: 2, Red: 6, Green: 3, Power: 36}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_02.DetermineGameMinimums(tt.args.game); !compareGameMinimums(got, tt.want) {
				t.Errorf("DetermineGameMinimums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareGameMinimums(got day_02.GameMinimum, want day_02.GameMinimum) bool {
	if got.Blue != want.Blue {
		return false
	}
	if got.Red != want.Red {
		return false
	}
	if got.Green != want.Green {
		return false
	}
	if got.Power != want.Power {
		return false
	}
	return true
}

func TestGetGameMinimums(t *testing.T) {
	type args struct {
		games []day_02.Game
	}
	tests := []struct {
		name string
		args args
		want []day_02.GameMinimum
	}{
		{"GetGameMinimums_1",
			args{[]day_02.Game{
				{Id: 1, Sets: []day_02.Set{{Blue: 3, Red: 4, Green: 0, IsValid: true}, {Blue: 6, Red: 1, Green: 2, IsValid: true}, {Blue: 0, Red: 0, Green: 2, IsValid: true}}, IsValid: true},
				{Id: 2, Sets: []day_02.Set{{Blue: 1, Red: 0, Green: 2, IsValid: true}, {Blue: 4, Red: 1, Green: 3, IsValid: true}, {Blue: 1, Red: 0, Green: 1, IsValid: true}}, IsValid: true},
				{Id: 3, Sets: []day_02.Set{{Blue: 6, Red: 20, Green: 8, IsValid: false}, {Blue: 5, Red: 4, Green: 13, IsValid: true}, {Blue: 0, Red: 1, Green: 5, IsValid: true}}, IsValid: false},
				{Id: 4, Sets: []day_02.Set{{Blue: 6, Red: 3, Green: 1, IsValid: true}, {Blue: 0, Red: 6, Green: 3, IsValid: true}, {Blue: 15, Red: 14, Green: 3, IsValid: false}}, IsValid: false},
				{Id: 5, Sets: []day_02.Set{{Blue: 1, Red: 6, Green: 3, IsValid: true}, {Blue: 2, Red: 1, Green: 2, IsValid: true}}, IsValid: true},
			}}, []day_02.GameMinimum{{Blue: 6, Red: 4, Green: 2, Power: 48}, {Blue: 4, Red: 1, Green: 3, Power: 12}, {Blue: 6, Red: 20, Green: 13, Power: 1560}, {Blue: 15, Red: 14, Green: 3, Power: 630}, {Blue: 2, Red: 6, Green: 3, Power: 36}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_02.GetGameMinimums(tt.args.games); !compareGameMinimumsArray(got, tt.want) {
				t.Errorf("GetGameMinimums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareGameMinimumsArray(got []day_02.GameMinimum, want []day_02.GameMinimum) bool {
	for i := range got {
		if !compareGameMinimums(got[i], want[i]) {
			return false
		}
	}
	return true
}

func TestSumGameMinimumPowers(t *testing.T) {
	type args struct {
		gameMinimums []day_02.GameMinimum
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"SumGameMinimumPowers_1", args{[]day_02.GameMinimum{{Blue: 6, Red: 4, Green: 2, Power: 48}, {Blue: 4, Red: 1, Green: 3, Power: 12}, {Blue: 6, Red: 20, Green: 13, Power: 1560}, {Blue: 15, Red: 14, Green: 3, Power: 630}, {Blue: 2, Red: 6, Green: 3, Power: 36}}}, 2286},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_02.SumGameMinimumPowers(tt.args.gameMinimums); got != tt.want {
				t.Errorf("SumGameMinimumPowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
