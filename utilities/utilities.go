package utilities

import (
	"container/list"
	"fmt"
	"os"
)

func ReadInput(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		// Handle error
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	// Read the entire file into memory
	data, err := os.ReadFile(filename)
	if err != nil {
		// Handle error
		fmt.Println("Error reading file:", err)
		return ""
	}

	// Print the contents
	//fmt.Println(string(data))
	return string(data)
}

func AbsoluteValue(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Coordinate struct to hold x and y values
type Coordinate struct {
	X int
	Y int
}

type Direction struct {
	X, Y int
}

var (
	N = Direction{X: 0, Y: -1}
	S = Direction{X: 0, Y: 1}
	E = Direction{X: 1, Y: 0}
	W = Direction{X: -1, Y: 0}
)

func FindShortestPathFromAToB(start, end Coordinate, canIMoveHere func(Coordinate) bool, movementDirections []Direction) []Coordinate {
	visited := make(map[Coordinate]bool)
	queue := list.New()
	queue.PushBack([]Coordinate{start})

	for queue.Len() > 0 {
		element := queue.Front()
		path := element.Value.([]Coordinate)
		queue.Remove(element)
		current := path[len(path)-1]

		if current == end {
			return path
		}

		if !visited[current] {
			visited[current] = true
			for _, direction := range movementDirections {
				neighbor := Coordinate{X: current.X + direction.X, Y: current.Y + direction.Y}
				if canIMoveHere(neighbor) && !visited[neighbor] {
					newPath := append([]Coordinate{}, path...)
					newPath = append(newPath, neighbor)
					queue.PushBack(newPath)
				}
			}
		}
	}

	return nil
}
