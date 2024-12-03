package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"aoc/utils"
)

func getValidMulExpFromPattern(pattern string) [][]string {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	return re.FindAllStringSubmatch(pattern, -1)
}

func getMuls(matches [][]string) ([]int, error) {
	var muls []int
	for _, match := range matches {
		num1Str, num2Str := match[1], match[2]

		num1, err := strconv.Atoi(num1Str)
		if err != nil {
			return nil, err
		}

		num2, err := strconv.Atoi(num2Str)
		if err != nil {
			return nil, err
		}

		muls = append(muls, num1*num2)
	}

	return muls, nil
}

func getAddUpOfMuls(muls []int) int {
	sum := 0
	for _, mul := range muls {
		sum += mul
	}

	return sum
}

func getTotalSumOfNormalMuls(lines []string) (int, error) {
	totalSum := 0
	for _, line := range lines {
		matches := getValidMulExpFromPattern(line)
		muls, err := getMuls(matches)
		if err != nil {
			return 0, err
		}

		totalSum += getAddUpOfMuls(muls)
	}

	return totalSum, nil
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	totalSum, err := getTotalSumOfNormalMuls(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[PART 1] total sum: ", totalSum)
}
