package main

import (
	"path/filepath"
	"testing"

	"../utils"
)

func TestPart1(t *testing.T) {
	example := "sample"
	preambleLength := 5
	expected := 127

	t.Run(example, func(t *testing.T) {
		input := utils.ReadInts(filepath.FromSlash("examples/" + example))
		result := Part1(input, preambleLength)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func TestPart2(t *testing.T) {
	example := "sample"
	target := 127
	expected := 62

	t.Run(example, func(t *testing.T) {
		input := utils.ReadInts(filepath.FromSlash("examples/" + example))
		result := Part2(input, target)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
