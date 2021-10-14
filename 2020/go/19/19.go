package main

import (
	"aoc2020/utils"
	"fmt"
	"regexp"
	"strings"
)

var rules map[string]messageRule
var data []string

var ruleRegex *regexp.Regexp = regexp.MustCompile(`(\d+): (.*)`)
var messageRegex *regexp.Regexp = regexp.MustCompile(`[ab]+`)

func Part1(input []string) (count int) {
	rules = make(map[string]messageRule)
	data = make([]string, 0)

	for _, line := range input {
		parseLine(line)
	}

	rule := rules["0"]
	for _, message := range data {
		if rule.Matches(message) == len(message) {
			count++
		}
	}

	return
}

func Part2(input []string) (count int) {
	rules = make(map[string]messageRule)
	data = make([]string, 0)

	for _, line := range input {
		parseLine(line)
	}

	rule42 := rules["42"]
	rule31 := rules["31"]

	for _, message := range data {
		var matches42, matches31 int
		match := rule42.Matches(message)
		for match > 0 {
			matches42++
			message = message[match:]
			match = rule42.Matches(message)
		}
		if matches42 == 0 {
			continue
		}
		match = rule31.Matches(message)
		for match > 0 {
			matches31++
			message = message[match:]
			match = rule31.Matches(message)
		}
		if matches31 > 0 && matches42 > matches31 && len(message) == 0 {
			count++
		}
	}

	return
}

func parseLine(line string) {
	tokens := ruleRegex.FindStringSubmatch(line)
	if tokens == nil {
		message := messageRegex.FindStringSubmatch(line)
		if message != nil {
			data = append(data, line)
		}

		return
	}

	tokens[2] = strings.Trim(tokens[2], "\"")
	if tokens[2] == "a" || tokens[2] == "b" {
		rules[tokens[1]] = &literalRule{tokens[2]}
	} else {
		rules[tokens[1]] = NewSequenceRule(tokens[2])
	}
}

type messageRule interface {
	Matches(input string) int
}

type literalRule struct {
	literal string
}

func (rule *literalRule) Matches(input string) int {
	if len(input) > 0 && input[0] == rule.literal[0] {
		return 1
	}

	return 0
}

type sequenceRule struct {
	conditions [][]string
}

func (rule *sequenceRule) Matches(input string) int {
	for _, sequence := range rule.conditions {
		test := input
		total := 0
		for _, condition := range sequence {
			subrule := rules[condition]
			match := subrule.Matches(test)
			if match > 0 {
				total += match
				test = test[match:]
			} else {
				total = 0
				break
			}
		}

		if total > 0 {
			return total
		}
	}

	return 0
}

func NewSequenceRule(input string) *sequenceRule {
	rule := &sequenceRule{make([][]string, 0)}
	subs := strings.SplitN(input, "|", -1)
	for _, sub := range subs {
		rule.conditions = append(rule.conditions, strings.Fields(sub))
	}
	return rule
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
