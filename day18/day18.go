package day18

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {
	i := utilities.ReadInput("day18/day18-input.txt")
	fallingByteCoordinates := parseInput(i)
	grid := simulateNumberOfFallingBytes(1024, fallingByteCoordinates, 71)
	//printGrid(grid)
	shortestPath := findShortestPathFromPointAtoPointB(grid, utilities.Coordinate{X: 0, Y: 0}, utilities.Coordinate{X: 70, Y: 70})
	fmt.Println("Shortest Path: ", shortestPath)

}

func Part2() {
	gridSize := 71
	knownBytesStartPoint := 1024
	i := utilities.ReadInput("day18/day18-input.txt")
	fallingByteCoordinates := parseInput(i)
	for i := knownBytesStartPoint; i < len(fallingByteCoordinates); i++ {
		grid := simulateNumberOfFallingBytes(i, fallingByteCoordinates, gridSize)
		shortestPath := findShortestPathFromPointAtoPointB(grid, utilities.Coordinate{X: 0, Y: 0}, utilities.Coordinate{X: gridSize - 1, Y: gridSize - 1})
		if shortestPath == -1 {
			fmt.Println("First byte to block exist: ", fallingByteCoordinates[i-1])
			break
		}
	}
}

func parseInput(i string) []utilities.Coordinate {
	lines := strings.Split(i, "\n")
	fallingByteCoordinates := []utilities.Coordinate{}

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		fallingByteCoordinates = append(fallingByteCoordinates, utilities.Coordinate{X: x, Y: y})
	}

	return fallingByteCoordinates
}

func simulateNumberOfFallingBytes(bytesFallen int, fallingByteCoordinates []utilities.Coordinate, gridSize int) [][]string {
	grid := make([][]string, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]string, gridSize)
		for j := 0; j < gridSize; j++ {
			grid[i][j] = "."
		}
	}
	for i := 0; i < bytesFallen; i++ {
		byteFallen := fallingByteCoordinates[i]
		grid[byteFallen.Y][byteFallen.X] = "#"
	}
	return grid
}

func printGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
}

func findShortestPathFromPointAtoPointB(grid [][]string, pointA utilities.Coordinate, pointB utilities.Coordinate) int {
	stepsTaken := utilities.FindShortestPathFromAToB(pointA, pointB, func(c utilities.Coordinate) bool { return canIMoveHere(c, grid) }, []utilities.Direction{utilities.N, utilities.S, utilities.E, utilities.W})
	//fmt.Println("Steps Taken: ", stepsTaken)
	//visualizeSteps(grid, stepsTaken)
	return len(stepsTaken) - 1
}

func canIMoveHere(c utilities.Coordinate, grid [][]string) bool {
	if c.X < 0 || c.Y < 0 || c.X >= len(grid[0]) || c.Y >= len(grid) {
		return false
	}
	return grid[c.Y][c.X] == "."
}

func visualizeSteps(grid [][]string, steps []utilities.Coordinate) {
	for _, step := range steps {
		grid[step.Y][step.X] = "0"
	}
	printGrid(grid)
}
