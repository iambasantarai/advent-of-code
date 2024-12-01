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

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lParts, rParts := getSortedArrays(lines)
	totaldistance := getTotalDistance(lParts, rParts)
	fmt.Printf("[PART 1] total distance: %d\n", totaldistance)
}
