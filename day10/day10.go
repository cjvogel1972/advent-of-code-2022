package day10

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

type Puzzle struct{}

// Solve solves day 10's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day10/day10-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	result := solvePart2(lines)
	for _, l := range result {
		fmt.Printf("Part 2: %s\n", l)
	}
}

func solvePart1(lines []string) int {
	x := 1
	cycleNo := 0
	sumSignalStrength := 0

	for _, instruction := range lines {
		if instruction == "noop" {
			cycleNo++
			if cycleNo == 20 || (cycleNo-20)%40 == 0 {
				sumSignalStrength += x * cycleNo
			}
		} else {
			val := utils.ToInt(strings.Split(instruction, " ")[1])
			cycleNo++
			if cycleNo == 20 || (cycleNo-20)%40 == 0 {
				sumSignalStrength += x * cycleNo
			}
			cycleNo++
			if cycleNo == 20 || (cycleNo-20)%40 == 0 {
				sumSignalStrength += x * cycleNo
			}
			x += val
		}
	}

	return sumSignalStrength
}

func solvePart2(lines []string) []string {
	x := 1
	cycleNo := 0
	result := make([]string, 0)
	crtLine := make([]rune, 40)

	for _, instruction := range lines {
		if instruction == "noop" {
			if x-1 <= cycleNo%40 && cycleNo%40 <= x+1 {
				crtLine[cycleNo%40] = '#'
			} else {
				crtLine[cycleNo%40] = '.'
			}
			cycleNo++
			if cycleNo%40 == 0 {
				result = append(result, string(crtLine))
				crtLine = make([]rune, 40)
			}
		} else {
			val := utils.ToInt(strings.Split(instruction, " ")[1])
			if x-1 <= cycleNo%40 && cycleNo%40 <= x+1 {
				crtLine[cycleNo%40] = '#'
			} else {
				crtLine[cycleNo%40] = '.'
			}
			cycleNo++
			if cycleNo%40 == 0 {
				result = append(result, string(crtLine))
				crtLine = make([]rune, 40)
			}
			if x-1 <= cycleNo%40 && cycleNo%40 <= x+1 {
				crtLine[cycleNo%40] = '#'
			} else {
				crtLine[cycleNo%40] = '.'
			}
			cycleNo++
			if cycleNo%40 == 0 {
				result = append(result, string(crtLine))
				crtLine = make([]rune, 40)
			}
			x += val
		}
	}

	return result
}
