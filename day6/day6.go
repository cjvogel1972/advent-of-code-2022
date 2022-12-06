package day6

import (
	"advent-of-code-2022/utils"
	"fmt"
)

type Puzzle struct{}

// Solve solves day 6's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day6/day6-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines[0]))
	fmt.Printf("Part 2: %d\n", solvePart2(lines[0]))
}

func solvePart1(line string) int {
	for i := 0; i < len(line); i++ {
		if line[i] != line[i+1] && line[i] != line[i+2] && line[i] != line[i+3] && line[i+1] != line[i+2] && line[i+1] != line[i+3] && line[i+2] != line[i+3] {
			return i + 4
		}
	}
	return 0
}

func solvePart2(line string) int {
next:
	for i := 0; i < len(line); i++ {
		unique := true
		for j := 0; j < 14; j++ {
			for k := j + 1; k < 14; k++ {
				if line[i+j] == line[i+k] {
					unique = false
					continue next
				}
			}
		}
		if unique {
			return i + 14
		}
	}
	return 0
}
