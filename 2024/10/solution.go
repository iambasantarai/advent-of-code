package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strings"
)

func parseMap(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	heightMap := make([][]int, len(lines))
	for i, line := range lines {
		heightMap[i] = make([]int, len(line))
		for j, ch := range line {
			heightMap[i][j] = int(ch - '0')
		}
	}

	return heightMap
}

func findTrailheads(heightMap [][]int) [][2]int {
	trailheads := [][2]int{}
	for r, row := range heightMap {
		for c, height := range row {
			if height == 0 {
				trailheads = append(trailheads, [2]int{r, c})
			}
		}
	}

	return trailheads
}

func dfs(heightMap [][]int, r, c int, visited map[[2]int]bool) map[[2]int]bool {
	stack := [][2]int{{r, c}}
	reachableNines := make(map[[2]int]bool)

	for len(stack) > 0 {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r, c = pos[0], pos[1]

		if visited[pos] {
			continue
		}
		visited[pos] = true

		if heightMap[r][c] == 9 {
			reachableNines[pos] = true
		}

		directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr >= 0 && nr < len(heightMap) && nc >= 0 && nc < len(heightMap[0]) {
				if heightMap[nr][nc] == heightMap[r][c]+1 {
					stack = append(stack, [2]int{nr, nc})
				}
			}
		}
	}

	return reachableNines
}

func calculateScores(lines []string) int {
	input := strings.Join(lines, "\n")
	heightMap := parseMap(input)

	trailheads := findTrailheads(heightMap)

	totalScore := 0
	for _, trailhead := range trailheads {
		visited := make(map[[2]int]bool)
		reachableNines := dfs(heightMap, trailhead[0], trailhead[1], visited)
		totalScore += len(reachableNines)
	}

	return totalScore
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	totalScore := calculateScores(lines)

	fmt.Println("[PART 1] total score: ", totalScore)
}
