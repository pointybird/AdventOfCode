package main

import (
	"testing"

	"../utils"
)

func TestPart1(t *testing.T) {
	utils.TestExamples(t, Part1, map[string]int{"1a": 2})
}

func TestPart2(t *testing.T) {
	utils.TestExamples(t, Part2, map[string]int{"2a": 0, "2b": 4})
}
