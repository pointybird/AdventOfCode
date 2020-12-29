package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

const InputFile = "input"

func ReadLines(path string) (lines []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func ReadInts(path string) (nums []int) {
	lines := ReadLines(path)

	for _, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, value)
	}

	return
}

func TestExamples(t *testing.T, f func([]string) int, examples map[string]int) {
	for example, expected := range examples {
		t.Run(example, func(t *testing.T) {
			input := ReadLines(filepath.FromSlash("examples/" + example))
			result := f(input)
			if result != expected {
				t.Errorf("Expected %d, got %d", expected, result)
			}
		})
	}
}

func TestExampleInts(t *testing.T, f func([]int) int, examples map[string]int) {
	for example, expected := range examples {
		t.Run(example, func(t *testing.T) {
			input := ReadInts(filepath.FromSlash("examples/" + example))
			result := f(input)
			if result != expected {
				t.Errorf("Expected %d, got %d", expected, result)
			}
		})
	}
}

func TestInput(t *testing.T, f func(string) int, examples map[string]int) {
	for example, expected := range examples {
		t.Run(example, func(t *testing.T) {
			result := f(example)
			if result != expected {
				t.Errorf("Expected %d, got %d", expected, result)
			}
		})
	}
}
