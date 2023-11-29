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
				{Compartment: [2]helper.Compartment{
					{
						Items:    "vJrwpWtwJgWr",
						Contents: []string{"v", "J", "r", "w", "p", "W", "t", "w", "J", "g", "W", "r"},
					},
					{
						Items:    "hcsFMMfFFhFp",
						Contents: []string{"h", "c", "s", "F", "M", "M", "f", "F", "F", "h", "F", "p"},
					}}},
				{Compartment: [2]helper.Compartment{
					{
						Items:    "jqHRNqRjqzjGDLGL",
						Contents: []string{"j", "q", "H", "R", "N", "q", "R", "j", "q", "z", "j", "G", "D", "L", "G", "L"},
					},
					{
						Items:    "rsFMfFZSrLrFZsSL",
						Contents: []string{"r", "s", "F", "M", "f", "F", "Z", "S", "r", "L", "r", "F", "Z", "s", "S", "L"},
					}}},
				{Compartment: [2]helper.Compartment{
					{
						Items:    "PmmdzqPrV",
						Contents: []string{"P", "m", "m", "d", "z", "q", "P", "r", "V"},
					},
					{
						Items:    "vPwwTWBwg",
						Contents: []string{"v", "P", "w", "w", "T", "W", "B", "w", "g"},
					}}}}}}, "r"},
		{"DetermineBadge", args{helper.Group{
			Badge: "unknown",
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
