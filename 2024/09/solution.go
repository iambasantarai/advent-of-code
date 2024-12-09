package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines, err := utils.ReadFileLineByLine("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	checkSum1 := compactDiskByMovingBlocksAndCalculateChecksum(strings.Split(lines[0], ""))
	fmt.Println("[PART 1] checksum: ", checkSum1)

	checkSum2 := compactDiskByMovingFilesAndCalculateChecksum(strings.Split(lines[0], ""))
	fmt.Println("[PART 2] checksum: ", checkSum2)
}

func compactDiskByMovingBlocksAndCalculateChecksum(diskMapInput []string) int64 {
	checksum := int64(0)
	diskMap := make([]string, 0)

	for i := range diskMapInput {
		if i%2 == 0 {
			fileSpace, _ := strconv.ParseInt(diskMapInput[i], 10, 32)
			for j := 0; j < int(fileSpace); j++ {
				diskMap = append(diskMap, fmt.Sprintf("%d", i/2))
			}
		} else {
			freeSpace, _ := strconv.ParseInt(diskMapInput[i], 10, 32)
			for j := 0; j < int(freeSpace); j++ {
				diskMap = append(diskMap, ".")
			}
		}
	}

	for i, j := 0, len(diskMap)-1; i <= j; {
		if diskMap[i] == "." && diskMap[j] != "." {
			diskMap[i], diskMap[j] = diskMap[j], diskMap[i]
			i++
			j--
		} else {
			if diskMap[i] != "." {
				i++
			} else if diskMap[j] != "." {
				j--
			} else if diskMap[i] == "." && diskMap[j] == "." {
				j--
			}
		}
	}

	for i := range diskMap {
		if diskMap[i] != "." {
			fileID, _ := strconv.ParseInt(diskMap[i], 10, 32)
			checksum += int64(i) * fileID
		}
	}

	return checksum
}

func compactDiskByMovingFilesAndCalculateChecksum(diskMapInput []string) int64 {
	solution := int64(0)
	diskMap := make([]string, 0)
	files := make([][]int64, 0)
	spaces := make([][]int64, 0)
	for i := range diskMapInput {
		if i%2 == 0 {
			id := i / 2
			fileSpace, _ := strconv.ParseInt(diskMapInput[i], 10, 32)
			files = append(files, []int64{int64(len(diskMap)), int64(len(diskMap)-1) + fileSpace})
			for j := 0; j < int(fileSpace); j++ {
				diskMap = append(diskMap, fmt.Sprintf("%d", id))
			}
		} else {
			space, _ := strconv.ParseInt(diskMapInput[i], 10, 32)
			spaces = append(spaces, []int64{int64(len(diskMap)), int64(len(diskMap)-1) + space})
			for j := 0; j < int(space); j++ {
				diskMap = append(diskMap, ".")
			}
		}
	}

	for j := len(files) - 1; j >= 0; j-- {
		for i := 0; i < len(spaces); i++ {
			space := spaces[i]
			if space[1] < files[j][1] && space[1]-space[0] >= files[j][1]-files[j][0] {
				for k := space[0]; k <= space[0]+files[j][1]-files[j][0]; k++ {
					diskMap[k] = fmt.Sprintf("%d", j)
				}
				for k := files[j][0]; k <= files[j][1]; k++ {
					diskMap[k] = "."
				}
				space[0] = space[0] + files[j][1] - files[j][0] + 1
				break
			}
		}
	}

	for i := range diskMap {
		if diskMap[i] != "." {
			x, _ := strconv.ParseInt(diskMap[i], 10, 32)
			solution += int64(i) * x
		}
	}

	return solution
}
