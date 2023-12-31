package helper

import (
	"2022/golang/day_03/structs"
	"regexp"
	"sort"
)

func DayThree_B(data []string) int {

	groups := GetGroups(data)

	result := 0
	badgeValues := []int{}

	for _, group := range groups {
		group.Badge = DetermineBadge(group)

		if !ValidateBadge(group.Badge) {
			panic("invalid badge")
		}

		badgeValues = append(badgeValues, DeterminePriority(group.Badge))
	}

	result = SumBadges(badgeValues)

	return result
}

func ValidateBadge(badge string) bool {
	if badge == "badge not found" || badge == "" || badge == " " || len(badge) != 1 {
		return false
	}

	match, _ := regexp.MatchString("[^a-zA-Z]", badge)

	return !match
}

func SumBadges(badges []int) int {

	sum := 0

	for _, badge := range badges {
		sum += badge
	}

	return sum
}

func DetermineBadge(group structs.Group) string {

	uniqueItems := []string{}

	for i := 0; i < len(group.Rucksacks); i++ {
		uniqueItems = append(uniqueItems, GetUniqueContents(group.Rucksacks[i])...)
	}

	sort.Slice(uniqueItems, func(i, j int) bool {
		return uniqueItems[i] < uniqueItems[j]
	})

	for i := 0; i < len(uniqueItems); i++ {
		if uniqueItems[i] == uniqueItems[i+1] && uniqueItems[i] == uniqueItems[i+2] {
			return uniqueItems[i]
		}
	}
	return "badge not found"
}

func GetUniqueContents(rucksack structs.Rucksack) []string {
	unique := []string{}

	for _, item := range rucksack.Compartment {
		for _, compartment := range item.Contents {
			if !Contains(unique, compartment) {
				unique = append(unique, compartment)
			}
		}
	}

	return unique
}

func Contains(unique []string, item string) bool {
	for _, u := range unique {
		if u == item {
			return true
		}
	}
	return false
}

func GetGroups(rucksacks []string) []structs.Group {

	if !structs.ValidateNumberOfRucksacks(rucksacks) {
		panic("invalid number of rucksacks")
	}

	groups := make([]structs.Group, len(rucksacks)/3)
	index := 0

	for g := 0; g < len(groups); g++ {
		for i := 0; i < structs.MaxRucksacks && index < len(rucksacks); i++ {
			groups[g].Rucksacks[i] = structs.GetRucksack(rucksacks[index])
			index++
		}
		groups[g].Badge = DetermineBadge(groups[g])
	}

	return groups
}
