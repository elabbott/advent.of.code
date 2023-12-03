package day_01_test

import (
	day_01 "2023/day_01/helpers"
	"testing"
)

func TestTransformStringToInt(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TransformStringToInt_1abc2", args{"1abc2"}, 12},
		{"TransformStringToInt_pqr3stu8vwx", args{"pqr3stu8vwx"}, 38},
		{"TransformStringToInt_a1b2c3d4e5f", args{"a1b2c3d4e5f"}, 15},
		{"TransformStringToInt_treb7uchet", args{"treb7uchet"}, 77},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_01.TransformStringToInt(tt.args.data); got != tt.want {
				t.Errorf("TransformStringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecipherCalibration(t *testing.T) {
	type args struct {
		calibrationStrings []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"DecipherCalibration_1abc2", args{[]string{"1abc2"}}, []int{12}},
		{"DecipherCalibration_pqr3stu8vwx", args{[]string{"pqr3stu8vwx"}}, []int{38}},
		{"DecipherCalibration_a1b2c3d4e5f", args{[]string{"a1b2c3d4e5f"}}, []int{15}},
		{"DecipherCalibration_treb7uchet", args{[]string{"treb7uchet"}}, []int{77}},
		{"DecipherCalibration_1abc2_pqr3stu8vwx_a1b2c3d4e5f_treb7uchet", args{[]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}}, []int{12, 38, 15, 77}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_01.DecipherCalibration(tt.args.calibrationStrings); !compareIntSlices(got, tt.want) {
				t.Errorf("DecipherCalibration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumCalibration(t *testing.T) {
	type args struct {
		calibrationNumbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"SumCalibration_12", args{[]int{12}}, 12},
		{"SumCalibration_38", args{[]int{38}}, 38},
		{"SumCalibration_15", args{[]int{15}}, 15},
		{"SumCalibration_77", args{[]int{77}}, 77},
		{"SumCalibration_12_38_15_77", args{[]int{12, 38, 15, 77}}, 142},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day_01.SumCalibration(tt.args.calibrationNumbers); got != tt.want {
				t.Errorf("SumCalibration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareIntSlices(a, b []int) bool {
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
