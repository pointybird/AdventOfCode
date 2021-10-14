package main

import (
	"aoc2020/utils"
	"fmt"
	"sort"
)

func Part1(input []int, preambleLength int) int {
	for i := preambleLength; i < len(input); i++ {
		if !sumInPreviousWindow(input, i, preambleLength) {
			return input[i]
		}
	}

	return 0
}

func Part2(input []int, target int) int {
	min, max := findRangeSumsToTarget(input, target)

	return min + max
}

func sumInPreviousWindow(input []int, index, preambleLength int) bool {
	window := make([]int, preambleLength)
	copy(window, input[index-preambleLength:index])
	sort.Ints(window)

	target := input[index]
	var i int = 0
	var j int = len(window) - 1

	for i < j {
		sum := window[i] + window[j]
		switch {
		case sum == target:
			return true
		case sum < target:
			i++
		case sum > target:
			j--
		}
	}

	return false
}

func findRangeSumsToTarget(input []int, target int) (max int, min int) {
	for i := 0; i < len(input)-1; i++ {
		sum := input[i]
		max = input[i]
		min = input[i]
		for j := i + 1; j < len(input); j++ {
			if input[j] > max {
				max = input[j]
			}
			if input[j] < min {
				min = input[j]
			}
			sum += input[j]
			if sum == target {
				return
			} else if sum > target {
				break
			}
		}
	}

	return
}

func main() {
	input := utils.ReadInts(utils.InputFile)
	preambleLength := 25

	part1 := Part1(input, preambleLength)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", Part2(input, part1))
}
