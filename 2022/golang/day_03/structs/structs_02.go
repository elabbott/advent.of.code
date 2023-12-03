package structs

const MaxRucksacks int = 3

type Group struct {
	Badge     string
	Rucksacks [MaxRucksacks]Rucksack
}

func ValidateNumberOfRucksacks(rucksacks []string) bool {
	return len(rucksacks)%3 == 0
}
