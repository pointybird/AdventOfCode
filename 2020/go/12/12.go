package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"../utils"
)

var x, y, heading int
var orderRegex *regexp.Regexp = regexp.MustCompile(`([NSEWLRF])(\d+)`)

var waypoint [2]int = [2]int{10, 1}

func Part1(input []string) int {
	return elapsedDistance(input, movePosition)
}

func Part2(input []string) int {
	return elapsedDistance(input, moveWaypoint)
}

func elapsedDistance(input []string, movementFunc func(string)) int {
	x, y = 0, 0
	for _, order := range input {
		movementFunc(order)
	}

	return manhattanDistance(x, y)
}

func movePosition(order string) {
	tokens := orderRegex.FindStringSubmatch(order)
	if tokens == nil {
		return
	}
	units, _ := strconv.Atoi(tokens[2])
	switch tokens[1] {
	case "N":
		y += units
	case "S":
		y -= units
	case "E":
		x += units
	case "W":
		x -= units
	case "L":
		heading += units
		heading %= 360
	case "R":
		heading -= units
		heading = (heading + 360) % 360
	case "F":
		switch heading {
		case 0:
			x += units
		case 90:
			y += units
		case 180:
			x -= units
		case 270:
			y -= units
		}
	}
}

func moveWaypoint(order string) {
	tokens := orderRegex.FindStringSubmatch(order)
	if tokens == nil {
		return
	}
	units, _ := strconv.Atoi(tokens[2])
	switch tokens[1] {
	case "N":
		waypoint[1] += units
	case "S":
		waypoint[1] -= units
	case "E":
		waypoint[0] += units
	case "W":
		waypoint[0] -= units
	case "L":
		updateWaypoint(units)
	case "R":
		updateWaypoint(-units)
	case "F":
		x += units * waypoint[0]
		y += units * waypoint[1]
	}
}

func updateWaypoint(angle int) {
	angle = (360 + angle) % 360
	switch angle {
	case 0:
		break
	case 90:
		waypoint[0], waypoint[1] = -waypoint[1], waypoint[0]
	case 180:
		waypoint[0], waypoint[1] = -waypoint[0], -waypoint[1]
	case 270:
		waypoint[0], waypoint[1] = waypoint[1], -waypoint[0]
	}
}

func manhattanDistance(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
