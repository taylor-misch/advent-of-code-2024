package day04

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strings"
)

func Part1() {
	fmt.Println("Part 1 Begin")
	i := utilities.ReadInput("day04/day04-input.txt")
	wordSearch := parseInput(i)
	count := getWordCountPart1("XMAS", wordSearch)
	fmt.Println("Count: ", count)
	fmt.Println("Part 1 End")
}

func Part2() {
	fmt.Println("Part 2 Begin")
	i := utilities.ReadInput("day04/day04-input.txt")
	wordSearch := parseInput(i)
	count := getWordCountPart2("MAS", wordSearch)
	fmt.Println("Count: ", count)
	fmt.Println("Part 2 End")
}

func parseInput(data string) [][]string {
	wordSearch := make([][]string, 0)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		parts := make([]string, len(line))
		for i, char := range line {
			parts[i] = string(char)
		}
		wordSearch = append(wordSearch, parts)
	}
	fmt.Println("Input: ", wordSearch)
	return wordSearch
}

func getWordCountPart1(word string, wordSearch [][]string) int {
	count := 0
	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[i]); j++ {
			if wordSearch[i][j] != string(word[0]) {
				continue
			}
			count += searchForwards(j, i, word, wordSearch)
			count += searchBackwards(j, i, word, wordSearch)
			count += searchUpwards(j, i, word, wordSearch)
			count += searchDownwards(j, i, word, wordSearch)
			count += searchDiagonallyUpLeftPart1(j, i, word, wordSearch)
			count += searchDiagonallyUpRightPart1(j, i, word, wordSearch)
			count += searchDiagonallyDownLeftPart1(j, i, word, wordSearch)
			count += searchDiagonallyDownRightPart1(j, i, word, wordSearch)
		}
	}
	return count
}

func getWordCountPart2(word string, wordSearch [][]string) int {
	count := 0
	indexOfA := make([]utilities.Coordinate, 0)
	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[i]); j++ {
			if wordSearch[i][j] != string(word[0]) {
				continue
			}
			indexOfA = append(indexOfA, searchDiagonallyUpLeftPart2(j, i, word, wordSearch))
			indexOfA = append(indexOfA, searchDiagonallyUpRightPart2(j, i, word, wordSearch))
			indexOfA = append(indexOfA, searchDiagonallyDownLeftPart2(j, i, word, wordSearch))
			indexOfA = append(indexOfA, searchDiagonallyDownRightPart2(j, i, word, wordSearch))
		}
	}
	countOfEachIndexOccurrence := map[utilities.Coordinate]int{}
	for _, index := range indexOfA {
		countOfEachIndexOccurrence[index]++
	}

	for key, value := range countOfEachIndexOccurrence {
		if value == 2 {
			fmt.Println("Key: ", key, " has an X-MAS")
			count++
		}
	}
	return count
}

func searchForwards(xPos int, yPos int, word string, wordSearch [][]string) int {
	if xPos+len(word) <= len(wordSearch[yPos]) {
		for i := 0; i < len(word); i++ {
			if wordSearch[yPos][xPos+i] != string(word[i]) {
				return 0
			}
		}
		fmt.Println("Found word FORWARDS at position: ", yPos, xPos)
		return 1
	}
	return 0
}

func searchBackwards(xPos int, yPos int, word string, wordSearch [][]string) int {
	if xPos-len(word) >= -1 {
		for i := 0; i < len(word); i++ {
			if wordSearch[yPos][xPos-i] != string(word[i]) {
				return 0
			}
		}
		fmt.Println("Found word BACKWARDS at position: ", yPos, xPos)
		return 1
	}
	return 0
}

func searchUpwards(xPos int, yPos int, word string, wordSearch [][]string) int {
	if yPos-len(word) >= -1 {
		for i := 0; i < len(word); i++ {
			if wordSearch[yPos-i][xPos] != string(word[i]) {
				return 0
			}
		}
		fmt.Println("Found word UPWARDS at position: ", yPos, xPos)
		return 1
	}
	return 0
}

func searchDownwards(xPos int, yPos int, word string, wordSearch [][]string) int {
	if yPos+len(word) <= len(wordSearch) {
		for i := 0; i < len(word); i++ {
			//fmt.Println("Checking: ", wordSearch[xPos][yPos+i], " is equal to ", string(word[i]))
			if wordSearch[yPos+i][xPos] != string(word[i]) {
				return 0
			}
		}
		fmt.Println("Found word DOWNWARDS at position: ", yPos, xPos)
		return 1
	}
	return 0
}

