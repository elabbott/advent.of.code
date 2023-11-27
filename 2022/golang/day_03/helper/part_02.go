package helper

func DayThree_B(data []string) int {

	return 0
}

func DetermineBadge(group Group) string {

	badge := "unknown"

	items := []string{}

	foundItems := []string{}

	for i := 0; i < len(group.Rucksacks); i++ {
		group.Rucksacks[i].Contents = RemoveDuplicateItems(group.Rucksacks[i])
		for _, item := range group.Rucksacks[i].Contents {
			if !Contains(items, item) {
				items = append(items, item)
			} else if Contains(foundItems, item) {
				badge = item
				break
			} else {
				foundItems = append(foundItems, item)
				continue
			}
		}
	}
	return badge
}

func RemoveDuplicateItems(rucksack Rucksack) []string {
	unique := []string{}

	for _, item := range rucksack.Contents {
		if !Contains(unique, item) {
			unique = append(unique, item)
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
