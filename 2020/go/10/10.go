package main

import (
	"aoc2020/utils"
	"fmt"
	"sort"
)

var paths map[int]int

func Part1(adapters []int) int {
	adapters = append(adapters, 0) // Add 0-jolt outlet
	sort.Ints(adapters)

	diffs := make(map[int]int)

	for i := 1; i < len(adapters); i++ {
		diff := adapters[i] - adapters[i-1]
		diffs[diff]++
	}

	return diffs[1] * (diffs[3] + 1) // Add 1 to difference of 3s for final device
}

func Part2(adapters []int) int {
	adapters = append(adapters, 0) // Add 0-jolt outlet
	sort.Ints(adapters)

	paths = make(map[int]int)

	return pathsToEnd(adapters, 0)
}

func pathsToEnd(adapters []int, start int) (count int) {
	if value, ok := paths[start]; ok {
		return value
	}

	if start >= len(adapters)-1 {
		count = 1
	} else {
		for i := start + 1; i < len(adapters); i++ {
			if adapters[i]-adapters[start] > 3 {
				break
			}

			count += pathsToEnd(adapters, i)
		}
	}

	paths[start] = count
	return
}

func main() {
	input := utils.ReadInts(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
