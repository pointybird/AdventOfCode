package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 436
	start := []int{0, 3, 6}

	t.Run("Part1", func(t *testing.T) {
		result := Part1(start)
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}
