package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../utils"
)

var containsPattern *regexp.Regexp = regexp.MustCompile(`(\d+) (\w+ \w+) bag[s]?[,.]`)

var targetBag string = "shiny gold"

var rules map[string]baggageRule

func Part1(input []string) int {
	readRules(input)

	holders := make(map[string]bool)
	for _, rule := range rules {
		if canHold(rule, targetBag) {
			holders[rule.name] = true
		}
	}

	return len(holders)
}

func Part2(input []string) int {
	readRules(input)

	return numBags(rules[targetBag])
}

func readRules(input []string) map[string]baggageRule {
	rules = make(map[string]baggageRule)
	for _, line := range input {
		tokens := strings.Split(line, " bags contain ")
		name, contents := tokens[0], tokens[1]
		rule := findOrAddBaggageRule(name)
		parseContents(rule, contents)
	}

	return rules
}

func parseContents(rule baggageRule, contains string) {
	matches := containsPattern.FindAllStringSubmatch(contains, -1)
	for _, match := range matches {
		contained := match[2]
		count, _ := strconv.Atoi(match[1])
		rule.contains[contained] = count
		containedRule := findOrAddBaggageRule(contained)
		containedRule.containedBy[rule.name] = true
	}
}

type baggageRule struct {
	name        string
	contains    map[string]int
	containedBy map[string]bool
}

func NewBaggageRule(name string) *baggageRule {
	rule := new(baggageRule)
	rule.name = name
	rule.contains = make(map[string]int)
	rule.containedBy = make(map[string]bool)
	return rule
}

func findOrAddBaggageRule(name string) (rule baggageRule) {
	if rule, ok := rules[name]; !ok {
		rule = *NewBaggageRule(name)
		rules[name] = rule
	}
	return rules[name]
}

func canHold(bag baggageRule, target string) bool {
	if _, ok := bag.contains[target]; ok {
		return true
	}

	for sub := range bag.contains {
		subBag := rules[sub]
		if canHold(subBag, target) {
			return true
		}
	}

	return false
}

func numBags(bag baggageRule) (count int) {
	for sub, num := range bag.contains {
		subBag := rules[sub]
		count += num * (1 + numBags(subBag))
	}

	return
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
