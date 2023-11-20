package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	filePath := "./input.txt"

	numbers := loadArray(filePath)

	if numbers == nil {
		fmt.Println("Error loading file")
		return
	}

	answer := findNLargest(numbers, 2)
	fmt.Println(answer)

	// Reset the numbers array
	numbers = loadArray(filePath)

	answer2 := findNLargest(numbers, 3)
	fmt.Println(answer2)
}

func findNLargest(numbers [][]int, n int) int {
	var biggestNumberGroup []int
	var savedNumbers [][]int
	var zeroArray []int

	zeroArray = append(zeroArray, 0)

	for i := 0; i < n; i++ {
		biggestJ := 0
		biggestJ = findIndexOfLargestArraySum(numbers)
		sum := sumArray(numbers[biggestJ])
		biggestNumberGroup = append(biggestNumberGroup, sum)
		savedNumbers = append(savedNumbers, numbers[biggestJ])

		fmt.Println(savedNumbers)
		numbers[biggestJ] = zeroArray
	}

	return sumArray(biggestNumberGroup)
}

func sumArray(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func findIndexOfLargestArraySum(numbers [][]int) int {
	sum := 0
	innerSum := 0
	i := 0
	biggestI := 0

	for _, numberGroup := range numbers {
		innerSum = sumArray(numberGroup)

		if innerSum > sum {
			sum = innerSum
			biggestI = i
		}
		i++
		innerSum = 0
	}

	return biggestI
}

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
