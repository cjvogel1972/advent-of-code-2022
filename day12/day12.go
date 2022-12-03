package day12

import (
	"advent-of-code-2022/utils"
	"fmt"
)

type Puzzle struct{}

// Solve solves day 12's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day12/day12-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	return 0
}

func solvePart2(lines []string) int {
	return 0
}
