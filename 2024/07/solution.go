package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
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

func evaluateExpression(nums []int, ops []string) (int, bool) {
	expr := strconv.Itoa(nums[0])
	for i := 1; i < len(nums); i++ {
		expr += " " + ops[i-1] + " " + strconv.Itoa(nums[i])
	}
	return evalExprWithConcat(expr)
}

func evalExprWithConcat(expression string) (int, bool) {
	parts := strings.Fields(expression)
	current := parts[0]
	for i := 1; i < len(parts); i += 2 {
		op := parts[i]
		num := parts[i+1]
		if op == "||" {
			current += num
		} else {
			currVal, err1 := strconv.Atoi(current)
			nextVal, err2 := strconv.Atoi(num)
			if err1 != nil || err2 != nil {
				return 0, false
			}
			if op == "+" {
				currVal += nextVal
			} else if op == "*" {
				currVal *= nextVal
			}
			current = strconv.Itoa(currVal)
		}
	}
	result, err := strconv.Atoi(current)
	if err != nil {
		return 0, false
	}
	return result, true
}

func checkCombinations(testValue int, nums []int, operators []string) bool {
	numCount := len(nums)
	combos := generateOperatorCombinations(numCount, operators)

	for _, combo := range combos {
		result, valid := evaluateExpression(nums, combo)
		if valid && result == testValue {
			return true
		}
	}
	return false
}

func totalResult(lines []string, operators []string) int {
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		testValue, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		numStrings := strings.Fields(strings.TrimSpace(parts[1]))

		var nums []int
		for _, numStr := range numStrings {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		if checkCombinations(testValue, nums, operators) {
			total += testValue
		}
	}
	return total
}

func checkCombinationsWithConcat(testValue int, nums []int, operators []string) bool {
	numCount := len(nums)
	combos := generateOperatorCombinations(numCount, operators)

	for _, combo := range combos {
		result, valid := evaluateExpression(nums, combo)
		if valid && result == testValue {
			return true
		}
	}
	return false
}

func totalResultWithConcat(lines []string, operators []string) int {
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		testValue, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		numStrings := strings.Fields(strings.TrimSpace(parts[1]))

		var nums []int
		for _, numStr := range numStrings {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		if checkCombinationsWithConcat(testValue, nums, operators) {
			total += testValue
		}
	}
	return total
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	operators := []string{"+", "*"}
	part1Result := totalResult(lines, operators)
	fmt.Println("[PART1] result:", part1Result)

	operatorsWithConcat := []string{"+", "*", "||"}
	part2Result := totalResultWithConcat(lines, operatorsWithConcat)
	fmt.Println("[PART2] result:", part2Result)
}
