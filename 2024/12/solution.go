package main

import (
	"fmt"
	"log"

	"aoc/utils"
)

type Point struct {
	x, y int
}

type polynomial struct {
	area      int
	perimeter int
	sides     int
}

func checkAll4(input [][]string, current Point) []Point {
	sameAround := []Point{}

	if current.x > 0 && input[current.y][current.x-1] == input[current.y][current.x] {
		sameAround = append(sameAround, Point{current.x - 1, current.y})
	}

	if current.x < len(input[0])-1 && input[current.y][current.x+1] == input[current.y][current.x] {
		sameAround = append(sameAround, Point{current.x + 1, current.y})
	}

	if current.y > 0 && input[current.y-1][current.x] == input[current.y][current.x] {
		sameAround = append(sameAround, Point{current.x, current.y - 1})
	}

	if current.y < len(input)-1 && input[current.y+1][current.x] == input[current.y][current.x] {
		sameAround = append(sameAround, Point{current.x, current.y + 1})
	}

	return sameAround
}

func ans(input [][]string) (int, int) {
	totalCostWithPerimeter, totalCostWithSides := 0, 0

	visited := make(map[Point]struct{})
	for j, row := range input {
		for i := range row {
			if _, ok := visited[Point{i, j}]; ok {
				continue
			}
			shape := findAllGardensRecursive(input, Point{i, j}, polynomial{}, visited)
			totalCostWithPerimeter += shape.area * shape.perimeter
			totalCostWithSides += shape.area * shape.sides
		}
	}

	return totalCostWithPerimeter, totalCostWithSides
}

func findAllGardensRecursive(
	input [][]string,
	current Point,
	shape polynomial,
	visited map[Point]struct{},
) polynomial {
	if _, ok := visited[current]; ok {
		return shape
	}

	checkNext := checkAll4(input, current)

	if len(checkNext) == 0 {
		if shape.area == 0 {
			shape.area = 1
			shape.perimeter = 4
			visited[current] = struct{}{}
			shape.sides = checkCorners(input, current)
			return shape
		}
		return shape
	}

	shape.perimeter += 4 - len(checkNext)
	shape.area += 1
	visited[current] = struct{}{}
	shape.sides += checkCorners(input, current)

	for _, next := range checkNext {
		shape = findAllGardensRecursive(input, next, shape, visited)
	}

	return shape
}

func checkCorners(input [][]string, current Point) int {
	count := 0
	gardenType := input[current.y][current.x]
	x, y := current.x, current.y

	if x == 0 && y == 0 {
		count += 1
	}

	if x == 0 && y == len(input)-1 {
		count += 1
	}

	if x == len(input[0])-1 && y == len(input)-1 {
		count += 1
	}

	if x == len(input[0])-1 && y == 0 {
		count += 1
	}

	if (x > 0 && y > 0 && input[y][x-1] != gardenType && input[y-1][x] != gardenType) ||
		(x > 0 && y == 0 && input[y][x-1] != gardenType) || (x == 0 && y > 0 && input[y-1][x] != gardenType) {
		count += 1
	}

	if x < len(input[0])-1 && y < len(input)-1 && input[y][x+1] == gardenType &&
		input[y+1][x] == gardenType &&
		input[y+1][x+1] != gardenType {
		count += 1
	}

	if (x < len(input[0])-1 && y > 0 && input[y][x+1] != gardenType && input[y-1][x] != gardenType) ||
		(x < len(input[0])-1 && y == 0 && input[y][x+1] != gardenType) ||
		(x == len(input[0])-1 && y > 0 && input[y-1][x] != gardenType) {
		count += 1
	}

	if x > 0 && y < len(input)-1 && input[y][x-1] == gardenType && input[y+1][x] == gardenType &&
		input[y+1][x-1] != gardenType {
		count += 1
	}

	if (x > 0 && y < len(input)-1 && input[y][x-1] != gardenType && input[y+1][x] != gardenType) ||
		(x > 0 && y == len(input)-1 && input[y][x-1] != gardenType) || (x == 0 && y < len(input)-1 && input[y+1][x] != gardenType) {
		count += 1
	}

	if x < len(input[0])-1 && y > 0 && input[y][x+1] == gardenType && input[y-1][x] == gardenType &&
		input[y-1][x+1] != gardenType {
		count += 1
	}

	if (x < len(input[0])-1 && y < len(input)-1 && input[y][x+1] != gardenType && input[y+1][x] != gardenType) ||
		(x < len(input[0])-1 && y == len(input)-1 && input[y][x+1] != gardenType) ||
		(x == len(input[0])-1 && y < len(input)-1 && input[y+1][x] != gardenType) {
		count += 1
	}

	if x > 0 && y > 0 && input[y][x-1] == gardenType && input[y-1][x] == gardenType &&
		input[y-1][x-1] != gardenType {
		count += 1
	}

	return count
}

func findAllGardensNonRecursively(
	input [][]string,
	current Point,
	shape polynomial,
	visited map[Point]struct{},
) (polynomial, []Point) {
	if _, ok := visited[current]; ok {
		return shape, []Point{}
	}

	checkNext := checkAll4(input, current)

	if len(checkNext) == 0 {
		if shape.area == 0 {
			visited[current] = struct{}{}
			shape = polynomial{
				area: 1, perimeter: 4, sides: 4,
			}
			return shape, []Point{}
		}
		return shape, []Point{}
	}

	shape.perimeter += 4 - len(checkNext)
	shape.area += 1
	visited[current] = struct{}{}
	shape.sides += checkCorners(input, current)

	return shape, checkNext
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := utils.Build2DGrid(lines)

	totalCost, totalCostWithSides := ans(grid)
	fmt.Println("[PART 1] total cost: ", totalCost)
	fmt.Println("[PART 2] total cost: ", totalCostWithSides)
}
