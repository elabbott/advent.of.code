package globalhelpers

import (
	"bufio"
	"fmt"
	"os"
)

func LoadInput(filePath string) []string {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return input
}
