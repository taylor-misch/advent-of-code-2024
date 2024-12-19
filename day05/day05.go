package day05

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1() {
	i := utilities.ReadInput("day05/day05-input.txt")
	rules, printUpdates := parseInput(i)
	fmt.Println("Rules: ", rules)
	fmt.Println("Print updates: ", printUpdates)

	ruleMap := createRuleMap(rules)
	fmt.Println("Rule Map: ", ruleMap)
	validPrintUpdates := make([]PrintUpdate, 0)
	for _, printUpdate := range printUpdates {
		if validatePrintOrder(ruleMap, printUpdate) {
			validPrintUpdates = append(validPrintUpdates, printUpdate)
		}
	}

	fmt.Println("Valid Print Updates: ", validPrintUpdates)

	sumOfMiddleNumbers := 0
	for _, printUpdate := range validPrintUpdates {
		middleIndex := math.Ceil(float64(len(printUpdate.updates) / 2))
		sumOfMiddleNumbers += printUpdate.updates[int(middleIndex)]
	}

	fmt.Println("Sum of middle numbers: ", sumOfMiddleNumbers)
}

func Part2() {

}

func parseInput(data string) ([]Rule, []PrintUpdate) {
	lines := strings.Split(data, "\n")
	rules := []Rule{}
	printUpdates := []PrintUpdate{}

	for _, line := range lines {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			fmt.Println("Parts: ", parts)
			rule := Rule{}
			rule.pageInFront, _ = strconv.Atoi(parts[0])
			rule.pageBehind, _ = strconv.Atoi(parts[1])
			rules = append(rules, rule)
		} else if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			printUpdate := PrintUpdate{}
			for _, part := range parts {
				num, err := strconv.Atoi(part)
				if err != nil {
					continue
				}
				printUpdate.updates = append(printUpdate.updates, num)
			}
			printUpdates = append(printUpdates, printUpdate)
		} else {
			continue
		}
	}
	return rules, printUpdates
}

// for each rule, create an entry in the map with a list of numbers it can't appear ahead of
func createRuleMap(rules []Rule) map[int][]int {
	ruleMap := make(map[int][]int)
	for _, rule := range rules {
		if _, ok := ruleMap[rule.pageBehind]; ok {
			ruleMap[rule.pageBehind] = append(ruleMap[rule.pageBehind], rule.pageInFront)
		} else {
			ruleMap[rule.pageBehind] = []int{rule.pageInFront}
		}
	}
	return ruleMap
}

// given the rule map and the print updates, iterate through all print updates and validate order
// by checking if the current number exists in the rule map, and if it does, make sure it's corresponding list
// does not contain any numbers that occur after it in the printUpdate list
func validatePrintOrder(ruleMap map[int][]int, printUpdate PrintUpdate) bool {
	fmt.Println("Print Update: ", printUpdate)
	for i, number := range printUpdate.updates {
		if _, ok := ruleMap[number]; ok {
			for j := i + 1; j < len(printUpdate.updates); j++ {
				if listContains(ruleMap[number], printUpdate.updates[j]) {
					fmt.Println("Invalid Order: ", printUpdate.updates)
					return false
				}
			}
		}
	}
	return true
}

func listContains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

type Rule struct {
	pageInFront int
	pageBehind  int
}

type PrintUpdate struct {
	updates []int
}
