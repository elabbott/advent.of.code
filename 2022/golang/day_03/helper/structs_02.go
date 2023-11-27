package helper

type Group struct {
	Badge     string
	Rucksacks [3]Rucksack
}

func GetGroups(rucksacks []string) []Group {
	groups := make([]Group, len(rucksacks))

	for i, rucksack := range rucksacks {
		for j := 0; j < 3; j++ {
			groups[i].Rucksacks[j] = GetRucksack(rucksack)
		}
		groups[i].Badge = "unknown"
	}

	return groups
}