func searchDiagonallyUpLeftPart1(xPos int, yPos int, word string, wordSearch [][]string) int {
	if xPos-len(word) >= -1 && yPos-len(word) >= -1 {
		for i := 0; i < len(word); i++ {
			if wordSearch[yPos-i][xPos-i] != string(word[i]) {
				return 0
			}
		}
		fmt.Println("Found word UP_LEFT at position: ", yPos, xPos)
		return 1
	}
	return 0
}

func searchDiagonallyUpRightPart1(xPos int, yPos int, word string, wordSearch [][]string) int {
	if xPos+len(word) <= len(wordSearch) && yPos-len(word) >= -1 {
		for i := 0; i < len(word); i++ {
			if wordSearch[yPos-i][xPos+i] != string(word[i]) {
				return 0
			}
		}
		fmt.Println("Found word UP_RIGHT at position: ", yPos, xPos)
		return 1
	}
	return 0
}

func searchDiagonallyDownLeftPart1(xPos int, yPos int, word string, wordSearch [][]string) int {
	if xPos-len(word) >= -1 && yPos+len(word) <= len(wordSearch) {
		for i := 0; i < len(word); i++ {
			if wordSearch[yPos+i][xPos-i] != string(word[i]) {
				return 0
			}
		}
		fmt.Println("Found word DOWN_LEFT at position: ", yPos, xPos)
		return 1
	}
	return 0
}

func searchDiagonallyDownRightPart1(xPos int, yPos int, word string, wordSearch [][]string) int {
	if xPos+len(word) <= len(wordSearch) && yPos+len(word) <= len(wordSearch) {
		for i := 0; i < len(word); i++ {
			if wordSearch[yPos+i][xPos+i] != string(word[i]) {
				return 0
			}
		}
		fmt.Println("Found word DOWN_RIGHT at position: ", yPos, xPos)
		return 1
	}
	return 0
}

func searchDiagonallyUpLeftPart2(xPos int, yPos int, word string, wordSearch [][]string) utilities.Coordinate {
	if xPos-len(word) >= -1 && yPos-len(word) >= -1 {
		for i := 0; i < len(word); i++ {
			if wordSearch[yPos-i][xPos-i] != string(word[i]) {
				return utilities.Coordinate{X: -1, Y: -1}
			}
		}
		fmt.Println("Found word UP_LEFT at position: ", yPos, xPos)
		return utilities.Coordinate{X: xPos - 1, Y: yPos - 1}
	}
	return utilities.Coordinate{X: -1, Y: -1}
}

func searchDiagonallyUpRightPart2(xPos int, yPos int, word string, wordSearch [][]string) utilities.Coordinate {
	if xPos+len(word) <= len(wordSearch) && yPos-len(word) >= -1 {
		for i := 0; i < len(word); i++ {
			if wordSearch[yPos-i][xPos+i] != string(word[i]) {
				return utilities.Coordinate{X: -1, Y: -1}
			}
		}
		fmt.Println("Found word UP_RIGHT at position: ", yPos, xPos)
		return utilities.Coordinate{X: xPos + 1, Y: yPos - 1}
	}
	return utilities.Coordinate{X: -1, Y: -1}
}

func searchDiagonallyDownLeftPart2(xPos int, yPos int, word string, wordSearch [][]string) utilities.Coordinate {
	if xPos-len(word) >= -1 && yPos+len(word) <= len(wordSearch) {
		for i := 0; i < len(word); i++ {
			if wordSearch[yPos+i][xPos-i] != string(word[i]) {
				return utilities.Coordinate{X: -1, Y: -1}
			}
		}
		fmt.Println("Found word DOWN_LEFT at position: ", yPos, xPos)
		return utilities.Coordinate{X: xPos - 1, Y: yPos + 1}
	}
	return utilities.Coordinate{X: -1, Y: -1}
}

func searchDiagonallyDownRightPart2(xPos int, yPos int, word string, wordSearch [][]string) utilities.Coordinate {
	if xPos+len(word) <= len(wordSearch) && yPos+len(word) <= len(wordSearch) {
		for i := 0; i < len(word); i++ {
			if wordSearch[yPos+i][xPos+i] != string(word[i]) {
				return utilities.Coordinate{X: -1, Y: -1}
			}
		}
		fmt.Println("Found word DOWN_RIGHT at position: ", yPos, xPos)
		return utilities.Coordinate{X: xPos + 1, Y: yPos + 1}
	}
	return utilities.Coordinate{X: -1, Y: -1}
}
