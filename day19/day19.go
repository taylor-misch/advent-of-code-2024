package day19

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strings"
)

func Part1() []string {
	i := utilities.ReadInput("day19/day19-input.txt")
	availablePatterns, patternsToMake := parseInput(i)
	fmt.Println("Available Patterns: ", availablePatterns)
	fmt.Println("Patterns to make: ", patternsToMake)

	possibleToMakeCount := 0
	possiblePatterns := []string{}
	for _, pattern := range patternsToMake {
		if makePatternPart1(pattern, availablePatterns) {
			possibleToMakeCount++
			possiblePatterns = append(possiblePatterns, pattern)
			fmt.Println("Pattern: ", pattern, " is possible to make")
		} else {
			fmt.Println("Pattern: ", pattern, " is not possible to make")

		}
	}

	fmt.Println("answer to part 1: ", possibleToMakeCount)
	return possiblePatterns
}

func Part2() {
	i := utilities.ReadInput("day19/day19-input.txt")
	availablePatterns, patternsToMake := parseInput(i)
	fmt.Println("Available Patterns: ", availablePatterns)
	fmt.Println("Patterns to make: ", patternsToMake)

	possibleToMakeCount := 0
	possiblePatterns := Part1()
	solutions := make(map[string]int)
	for _, pattern := range possiblePatterns {
		possibleToMakeCount += makePatternPart2(pattern, availablePatterns, solutions)
	}

	fmt.Println("answer to part 2: ", possibleToMakeCount)
}

func parseInput(i string) ([]string, []string) {
	lines := strings.Split(i, "\n")

	availablePatterns := strings.Split(lines[0], ", ")

	return availablePatterns, lines[2 : len(lines)-1]
}

func makePatternPart1(pattern string, availablePatterns []string) bool {
	if len(pattern) == 0 {
		return true
	}

	for _, availablePattern := range availablePatterns {
		if strings.HasPrefix(pattern, availablePattern) {
			fmt.Println("For pattern: ", pattern, " ", availablePattern, " works")
			if makePatternPart1(pattern[len(availablePattern):], availablePatterns) {
				return true
			}
		}
	}
	return false
}

func makePatternPart2(pattern string, availableSubPatterns []string, solutions map[string]int) int {
	if countOfWays, ok := solutions[pattern]; ok {
		return countOfWays
	}

	countOfWays := 0

	for _, subPattern := range availableSubPatterns {
		if len(subPattern) > len(pattern) {
			continue
		}
		if strings.HasPrefix(pattern, subPattern) {
			if len(subPattern) == len(pattern) {
				countOfWays++
				continue
			}
			countOfWays += makePatternPart2(pattern[len(subPattern):], availableSubPatterns, solutions)
		}
	}
	solutions[pattern] = countOfWays
	return countOfWays

}
