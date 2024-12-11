package main

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
