package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc/utils"
)

func getLevels(line string) ([]int, error) {
	var levels []int

	parts := strings.Fields(line)

	for i := 0; i < len(parts); i++ {
		num, err := strconv.Atoi(parts[i])
		if err != nil {
			return nil, err
		}

		levels = append(levels, num)
	}

	return levels, nil
}

func isReportSafe(levels []int) bool {
	increasing, decreasing := false, false

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff > 3 || diff < -3 {
			return false
		}

		if diff > 0 {
			increasing = true
		} else if diff < 0 {
			decreasing = true
		} else {
			return false
		}

		if decreasing && increasing {
			return false
		}
	}

	return true
}

func getSafeReports(lines []string) (int, error) {
	count := 0

	for _, line := range lines {
		levels, err := getLevels(line)
		if err != nil {
			return 0, err
		}

		if isReportSafe(levels) {
			count++
		}
	}

	return count, nil
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	safeReports, err := getSafeReports(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[PART 1] safe reports: ", safeReports)
}
