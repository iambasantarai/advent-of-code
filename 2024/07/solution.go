package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc/utils"
)

func generateOperatorCombinations(numCount int, operators []string) [][]string {
	if numCount <= 1 {
		return [][]string{}
	}
	if numCount == 2 {
		var combinations [][]string
		for _, op := range operators {
			combinations = append(combinations, []string{op})
		}
		return combinations
	}
	subCombos := generateOperatorCombinations(numCount-1, operators)
	var result [][]string
	for _, op := range operators {
		for _, subCombo := range subCombos {
			newCombo := append([]string{op}, subCombo...)
			result = append(result, newCombo)
		}
	}
	return result
}

func evaluateExpression(nums []int, ops []string) (int, bool, error) {
	expr := strconv.Itoa(nums[0])
	for i := 1; i < len(nums); i++ {
		expr += " " + ops[i-1] + " " + strconv.Itoa(nums[i])
	}
	return evalExprWithConcat(expr)
}

func evalExprWithConcat(expression string) (int, bool, error) {
	parts := strings.Fields(expression)
	current := parts[0]

	for i := 1; i < len(parts); i += 2 {
		op := parts[i]
		num := parts[i+1]

		if op == "||" {
			current += num
		} else {
			currVal, err := strconv.Atoi(current)
			if err != nil {
				return 0, false, fmt.Errorf("invalid number %q: %w", current, err)
			}

			nextVal, err := strconv.Atoi(num)
			if err != nil {
				return 0, false, fmt.Errorf("invalid number %q: %w", num, err)
			}

			if op == "+" {
				currVal += nextVal
			} else if op == "*" {
				currVal *= nextVal
			} else {
				return 0, false, fmt.Errorf("unsupported operator %q", op)
			}

			current = strconv.Itoa(currVal)
		}
	}

	result, err := strconv.Atoi(current)
	if err != nil {
		return 0, false, fmt.Errorf("failed to parse final result %q: %w", current, err)
	}

	return result, true, nil
}

func checkCombinations(testValue int, nums []int, operators []string) (bool, error) {
	numCount := len(nums)
	combos := generateOperatorCombinations(numCount, operators)

	for _, combo := range combos {
		result, valid, err := evaluateExpression(nums, combo)
		if err != nil {
			return false, err
		}
		if valid && result == testValue {
			return true, nil
		}
	}
	return false, nil
}

func totalResult(lines []string, operators []string) (int, error) {
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		testValue, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return 0, fmt.Errorf("invalid test value in line %q: %w", line, err)
		}
		numStrings := strings.Fields(strings.TrimSpace(parts[1]))

		var nums []int
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return 0, fmt.Errorf("invalid number %q in line %q: %w", numStr, line, err)
			}
			nums = append(nums, num)
		}

		valid, err := checkCombinations(testValue, nums, operators)
		if err != nil {
			return 0, err
		}

		if valid {
			total += testValue
		}
	}
	return total, nil
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- Day 7: Bridge Repair ---")
	operators := []string{"+", "*"}
	part1Result, err := totalResult(lines, operators)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[PART1] result:", part1Result)

	operatorsWithConcat := []string{"+", "*", "||"}
	part2Result, err := totalResult(lines, operatorsWithConcat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[PART2] result:", part2Result)
}
