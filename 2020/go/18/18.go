package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"../utils"
)

func Part1(input []string) (sum int64) {
	for _, line := range input {
		tokens := strings.Split(strings.ReplaceAll(line, " ", ""), "")
		result, _ := strconv.Atoi(eval(tokens)[0])
		// fmt.Println(result)
		sum += int64(result)
	}

	return
}

func Part2(input []string) (sum *big.Int) {
	sum = new(big.Int)

	for _, line := range input {
		tokens := strings.Split(strings.ReplaceAll(line, " ", ""), "")
		ev := eval2(tokens)
		for len(ev) > 1 {
			ev = eval2(ev)
		}
		result, _ := new(big.Int).SetString(ev[0], 10)
		// fmt.Println(result)
		sum.Add(sum, result)
	}

	return
}

func eval(tokens []string) []string {
	fmt.Printf("Evaluating %s\n", strings.Join(tokens, ""))
	if len(tokens) == 1 {
		return tokens
	}

	stack := make([]string, 0)

	for _, token := range tokens {
		switch token {
		case ")":
			index := len(stack)
			for j := index - 1; j >= 0; j-- {
				if stack[j] == "(" {
					inner := stack[j+1 : index]
					stack = stack[:j]
					for _, k := range eval(inner) {
						stack = append(stack, k)
					}
					res := eval(stack)
					stack = []string{}
					for _, k := range res {
						stack = append(stack, k)
					}
					fmt.Println(strings.Join(stack, " "))
					break
				}
			}
		case "(", "+", "*":
			stack = append(stack, token)
		default:
			len := len(stack) - 1
			if len > 0 {
				switch stack[len] {
				case "+":
					num1, _ := strconv.Atoi(stack[len-1])
					num2, _ := strconv.Atoi(token)
					stack = stack[:len-1]
					stack = append(stack, strconv.Itoa(num1+num2))
					fmt.Println(strings.Join(stack, " "))
				case "*":
					num1, _ := strconv.Atoi(stack[len-1])
					num2, _ := strconv.Atoi(token)
					stack = stack[:len-1]
					stack = append(stack, strconv.Itoa(num1*num2))
					fmt.Println(strings.Join(stack, " "))
				default:
					stack = append(stack, token)
					fmt.Println(strings.Join(stack, " "))
				}
			} else {
				stack = append(stack, token)
				fmt.Println(strings.Join(stack, " "))
			}
		}
	}

	if len(stack)%2 == 1 {
		return eval(stack)
	}

	return stack
}

func eval2(tokens []string) []string {
	// fmt.Printf("Evaluating %s\n", strings.Join(tokens, ""))
	if len(tokens) == 1 {
		return tokens
	}

	evalMul := !contains(tokens, "+") && !contains(tokens, "(")

	stack := make([]string, 0)

	for _, token := range tokens {
		switch token {
		case ")":
			index := len(stack)
			for j := index - 1; j >= 0; j-- {
				if stack[j] == "(" {
					inner := stack[j+1 : index]
					stack = stack[:j]
					for _, k := range eval2(inner) {
						stack = append(stack, k)
					}
					res := eval2(stack)
					stack = []string{}
					for _, k := range res {
						stack = append(stack, k)
					}
					// fmt.Println(strings.Join(stack, " "))
					break
				}
			}
		case "(", "+", "*":
			stack = append(stack, token)
		default:
			len := len(stack) - 1
			if len > 0 {
				switch stack[len] {
				case "+":
					num1, _ := new(big.Int).SetString(stack[len-1], 10)
					num2, _ := new(big.Int).SetString(token, 10)
					num1.Add(num1, num2)
					stack = stack[:len-1]
					stack = append(stack, num1.String())
					// fmt.Println(strings.Join(stack, " "))
				case "*":
					if evalMul {
						num1, _ := new(big.Int).SetString(stack[len-1], 10)
						num2, _ := new(big.Int).SetString(token, 10)
						num1.Mul(num1, num2)
						stack = stack[:len-1]
						stack = append(stack, num1.String())
						// fmt.Println(strings.Join(stack, " "))
					} else {
						stack = append(stack, token)
					}
				default:
					stack = append(stack, token)
					// fmt.Println(strings.Join(stack, " "))
				}
			} else {
				stack = append(stack, token)
				// fmt.Println(strings.Join(stack, " "))
			}
		}
	}

	if len(stack)%2 == 1 {
		return eval2(stack)
	} else {
		return stack
	}
}

func contains(arr []string, target string) bool {
	for _, val := range arr {
		if val == target {
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
