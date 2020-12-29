package main

import (
	"fmt"
)

var history map[int][]int
var start []int = []int{0, 3, 1, 6, 7, 5}

func Part1(start []int) int {
	return nthWordSpoken(2020, start)
}

func Part2(start []int) int {
	return nthWordSpoken(30000000, start)
}

func nthWordSpoken(n int, start []int) int {
	history = make(map[int][]int)
	var lastSpoken int
	turn := 1

	for _, num := range start {
		lastSpoken = num
		history[num] = append(history[num], turn)
		turn++
	}

	for turn <= n {
		timesSpoken := len(history[lastSpoken])
		if timesSpoken == 1 {
			lastSpoken = 0
		} else {
			lastSpoken = history[lastSpoken][timesSpoken-1] - history[lastSpoken][timesSpoken-2]
		}
		history[lastSpoken] = append(history[lastSpoken], turn)
		turn++
	}

	return lastSpoken
}

func main() {
	fmt.Printf("Part 1: %d\n", Part1(start))
	fmt.Printf("Part 2: %d\n", Part2(start))
}
