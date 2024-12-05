package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc/utils"
)

func getRulesAndUpdates(lines []string) ([][]int, [][]int, error) {
	isRule := true
	var rules, updates [][]int

	for _, line := range lines {
		if len(line) == 0 {
			isRule = false
			continue
		}

		if isRule {
			parts := strings.Split(line, "|")
			lPart, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, nil, err
			}

			rPart, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, err
			}

			rules = append(rules, []int{lPart, rPart})
		} else {
			var updatesSlice []int
			parts := strings.Split(line, ",")

			for _, part := range parts {
				num, err := strconv.Atoi(string(part))
				if err != nil {
					return nil, nil, err
				}

				updatesSlice = append(updatesSlice, num)
			}

			updates = append(updates, updatesSlice)
		}
	}

	return rules, updates, nil
}

func isUpdateValid(update []int, rules [][]int) bool {
	pageIndices := make(map[int]int)
	for i, page := range update {
		pageIndices[page] = i
	}

	for _, rule := range rules {
		x, y := rule[0], rule[1]
		xIndex, xExists := pageIndices[x]
		yIndex, yExists := pageIndices[y]

		if xExists && yExists && xIndex >= yIndex {
			return false
		}
	}

	return true
}

func updatesFollowsRules(updates [][]int, rules [][]int) int {
	sumOfMids := 0

	for _, update := range updates {
		if isUpdateValid(update, rules) {
			midIndex := len(update) / 2
			sumOfMids += update[midIndex]
		}
	}

	return sumOfMids
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rules, updates, err := getRulesAndUpdates(lines)
	if err != nil {
		log.Fatal(err)
	}

	sumOfMids := updatesFollowsRules(updates, rules)

	fmt.Println("[PART1] sum of mids: ", sumOfMids)
}
