package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/utils"
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

func dfsForRating(heightMap [][]int, r, c int, path map[[2]int]bool) int {
	if heightMap[r][c] == 9 {
		return 1
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	path[[2]int{r, c}] = true
	count := 0

	for _, dir := range directions {
		nr, nc := r+dir[0], c+dir[1]
		if nr >= 0 && nr < len(heightMap) && nc >= 0 && nc < len(heightMap[0]) {
			if !path[[2]int{nr, nc}] && heightMap[nr][nc] == heightMap[r][c]+1 {
				count += dfsForRating(heightMap, nr, nc, path)
			}
		}
	}

	delete(path, [2]int{r, c})

	return count
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

func calculateRatings(lines []string) int {
	input := strings.Join(lines, "\n")
	heightMap := parseMap(input)
	trailheads := findTrailheads(heightMap)

	totalRating := 0
	for _, trailhead := range trailheads {
		path := make(map[[2]int]bool)
		totalRating += dfsForRating(heightMap, trailhead[0], trailhead[1], path)
	}

	return totalRating
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- Day 10: Hoof It ---")
	totalScore := calculateScores(lines)
	fmt.Println("[PART 1] total score: ", totalScore)

	totalRating := calculateRatings(lines)
	fmt.Println("[PART 2] total rating: ", totalRating)
}
