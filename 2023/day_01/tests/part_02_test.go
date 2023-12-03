package day_01_test

import (
	day_01 "2023/day_01/helpers"
	"testing"
)

func TestDecipherStringNumber(t *testing.T) {
	type args struct {
		calibrationString string
		lastLetter        string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"DecipherStringNumber_zero", args{"zero", ""}, "0"},
		{"DecipherStringNumber_one", args{"one", ""}, "1"},
		{"DecipherStringNumber_two", args{"two", ""}, "2"},
		{"DecipherStringNumber_three", args{"three", ""}, "3"},
		{"DecipherStringNumber_four", args{"four", ""}, "4"},
		{"DecipherStringNumber_five", args{"five", ""}, "5"},
		{"DecipherStringNumber_six", args{"six", ""}, "6"},
		{"DecipherStringNumber_seven", args{"seven", ""}, "7"},
		{"DecipherStringNumber_eight", args{"eight", ""}, "8"},
		{"DecipherStringNumber_nine", args{"nine", ""}, "9"},
		{"DecipherStringNumber_abc", args{"abc", ""}, ""},
		{"DecipherStringNumber_123", args{"123", ""}, ""},
		{"DecipherStringNumber_abc123", args{"ten", ""}, ""},
		{"DecipherStringNumber_two", args{"wo", "t"}, "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_01.DecipherStringNumber(tt.args.calibrationString, tt.args.lastLetter); got != tt.want {
				t.Errorf("DecipherStringNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartsWithNumber(t *testing.T) {
	type args struct {
		previousLastLetter string
		calibrationString  string
	}
	tests := []struct {
		name           string
		args           args
		wantIsNumber   bool
		wantNumDigits  int
		wantLastLetter string
	}{
		{"StartsWithNumber_zero", args{"", "zero"}, false, 0, ""},
		{"StartsWithNumber_one", args{"", "one"}, true, 3, "e"},
		{"StartsWithNumber_two", args{"", "two"}, true, 3, "o"},
		{"StartsWithNumber_three", args{"", "three"}, true, 5, "e"},
		{"StartsWithNumber_four", args{"", "four"}, true, 4, "r"},
		{"StartsWithNumber_five", args{"", "five"}, true, 4, "e"},
		{"StartsWithNumber_six", args{"", "six"}, true, 3, "x"},
		{"StartsWithNumber_seven", args{"", "seven"}, true, 5, "n"},
		{"StartsWithNumber_eight", args{"", "eight"}, true, 5, "t"},
		{"StartsWithNumber_nine", args{"", "nine"}, true, 4, "e"},
		{"StartsWithNumber_abc", args{"", "abc"}, false, 0, ""},
		{"StartsWithNumber_123", args{"", "123"}, false, 0, ""},
		{"StartsWithNumber_abc123", args{"", "abc123"}, false, 0, ""},
		{"StartsWithNumber_123abc", args{"", "123abc"}, false, 0, ""},
		{"StartsWithNumber_two", args{"t", "wo"}, true, 2, "o"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsNumber, gotNumDigits, gotLastLetter := day_01.StartsWithNumber(tt.args.previousLastLetter, tt.args.calibrationString)
			if gotIsNumber != tt.wantIsNumber {
				t.Errorf("StartsWithNumber() gotIsNumber = %v, want %v", gotIsNumber, tt.wantIsNumber)
			}
			if gotNumDigits != tt.wantNumDigits {
				t.Errorf("StartsWithNumber() gotNumDigits = %v, want %v", gotNumDigits, tt.wantNumDigits)
			}
			if gotLastLetter != tt.wantLastLetter {
				t.Errorf("StartsWithNumber() gotLastLetter = %v, want %v", gotLastLetter, tt.wantLastLetter)
			}
		})
	}
}

func TestGetNumberStringArray(t *testing.T) {
	type args struct {
		calibrationString string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"GetNumberStringArray_two1nine", args{"two1nine"}, []string{"2", "1", "9"}},
		// {"GetNumberStringArray_eightwothree", args{"eightwothree"}, []string{"8", "3"}},
		{"GetNumberStringArray_eightwothree", args{"eightwothree"}, []string{"8", "2", "3"}},
		{"GetNumberStringArray_abcone2threexyz", args{"abcone2threexyz"}, []string{"1", "2", "3"}},
		// {"GetNumberStringArray_xtwone3four", args{"xtwone3four"}, []string{"2", "3", "4"}},
		{"GetNumberStringArray_xtwone3four", args{"xtwone3four"}, []string{"2", "1", "3", "4"}},
		{"GetNumberStringArray_4nineeightseven2", args{"4nineeightseven2"}, []string{"4", "9", "8", "7", "2"}},
		// {"GetNumberStringArray_zoneight234", args{"zoneight234"}, []string{"1", "2", "3", "4"}},
		{"GetNumberStringArray_zoneight234", args{"zoneight234"}, []string{"1", "8", "2", "3", "4"}},
		{"GetNumberStringArray_7pqrstsixteen", args{"7pqrstsixteen"}, []string{"7", "6"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_01.GetNumberStringArray(tt.args.calibrationString); !compareStringSlices(got, tt.want) {
				t.Errorf("GetNumberStringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareStringSlices(a, b []string) bool {
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

func TestDecipherCalibration2(t *testing.T) {
	type args struct {
		calibrationStrings []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"DecipherCalibration2_two1nine", args{[]string{"two1nine"}}, []int{29}},
		{"DecipherCalibration2_eightwothree", args{[]string{"eightwothree"}}, []int{83}},
		{"DecipherCalibration2_abcone2threexyz", args{[]string{"abcone2threexyz"}}, []int{13}},
		{"DecipherCalibration2_xtwone3four", args{[]string{"xtwone3four"}}, []int{24}},
		{"DecipherCalibration2_4nineeightseven2", args{[]string{"4nineeightseven2"}}, []int{42}},
		{"DecipherCalibration2_zoneight234", args{[]string{"zoneight234"}}, []int{14}},
		{"DecipherCalibration2_7pqrstsixteen", args{[]string{"7pqrstsixteen"}}, []int{76}},
		{"DecipherCalibration2_two1nine_eightwothree_abcone2threexyz_xtwone3four_4nineeightseven2_zoneight234_7pqrstsixteen", args{[]string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}}, []int{29, 83, 13, 24, 42, 14, 76}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_01.DecipherCalibration2(tt.args.calibrationStrings); !compareIntSlices(got, tt.want) {
				t.Errorf("DecipherCalibration2() = %v, want %v", got, tt.want)
			}
		})
	}
}
