package day_03_test

import (
	day_03 "2023/day_03/helpers"
	"testing"
)

func TestFindEnginePartNumbers(t *testing.T) {
	type args struct {
		schematic [][]string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// 467..114..
		// ...*......
		// ..35..633.
		// ......#...
		// 617*......
		// .....+.58.
		// ..592.....
		// ......755.
		// ...$.*....
		// .664.598..
		{"FindEnginePartNumbers_467..114.._...*......_..35..633._......#..._617*......_.....+.58._..592....._......755._...$.*...._.664.598..", args{[][]string{
			{"4", "6", "7", ".", ".", "1", "1", "4", ".", "."},
			{".", ".", ".", "*", ".", ".", ".", ".", ".", "."},
			{".", ".", "3", "5", ".", ".", "6", "3", "3", "."},
			{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
			{"6", "1", "7", "*", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", "+", ".", "5", "8", "."},
			{".", ".", "5", "9", "2", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "7", "5", "5", "."},
			{".", ".", ".", "$", ".", "*", ".", ".", ".", "."},
			{".", "6", "6", "4", ".", "5", "9", "8", ".", "."},
		}}, []int{0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_03.FindEnginePartNumbers(tt.args.schematic); !compareSlice(got, tt.want) {
				t.Errorf("FindEnginePartNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range b {
		if a[i] != v {
			return false
		}
	}
	return true
}
