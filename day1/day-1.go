package day1

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	fmt.Println("Part 1 Begin")
	i := utilities.ReadInput("day1/day-1-input.txt")
	list1, list2 := parseInput(i)
	sortedList1 := sortList(list1)
	sortedList2 := sortList(list2)

	totalDistance := 0
	for i := 0; i < len(sortedList1); i++ {
		totalDistance += absoluteValue(sortedList2[i] - sortedList1[i])
	}
	fmt.Println("Total Distance: ", totalDistance)
	fmt.Println("Part 1 End")
}

func Part2() {
	fmt.Println("Part 2 Begin")
	i := utilities.ReadInput("day1/day-1-input.txt")
	list1, list2 := parseInput(i)

	mapOfNumberCount := make(map[int]int)
	for i := 0; i < len(list1); i++ {
		mapOfNumberCount[list2[i]]++
	}

	totalDistance := 0
	for i := 0; i < len(list1); i++ {
		totalDistance += list1[i] * mapOfNumberCount[list1[i]]
	}

	fmt.Println("Total Distance: ", totalDistance)
	fmt.Println("Part 2 End")

}

func parseInput(data string) ([]int, []int) {
	lines := strings.Split(data, "\n")
	list1 := []int{}
	list2 := []int{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	return list1, list2
}

func sortList(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				temp := list[i]
				list[i] = list[j]
				list[j] = temp
			}
		}
	}
	return list
}

func absoluteValue(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
