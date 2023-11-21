package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

}

const (
	rock     = 1
	paper    = 2
	scissors = 3
	win      = 6
	lose     = 0
	draw     = 3
)

func loadArray(filePath string) [][]int {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers [][]int
	var numberGroups []int
	i := 0
	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Convert the line to an integer
		number, err := strconv.Atoi(line)
		if err != nil {
			i++
			fmt.Sprintf("Error converting line to integer: %s. Starting group %d", err, i)

			numbers = append(numbers, numberGroups)
			numberGroups = []int{}
			continue
		}
		// Append the number to the array
		numberGroups = append(numberGroups, number)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return numbers
}
