package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	filePath := "./input.txt"

	numbers := loadShard(filePath)

	if numbers == nil {
		fmt.Println("Error loading file")
		return
	}

	var array = convertShardToArray(numbers)

	findNLargest(numbers, 1)

	// Reset the numbers array
	// numbers = loadShard(filePath)

	numbers = array

	findNLargest(numbers, 3)

}

func findNLargest(numbers [][]int, n int) {
	var biggestNumberGroup []int
	var zeroArray []int

	zeroArray = append(zeroArray, 0)

	for i := 0; i < n; i++ {
		biggestJ := 0
		biggestJ = findIndexOfLargestArraySum(numbers)
		sum := sumArray(numbers[biggestJ])
		biggestNumberGroup = append(biggestNumberGroup, sum)

		numbers[biggestJ] = zeroArray
	}

	result := sumArray(biggestNumberGroup)

	fmt.Println(formatResult(result, n))
}

func formatResult(result int, n int) string {
	return fmt.Sprintf("Sum of top %d is %d", n, result)
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

func loadShard(filePath string) [][]int {
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

func convertShardToArray(shard [][]int) [][]int {
	array := make([][]int, len(shard))

	for i, subArray := range shard {
		array[i] = make([]int, len(subArray))
		copy(array[i], subArray)
	}

	return array
}
