package day06

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strings"
)

func Part1() {
	i := utilities.ReadInput("day06/day06-input.txt")
	guardsPosition, obstacles, xSize, ySize := parseInput(i)

	guard := Guard{guardsPosition, utilities.Coordinate{0, -1}}
	fmt.Println(guardsPosition)
	fmt.Println(obstacles)

	positionsVisited := map[utilities.Coordinate]bool{}
	positionsVisited[guard.position] = true
	isGuardInGrid := true
	for isGuardInGrid {
		guard, _ = walk(guard, obstacles)
		if hasGuardLeftGrid(guard, xSize, ySize) {
			isGuardInGrid = false
		} else {
			positionsVisited[guard.position] = true
		}
	}

	fmt.Println("Positions visited: ", len(positionsVisited))

}

// for this part, I basically follow the same logic and have the guard walk the path
// however, for every position, check for the obstacle in front,
//if no obstacle check if a right turn would place the guard back on a path he has previously visited
// if yes, save the position of the guards next step in a list of infinite loops possibilities

func Part2() {
	i := utilities.ReadInput("day06/day06-input.txt")
	guardsPosition, obstacles, xSize, ySize := parseInput(i)

	guard := Guard{guardsPosition, utilities.Coordinate{0, -1}}
	fmt.Println(guardsPosition)
	fmt.Println(obstacles)

	infiniteLoopObstacles := make([]utilities.Coordinate, 0)
	positionsVisited := map[utilities.Coordinate]VisitedPosition{}
	positionsVisited[guard.position] = VisitedPosition{true, []utilities.Coordinate{guard.direction}}
	isGuardInGrid := true
	for isGuardInGrid {
		if obstaclePlacement, loopOccurs := castRayAndCheckForPreviousVisitOrObstacle(guard, obstacles, positionsVisited, xSize, ySize); loopOccurs {
			infiniteLoopObstacles = append(infiniteLoopObstacles, obstaclePlacement)
		}
		guard, _ = walk(guard, obstacles)
		if hasGuardLeftGrid(guard, xSize, ySize) {
			isGuardInGrid = false
		} else {
			if visitedPosition, exists := positionsVisited[guard.position]; exists {
				visitedPosition.directionWhenVisited = append(visitedPosition.directionWhenVisited, guard.direction)
				positionsVisited[guard.position] = visitedPosition
			} else {
				positionsVisited[guard.position] = VisitedPosition{true, []utilities.Coordinate{guard.direction}}
			}
		}
	}

	fmt.Println("Infinite loop obstacles: ", infiniteLoopObstacles)
	fmt.Println("Infinite loop possibilities: ", len(infiniteLoopObstacles))

}

func parseInput(input string) (utilities.Coordinate, []utilities.Coordinate, int, int) {
	obstacles := make([]utilities.Coordinate, 0)
	guardPosition := utilities.Coordinate{}
	lines := strings.Split(input, "\n")
	ySize := len(lines)
	xSize := len(lines[0])
	for currentRow, line := range lines {
		fmt.Println(line)
		for currentColumn, char := range line {
			if char == '#' {
				obstacles = append(obstacles, utilities.Coordinate{currentColumn, currentRow})
			} else if char == '^' {
				guardPosition = utilities.Coordinate{currentColumn, currentRow}
			}
		}
	}

	return guardPosition, obstacles, xSize, ySize

}

func turnRight(direction utilities.Coordinate) utilities.Coordinate {
	switch direction {
	case utilities.Coordinate{1, 0}:
		return utilities.Coordinate{0, 1}
	case utilities.Coordinate{0, 1}:
		return utilities.Coordinate{-1, 0}
	case utilities.Coordinate{-1, 0}:
		return utilities.Coordinate{0, -1}
	case utilities.Coordinate{0, -1}:
		return utilities.Coordinate{1, 0}
	default:
		fmt.Println("Invalid")
		return utilities.Coordinate{}
	}
}

func walk(guard Guard, obstacles []utilities.Coordinate) (Guard, bool) {
	nextPosition := utilities.Coordinate{guard.position.X + guard.direction.X, guard.position.Y + guard.direction.Y}
	if listContains(obstacles, nextPosition) {
		guard.direction = turnRight(guard.direction)
		return guard, true
	} else {
		guard.position = nextPosition
	}
	return guard, false
}

func castRayAndCheckForPreviousVisitOrObstacle(guard Guard, obstacles []utilities.Coordinate, positionsVisited map[utilities.Coordinate]VisitedPosition, xSize int, ySize int) (utilities.Coordinate, bool) {
	obstaclePlacement := utilities.Coordinate{guard.position.X + guard.direction.X, guard.position.Y + guard.direction.Y}
	//fmt.Println("Guard position: ", guard.position, "Obstacle placement consideration: ", obstaclePlacement)
	if listContains(obstacles, obstaclePlacement) {
		fmt.Println("Raycast: Obstacle already in front")
		return obstaclePlacement, false
	}
	guard.direction = turnRight(guard.direction)
	nextPosition := guard.position
	for {
		//fmt.Println("Position considered: ", nextPosition)
		//fmt.Println("Was position visited? ", positionsVisited[nextPosition])
		//if positionsVisited[nextPosition].visited {
		//	fmt.Println("Guard direction when visited?: ", guard.direction, " vs current direction: ", positionsVisited[nextPosition].directionWhenVisited)
		//}
		if listContains(obstacles, nextPosition) {
			fmt.Println("Raycast would hit obstacle")
			return obstaclePlacement, false
		} else if positionsVisited[nextPosition].visited && listContains(positionsVisited[nextPosition].directionWhenVisited, guard.direction) {
			fmt.Println("Raycast: Success: Guard would hit previous visit - position: ", obstaclePlacement)
			return obstaclePlacement, true
		} else {
			guard.position = nextPosition
		}
		nextPosition = utilities.Coordinate{guard.position.X + guard.direction.X, guard.position.Y + guard.direction.Y}
		if hasGuardLeftGrid(guard, xSize, ySize) {
			fmt.Println("Raycast left grid")
			return obstaclePlacement, false
		}
	}
}

func listContains(slice []utilities.Coordinate, value utilities.Coordinate) bool {
	for _, v := range slice {
		if v.X == value.X && v.Y == value.Y {
			fmt.Println("Obstacle hit at: ", value)
			return true
		}
	}
	return false
}

func hasGuardLeftGrid(guard Guard, gridX int, gridY int) bool {
	return guard.position.X < 0 || guard.position.X >= gridX || guard.position.Y < 0 || guard.position.Y >= gridY
}

type Guard struct {
	position  utilities.Coordinate
	direction utilities.Coordinate
}

type VisitedPosition struct {
	visited              bool
	directionWhenVisited []utilities.Coordinate
}
