package main

import (
	"aoc2020/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	utils.TestExamples(t, Part1, map[string]int{"1a": 4})
}

func TestPart2(t *testing.T) {
	utils.TestExamples(t, Part2, map[string]int{
		"2a": 32,
		"2b": 126})
}
