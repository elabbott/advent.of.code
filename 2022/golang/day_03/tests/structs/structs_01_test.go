package structs_test

import (
	"2022/global"
	"2022/golang/day_03/structs"
	"testing"
)

func TestGetInput(t *testing.T) {
	type args struct {
		rucksacks []string
	}
	tests := []struct {
		name string
		args args
		want structs.Input
	}{
		{"GetInput", args{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"}}, structs.Input{Rucksacks: []structs.Rucksack{{Compartment: [2]structs.Compartment{{Items: "vJrwpWtwJgWr"}, {Items: "hcsFMMfFFhFp"}}}, {Compartment: [2]structs.Compartment{{Items: "jqHRNqRjqzjGDLGL"}, {Items: "rsFMfFZSrLrFZsSL"}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := structs.GetInput(tt.args.rucksacks); got.Rucksacks[0].Compartment[0].Items != tt.want.Rucksacks[0].Compartment[0].Items {
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
		want structs.Rucksack
	}{
		{"GetRucksack",
			args{"vJrwpWtwJgWrhcsFMMfFFhFp"},
			structs.Rucksack{
				Compartment: [2]structs.Compartment{
					{
						Items:    "vJrwpWtwJgWr",
						Contents: []string{"v", "J", "r", "w", "p", "W", "t", "w", "J", "g", "W", "r"},
					},
					{
						Items:    "hcsFMMfFFhFp",
						Contents: []string{"h", "c", "s", "F", "M", "M", "f", "F", "F", "h", "F", "p"},
					}},
			}},
		{"GetRucksack",
			args{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
			structs.Rucksack{
				Compartment: [2]structs.Compartment{
					{
						Items:    "jqHRNqRjqzjGDLGL",
						Contents: []string{"j", "q", "H", "R", "N", "q", "R", "j", "q", "z", "j", "G", "D", "L", "G", "L"},
					},
					{
						Items:    "rsFMfFZSrLrFZsSL",
						Contents: []string{"r", "s", "F", "M", "f", "F", "Z", "S", "r", "L", "r", "F", "Z", "s", "S", "L"},
					}},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := structs.GetRucksack(tt.args.rucksack)

			TestResult := compareCompartments(got.Compartment, tt.want.Compartment)
			if !TestResult.IsSuccess { //got.Compartment[0].Items != tt.want.Compartment[0].Items {
				t.Error(TestResult.Message)
				// t.Errorf("GetRucksack() = %v, want %v", got, tt.want)
			}
		})
	}

}

func compareCompartments(got [2]structs.Compartment, want [2]structs.Compartment) global.TestResult {
	TestResult := global.TestResult{IsSuccess: true, Message: "success"}
	if len(got) != len(want) {
		return TestResult
	}

	for i := 0; i < len(got); i++ {
		if got[i].Items != want[i].Items {
			TestResult.IsSuccess = false
			TestResult.Message = "items do not match, got: " + got[i].Items + ", want: " + want[i].Items
			return TestResult
		}

		TestResult = compareArrays(got[i].Contents, want[i].Contents)
		if !TestResult.IsSuccess {
			return TestResult
		}
	}

	return TestResult
}

func compareArrays(got []string, want []string) global.TestResult {
	TestResult := global.TestResult{IsSuccess: true, Message: "success"}
	if len(got) != len(want) {
		TestResult.IsSuccess = false
		TestResult.Message = "lengths do not match, got: " + string(rune(len(got))) + ", want: " + string(rune(len(want)))
		return TestResult
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			TestResult.IsSuccess = false
			TestResult.Message = "items do not match, got: " + got[i] + ", want: " + want[i]
			return TestResult
		}
	}

	return TestResult
}

func TestGetCompartment(t *testing.T) {
	type args struct {
		compartment string
	}
	tests := []struct {
		name string
		args args
		want structs.Compartment
	}{
		{"GetCompartment", args{"vJrwpWtwJgWr"}, structs.Compartment{Items: "vJrwpWtwJgWr"}},
		{"GetCompartment", args{"hcsFMMfFFhFp"}, structs.Compartment{Items: "hcsFMMfFFhFp"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := structs.GetCompartment(tt.args.compartment); got.Items != tt.want.Items {
				t.Errorf("GetCompartment() = %v, want %v", got, tt.want)
			}
		})
	}
}
