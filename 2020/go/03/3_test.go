package main

import (
	"testing"

	"../utils"
)

func TestPart1(t *testing.T) {
	utils.TestExamples(t, Part1, map[string]int{"sample": 7})
}

func TestPart2(t *testing.T) {
	utils.TestExamples(t, Part2, map[string]int{"sample": 336})
}
