package main

import (
	"fmt"

	"../utils"
)

var target int = 2020

func Part1(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return nums[i] * nums[j]
			}
		}
	}

	return 0
}

func Part2(nums []int) int {
	length := len(nums)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == target {
					return nums[i] * nums[j] * nums[k]
				}
			}
		}
	}

	return 0
}

func main() {
	input := utils.ReadInts(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
