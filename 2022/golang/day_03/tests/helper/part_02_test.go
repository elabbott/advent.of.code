package helper_test

import (
	"2022/global"
	"2022/golang/day_03/helper"
	"2022/golang/day_03/structs"
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

func TestValidateBadge(t *testing.T) {
	t.Run("Valid Badge", func(t *testing.T) {
		badge := "A"
		want := true

		got := helper.ValidateBadge(badge)

		if got != want {
			t.Errorf("ValidateBadge(%q) = %v, want %v", badge, got, want)
		}
	})

	t.Run("Invalid Badge - Empty String", func(t *testing.T) {
		badge := ""
		want := false

		got := helper.ValidateBadge(badge)

		if got != want {
			t.Errorf("ValidateBadge(%q) = %v, want %v", badge, got, want)
		}
	})

	t.Run("Invalid Badge - Length != 1", func(t *testing.T) {
		badge := "AB"
		want := false

		got := helper.ValidateBadge(badge)

		if got != want {
			t.Errorf("ValidateBadge(%q) = %v, want %v", badge, got, want)
		}
	})

	t.Run("Invalid Badge - Contains Non-Alphabetic Characters", func(t *testing.T) {
		badge := "1"
		want := false

		got := helper.ValidateBadge(badge)

		if got != want {
			t.Errorf("ValidateBadge(%q) = %v, want %v", badge, got, want)
		}
	})

	t.Run("Invalid Badge - 'badge not found'", func(t *testing.T) {
		badge := "badge not found"
		want := false

		got := helper.ValidateBadge(badge)

		if got != want {
			t.Errorf("ValidateBadge(%q) = %v, want %v", badge, got, want)
		}
	})

	t.Run("Invalid Badge - Space", func(t *testing.T) {
		badge := " "
		want := false

		got := helper.ValidateBadge(badge)

		if got != want {
			t.Errorf("ValidateBadge(%q) = %v, want %v", badge, got, want)
		}
	})
}

func TestDetermineBadge(t *testing.T) {
	type args struct {
		group structs.Group
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"DetermineBadge", args{structs.Group{
			Badge: "unknown",
			Rucksacks: [3]structs.Rucksack{
				{Compartment: [2]structs.Compartment{
					{
						Items:    "vJrwpWtwJgWr",
						Contents: []string{"v", "J", "r", "w", "p", "W", "t", "w", "J", "g", "W", "r"},
					},
					{
						Items:    "hcsFMMfFFhFp",
						Contents: []string{"h", "c", "s", "F", "M", "M", "f", "F", "F", "h", "F", "p"},
					}}},
				{Compartment: [2]structs.Compartment{
					{
						Items:    "jqHRNqRjqzjGDLGL",
						Contents: []string{"j", "q", "H", "R", "N", "q", "R", "j", "q", "z", "j", "G", "D", "L", "G", "L"},
					},
					{
						Items:    "rsFMfFZSrLrFZsSL",
						Contents: []string{"r", "s", "F", "M", "f", "F", "Z", "S", "r", "L", "r", "F", "Z", "s", "S", "L"},
					}}},
				{Compartment: [2]structs.Compartment{
					{
						Items:    "PmmdzqPrV",
						Contents: []string{"P", "m", "m", "d", "z", "q", "P", "r", "V"},
					},
					{
						Items:    "vPwwTWBwg",
						Contents: []string{"v", "P", "w", "w", "T", "W", "B", "w", "g"},
					}}}}}}, "r"},
		{"DetermineBadge", args{structs.Group{
			Badge: "unknown",
			Rucksacks: [3]structs.Rucksack{
				{Compartment: [2]structs.Compartment{
					{
						Items:    "wMqvLMZHhHMvwLH",
						Contents: []string{"w", "M", "q", "v", "L", "M", "Z", "H", "h", "H", "M", "v", "w", "L", "H"},
					},
					{
						Items:    "jbvcjnnSBnvTQFn",
						Contents: []string{"j", "b", "v", "c", "j", "n", "n", "S", "B", "n", "v", "T", "Q", "F", "n"},
					}}},
				{Compartment: [2]structs.Compartment{
					{
						Items:    "ttgJtRGJ",
						Contents: []string{"t", "t", "g", "J", "t", "R", "G", "J"},
					},
					{
						Items:    "QctTZtZT",
						Contents: []string{"Q", "c", "t", "T", "Z", "t", "Z", "T"},
					}}},
				{Compartment: [2]structs.Compartment{
					{
						Items:    "CrZsJsPPZsGz",
						Contents: []string{"C", "r", "Z", "s", "J", "s", "P", "Z", "s", "G", "z"}},
					{
						Items:    "wwsLwLmpwMDw",
						Contents: []string{"w", "w", "s", "L", "w", "L", "m", "p", "w", "M", "D", "w"},
					}}}}}}, "Z"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper.DetermineBadge(tt.args.group); got != tt.want {
				t.Errorf("DetermineBadge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGroups(t *testing.T) {
	type args struct {
		rucksacks []string
	}
	tests := []struct {
		name string
		args args
		want []structs.Group
	}{
		{
			"GetGroups",
			args{[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}},
			[]structs.Group{
				{
					Badge: "r",
					Rucksacks: [3]structs.Rucksack{
						{Compartment: [2]structs.Compartment{{
							Items:    "vJrwpWtwJgWr",
							Contents: []string{"v", "J", "r", "w", "p", "W", "t", "w", "J", "g", "W", "r"},
						},
							{
								Items:    "hcsFMMfFFhFp",
								Contents: []string{"h", "c", "s", "F", "M", "M", "f", "F", "F", "h", "F", "p"},
							}}},
						{Compartment: [2]structs.Compartment{{
							Items:    "jqHRNqRjqzjGDLGL",
							Contents: []string{"j", "q", "H", "R", "N", "q", "R", "j", "q", "z", "j", "G", "D", "L", "G", "L"},
						},
							{
								Items:    "rsFMfFZSrLrFZsSL",
								Contents: []string{"r", "s", "F", "M", "f", "F", "Z", "S", "r", "L", "r", "F", "Z", "s", "S", "L"},
							}}},
						{Compartment: [2]structs.Compartment{{
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
					Rucksacks: [3]structs.Rucksack{
						{Compartment: [2]structs.Compartment{
							{
								Items:    "wMqvLMZHhHMvwLH",
								Contents: []string{"w", "M", "q", "v", "L", "M", "Z", "H", "h", "H", "M", "v", "w", "L", "H"},
							},
							{
								Items:    "jbvcjnnSBnvTQFn",
								Contents: []string{"j", "b", "v", "c", "j", "n", "n", "S", "B", "n", "v", "T", "Q", "F", "n"},
							}}},
						{Compartment: [2]structs.Compartment{
							{
								Items:    "ttgJtRGJ",
								Contents: []string{"t", "t", "g", "J", "t", "R", "G", "J"},
							},
							{
								Items:    "QctTZtZT",
								Contents: []string{"Q", "c", "t", "T", "Z", "t", "Z", "T"},
							}}},
						{Compartment: [2]structs.Compartment{
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
			if !groupResults.IsSuccess {
				t.Error(groupResults.Message)
			}
		})
	}
}

func compareGroups(got, want []structs.Group) global.TestResult {

	for i := 0; i < len(got); i++ {
		// badges
		if got[i].Badge != want[i].Badge {
			return global.TestResult{IsSuccess: false, Message: "badges do not match, got: " + got[i].Badge + ", want: " + want[i].Badge}
		}

		// rucksacks
		for j := 0; j < len(got[i].Rucksacks); j++ {
			// compartments
			for k := 0; k < len(got[i].Rucksacks[j].Compartment); k++ {
				// items
				if got[i].Rucksacks[j].Compartment[k].Items != want[i].Rucksacks[j].Compartment[k].Items {
					return global.TestResult{IsSuccess: false, Message: "items do not match, got: " + got[i].Rucksacks[j].Compartment[k].Items + ", want: " + want[i].Rucksacks[j].Compartment[k].Items}
				}

				// contents
				for l := 0; l < len(got[i].Rucksacks[j].Compartment[k].Contents); l++ {
					if got[i].Rucksacks[j].Compartment[k].Contents[l] != want[i].Rucksacks[j].Compartment[k].Contents[l] {
						return global.TestResult{IsSuccess: false, Message: "contents do not match, got: " + got[i].Rucksacks[j].Compartment[k].Contents[l] + ", want: " + want[i].Rucksacks[j].Compartment[k].Contents[l]}
					}
				}
			}
		}
	}

	return global.TestResult{IsSuccess: true, Message: "success"}
}
