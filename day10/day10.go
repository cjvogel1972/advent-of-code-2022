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
	totalSignalStrength := 0

	for _, instruction := range lines {
		if instruction == "noop" {
			cycleNo, totalSignalStrength = part1Cycle(cycleNo, x, totalSignalStrength)
		} else {
			val := utils.ToInt(strings.Split(instruction, " ")[1])
			cycleNo, totalSignalStrength = part1Cycle(cycleNo, x, totalSignalStrength)
			cycleNo, totalSignalStrength = part1Cycle(cycleNo, x, totalSignalStrength)
			x += val
		}
	}

	return totalSignalStrength
}

func part1Cycle(cycleNo int, x int, totalSignalStrength int) (int, int) {
	cycleNo++
	if cycleNo == 20 || (cycleNo-20)%40 == 0 {
		totalSignalStrength += x * cycleNo
	}

	return cycleNo, totalSignalStrength
}

func solvePart2(lines []string) []string {
	x := 1
	cycleNo := 0
	crtScreen := make([]string, 0)
	crtLine := make([]rune, 40)
	initCrtLine(crtLine)

	for _, instruction := range lines {
		if instruction == "noop" {
			cycleNo, crtScreen = part2Cycle(x, cycleNo, crtLine, crtScreen)
		} else {
			val := utils.ToInt(strings.Split(instruction, " ")[1])
			cycleNo, crtScreen = part2Cycle(x, cycleNo, crtLine, crtScreen)
			cycleNo, crtScreen = part2Cycle(x, cycleNo, crtLine, crtScreen)
			x += val
		}
	}

	return crtScreen
}

func part2Cycle(x int, cycleNo int, crtLine []rune, crtScreen []string) (int, []string) {
	drawPixel(x, cycleNo, crtLine)
	cycleNo++
	crtScreen = checkIfAddCrtLine(cycleNo, crtScreen, crtLine)

	return cycleNo, crtScreen
}

func initCrtLine(crtLine []rune) {
	for i := 0; i < len(crtLine); i++ {
		crtLine[i] = '.'
	}
}

func drawPixel(x int, cycleNo int, crtLine []rune) {
	if x-1 <= cycleNo%40 && cycleNo%40 <= x+1 {
		crtLine[cycleNo%40] = '#'
	}
}

func checkIfAddCrtLine(cycleNo int, crtScreen []string, crtLine []rune) []string {
	if cycleNo%40 == 0 {
		crtScreen = append(crtScreen, string(crtLine))
		initCrtLine(crtLine)
	}

	return crtScreen
}
