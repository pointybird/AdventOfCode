package main

import (
	"testing"

	"../utils"
)

func TestPart1(t *testing.T) {
	utils.TestInput(t, DecodeBoardingPass, map[string]int{
		"FBFBBFFRLR": 357,
		"BFFFBBFRRR": 567,
		"FFFBBBFRRR": 119,
		"BBFFBBFRLL": 820})
}
