package main

import (
	"fmt"
	"log"

	"aoc/utils"
)

// NOTE: https://dev.to/ankitmalikg/how-use-iota-in-golang-3hcb
const (
	Up = iota
	Right
	Down
	Left
)

type Point struct {
	x         int
	y         int
	direction int
}

type Coordinate struct {
	x int
	y int
}

var directions = []Coordinate{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func findStartingPoint(labMap [][]string) Point {
	for i, row := range labMap {
		for j, char := range row {
			switch char {
			case "^":
				return Point{j, i, Up}
			case ">":
				return Point{j, i, Right}
			case "v":
				return Point{j, i, Down}
			case "<":
				return Point{j, i, Left}
			}
		}
	}
	return Point{-1, -1, Up}
}

func getNextStepWithDirectionPreserved(labMap [][]string, current Point) (bool, Point) {
	switch current.direction {
	case Up:
		current.y -= 1
	case Down:
		current.y += 1
	case Right:
		current.x += 1
	case Left:
		current.x -= 1
	}
	if current.x < 0 || current.y < 0 || current.x >= len(labMap[0]) || current.y >= len(labMap) {
		return false, current
	}
	return true, current
}

func turnRight(current Point) Point {
	switch current.direction {
	case Up:
		current.direction = Right
	case Down:
		current.direction = Left
	case Right:
		current.direction = Down
	case Left:
		current.direction = Up
	}

	return current
}

func findNextStep(labMap [][]string, current Point) (bool, Point) {
	valid, possibleNext := getNextStepWithDirectionPreserved(labMap, current)
	if !valid {
		return false, possibleNext
	}

	switch labMap[possibleNext.y][possibleNext.x] {
	case "#":
		return findNextStep(labMap, turnRight(current))
	case ".":
		return true, possibleNext
	case "^":
		return true, possibleNext
	}
	return false, possibleNext
}

func patrol(labMap [][]string, gate Point) (int, map[Coordinate]int) {
	path := make(map[Coordinate]int)
	count := 0
	current := gate

	for {
		if _, ok := path[Coordinate{current.x, current.y}]; !ok {
			count++
			path[Coordinate{current.x, current.y}] = current.direction
		}

		isValid, newCurrent := findNextStep(labMap, current)
		if !isValid {
			return count, path
		}

		current = newCurrent
	}
}

func isLoop(input [][]string, gate Point) bool {
	path := make(map[Point]struct{})
	path2 := make(map[Coordinate]struct{})
	current := gate
	for {
		if _, ok := path[current]; !ok {
			path[current] = struct{}{}
		} else {
			return true
		}

		if _, ok := path2[Coordinate{current.x, current.y}]; !ok {
			path2[Coordinate{current.x, current.y}] = struct{}{}
		}

		valid, newCurrent := findNextStep(input, current)
		if !valid {
			return false
		}

		current = newCurrent
	}
}

func findNewObstacleCount(input [][]string, gate Point, path map[Coordinate]int) int {
	count := 0
	obstanceMap := make(map[Coordinate]struct{})
	for step := range path {
		if step.x == gate.x && step.y == gate.y {
			continue
		}
		if input[step.y][step.x] == "." {

			input[step.y][step.x] = "#"
			if isLoop(input, gate) {
				if _, ok := obstanceMap[step]; !ok {
					count++
					obstanceMap[step] = struct{}{}
				}
			}
			input[step.y][step.x] = "."
		}
	}
	return count
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	labMap := utils.Build2DGrid(lines)
	startingPoint := findStartingPoint(labMap)

	visitedPositions, path := patrol(labMap, startingPoint)
	differentPosition := findNewObstacleCount(labMap, startingPoint, path)

	fmt.Println("--- Day 6: Guard Gallivant ---")
	fmt.Println("[PART 1] visited positions:", visitedPositions)
	fmt.Println("[PART 2] different positions:", differentPosition)
}
