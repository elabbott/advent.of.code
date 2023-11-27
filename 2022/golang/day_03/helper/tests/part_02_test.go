package helper_test

import (
	"day_03/helper"
	"testing"
)

func TestDayThree_B(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"DayThree_B", args{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}}, 70},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper.DayThree_B(tt.args.data); got != tt.want {
				t.Errorf("DayThree_B() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetermineBadge(t *testing.T) {
	type args struct {
		group helper.Group
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"DetermineBadge", args{helper.Group{
			Badge: "unknown",
			Rucksacks: [3]helper.Rucksack{
				{Compartment: [2]helper.Compartment{{Items: "vJrwpWtwJgWr"}, {Items: "hcsFMMfFFhFp"}}},
				{Compartment: [2]helper.Compartment{{Items: "jqHRNqRjqzjGDLGL"}, {Items: "rsFMfFZSrLrFZsSL"}}},
				{Compartment: [2]helper.Compartment{{Items: "PmmdzqPrV"}, {Items: "vPwwTWBwg"}}},
			},
		}}, "r"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper.DetermineBadge(tt.args.group); got != tt.want {
				t.Errorf("DetermineBadge() = %v, want %v", got, tt.want)
			}
		})
	}
}
