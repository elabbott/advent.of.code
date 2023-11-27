package helper

import "strings"

func GetInput(rucksacks []string) Input {
	input := Input{}
	for _, rucksack := range rucksacks {
		input.Rucksacks = append(input.Rucksacks, GetRucksack(rucksack))
	}
	return input
}

func GetRucksack(rucksack string) Rucksack {
	rucksackStruct := Rucksack{}

	length := len(rucksack) / 2

	rucksackStruct.Compartment[0] = GetCompartment(rucksack[:length])
	rucksackStruct.Compartment[1] = GetCompartment(rucksack[length:])
	rucksackStruct.Contents = make([]string, len(rucksack)) // added for part 2

	for i, c := range rucksack {
		rucksackStruct.Contents[i] = string(c)
	}

	return rucksackStruct
}

func GetCompartment(compartment string) Compartment {
	compartmentStruct := Compartment{}
	compartmentStruct.Items = compartment
	return compartmentStruct
}

func GetItems(compartment string) []string {
	return strings.Split(compartment, "")
}

type Input struct {
	Rucksacks []Rucksack
}

type Rucksack struct {
	Compartment [2]Compartment
	Contents    []string // added for part 2
}

type Compartment struct {
	Items string
}
