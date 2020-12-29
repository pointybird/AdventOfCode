package main

import (
	"testing"

	"../utils"
)

func TestPart1(t *testing.T) {
	utils.TestExamples(t, Part1, map[string]int{
		"a": 71,
		"b": 51,
		"c": 26,
		"d": 437,
		"e": 12240,
		"f": 13632,
	})
}

func TestPart2(t *testing.T) {
	utils.TestExamples(t, Part2, map[string]int{
		"a": 231,
		"b": 51,
		"c": 46,
		"d": 1445,
		"e": 669060,
		"f": 23340,
	})
}
