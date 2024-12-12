package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc/utils"
)

func getIntArray(line string) ([]int, error) {
	var nums []int

	numbers := strings.Fields(line)

	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

func splitNumber(num int) (int, int, error) {
	str := strconv.Itoa(num)
	mid := len(str) / 2

	left, err := strconv.Atoi(str[:mid])
	if err != nil {
		return 0, 0, err
	}

	right, err := strconv.Atoi(str[mid:])
	if err != nil {
		return 0, 0, err
	}

	return left, right, nil
}

func getStoneCountAfterBlink(arrangement []int, blinks int) (int, error) {
	queue := append([]int{}, arrangement...)
	for i := 0; i < blinks; i++ {
		nextQueue := []int{}
		for _, stone := range queue {
			switch {
			case stone == 0:
				nextQueue = append(nextQueue, 1)
			case len(strconv.Itoa(stone))%2 == 0:
				left, right, err := splitNumber(stone)
				if err != nil {
					return 0, err
				}
				nextQueue = append(nextQueue, left, right)
			default:
				nextQueue = append(nextQueue, stone*2024)
			}
		}
		queue = nextQueue
	}
	return len(queue), nil
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	arrangement, err := getIntArray(lines[0])
	if err != nil {
		log.Fatal(err)
	}

	stoneCountAfter25Blinks, err := getStoneCountAfterBlink(arrangement, 25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[PART 1] stones: ", stoneCountAfter25Blinks)
}
