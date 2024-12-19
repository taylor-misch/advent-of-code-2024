package day07

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	i := utilities.ReadInput("day07/day07-input.txt")
	functions := parseInput(i)
	fmt.Println(functions)

	runningTotal := int64(0)
	for _, function := range functions {
		if checkCalculationPart1(function.numbers, function.result, 1, function.numbers[0]) {
			runningTotal += function.result
		}
	}
	fmt.Println("Running total: ", runningTotal)
}

func Part2() {
	i := utilities.ReadInput("day07/day07-input.txt")
	functions := parseInput(i)
	fmt.Println(functions)

	runningTotal := int64(0)
	for _, function := range functions {
		if checkCalculationPart2(function.numbers, function.result, 1, function.numbers[0]) {
			runningTotal += function.result
		}
	}
	fmt.Println("Running total: ", runningTotal)
}

func parseInput(data string) []Function {
	lines := strings.Split(data, "\n")
	functions := []Function{}
	for _, line := range lines {
		lineParts := strings.Split(line, ":")
		result, _ := strconv.ParseInt(lineParts[0], 10, 64)
		parts := strings.Fields(lineParts[1])
		numbers := []int64{}
		for _, part := range parts {
			number, _ := strconv.ParseInt(part, 10, 64)
			numbers = append(numbers, number)
		}
		functions = append(functions, Function{result, numbers})
	}

	return functions
}

func checkCalculationPart1(numbers []int64, result int64, index int, runningTotal int64) bool {
	if index == len(numbers) {
		return runningTotal == result
	}
	runningTotal1 := runningTotal + numbers[index]
	runningTotal2 := runningTotal * numbers[index]
	index++
	checkCalc1 := checkCalculationPart1(numbers, result, index, runningTotal1)
	checkCalc2 := checkCalculationPart1(numbers, result, index, runningTotal2)
	return checkCalc1 || checkCalc2

}

func checkCalculationPart2(numbers []int64, result int64, index int, runningTotal int64) bool {
	if index == len(numbers) {
		return runningTotal == result
	}
	runningTotal1 := runningTotal + numbers[index]
	runningTotal2 := runningTotal * numbers[index]
	runningTotal3, _ := strconv.ParseInt(strconv.FormatInt(runningTotal, 10)+strconv.FormatInt(numbers[index], 10), 10, 64)
	index++
	checkCalc1 := checkCalculationPart2(numbers, result, index, runningTotal1)
	checkCalc2 := checkCalculationPart2(numbers, result, index, runningTotal2)
	checkCalc3 := checkCalculationPart2(numbers, result, index, runningTotal3)
	return checkCalc1 || checkCalc2 || checkCalc3

}

type Function struct {
	result  int64
	numbers []int64
}
