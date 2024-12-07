package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func evaluateExpression(numbers []int, operators []rune) int {
	result := numbers[0]
	for i, op := range operators {
		switch op {
		case '+':
			result += numbers[i+1]
		case '*':
			result *= numbers[i+1]
		}
	}
	return result
}

func isValidEquation(target int, numbers []int) bool {
	operatorsCombinations := getOperatorsCombinations(len(numbers) - 1)
	for _, operators := range operatorsCombinations {
		if evaluateExpression(numbers, operators) == target {
			return true
		}
	}
	return false
}

func getOperatorsCombinations(numOperators int) [][]rune {
	combinations := make([][]rune, 0)
	for i := 0; i < (1 << numOperators); i++ {
		combination := make([]rune, numOperators)
		for j := 0; j < numOperators; j++ {
			if (i>>j)&1 == 0 {
				combination[j] = '+'
			} else {
				combination[j] = '*'
			}
		}
		combinations = append(combinations, combination)
	}
	return combinations
}

func totalCalibrationResult(inputLines []string) int {
	total := 0
	for _, line := range inputLines {
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		numStrs := strings.Fields(parts[1])

		numbers := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			numbers[i], _ = strconv.Atoi(numStr)
		}

		if isValidEquation(target, numbers) {
			total += target
		}
	}
	return total
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := totalCalibrationResult(lines)
	fmt.Println("[PART1] result: ", result)
}
