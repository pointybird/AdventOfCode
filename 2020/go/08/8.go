package main

import (
	"fmt"
	"strconv"
	"strings"

	"../utils"
)

func Part1(input []string) int {
	m := NewMachine(input)
	m.execute()

	return m.accumulator
}

func Part2(input []string) int {
	for i := 0; i < len(input); i++ {
		m := NewMachine(input)
		if m.mutateInstruction(i) {
			m.execute()
			if !m.loopError {
				return m.accumulator
			}
		}
	}

	return -1
}

type machine struct {
	accumulator int
	instruction int
	loopError   bool

	program []string

	executed map[int]bool
}

func NewMachine(program []string) *machine {
	m := new(machine)
	m.program = make([]string, len(program))
	copy(m.program, program)
	m.executed = make(map[int]bool)
	return m
}

func (m *machine) execute() {
	for true {
		instruction := m.instruction
		line := m.program[m.instruction]
		tokens := strings.Split(line, " ")
		value, _ := strconv.Atoi(tokens[1])
		switch tokens[0] {
		case "nop":
			m.instruction++
		case "acc":
			m.accumulator += value
			m.instruction++
		case "jmp":
			m.instruction += value
		}

		if _, ok := m.executed[m.instruction]; ok {
			m.loopError = true
			break
		}

		if m.instruction >= len(m.program) {
			break
		}

		m.executed[instruction] = true
	}
}

func (m *machine) mutateInstruction(instruction int) bool {
	line := m.program[instruction]
	swap := []string{"nop", "jmp"}

	for i, op := range swap {
		if strings.HasPrefix(line, op) {
			m.program[instruction] = strings.Replace(line, op, swap[(i+1)%2], 1)
			return true
		}
	}

	return false
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
