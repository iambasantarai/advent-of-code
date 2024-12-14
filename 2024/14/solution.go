package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"aoc/utils"
)

const (
	width, height = 101, 103
	time          = 100
)

type Robot struct {
	x, y   int
	vx, vy int
}

func getRobotPositions(lines []string) []Robot {
	var robots []Robot

	re := regexp.MustCompile(`p=([0-9]*),([0-9]*) v=(-?[0-9]*),(-?[0-9]*)`)

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		x, _ := strconv.Atoi(matches[0][1])
		y, _ := strconv.Atoi(matches[0][2])
		vx, _ := strconv.Atoi(matches[0][3])
		vy, _ := strconv.Atoi(matches[0][4])

		robots = append(robots, Robot{x: x, y: y, vx: vx, vy: vy})
	}

	return robots
}

func simulate(robots []Robot, time int) []Robot {
	for i := 0; i < time; i++ {
		for j := range robots {
			robots[j].x = (robots[j].x + robots[j].vx + width) % width
			robots[j].y = (robots[j].y + robots[j].vy + height) % height
		}
	}
	return robots
}

func countRobotsInQuadrants(robots []Robot) (int, int, int, int) {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, robot := range robots {
		if robot.x == width/2 || robot.y == height/2 {
			continue
		}
		if robot.x > width/2 && robot.y < height/2 {
			q1++
		} else if robot.x < width/2 && robot.y < height/2 {
			q2++
		} else if robot.x < width/2 && robot.y > height/2 {
			q3++
		} else if robot.x > width/2 && robot.y > height/2 {
			q4++
		}
	}

	return q1, q2, q3, q4
}

func computeSafetyFactor(q1, q2, q3, q4 int) int {
	return q1 * q2 * q3 * q4
}

func calculateTotalSafetyFactor(robots []Robot) int {
	movements := simulate(robots, time)
	q1, q2, q3, q4 := countRobotsInQuadrants(movements)

	safetyFactor := computeSafetyFactor(q1, q2, q3, q4)

	return safetyFactor
}

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	robots := getRobotPositions(lines)
	safetyFactor := calculateTotalSafetyFactor(robots)

	fmt.Println("--- Day 14: Restroom Redoubt ---")
	fmt.Println("[PART 1] safety factor: ", safetyFactor)
}
