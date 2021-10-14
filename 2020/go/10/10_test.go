package main

import (
	"aoc2020/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	utils.TestExampleInts(t, Part1, map[string]int{
		"sample1": 35,
		"sample2": 220})
}

func TestPart2(t *testing.T) {
	utils.TestExampleInts(t, Part2, map[string]int{
		"sample1": 8,
		"sample2": 19208})
}
