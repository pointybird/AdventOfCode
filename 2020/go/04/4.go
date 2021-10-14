package main

import (
	"aoc2020/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var required []string = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var heightRegex *regexp.Regexp = regexp.MustCompile(`^(\d+)(cm|in)$`)
var hairColorRegex *regexp.Regexp = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var eyeColorRegex *regexp.Regexp = regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
var passportIDRegex *regexp.Regexp = regexp.MustCompile(`^\d{9}$`)

func Part1(input []string) int {
	return process(input, false)
}

func Part2(input []string) int {
	return process(input, true)
}

func process(input []string, validateFields bool) (validCount int) {
	doc := make(map[string]string)
	for _, line := range input {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			if validate(doc, validateFields) {
				validCount++
			}
			doc = make(map[string]string)
			continue
		}
		for _, field := range fields {
			tokens := strings.Split(field, ":")
			if tokens[0] == "cid" {
				continue
			}
			doc[tokens[0]] = tokens[1]
		}
	}

	if validate(doc, validateFields) {
		validCount++
	}

	return
}

func validate(doc map[string]string, validateFields bool) bool {
	var value string
	var ok bool

	for _, field := range required {
		if value, ok = doc[field]; !ok {
			return false
		}

		if validateFields {
			switch field {
			case "byr":
				if !validateRange(value, 1920, 2002) {
					return false
				}
			case "iyr":
				if !validateRange(value, 2010, 2020) {
					return false
				}
			case "eyr":
				if !validateRange(value, 2020, 2030) {
					return false
				}
			case "hgt":
				tokens := heightRegex.FindStringSubmatch(value)
				if tokens == nil {
					return false
				}
				height, units := tokens[1], tokens[2]
				if (units == "cm" && !validateRange(height, 150, 193)) || (units == "in" && !validateRange(height, 59, 76)) {
					return false
				}
			case "hcl":
				if !hairColorRegex.Match([]byte(value)) {
					return false
				}
			case "ecl":
				if !eyeColorRegex.Match([]byte(value)) {
					return false
				}
			case "pid":
				if !passportIDRegex.Match([]byte(value)) {
					return false
				}
			}
		}
	}

	return true
}

func validateRange(value string, min, max int) bool {
	year, err := strconv.Atoi(value)
	return err == nil && year >= min && year <= max
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
