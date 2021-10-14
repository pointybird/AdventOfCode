package main

import (
	"aoc2020/utils"
	"fmt"
)

func Part1(input []string) int {
	return sled(input, 3, 1)
}

func Part2(input []string) int {
	slopes := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	total := 1
	for _, slope := range slopes {
		total *= sled(input, slope[0], slope[1])
	}

	return total
}

func sled(input []string, right, down int) (trees int) {
	pos := 0
	for i := down; i < len(input); i += down {
		pos += right
		row := input[i]
		if string(row[pos%len(row)]) == "#" {
			trees++
		}
	}

	return
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
