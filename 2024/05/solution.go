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

func topologicalSort(graph map[int][]int, inDegree map[int]int, nodes []int) []int {
	var sorted []int
	queue := []int{}

	for _, node := range nodes {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		sorted = append(sorted, curr)

		for _, neighbor := range graph[curr] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func orderCorrectly(update []int, rules [][]int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	pageSet := make(map[int]bool)

	for _, page := range update {
		pageSet[page] = true
		inDegree[page] = 0
	}

	for _, rule := range rules {
		x, y := rule[0], rule[1]
		if pageSet[x] && pageSet[y] {
			graph[x] = append(graph[x], y)
			inDegree[y]++
		}
	}

	sorted := topologicalSort(graph, inDegree, update)

	seen := make(map[int]bool)
	for _, page := range sorted {
		seen[page] = true
	}

	for _, page := range update {
		if !seen[page] {
			sorted = append(sorted, page)
		}
	}

	return sorted
}

func updatesFollowsRules(updates [][]int, rules [][]int) (int, int) {
	sumOfMids := 0
	sumOfOrderCorrectedMids := 0

	for _, update := range updates {
		if isUpdateValid(update, rules) {
			midIndex := len(update) / 2
			sumOfMids += update[midIndex]
		} else {
			orderCorrectedUpdate := orderCorrectly(update, rules)
			midIndex := len(update) / 2
			sumOfOrderCorrectedMids += orderCorrectedUpdate[midIndex]
		}
	}

	return sumOfMids, sumOfOrderCorrectedMids
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

	sumOfMids, sumOfOrderCorrectedMids := updatesFollowsRules(updates, rules)

	fmt.Println("--- Day 5: Print Queue ---")
	fmt.Println("[PART1] sum of mids: ", sumOfMids)
	fmt.Println("[PART2] sum of mids: ", sumOfOrderCorrectedMids)
}
