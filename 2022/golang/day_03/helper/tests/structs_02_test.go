package helper_test

import (
	"day_03/helper"
	"testing"
)

func TestGetGroups(t *testing.T) {
	type args struct {
		rucksacks []string
	}
	tests := []struct {
		name string
		args args
		want []helper.Group
	}{
		{
			"GetGroups",
			args{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}},
			[]helper.Group{
				{
					Badge: "unknown",
					Rucksacks: [3]helper.Rucksack{
						{Compartment: [2]helper.Compartment{{Items: "vJrwpWtwJgWr"}, {Items: "hcsFMMfFFhFp"}}},
						{Compartment: [2]helper.Compartment{{Items: "jqHRNqRjqzjGDLGL"}, {Items: "rsFMfFZSrLrFZsSL"}}},
						{Compartment: [2]helper.Compartment{{Items: "PmmdzqPrV"}, {Items: "vPwwTWBwg"}}},
					},
				},
				{
					Badge: "unknown",
					Rucksacks: [3]helper.Rucksack{
						{Compartment: [2]helper.Compartment{{Items: "wMqvLMZHhHMvwLH"}, {Items: "jbvcjnnSBnvTQFn"}}},
						{Compartment: [2]helper.Compartment{{Items: "ttgJtRGJ"}, {Items: "QctTZtZT"}}},
						{Compartment: [2]helper.Compartment{{Items: "CrZsJsPPZsGz"}, {Items: "wwsLwLmpwMDw"}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper.GetGroups(tt.args.rucksacks); got[0].Rucksacks[0].Compartment[0].Items != tt.want[0].Rucksacks[0].Compartment[0].Items {
				t.Errorf("GetGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
