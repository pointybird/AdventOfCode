package main

import (
	"aoc2020/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func validate(pattern *regexp.Regexp, input []string, passwordValidator func([]string) bool) (count int) {
	for _, line := range input {
		tokens := pattern.FindStringSubmatch(line)
		if passwordValidator(tokens) {
			count++
		}
	}

	return
}

func policy1(tokens []string) bool {
	min, _ := strconv.Atoi(tokens[1])
	max, _ := strconv.Atoi(tokens[2])
	required := tokens[3]
	password := tokens[4]

	count := strings.Count(password, required)

	return min <= count && count <= max
}

func policy2(tokens []string) bool {
	index1, _ := strconv.Atoi(tokens[1])
	index2, _ := strconv.Atoi(tokens[2])
	required := tokens[3][0]
	password := tokens[4]

	index1Matches := password[index1-1] == required
	index2Matches := password[index2-1] == required

	return index1Matches != index2Matches
}

func Part1(input []string) int {
	pattern := regexp.MustCompile(`(\d+)\-(\d+) ([a-z]): ([a-z]*)`)

	return validate(pattern, input, policy1)
}

func Part2(input []string) int {
	pattern := regexp.MustCompile(`(\d+)\-(\d+) ([a-z]): ([a-z]*)`)

	return validate(pattern, input, policy2)
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
