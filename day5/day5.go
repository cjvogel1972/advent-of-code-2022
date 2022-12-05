package day5

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

type Puzzle struct{}

// Solve solves day 5's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day5/day5-input.txt")

	fmt.Printf("Part 1: %s\n", solvePart1(lines))
	fmt.Printf("Part 2: %s\n", solvePart2(lines))
}

func solvePart1(lines []string) string {
	startCrateStackSize := 0
	for i := 0; len(lines[i]) > 0; i++ {
		startCrateStackSize++
	}
	startCrateStackSize--
	numColumns := utils.ToInt(lines[startCrateStackSize][len(lines[startCrateStackSize])-1:])
	crates := make([][]string, numColumns)
	for i := 0; i < numColumns; i++ {
		crates[i] = make([]string, 0)
	}
	for i := 0; i < startCrateStackSize; i++ {
		for j := 0; j*4 < len(lines[i]); j++ {
			crate := lines[i][j*4 : j*4+3]
			if strings.ContainsRune(crate, '[') {
				crates[j] = append(crates[j], string(crate[1]))
			}
		}
	}

	for i := startCrateStackSize + 2; i < len(lines); i++ {
		var numCrates, column1, column2 int
		_, err := fmt.Sscanf(lines[i], "move %d from %d to %d", &numCrates, &column1, &column2)
		if err != nil {
			panic(err)
		}
		column1--
		column2--
		top := make([]string, numCrates)
		for j := 0; j < numCrates; j++ {
			top[j] = crates[column1][numCrates-1-j]
		}
		crates[column2] = append(top, crates[column2]...)
		crates[column1] = crates[column1][numCrates:]
	}
	topCrates := ""
	for i := 0; i < len(crates); i++ {
		topCrates = topCrates + crates[i][0]
	}
	return topCrates
}

func solvePart2(lines []string) string {
	startCrateStackSize := 0
	for i := 0; len(lines[i]) > 0; i++ {
		startCrateStackSize++
	}
	startCrateStackSize--
	numColumns := utils.ToInt(lines[startCrateStackSize][len(lines[startCrateStackSize])-1:])
	crates := make([][]string, numColumns)
	for i := 0; i < numColumns; i++ {
		crates[i] = make([]string, 0)
	}
	for i := 0; i < startCrateStackSize; i++ {
		for j := 0; j*4 < len(lines[i]); j++ {
			crate := lines[i][j*4 : j*4+3]
			if strings.ContainsRune(crate, '[') {
				crates[j] = append(crates[j], string(crate[1]))
			}
		}
	}

	for i := startCrateStackSize + 2; i < len(lines); i++ {
		var numCrates, column1, column2 int
		_, err := fmt.Sscanf(lines[i], "move %d from %d to %d", &numCrates, &column1, &column2)
		if err != nil {
			panic(err)
		}
		column1--
		column2--
		top := make([]string, numCrates)
		for j := 0; j < numCrates; j++ {
			top[j] = crates[column1][j]
		}
		crates[column2] = append(top, crates[column2]...)
		crates[column1] = crates[column1][numCrates:]
	}
	topCrates := ""
	for i := 0; i < len(crates); i++ {
		topCrates = topCrates + crates[i][0]
	}
	return topCrates
}
