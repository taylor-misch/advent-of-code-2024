package day3

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1() {
	fmt.Println("Day 3 Part 1 Begin")
	i := utilities.ReadInput("day3/day03-input.txt")
	r, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	matches := r.FindAllString(i, -1)
	result := 0
	for _, match := range matches {
		fmt.Println(match)
		result += parseNumbers(match)
	}
	fmt.Println("Result: ", result)
	fmt.Println("Day 3 Part 1 End")
}

func Part2() {
	fmt.Println("Day 3 Part 2 Begin")
	i := utilities.ReadInput("day3/day03-input-test02.txt")
	//r, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	//matches := r.FindAllString(i, -1)
	doDontRegex, _ := regexp.Compile("don't\\(\\)")
	doRegex, _ := regexp.Compile("do\\(\\)")
	doNotRegexMatches := doDontRegex.FindAllStringIndex(i, -1)
	doRegexMatches := doRegex.FindAllStringIndex(i, -1)

	// given these matches, I can get rid of all parts of the string between a don't and a do
	// then I can process the final string through part1's solution again

	fmt.Println(doNotRegexMatches)
	fmt.Println(doRegexMatches)
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
