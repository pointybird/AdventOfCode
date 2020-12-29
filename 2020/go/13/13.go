package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"../utils"
)

func Part1(input []string) int {
	target, _ := strconv.Atoi(input[0])
	tokens := strings.Split(input[1], ",")
	buses := make([]int, 0)
	earliest := math.MaxInt32
	route := 0
	var wait int
	for _, token := range tokens {
		if token != "x" {
			bus, _ := strconv.Atoi(token)
			buses = append(buses, bus)

			n := target / bus
			if target%bus == 0 {
				route = bus
				wait = 0
				earliest = target * route
				break
			}

			departure := (n + 1) * bus
			if departure < earliest {
				earliest = departure
				wait = departure - target
				route = bus
			}
		}
	}

	return route * wait
}

func Part2(input []string) int {
	tokens := strings.Split(input[1], ",")

	nums := make([]int, 0)
	remainders := make([]int, 0)

	for i, token := range tokens {
		if token != "x" {
			val, _ := strconv.Atoi(token)
			remainders = append(remainders, i)
			nums = append(nums, val)
		}
	}

	return chineseRemainderTheorem(nums, remainders)
}

// https://shainer.github.io/crypto/math/2017/10/22/chinese-remainder-theorem.html
func chineseRemainderTheorem(n, a []int) int {
	prod := 1
	for _, val := range n {
		prod *= val
	}

	var sum int
	for i, val := range n {
		ai := a[i]

		_, _, si := extendedEuclidean(val, prod/val)
		sum += ai * si * (prod / val)
	}

	return leastPositiveEquivalent(sum, prod)
}

func leastPositiveEquivalent(sum, prod int) int {
	if sum > 0 {
		return sum
	}

	sum *= -1
	return (sum + prod) % prod
}

func extendedEuclidean(x, y int) (gcd, x0, y0 int) {
	x0, x1, y0, y1 := 1, 0, 0, 1

	for y > 0 {
		gcd, x, y = x/y, y, x%y
		x0, x1 = x1, x0-gcd*x1
		y0, y1 = y1, y0-gcd*y1
	}

	return gcd, x0, y0
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
