package day2

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	fmt.Println("Part 1 Begin")
	i := utilities.ReadInput("day2/day02-input.txt")
	reports := parseInput(i)

	safeReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		} else {
			fmt.Println("Unsafe Report: ", report)
		}
	}

	fmt.Println("Safe Reports: ", safeReports)
	fmt.Println("Part 1 End")
}

func Part2() {
	fmt.Println("Part 2 Begin")
	i := utilities.ReadInput("day2/day02-input-test.txt")
	reports := parseInput(i)

	safeReports := 0
	for _, report := range reports {
		if isReportSafeV2(report) {
			safeReports++
		} else {
			fmt.Println("Unsafe Report: ", report)
		}
	}

	fmt.Println("Safe Reports: ", safeReports)
	fmt.Println("Part 2 End")
}

func parseInput(data string) [][]int {
	lines := strings.Split(data, "\n")
	reports := [][]int{}

	for _, line := range lines {
		parts := strings.Fields(line)
		report := []int{}
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				continue
			}
			report = append(report, num)
		}

		reports = append(reports, report)
	}
	return reports
}

func isReportSafe(report []int) bool {
	reportLength := len(report)
	if !isIncreasing(report) && !isDecreasing(report) {
		return false
	}
	for i := 0; i < reportLength-1; i++ {
		absoluteDiff := utilities.AbsoluteValue(report[i] - report[i+1])
		if absoluteDiff < 1 || absoluteDiff > 3 {
			return false
		}
	}
	return true
}

func isIncreasing(report []int) bool {
	reportLength := len(report)
	for i := 0; i < reportLength-1; i++ {
		if report[i] > report[i+1] {
			return false
		}
	}
	return true
}

func isDecreasing(report []int) bool {
	reportLength := len(report)
	for i := 0; i < reportLength-1; i++ {
		if report[i] < report[i+1] {
			return false
		}
	}
	return true
}

// part2
func isReportSafeV2(report []int) bool {
	reportLength := len(report)
	issuesFound := 0

	// it can only be one direction or the other can't count issues from both
	//issuesFound += isIncreasingV2(report)
	//issuesFound += isDecreasingV2(report)

	// there could be overlap in the issues found causing an over count

	increasing := report[0] < report[reportLength-1]

	for i := 0; i < reportLength-1; i++ {
		trendingCorrectly := increasing == (report[i] < report[i+1])
		if !trendingCorrectly {
			fmt.Println("Trending Incorrectly: ", report[i], report[i+1])
		}
		absoluteDiff := utilities.AbsoluteValue(report[i] - report[i+1])
		if absoluteDiff < 1 || absoluteDiff > 3 || !trendingCorrectly {
			fmt.Println("Issue found: ", report[i], report[i+1])
			issuesFound++
		}
	}
	return issuesFound < 2
}

func isIncreasingV2(report []int) int {
	issuesFound := 0
	reportLength := len(report)
	for i := 0; i < reportLength-1; i++ {
		if report[i] > report[i+1] {
			issuesFound++
		}
	}
	return issuesFound
}

func isDecreasingV2(report []int) int {
	issuesFound := 0
	reportLength := len(report)
	for i := 0; i < reportLength-1; i++ {
		if report[i] < report[i+1] {
			issuesFound++
		}
	}
	return issuesFound
}
