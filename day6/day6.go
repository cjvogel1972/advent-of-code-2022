package day6

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

type Puzzle struct{}

// Solve solves day 6's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day6/day6-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines[0]))
	fmt.Printf("Part 2: %d\n", solvePart2(lines[0]))
}

func solvePart1(line string) int {
	return computeMarkerLocation(line, 4)
}

func solvePart2(line string) int {
	return computeMarkerLocation(line, 14)
}

func computeMarkerLocation(line string, markerSize int) int {
	for i := 0; i < len(line)-markerSize; i++ {
		startMarker := line[i : i+markerSize]
		unique := true
		for j := 0; j < markerSize; j++ {
			if strings.Count(startMarker, string(startMarker[j])) != 1 {
				unique = false
				break
			}
		}
		if unique {
			return i + markerSize
		}
	}
	panic(-1)
}
