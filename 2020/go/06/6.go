package main

import (
	"aoc2020/utils"
	"fmt"
)

func Part1(input []string) (total int) {
	groupAnswers := make(map[rune]bool)
	for _, line := range input {
		if len(line) == 0 {
			total += len(groupAnswers)
			groupAnswers = make(map[rune]bool)
			continue
		}
		for _, ch := range line {
			groupAnswers[ch] = true
		}
	}

	total += len(groupAnswers)

	return
}

func Part2(input []string) (total int) {
	groupAnswers := make(map[rune]int)
	groupSize := 0
	for _, line := range input {
		if len(line) == 0 {
			total += checkAnsweredByAll(groupAnswers, groupSize)
			groupAnswers = make(map[rune]int)
			groupSize = 0
			continue
		}
		groupSize++
		for _, ch := range line {
			groupAnswers[ch]++
		}
	}

	total += checkAnsweredByAll(groupAnswers, groupSize)

	return
}

func checkAnsweredByAll(answers map[rune]int, groupSize int) (total int) {
	for _, v := range answers {
		if v == groupSize {
			total++
		}
	}

	return
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
