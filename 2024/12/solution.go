package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/utils"
)

type Point struct {
	x, y int
}

func build2DGrid(lines []string) [][]string {
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}
	return grid
}

func floodFill(grid [][]string, visited map[Point]bool, start Point) (int, int) {
	rows, cols := len(grid), len(grid[0])
	stack := []Point{start}
	plantType := grid[start.x][start.y]
	area := 0
	perimeter := 0

	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[current] {
			continue
		}

		visited[current] = true
		area++

		// Check all neighbors
		for _, d := range directions {
			nx, ny := current.x+d.x, current.y+d.y

			if nx < 0 || ny < 0 || nx >= rows || ny >= cols {
				perimeter++
			} else if grid[nx][ny] != plantType {
				perimeter++
			} else if !visited[Point{nx, ny}] {
				stack = append(stack, Point{nx, ny})
			}
		}
	}

	return area, perimeter
}

func calculateFencingCost(grid [][]string) int {
	visited := make(map[Point]bool)
	totalCost := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if !visited[Point{x, y}] {
				area, perimeter := floodFill(grid, visited, Point{x, y})
				totalCost += area * perimeter
			}
		}
	}

	return totalCost
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := build2DGrid(lines)

	totalCost := calculateFencingCost(grid)
	fmt.Println("[PART 1] total cost: ", totalCost)
}
