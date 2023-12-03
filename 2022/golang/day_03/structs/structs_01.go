package structs

import (
	"strings"
)

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

	return rucksackStruct
}

func GetCompartment(compartment string) Compartment {
	compartmentStruct := Compartment{}
	compartmentStruct.Items = compartment
	compartmentStruct.Contents = make([]string, len(compartment)) // added for part 2

	for i, c := range compartment {
		compartmentStruct.Contents[i] = string(c)
	}
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
}

type Compartment struct {
	Items    string
	Contents []string // added for part 2
}
