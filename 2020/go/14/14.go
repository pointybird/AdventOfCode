package main

import (
	"aoc2020/utils"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

var mask []string
var memory map[int]big.Int

var maskRegex *regexp.Regexp = regexp.MustCompile(`^mask = (\w+)$`)
var memRegex *regexp.Regexp = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func Part1(input []string) int {
	return int(evaluate(input, setMemory, setMask))
}

func Part2(input []string) int {
	return int(evaluate(input, setMemoryV2, setMaskV2))
}

func evaluate(input []string, setMemoryFunc func(int64, int64), setMaskFunc func(string)) (sum int64) {
	memory = make(map[int]big.Int)

	for _, line := range input {
		tokens := maskRegex.FindStringSubmatch(line)
		if tokens == nil {
			tokens = memRegex.FindStringSubmatch(line)
			index, _ := strconv.Atoi(tokens[1])
			value, _ := strconv.Atoi(tokens[2])

			setMemoryFunc(int64(index), int64(value))
		} else {
			setMaskFunc(tokens[1])
		}
	}

	for _, v := range memory {
		sum += v.Int64()
	}

	return
}

func setMemory(index, value int64) {
	newValue := big.NewInt(value)

	for i := 0; i < len(mask); i++ {
		switch mask[i] {
		case "1":
			newValue.SetBit(newValue, i, 1)
		case "0":
			newValue.SetBit(newValue, i, 0)
		}
	}
	memory[int(index)] = *newValue
}

func setMemoryV2(index, value int64) {
	b := strconv.FormatInt(index, 2)
	addressBytes := []byte(strings.Repeat("0", 36-len(b)) + b)

	for i := 0; i < len(mask); i++ {
		bit := mask[i][0]
		if bit == '0' {
			continue
		}
		addressBytes[i] = bit
	}

	addressText := string(addressBytes)
	wild := strings.Count(addressText, "X")
	if wild > 0 {
		count := int(math.Pow(2, float64(wild)))
		for i := 0; i < count; i++ {
			address := substituteAddress(addressText, i)
			memory[int(address)] = *big.NewInt(value)
		}
	} else {
		address, _ := strconv.ParseInt(addressText, 2, 64)
		memory[int(address)] = *big.NewInt(value)
	}
}

func setMask(input string) {
	var m string
	for _, v := range input {
		m = string(v) + m
	}
	mask = strings.Split(m, "")
}

func setMaskV2(input string) {
	mask = strings.Split(input, "")
}

func substituteAddress(addressText string, current int) int64 {
	text := addressText
	val := strconv.FormatInt(int64(current), 2)
	v := strings.Repeat("0", 36-len(val)) + val
	var m string
	for _, bit := range v {
		m = string(bit) + m
	}
	ids := strings.Split(m, "")
	i := 0
	for strings.Contains(text, "X") {
		text = strings.Replace(text, "X", ids[i], 1)
		i++
	}

	sub, _ := strconv.ParseInt(text, 2, 64)
	return sub
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
