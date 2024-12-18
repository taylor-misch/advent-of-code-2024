package day03

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) int {
	fmt.Println("Day 3 Part 1 Begin")
	i := ""
	if input == "" {
		i = utilities.ReadInput("day03/day03-input.txt")
	} else {
		i = input
	}
	r, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	matches := r.FindAllString(i, -1)
	result := 0
	for _, match := range matches {
		fmt.Println(match)
		result += parseNumbers(match)
	}
	fmt.Println("Result: ", result)
	fmt.Println("Day 3 Part 1 End")
	return result
}

func Part2() {
	fmt.Println("Day 3 Part 2 Begin")
	i := utilities.ReadInput("day03/day03-input.txt")
	//r, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	//matches := r.FindAllString(i, -1)
	doDontRegex, _ := regexp.Compile("don't\\(\\)")
	doRegex, _ := regexp.Compile("do\\(\\)")
	doNotRegexMatches := doDontRegex.FindAllStringIndex(i, -1)
	doRegexMatches := doRegex.FindAllStringIndex(i, -1)

	// given these matches, I can get rid of all parts of the string between a don't and a do
	// then I can process the final string through part1's solution again

	result := ""
	currentPos := 0
	lastDoConsidered := 0
	lastDontConsidered := 0

	for j := 0; j < len(doNotRegexMatches); j++ {
		if doNotRegexMatches[j][0] >= currentPos {
			for (lastDoConsidered < len(doRegexMatches)) && (doRegexMatches[lastDoConsidered][1] <= doNotRegexMatches[j][0]) {
				lastDoConsidered++
			}
			if lastDoConsidered != len(doRegexMatches) {
				start := doNotRegexMatches[j][0]
				end := doRegexMatches[lastDoConsidered][1]
				result += i[currentPos:start]
				currentPos = end
				lastDontConsidered = j
			}
		}
	}

	if lastDontConsidered < len(doNotRegexMatches) {
		result += i[currentPos:doNotRegexMatches[lastDontConsidered+1][0]]
	} else {
		result += i[currentPos:]
	}

	fmt.Println("Processed String: ", result)
	fmt.Println("Result: ", Part1(result))

	fmt.Println("Day 3 Part 2 End")
}

func parseNumbers(match string) int {
	// this is ugly and I know there is a cleaner way
	num1 := strings.Split(strings.Split(match, "mul(")[1], ",")[0]
	num2 := strings.Split(strings.Split(strings.Split(match, "mul(")[1], ",")[1], ")")[0]

	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)

	return n1 * n2
}
