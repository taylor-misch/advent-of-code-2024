package utilities

import (
	"fmt"
	"os"
)

func ReadInput(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		// Handle error
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	// Read the entire file into memory
	data, err := os.ReadFile(filename)
	if err != nil {
		// Handle error
		fmt.Println("Error reading file:", err)
		return ""
	}

	// Print the contents
	//fmt.Println(string(data))
	return string(data)
}

func AbsoluteValue(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
