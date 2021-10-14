package main

import (
	"aoc2020/utils"
	"fmt"
	"sort"
)

func Part1(input []string) int {
	var maxID int = 0
	for _, line := range input {
		id := DecodeBoardingPass(line)
		if id > maxID {
			maxID = id
		}
	}

	return maxID
}

func Part2(input []string) int {
	ids := []int{}
	for _, line := range input {
		ids = append(ids, DecodeBoardingPass(line))
	}
	sort.Ints(ids)
	offset := ids[0]
	for pos, id := range ids {
		if pos+offset != id {
			return pos + offset
		}
	}
	return -1
}

func DecodeBoardingPass(pass string) int {
	var minRow, minCol int = 0, 0
	maxRow := 127
	maxCol := 7

	runes := []rune(pass)
	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case 'F':
			maxRow = (minRow+maxRow+1)/2 - 1
		case 'B':
			minRow = (minRow + maxRow + 1) / 2
		case 'L':
			maxCol = (minCol+maxCol+1)/2 - 1
		case 'R':
			minCol = (minCol + maxCol + 1) / 2
		}
	}

	return maxRow*8 + maxCol
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
