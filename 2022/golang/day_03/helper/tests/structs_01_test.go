package helper_test

import (
	"day_03/helper"
	"testing"
)

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
		{"GetRucksack",
			args{"vJrwpWtwJgWrhcsFMMfFFhFp"},
			helper.Rucksack{
				Compartment: [2]helper.Compartment{
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
			helper.Rucksack{
				Compartment: [2]helper.Compartment{
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
			got := helper.GetRucksack(tt.args.rucksack)

			testResult := compareCompartments(got.Compartment, tt.want.Compartment)
			if !testResult.isSuccess { //got.Compartment[0].Items != tt.want.Compartment[0].Items {
				t.Error(testResult.message)
				// t.Errorf("GetRucksack() = %v, want %v", got, tt.want)
			}
		})
	}

}

type testResult struct {
	isSuccess bool
	message   string
}

func compareCompartments(got [2]helper.Compartment, want [2]helper.Compartment) testResult {
	testResult := testResult{isSuccess: true, message: "success"}
	if len(got) != len(want) {
		return testResult
	}

	for i := 0; i < len(got); i++ {
		if got[i].Items != want[i].Items {
			testResult.isSuccess = false
			testResult.message = "items do not match, got: " + got[i].Items + ", want: " + want[i].Items
			return testResult
		}

		testResult = compareArrays(got[i].Contents, want[i].Contents)
		if !testResult.isSuccess {
			return testResult
		}
	}

	return testResult
}

func compareArrays(got []string, want []string) testResult {
	testResult := testResult{isSuccess: true, message: "success"}
	if len(got) != len(want) {
		testResult.isSuccess = false
		testResult.message = "lengths do not match, got: " + string(rune(len(got))) + ", want: " + string(rune(len(want)))
		return testResult
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			testResult.isSuccess = false
			testResult.message = "items do not match, got: " + got[i] + ", want: " + want[i]
			return testResult
		}
	}

	return testResult
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
