package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"aoc/utils"
)

func getSortedArrays(lines []string) ([]int, []int) {
	var lParts, rParts []int

	for _, line := range lines {
		parts := strings.Fields(line)

		lNum, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		rNum, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		lParts = append(lParts, lNum)
		rParts = append(rParts, rNum)
	}

	sort.Ints(lParts)
	sort.Ints(rParts)

	return lParts, rParts
}

func getTotalDistance(lParts, rParts []int) int {
	totalDistance := 0

	for i := 0; i < len(lParts); i++ {
		difference := lParts[i] - rParts[i]
		if difference < 0 {
			difference *= -1
		}

		totalDistance += difference
	}

	return totalDistance
}

func getSimilarityScore(lParts, rParts []int) int {
	similarityScore := 0

	recurrenceCount := make(map[int]int)

	for i := 0; i < len(lParts); i++ {
		recurrenceCount[rParts[i]] += 1
	}

	for _, key := range lParts {
		if val, ok := recurrenceCount[key]; ok {
			similarityScore += key * val
		}
	}

	return similarityScore
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lParts, rParts := getSortedArrays(lines)

	fmt.Println("--- Day 1: Historian Hysteria ---")
	totaldistance := getTotalDistance(lParts, rParts)
	fmt.Printf("[PART 1] total distance: %d\n", totaldistance)

	similarityScore := getSimilarityScore(lParts, rParts)
	fmt.Printf("[PART 2] similarity score: %d\n", similarityScore)
}
