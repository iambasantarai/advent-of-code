package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"aoc/utils"
)

type Coordinates struct {
	X, Y int
}

func fetchSliceOfIntsInString(s string) ([]int, error) {
	var nums []int

	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(s, -1)

	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

// Solve the system of equations using Cramer's Rule
func solveLinearEquation(a, b, target Coordinates) (bool, Coordinates) {
	det := a.X*b.Y - a.Y*b.X
	if det == 0 {
		return false, Coordinates{}
	}

	detX := target.X*b.Y - target.Y*b.X
	detY := a.X*target.Y - a.Y*target.X

	if detX%det != 0 || detY%det != 0 {
		return false, Coordinates{}
	}

	x := detX / det
	y := detY / det

	return x >= 0 && y >= 0, Coordinates{x, y}
}

func getPrizes(input []string) ([][]Coordinates, error) {
	var values []Coordinates
	var prizes [][]Coordinates

	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		nums, err := fetchSliceOfIntsInString(line)
		if err != nil {
			return nil, err
		}
		values = append(values, Coordinates{nums[0], nums[1]})

		if strings.Contains(line, "Prize") {
			prizes = append(prizes, values)
			values = []Coordinates{}
		}
	}

	return prizes, nil
}

func getButtonPressesIfValid(a, b, final Coordinates, add int) (bool, Coordinates) {
	final.X += add
	final.Y += add

	return solveLinearEquation(a, b, final)
}

func getTokenCount(input [][]Coordinates) (int, int) {
	var count, count1 int = 0, 0

	for _, prize := range input {
		isValid, tokens := getButtonPressesIfValid(prize[0], prize[1], prize[2], 0)
		if isValid {
			count += tokens.X*3 + tokens.Y
		}

		isValid1, tokens1 := getButtonPressesIfValid(prize[0], prize[1], prize[2], 10000000000000)
		if isValid1 {
			count1 += tokens1.X*3 + tokens1.Y
		}

	}

	return count, count1
}

func main() {
	lines, err := utils.ReadFileLineByLine("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	prizes, err := getPrizes(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- Day 13: Claw Contraption ---")
	tokens, tokens1 := getTokenCount(prizes)
	fmt.Println("[PART 1] tokens: ", tokens)
	fmt.Println("[PART 2] tokens: ", tokens1)
}
