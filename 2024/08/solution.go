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

func findAntinodes(grid []string) int {
	coordinates := findAntennas(grid)
	uniqueAntiNodes := make(map[Coordinate]bool)

	for _, coordinate := range coordinates {
		if len(coordinate) < 2 {
			continue
		}

		for i := 0; i < len(coordinate)-1; i++ {
			for j := i + 1; j < len(coordinate); j++ {

				dx := coordinate[j].x - coordinate[i].x
				dy := coordinate[j].y - coordinate[i].y

				steps := 1

				for k := 1; k <= steps; k++ {
					antinode1 := Coordinate{x: coordinate[i].x - k*dx, y: coordinate[i].y - k*dy}
					antinode2 := Coordinate{x: coordinate[j].x + k*dx, y: coordinate[j].y + k*dy}

					if isValid(grid, antinode1) {
						uniqueAntiNodes[antinode1] = true
					}
					if isValid(grid, antinode2) {
						uniqueAntiNodes[antinode2] = true
					}
				}
			}
		}
	}
	return len(uniqueAntiNodes)
}

func main() {
	grid, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	antinodes := findAntinodes(grid)
	fmt.Println("[PART 1] antinodes: ", antinodes)
}
