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
