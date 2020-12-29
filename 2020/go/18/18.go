package main

import (
	"fmt"
	"strconv"
	"strings"

	"../utils"
)

func Part1(input []string) (sum int) {
	for _, line := range input {
		tokens := strings.Split(strings.ReplaceAll(line, " ", ""), "")
		result, _ := strconv.Atoi(eval(tokens)[0])
		sum += result
	}

	return
}

func Part2(input []string) (sum int) {
	for _, line := range input {
		tokens := strings.Split(rewritePrecedence(strings.ReplaceAll(line, " ", "")), "")
		result, _ := strconv.Atoi(eval2(tokens)[0])
		sum += result
	}

	return
}

func eval(tokens []string) []string {
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
				case "*":
					num1, _ := strconv.Atoi(stack[len-1])
					num2, _ := strconv.Atoi(token)
					stack = stack[:len-1]
					stack = append(stack, strconv.Itoa(num1*num2))
				default:
					stack = append(stack, token)
				}
			} else {
				stack = append(stack, token)
			}
		}
	}

	if len(stack)%2 == 1 {
		return eval(stack)
	}

	return stack
}

func rewritePrecedence(expr string) (rewritten string) {
	rewritten = "(((("

	for _, ch := range expr {
		switch ch {
		case '(':
			rewritten += "(((("
		case ')':
			rewritten += "))))"
		case '*':
			rewritten += ")))*((("
		case '+':
			rewritten += "))+(("
		default:
			rewritten += string(ch)
		}
	}

	rewritten += "))))"

	return
}

func eval2(tokens []string) []string {
	if len(tokens) == 1 {
		return tokens
	}

	evalMul := !contains(tokens, "+") && !contains(tokens, "(")

	stack := make([]string, 0)

	for _, token := range tokens {
		switch token {
		case ")":
			index := len(stack)
			orig := utils.CopyStringSlice(stack)
			for j := index - 1; j >= 0; j-- {
				if stack[j] == "(" {
					inner := stack[j+1 : index]
					stack = stack[:j]
					for _, k := range eval2(inner) {
						stack = append(stack, k)
					}
					if len(orig) != len(stack) {
						result := eval2(stack)
						stack = []string{}
						for _, k := range result {
							stack = append(stack, k)
						}
					}
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
				case "*":
					if evalMul {
						num1, _ := strconv.Atoi(stack[len-1])
						num2, _ := strconv.Atoi(token)
						stack = stack[:len-1]
						stack = append(stack, strconv.Itoa(num1*num2))
					} else {
						stack = append(stack, token)
					}
				default:
					stack = append(stack, token)
				}
			} else {
				stack = append(stack, token)
			}
		}
	}

	return stack
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
