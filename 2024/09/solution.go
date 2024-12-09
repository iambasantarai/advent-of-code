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
