package helper

import (
	"bufio"
	"os"
	"strings"
)

type Input struct {
	Rucksacks []Rucksack
}

type Rucksack struct {
	Compartment [2]Compartment
}

type Compartment struct {
	Items string
}

func DayThree_A(data []string) int {

	rucksacks := GetInput(data)
	sum := 0
	for _, rucksack := range rucksacks.Rucksacks {
		if DeteremineDuplicateItem(rucksack) != "" {
			sum += DeterminePriority(DeteremineDuplicateItem(rucksack))
		}
	}

	return sum
}

func DeteremineDuplicateItem(rucksack Rucksack) string {
	compartment1 := rucksack.Compartment[0].Items
	compartment2 := rucksack.Compartment[1].Items

	items1 := GetItems(compartment1)
	items2 := GetItems(compartment2)

	for _, item1 := range items1 {
		for _, item2 := range items2 {
			if item1 == item2 {
				return item1
			}
		}
	}

	return ""
}

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
	return compartmentStruct
}

func GetItems(compartment string) []string {
	return strings.Split(compartment, "")
}

func DeterminePriority(item string) int {
	// switch item {
	switch item {
	case "a":
		return 1
	case "b":
		return 2
	case "c":
		return 3
	case "d":
		return 4
	case "e":
		return 5
	case "f":
		return 6
	case "g":
		return 7
	case "h":
		return 8
	case "i":
		return 9
	case "j":
		return 10
	case "k":
		return 11
	case "l":
		return 12
	case "m":
		return 13
	case "n":
		return 14
	case "o":
		return 15
	case "p":
		return 16
	case "q":
		return 17
	case "r":
		return 18
	case "s":
		return 19
	case "t":
		return 20
	case "u":
		return 21
	case "v":
		return 22
	case "w":
		return 23
	case "x":
		return 24
	case "y":
		return 25
	case "z":
		return 26
	case "A":
		return 27
	case "B":
		return 28
	case "C":
		return 29
	case "D":
		return 30
	case "E":
		return 31
	case "F":
		return 32
	case "G":
		return 33
	case "H":
		return 34
	case "I":
		return 35
	case "J":
		return 36
	case "K":
		return 37
	case "L":
		return 38
	case "M":
		return 39
	case "N":
		return 40
	case "O":
		return 41
	case "P":
		return 42
	case "Q":
		return 43
	case "R":
		return 44
	case "S":
		return 45
	case "T":
		return 46
	case "U":
		return 47
	case "V":
		return 48
	case "W":
		return 49
	case "X":
		return 50
	case "Y":
		return 51
	case "Z":
		return 52
	default:
		return 0
	}
}

func GetInputFromFile(filePath string) []string {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
