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

func checkReportSafetyWithDampener(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		isSafe := isReportSafeWithDampener(levels, i)
		if isSafe {
			return true
		}
	}

	return false
}

func isReportSafeWithDampener(levels []int, deleteIndex int) bool {
	copyReport := make([]int, len(levels))
	copy(copyReport, levels)

	if deleteIndex == len(copyReport)-1 {
		copyReport = copyReport[:deleteIndex]
	} else {
		copyReport = append(copyReport[:deleteIndex], copyReport[deleteIndex+1:]...)
	}
	return isReportSafe(copyReport)
}

func getSafeReports(lines []string) (int, int, error) {
	safeReports := 0
	safeReportsWithDampener := 0

	for _, line := range lines {
		levels, err := getLevels(line)
		if err != nil {
			return 0, 0, err
		}

		if isReportSafe(levels) {
			safeReports++
			safeReportsWithDampener++
		} else if checkReportSafetyWithDampener(levels) {
			safeReportsWithDampener++
		}
	}

	return safeReports, safeReportsWithDampener, nil
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	safeReports, safeReportsWithDampener, err := getSafeReports(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- Day 2: Red-Nosed Reports ---")
	fmt.Println("[PART 1] safe reports: ", safeReports)
	fmt.Println("[PART 2] safe reports with dampener: ", safeReportsWithDampener)
}
