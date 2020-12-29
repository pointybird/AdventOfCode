package main

import (
	"testing"

	"../utils"
)

func TestPart1(t *testing.T) {
	utils.TestExampleInts(t, Part1, map[string]int{"sample": 514579})
}

func TestPart2(t *testing.T) {
	utils.TestExampleInts(t, Part2, map[string]int{"sample": 241861950})
}
