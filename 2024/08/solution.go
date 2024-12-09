package main

import (
	"aoc/utils"
	"fmt"
	"log"
)

type Coordinate struct {
	x int
	y int
}

func findAntennas(grid []string) map[byte][]Coordinate {
	antennas := make(map[byte][]Coordinate)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != '.' {
				antennas[grid[y][x]] = append(antennas[grid[y][x]], Coordinate{x: x, y: y})
			}
		}
	}
	return antennas
}

func isValid(grid []string, antinode Coordinate) bool {
	return antinode.x >= 0 &&
		antinode.x < len(grid[0]) &&
		antinode.y >= 0 &&
		antinode.y < len(grid)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func findAntinodes(grid []string, includeHarmonics bool) int {
	coordinates := findAntennas(grid)
	uniqueAntinodes := make(map[Coordinate]bool)

	for _, coordinate := range coordinates {
		if len(coordinate) < 2 {
			continue
		}

		if includeHarmonics {
			for _, antenna := range coordinate {
				uniqueAntinodes[antenna] = true
			}
		}

		for i := 0; i < len(coordinate)-1; i++ {
			for j := i + 1; j < len(coordinate); j++ {
				dx := coordinate[j].x - coordinate[i].x
				dy := coordinate[j].y - coordinate[i].y

				steps := 1
				if includeHarmonics {
					steps = max(len(grid), len(grid[0]))
				}

				for k := 1; k <= steps; k++ {
					antinode1 := Coordinate{x: coordinate[i].x - k*dx, y: coordinate[i].y - k*dy}
					antinode2 := Coordinate{x: coordinate[j].x + k*dx, y: coordinate[j].y + k*dy}

					if isValid(grid, antinode1) {
						uniqueAntinodes[antinode1] = true
					}
					if isValid(grid, antinode2) {
						uniqueAntinodes[antinode2] = true
					}
				}
			}
		}
	}

	return len(uniqueAntinodes)
}

func main() {
	grid, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	antinodes := findAntinodes(grid, false)
	fmt.Println("[PART 1] antinodes:", antinodes)

	antinodesWithHarmonics := findAntinodes(grid, true)
	fmt.Println("[PART 2] antinodes:", antinodesWithHarmonics)
}
