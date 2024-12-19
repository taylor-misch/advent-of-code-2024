package day02

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	fmt.Println("Part 1 Begin")
	i := utilities.ReadInput("day02/day02-input.txt")
	reports := parseInput(i)

	safeReports := 0
	for _, report := range reports {
		if isReportSafePart1(report) {
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
	i := utilities.ReadInput("day02/day02-input.txt")
	reports := parseInput(i)

	safeReports := 0
	for _, report := range reports {
		if isReportSafePar2(report) {
			safeReports++
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

func isReportSafePart1(report []int) bool {
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
func isReportSafePar2(report []int) bool {

	if isReportSafePart1(report) {
		fmt.Println("Full report safe: ", report)
		return true
	} else {
		fmt.Println("Checking Pieces of Report: ", report)
		for i := 0; i < len(report); i++ {
			shortenedReport := append([]int{}, report[:i]...)
			shortenedReport = append(shortenedReport, report[i+1:]...)
			if isReportSafePart1(shortenedReport) {
				fmt.Println("Safe Report Found: ", shortenedReport)
				return true
			}
		}
	}
	return false
}
