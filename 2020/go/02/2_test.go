package main

import (
	"aoc2020/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	utils.TestExamples(t, Part1, map[string]int{"sample": 2})
}

func TestPart2(t *testing.T) {
	utils.TestExamples(t, Part2, map[string]int{"sample": 1})
}
