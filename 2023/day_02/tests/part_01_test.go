package day_02_test

import (
	day_02 "2023/day_02/helpers"
	"testing"
)

func TestGetSetFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want day_02.Set
	}{
		// 3 blue, 4 red
		// 1 red, 2 green, 6 blue
		// 2 green
		// 8 green, 6 blue, 20 red
		{"GetSetFromString_3_blue_4_red", args{"3 blue, 4 red"}, day_02.Set{Blue: 3, Red: 4, Green: 0, IsValid: true}},
		{"GetSetFromString_1_red_2_green_6_blue", args{"1 red, 2 green, 6 blue"}, day_02.Set{Blue: 6, Red: 1, Green: 2, IsValid: true}},
		{"GetSetFromString_2_green", args{"2 green"}, day_02.Set{Blue: 0, Red: 0, Green: 2, IsValid: true}},
		{"GetSetFromString_8_green_6_blue_20_red", args{"8 green, 6 blue, 20 red"}, day_02.Set{Blue: 6, Red: 20, Green: 8, IsValid: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_02.GetSetFromString(tt.args.s); got != tt.want {
				t.Errorf("GetSetFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSetsFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []day_02.Set
	}{
		// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		// 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		// 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		// 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		// 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
		{"GetSetsFromString_3_blue_4_red_1_red_2_green_6_blue_2_green", args{"3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"}, []day_02.Set{
			{Blue: 3, Red: 4, Green: 0, IsValid: true},
			{Blue: 6, Red: 1, Green: 2, IsValid: true},
			{Blue: 0, Red: 0, Green: 2, IsValid: true},
		}},
		{"GetSetsFromString_1_blue_2_green_3_green_4_blue_1_red_1_green_1_blue", args{"1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"}, []day_02.Set{
			{Blue: 1, Red: 0, Green: 2, IsValid: true},
			{Blue: 4, Red: 1, Green: 3, IsValid: true},
			{Blue: 1, Red: 0, Green: 1, IsValid: true},
		}},
		{"GetSetsFromString_8_green_6_blue_20_red_5_blue_4_red_13_green_5_green_1_red", args{"8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"}, []day_02.Set{
			{Blue: 6, Red: 20, Green: 8, IsValid: false},
			{Blue: 5, Red: 4, Green: 13, IsValid: true},
			{Blue: 0, Red: 1, Green: 5, IsValid: true},
		}},
		{"GetSetsFromString_1_green_3_red_6_blue_3_green_6_red_3_green_15_blue_14_red", args{"1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"}, []day_02.Set{
			{Blue: 6, Red: 3, Green: 1, IsValid: true},
			{Blue: 0, Red: 6, Green: 3, IsValid: true},
			{Blue: 15, Red: 14, Green: 3, IsValid: false},
		}},
		{"GetSetsFromString_6_red_1_blue_3_green_2_blue_1_red_2_green", args{"6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}, []day_02.Set{
			{Blue: 1, Red: 6, Green: 3, IsValid: true},
			{Blue: 2, Red: 1, Green: 2, IsValid: true},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_02.GetSetsFromString(tt.args.s); !compareSets(got, tt.want) {
				t.Errorf("GetSetsFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareSets(a, b []day_02.Set) bool {
	if len(a) != len(b) {
		return false
	}
	for index, value := range a {
		if value != b[index] {
			return false
		}
	}
	return true
}

func TestGetGameFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want day_02.Game
	}{
		//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
		{"GetGameFromString_Game_1", args{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"}, day_02.Game{
			Id:      1,
			Sets:    []day_02.Set{{Blue: 3, Red: 4, Green: 0, IsValid: true}, {Blue: 6, Red: 1, Green: 2, IsValid: true}, {Blue: 0, Red: 0, Green: 2, IsValid: true}},
			IsValid: true,
		}},
		{"GetGameFromString_Game_2", args{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"}, day_02.Game{
			Id:      2,
			Sets:    []day_02.Set{{Blue: 1, Red: 0, Green: 2, IsValid: true}, {Blue: 4, Red: 1, Green: 3, IsValid: true}, {Blue: 1, Red: 0, Green: 1, IsValid: true}},
			IsValid: true,
		}},
		{"GetGameFromString_Game_3", args{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"}, day_02.Game{
			Id:      3,
			Sets:    []day_02.Set{{Blue: 6, Red: 20, Green: 8, IsValid: false}, {Blue: 5, Red: 4, Green: 13, IsValid: true}, {Blue: 0, Red: 1, Green: 5, IsValid: true}},
			IsValid: false,
		}},
		{"GetGameFromString_Game_4", args{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"}, day_02.Game{
			Id:      4,
			Sets:    []day_02.Set{{Blue: 6, Red: 3, Green: 1, IsValid: true}, {Blue: 0, Red: 6, Green: 3, IsValid: true}, {Blue: 15, Red: 14, Green: 3, IsValid: false}},
			IsValid: false,
		}},
		{"GetGameFromString_Game_5", args{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}, day_02.Game{
			Id:      5,
			Sets:    []day_02.Set{{Blue: 1, Red: 6, Green: 3, IsValid: true}, {Blue: 2, Red: 1, Green: 2, IsValid: true}},
			IsValid: true,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day_02.GetGameFromString(tt.args.s)
			if got.Id != tt.want.Id {
				t.Errorf("GetGameFromString() = %v, want %v", got.Id, tt.want.Id)
			}
			if !compareSets(got.Sets, tt.want.Sets) {
				t.Errorf("GetGameFromString() = %v, want %v", got.Sets, tt.want.Sets)
			}
			if got.IsValid != tt.want.IsValid {
				t.Errorf("GetGameFromString() = %v, want %v", got.IsValid, tt.want.IsValid)
			}
		})
	}
}

func TestGetGamesFromStringArray(t *testing.T) {
	type args struct {
		sa []string
	}
	tests := []struct {
		name string
		args args
		want []day_02.Game
	}{
		{"GetGamesFromStringArray_Game_1_Game_2_Game_3_Game_4_Game_5", args{[]string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		}}, []day_02.Game{
			{Id: 1, Sets: []day_02.Set{{Blue: 3, Red: 4, Green: 0, IsValid: true}, {Blue: 6, Red: 1, Green: 2, IsValid: true}, {Blue: 0, Red: 0, Green: 2, IsValid: true}}, IsValid: true},
			{Id: 2, Sets: []day_02.Set{{Blue: 1, Red: 0, Green: 2, IsValid: true}, {Blue: 4, Red: 1, Green: 3, IsValid: true}, {Blue: 1, Red: 0, Green: 1, IsValid: true}}, IsValid: true},
			{Id: 3, Sets: []day_02.Set{{Blue: 6, Red: 20, Green: 8, IsValid: false}, {Blue: 5, Red: 4, Green: 13, IsValid: true}, {Blue: 0, Red: 1, Green: 5, IsValid: true}}, IsValid: false},
			{Id: 4, Sets: []day_02.Set{{Blue: 6, Red: 3, Green: 1, IsValid: true}, {Blue: 0, Red: 6, Green: 3, IsValid: true}, {Blue: 15, Red: 14, Green: 3, IsValid: false}}, IsValid: false},
			{Id: 5, Sets: []day_02.Set{{Blue: 1, Red: 6, Green: 3, IsValid: true}, {Blue: 2, Red: 1, Green: 2, IsValid: true}}, IsValid: true},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_02.GetGamesFromStringArray(tt.args.sa); !compareGames(got, tt.want) {
				t.Errorf("GetGamesFromStringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareGames(a, b []day_02.Game) bool {
	if len(a) != len(b) {
		return false
	}
	for index, value := range a {
		if value.Id != b[index].Id {
			return false
		}
		if !compareSets(value.Sets, b[index].Sets) {
			return false
		}
		if value.IsValid != b[index].IsValid {
			return false
		}
	}
	return true
}

func TestSumValidGameIds(t *testing.T) {
	type args struct {
		games []day_02.Game
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"SumValidGameIds_Game_1_Game_2_Game_3_Game_4_Game_5", args{[]day_02.Game{
			{Id: 1, Sets: []day_02.Set{{Blue: 3, Red: 4, Green: 0, IsValid: true}, {Blue: 6, Red: 1, Green: 2, IsValid: true}, {Blue: 0, Red: 0, Green: 2, IsValid: true}}, IsValid: true},
			{Id: 2, Sets: []day_02.Set{{Blue: 1, Red: 0, Green: 2, IsValid: true}, {Blue: 4, Red: 1, Green: 3, IsValid: true}, {Blue: 1, Red: 0, Green: 1, IsValid: true}}, IsValid: true},
			{Id: 3, Sets: []day_02.Set{{Blue: 6, Red: 20, Green: 8, IsValid: false}, {Blue: 5, Red: 4, Green: 13, IsValid: true}, {Blue: 0, Red: 1, Green: 5, IsValid: true}}, IsValid: false},
			{Id: 4, Sets: []day_02.Set{{Blue: 6, Red: 3, Green: 1, IsValid: true}, {Blue: 0, Red: 6, Green: 3, IsValid: true}, {Blue: 15, Red: 14, Green: 3, IsValid: false}}, IsValid: false},
			{Id: 5, Sets: []day_02.Set{{Blue: 1, Red: 6, Green: 3, IsValid: true}, {Blue: 2, Red: 1, Green: 2, IsValid: true}}, IsValid: true},
		}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_02.SumValidGameIds(tt.args.games); got != tt.want {
				t.Errorf("SumValidGameIds() = %v, want %v", got, tt.want)
			}
		})
	}
}
