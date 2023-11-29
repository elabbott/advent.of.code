package helper

const maxRucksacks int = 3

type Group struct {
	Badge     string
	Rucksacks [maxRucksacks]Rucksack
}

func GetGroups(rucksacks []string) []Group {

	if !ValidateNumberOfRucksacks(rucksacks) {
		panic("invalid number of rucksacks")
	}

	groups := make([]Group, len(rucksacks)/3)
	index := 0

	for g := 0; g < len(groups); g++ {
		for i := 0; i < maxRucksacks && index < len(rucksacks); i++ {
			groups[g].Rucksacks[i] = GetRucksack(rucksacks[index])
			index++
		}
		groups[g].Badge = DetermineBadge(groups[g])
	}

	return groups
}

func ValidateNumberOfRucksacks(rucksacks []string) bool {
	return len(rucksacks)%3 == 0
}
