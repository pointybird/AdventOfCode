package main

import (
	"aoc2020/utils"
	"fmt"
)

const empty = byte('L')
const occupied = byte('#')

var layout, tempLayout [][]byte

func Part1(input []string) int {
	return occupiedSeats(input, 4, adjacentNeighbors)
}

func Part2(input []string) int {
	return occupiedSeats(input, 5, visibleNeighbors)
}

func occupiedSeats(input []string, tolerance int, neighborsFunc func(int, int) int) int {
	initLayouts(input)

	var previous int

	occupied := updateLayout(tolerance, neighborsFunc)
	for previous != occupied {
		previous = occupied
		occupied = updateLayout(tolerance, neighborsFunc)
	}

	return occupied
}

func adjacentNeighbors(row, col int) (count int) {
	for r := row - 1; r < row+2; r++ {
		for c := col - 1; c < col+2; c++ {
			if r == row && c == col {
				continue
			}
			if validLocation(r, c) && layout[r][c] == occupied {
				count++
			}
		}
	}

	return
}

func visibleNeighbors(row, col int) (count int) {
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, dir := range dirs {
		r, c := row, col
		next := false
		for !next {
			r += dir[0]
			c += dir[1]

			if validLocation(r, c) {
				switch layout[r][c] {
				case occupied:
					count++
					next = true
					break
				case empty:
					next = true
					break
				}
			} else {
				next = true
				break
			}
		}
	}

	return
}

func validLocation(row, col int) bool {
	return row >= 0 && row < len(layout) && col >= 0 && col < len(layout[row])
}

func updateLayout(tolerance int, neighborsFunc func(int, int) int) (numOccupied int) {
	for row := 0; row < len(layout); row++ {
		for col := 0; col < len(layout[row]); col++ {
			seat := layout[row][col]
			count := neighborsFunc(row, col)
			if seat == empty && count == 0 {
				tempLayout[row][col] = occupied
			} else if seat == occupied && count >= tolerance {
				tempLayout[row][col] = empty
			} else {
				tempLayout[row][col] = seat
			}

			if tempLayout[row][col] == occupied {
				numOccupied++
			}
		}
	}

	copyLayout()

	return
}

func initLayouts(input []string) {
	layout = make([][]byte, len(input))
	tempLayout = make([][]byte, len(input))
	for i, line := range input {
		layout[i] = make([]byte, len(line))
		copy(layout[i], []byte(input[i]))
		tempLayout[i] = make([]byte, len(line))
	}
}

func copyLayout() {
	for r := 0; r < len(layout); r++ {
		for c := 0; c < len(layout[r]); c++ {
			layout[r][c] = tempLayout[r][c]
		}
	}
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
