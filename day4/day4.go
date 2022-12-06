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
		if elf1.completeOverlap(elf2) || elf2.completeOverlap(elf1) {
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
		if elf1.overlap(elf2) || elf2.overlap(elf1) {
			someOverlaps++
		}
	}

	return someOverlaps
}

func newAssignment(sectionRange string) assignment {
	sections := strings.Split(sectionRange, "-")
	return assignment{utils.ToInt(sections[0]), utils.ToInt(sections[1])}
}

type assignment struct {
	start int
	end   int
}

func (a assignment) completeOverlap(other assignment) bool {
	return utils.Within(other.start, a.start, a.end) && utils.Within(other.end, a.start, a.end)
}

func (a assignment) overlap(other assignment) bool {
	return utils.Within(other.start, a.start, a.end) || utils.Within(other.end, a.start, a.end)
}
