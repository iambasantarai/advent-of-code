package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/utils"
)

const word = "XMAS"

type Coordinate struct {
	dx int
	dy int
}

func checkCoordinate(grid [][]string, x, y, dx, dy int) bool {
	for i := 0; i < len(word); i++ {
		newX := x + i*dx
		newY := y + i*dy

		if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[0]) {
			return false
		}

		if grid[newX][newY] != string(word[i]) {
			return false
		}
	}
	return true
}

func countXMASOccurrences(grid [][]string) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	directions := []Coordinate{
		{0, 1},   // Horizontal (left to right)
		{0, -1},  // Horizontal (right to left)
		{1, 0},   // Vertical (top to bottom)
		{-1, 0},  // Vertical (bottom to top)
		{1, 1},   // Diagonal (top-left to bottom-right)
		{-1, -1}, // Diagonal (bottom-right to top-left)
		{1, -1},  // Diagonal (top-right to bottom-left)
		{-1, 1},  // Diagonal (bottom-left to top-right)}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				if checkCoordinate(grid, i, j, dir.dx, dir.dy) {
					count++
				}
			}
		}
	}

	return count
}

func countXShapedMASOccurrences(grid [][]string) int {
	count := 0

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] != "A" {
				continue
			}
			if ((grid[i-1][j-1] == "S" && grid[i+1][j+1] == "M") || (grid[i-1][j-1] == "M" && grid[i+1][j+1] == "S")) &&
				((grid[i-1][j+1] == "S" && grid[i+1][j-1] == "M") || (grid[i-1][j+1] == "M" && grid[i+1][j-1] == "S")) {
				count++
			}
		}
	}

	return count
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	occurrencesCount := countXMASOccurrences(grid)
	fiexedOccurrencesCount := countXShapedMASOccurrences(grid)

	fmt.Println("--- Day 4: Ceres Search ---")
	fmt.Println("[PART 1] total occurrencesCount: ", occurrencesCount)
	fmt.Println("[PART 2] total occurrencesCount: ", fiexedOccurrencesCount)
}
