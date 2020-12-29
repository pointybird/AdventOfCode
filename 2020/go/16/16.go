package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../utils"
)

var ticketRules map[string][4]int = make(map[string][4]int)
var ruleRegex *regexp.Regexp = regexp.MustCompile(`(.*): (\d+)\-(\d+) or (\d+)\-(\d+)`)
var ticketRegex *regexp.Regexp = regexp.MustCompile(`(\d+),?`)
var myTicket []int = []int{}
var otherTickets [][]int = make([][]int, 0)
var goodTickets [][]int = make([][]int, 0)
var fieldMapping map[int]map[string]bool = make(map[int]map[string]bool)

func Part1(input []string) (errorRate int) {
	readTickets(input)

	for i, ticket := range otherTickets {
		errorRate += checkError(i, ticket)
	}

	return
}

func Part2(input []string) int {
	readTickets(input)

	for i, ticket := range otherTickets {
		checkError(i, ticket)
	}

	goodTickets = append(goodTickets, myTicket)

	determineFields()

	departureIndexes := solveUnique()

	total := 1
	for _, index := range departureIndexes {
		total *= myTicket[index]
	}

	return total
}

func readTickets(input []string) {
	var mine, other bool

	for _, line := range input {
		tokens := ruleRegex.FindStringSubmatch(line)
		if tokens != nil {
			min1, _ := strconv.Atoi(tokens[2])
			max1, _ := strconv.Atoi(tokens[3])
			min2, _ := strconv.Atoi(tokens[4])
			max2, _ := strconv.Atoi(tokens[5])

			ticketRules[tokens[1]] = [4]int{min1, max1, min2, max2}
		} else if line == "your ticket:" {
			mine = true
		} else if line == "nearby tickets:" {
			other = true
		} else {
			tokens := ticketRegex.FindAllStringSubmatch(line, -1)
			if tokens != nil {
				if other {
					otherTickets = append(otherTickets, parseTicket(line))
				} else if mine {
					myTicket = parseTicket(line)
				}
			}
		}
	}
}

func parseTicket(line string) []int {
	ticket := []int{}

	nums := strings.Split(line, ",")
	for _, num := range nums {
		val, _ := strconv.Atoi(num)
		ticket = append(ticket, val)
	}

	return ticket
}

func checkError(index int, ticket []int) (errors int) {
	var error bool
	for _, num := range ticket {
		var valid bool
		for _, ranges := range ticketRules {
			if (num >= ranges[0] && num <= ranges[1]) || (num >= ranges[2] && num <= ranges[3]) {
				valid = true
				break
			}
		}
		if !valid {
			error = true
			errors += num
		}
	}

	if !error {
		goodTickets = append(goodTickets, ticket)
	}

	return
}

func determineFields() {
	for _, ticket := range goodTickets {
		for section, num := range ticket {
			for name, ranges := range ticketRules {
				sectionValid := (num >= ranges[0] && num <= ranges[1]) || (num >= ranges[2] && num <= ranges[3])
				if fields, ok := fieldMapping[section]; !ok {
					fields = make(map[string]bool)
					fieldMapping[section] = fields
					fields[name] = sectionValid
				} else {
					if possible, ok2 := fields[name]; !ok2 {
						fields[name] = sectionValid
					} else {
						fields[name] = possible && sectionValid
					}
				}
			}
		}
	}
}

func solveUnique() (departureIndexes []int) {
	var mappedFields int

	for mappedFields < len(fieldMapping) {
		for index, fields := range fieldMapping {
			if field, count := possibleCount(fields); count == 1 {
				mappedFields++
				if strings.HasPrefix(field, "departure") {
					departureIndexes = append(departureIndexes, index)
				}
				for _, v := range fieldMapping {
					delete(v, field)
				}
				break
			}
		}
	}

	return
}

func possibleCount(mapping map[string]bool) (possibleField string, count int) {
	for field, possible := range mapping {
		if possible {
			possibleField = field
			count++
		}
	}
	return
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
