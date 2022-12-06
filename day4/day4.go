package day4

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

type Puzzle struct{}

// Solve solves day 4's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day4/day4-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	totalOverlaps := 0
	for _, pair := range lines {
		assignments := strings.Split(pair, ",")
		elf1 := newAssignment(assignments[0])
		elf2 := newAssignment(assignments[1])
		if elf1.Within(elf2) || elf2.Within(elf1) {
			totalOverlaps++
		}
	}

	return totalOverlaps
}

func solvePart2(lines []string) int {
	someOverlaps := 0
	for _, pair := range lines {
		assignments := strings.Split(pair, ",")
		elf1 := newAssignment(assignments[0])
		elf2 := newAssignment(assignments[1])
		if elf1.Overlap(elf2) || elf2.Overlap(elf1) {
			someOverlaps++
		}
	}

	return someOverlaps
}

func newAssignment(sectionRange string) utils.Range {
	sections := strings.Split(sectionRange, "-")
	return utils.NewRange(utils.ToInt(sections[0]), utils.ToInt(sections[1]))
}
