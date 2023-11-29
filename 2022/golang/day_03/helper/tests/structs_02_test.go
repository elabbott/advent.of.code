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
					Badge: "r",
					Rucksacks: [3]helper.Rucksack{
						{Compartment: [2]helper.Compartment{{
							Items:    "vJrwpWtwJgWr",
							Contents: []string{"v", "J", "r", "w", "p", "W", "t", "w", "J", "g", "W", "r"},
						},
							{
								Items:    "hcsFMMfFFhFp",
								Contents: []string{"h", "c", "s", "F", "M", "M", "f", "F", "F", "h", "F", "p"},
							}}},
						{Compartment: [2]helper.Compartment{{
							Items:    "jqHRNqRjqzjGDLGL",
							Contents: []string{"j", "q", "H", "R", "N", "q", "R", "j", "q", "z", "j", "G", "D", "L", "G", "L"},
						},
							{
								Items:    "rsFMfFZSrLrFZsSL",
								Contents: []string{"r", "s", "F", "M", "f", "F", "Z", "S", "r", "L", "r", "F", "Z", "s", "S", "L"},
							}}},
						{Compartment: [2]helper.Compartment{{
							Items:    "PmmdzqPrV",
							Contents: []string{"P", "m", "m", "d", "z", "q", "P", "r", "V"},
						},
							{
								Items:    "vPwwTWBwg",
								Contents: []string{"v", "P", "w", "w", "T", "W", "B", "w", "g"},
							}}},
					},
				},
				{
					Badge: "Z",
					Rucksacks: [3]helper.Rucksack{
						{Compartment: [2]helper.Compartment{
							{
								Items:    "wMqvLMZHhHMvwLH",
								Contents: []string{"w", "M", "q", "v", "L", "M", "Z", "H", "h", "H", "M", "v", "w", "L", "H"},
							},
							{
								Items:    "jbvcjnnSBnvTQFn",
								Contents: []string{"j", "b", "v", "c", "j", "n", "n", "S", "B", "n", "v", "T", "Q", "F", "n"},
							}}},
						{Compartment: [2]helper.Compartment{
							{
								Items:    "ttgJtRGJ",
								Contents: []string{"t", "t", "g", "J", "t", "R", "G", "J"},
							},
							{
								Items:    "QctTZtZT",
								Contents: []string{"Q", "c", "t", "T", "Z", "t", "Z", "T"},
							}}},
						{Compartment: [2]helper.Compartment{
							{
								Items:    "CrZsJsPPZsGz",
								Contents: []string{"C", "r", "Z", "s", "J", "s", "P", "P", "Z", "s", "G", "z"}},
							{
								Items:    "wwsLwLmpwMDw",
								Contents: []string{"w", "w", "s", "L", "w", "L", "m", "p", "w", "M", "D", "w"},
							}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := helper.GetGroups(tt.args.rucksacks)

			groupResults := compareGroups(got, tt.want)
			if !groupResults.isSuccess {
				t.Error(groupResults.message)
			}
		})
	}
}

func compareGroups(got, want []helper.Group) testResult {

	for i := 0; i < len(got); i++ {
		// badges
		if got[i].Badge != want[i].Badge {
			return testResult{isSuccess: false, message: "badges do not match, got: " + got[i].Badge + ", want: " + want[i].Badge}
		}

		// rucksacks
		for j := 0; j < len(got[i].Rucksacks); j++ {
			// compartments
			for k := 0; k < len(got[i].Rucksacks[j].Compartment); k++ {
				// items
				if got[i].Rucksacks[j].Compartment[k].Items != want[i].Rucksacks[j].Compartment[k].Items {
					return testResult{isSuccess: false, message: "items do not match, got: " + got[i].Rucksacks[j].Compartment[k].Items + ", want: " + want[i].Rucksacks[j].Compartment[k].Items}
				}

				// contents
				for l := 0; l < len(got[i].Rucksacks[j].Compartment[k].Contents); l++ {
					if got[i].Rucksacks[j].Compartment[k].Contents[l] != want[i].Rucksacks[j].Compartment[k].Contents[l] {
						return testResult{isSuccess: false, message: "contents do not match, got: " + got[i].Rucksacks[j].Compartment[k].Contents[l] + ", want: " + want[i].Rucksacks[j].Compartment[k].Contents[l]}
					}
				}
			}
		}
	}

	return testResult{isSuccess: true, message: "success"}
}
