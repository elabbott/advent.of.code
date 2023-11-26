package helper_test

import (
	"day_03/helper"
	"testing"
)

func TestDayThree_A(t *testing.T) {
	type args struct {
		data []string
	}
	// test values:
	// 	vJrwpWtwJgWrhcsFMMfFFhFp
	// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	// PmmdzqPrVvPwwTWBwg
	// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	// ttgJtRGJQctTZtZT
	// CrZsJsPPZsGzwwsLwLmpwMDw
	tests := []struct {
		name string
		args args
		want int
	}{
		{"DayThree_A", args{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}}, 157},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper.DayThree_A(tt.args.data); got != tt.want {
				t.Errorf("DayThree_A() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeteremineDuplicateItem(t *testing.T) {
	type args struct {
		rucksack helper.Rucksack
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"DeteremineDuplicateItem", args{helper.Rucksack{Compartment: [2]helper.Compartment{{Items: "vJrwpWtwJgWr"}, {Items: "hcsFMMfFFhFp"}}}}, "p"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper.DeteremineDuplicateItem(tt.args.rucksack); got != tt.want {
				t.Errorf("DeteremineDuplicateItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInput(t *testing.T) {
	type args struct {
		rucksacks []string
	}
	tests := []struct {
		name string
		args args
		want helper.Input
	}{
		{"GetInput", args{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"}}, helper.Input{Rucksacks: []helper.Rucksack{{Compartment: [2]helper.Compartment{{Items: "vJrwpWtwJgWr"}, {Items: "hcsFMMfFFhFp"}}}, {Compartment: [2]helper.Compartment{{Items: "jqHRNqRjqzjGDLGL"}, {Items: "rsFMfFZSrLrFZsSL"}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper.GetInput(tt.args.rucksacks); got.Rucksacks[0].Compartment[0].Items != tt.want.Rucksacks[0].Compartment[0].Items {
				t.Errorf("GetInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRucksack(t *testing.T) {
	type args struct {
		rucksack string
	}
	tests := []struct {
		name string
		args args
		want helper.Rucksack
	}{
		{"GetRucksack", args{"vJrwpWtwJgWrhcsFMMfFFhFp"}, helper.Rucksack{Compartment: [2]helper.Compartment{{Items: "vJrwpWtwJgWr"}, {Items: "hcsFMMfFFhFp"}}}},
		{"GetRucksack", args{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"}, helper.Rucksack{Compartment: [2]helper.Compartment{{Items: "jqHRNqRjqzjGDLGL"}, {Items: "rsFMfFZSrLrFZsSL"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper.GetRucksack(tt.args.rucksack); got.Compartment[0].Items != tt.want.Compartment[0].Items {
				t.Errorf("GetRucksack() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestGetCompartment(t *testing.T) {
	type args struct {
		compartment string
	}
	tests := []struct {
		name string
		args args
		want helper.Compartment
	}{
		{"GetCompartment", args{"vJrwpWtwJgWr"}, helper.Compartment{Items: "vJrwpWtwJgWr"}},
		{"GetCompartment", args{"hcsFMMfFFhFp"}, helper.Compartment{Items: "hcsFMMfFFhFp"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper.GetCompartment(tt.args.compartment); got.Items != tt.want.Items {
				t.Errorf("GetCompartment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeterminePriority(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"DeterminePriority_a", args{"a"}, 1},
		{"DeterminePriority_b", args{"b"}, 2},
		{"DeterminePriority_c", args{"c"}, 3},
		{"DeterminePriority_d", args{"d"}, 4},
		{"DeterminePriority_e", args{"e"}, 5},
		{"DeterminePriority_f", args{"f"}, 6},
		{"DeterminePriority_g", args{"g"}, 7},
		{"DeterminePriority_h", args{"h"}, 8},
		{"DeterminePriority_i", args{"i"}, 9},
		{"DeterminePriority_j", args{"j"}, 10},
		{"DeterminePriority_k", args{"k"}, 11},
		{"DeterminePriority_l", args{"l"}, 12},
		{"DeterminePriority_m", args{"m"}, 13},
		{"DeterminePriority_n", args{"n"}, 14},
		{"DeterminePriority_o", args{"o"}, 15},
		{"DeterminePriority_p", args{"p"}, 16},
		{"DeterminePriority_q", args{"q"}, 17},
		{"DeterminePriority_r", args{"r"}, 18},
		{"DeterminePriority_s", args{"s"}, 19},
		{"DeterminePriority_t", args{"t"}, 20},
		{"DeterminePriority_u", args{"u"}, 21},
		{"DeterminePriority_v", args{"v"}, 22},
		{"DeterminePriority_w", args{"w"}, 23},
		{"DeterminePriority_x", args{"x"}, 24},
		{"DeterminePriority_y", args{"y"}, 25},
		{"DeterminePriority_z", args{"z"}, 26},
		{"DeterminePriority_A", args{"A"}, 27},
		{"DeterminePriority_B", args{"B"}, 28},
		{"DeterminePriority_C", args{"C"}, 29},
		{"DeterminePriority_D", args{"D"}, 30},
		{"DeterminePriority_E", args{"E"}, 31},
		{"DeterminePriority_F", args{"F"}, 32},
		{"DeterminePriority_G", args{"G"}, 33},
		{"DeterminePriority_H", args{"H"}, 34},
		{"DeterminePriority_I", args{"I"}, 35},
		{"DeterminePriority_J", args{"J"}, 36},
		{"DeterminePriority_K", args{"K"}, 37},
		{"DeterminePriority_L", args{"L"}, 38},
		{"DeterminePriority_M", args{"M"}, 39},
		{"DeterminePriority_N", args{"N"}, 40},
		{"DeterminePriority_O", args{"O"}, 41},
		{"DeterminePriority_P", args{"P"}, 42},
		{"DeterminePriority_Q", args{"Q"}, 43},
		{"DeterminePriority_R", args{"R"}, 44},
		{"DeterminePriority_S", args{"S"}, 45},
		{"DeterminePriority_T", args{"T"}, 46},
		{"DeterminePriority_U", args{"U"}, 47},
		{"DeterminePriority_V", args{"V"}, 48},
		{"DeterminePriority_W", args{"W"}, 49},
		{"DeterminePriority_X", args{"X"}, 50},
		{"DeterminePriority_Y", args{"Y"}, 51},
		{"DeterminePriority_Z", args{"Z"}, 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper.DeterminePriority(tt.args.item); got != tt.want {
				t.Errorf("DeterminePriority() = %v, want %v", got, tt.want)
			}
		})
	}
}
