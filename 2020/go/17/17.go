package main

import (
	"fmt"

	"../utils"
)

const maxW int = 20
const maxX int = 20
const maxY int = 20
const maxZ int = 20

const numCycles int = 6

const inactive rune = '.'
const active rune = '#'

var dimension3d map[[3]int]rune = make(map[[3]int]rune)
var dimension4d map[[4]int]rune = make(map[[4]int]rune)

func Part1(input []string) (count int) {
	for r, line := range input {
		for c, ch := range line {
			loc := [3]int{0, r, c}
			dimension3d[loc] = ch
		}
	}

	for i := 0; i < numCycles; i++ {
		count = step3d()
	}

	return
}

func Part2(input []string) (count int) {
	for r, line := range input {
		for c, ch := range line {
			loc := [4]int{0, 0, r, c}
			dimension4d[loc] = ch
		}
	}

	for i := 0; i < numCycles; i++ {
		count = step4d()
	}

	return
}

func neighbors3d(x, y, z int) (count int) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				loc := [3]int{z + i, y + j, x + k}
				if val, ok := dimension3d[loc]; ok {
					if val == active {
						count++
					}
				}
			}
		}
	}

	return
}

func neighbors4d(w, x, y, z int) (count int) {
	for h := -1; h < 2; h++ {
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				for k := -1; k < 2; k++ {
					if h == 0 && i == 0 && j == 0 && k == 0 {
						continue
					}
					loc := [4]int{w + h, z + i, y + j, x + k}
					if val, ok := dimension4d[loc]; ok {
						if val == active {
							count++
						}
					}
				}
			}
		}
	}

	return
}

func step3d() (count int) {
	temp := make(map[[3]int]rune)

	for z := -maxZ; z < maxZ; z++ {
		for y := -maxY; y < maxY; y++ {
			for x := -maxX; x < maxX; x++ {
				neighbors := neighbors3d(x, y, z)
				cell := inactive
				loc := [3]int{z, y, x}
				if val, ok := dimension3d[loc]; ok {
					cell = val
				}
				if cell == active {
					if !(neighbors == 2 || neighbors == 3) {
						cell = inactive
					}
				} else {
					if neighbors == 3 {
						cell = active
					}
				}
				temp[loc] = cell
				if cell == active {
					count++
				}
			}
		}
	}

	dimension3d = make(map[[3]int]rune)
	for k, v := range temp {
		dimension3d[k] = v
	}

	return
}

func step4d() (count int) {
	temp := make(map[[4]int]rune)

	for w := -maxW; w < maxW; w++ {
		for z := -maxZ; z < maxZ; z++ {
			for y := -maxY; y < maxY; y++ {
				for x := -maxX; x < maxX; x++ {
					neighbors := neighbors4d(w, x, y, z)
					cell := inactive
					loc := [4]int{w, z, y, x}
					if val, ok := dimension4d[loc]; ok {
						cell = val
					}
					if cell == active {
						if !(neighbors == 2 || neighbors == 3) {
							cell = inactive
						}
					} else {
						if neighbors == 3 {
							cell = active
						}
					}
					temp[loc] = cell
					if cell == active {
						count++
					}
				}
			}
		}
	}

	dimension4d = make(map[[4]int]rune)
	for k, v := range temp {
		dimension4d[k] = v
	}

	return
}

func main() {
	input := utils.ReadLines(utils.InputFile)

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
