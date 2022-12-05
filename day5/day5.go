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
	crates, linesRead := parseCrates(lines)

	for i := linesRead; i < len(lines); i++ {
		numCratesToMove, fromColumn, toColumn := parseAction(lines[i])

		top := make([]string, numCratesToMove)
		for j := 0; j < numCratesToMove; j++ {
			top[j] = crates[fromColumn][numCratesToMove-1-j]
		}

		crates[toColumn] = append(top, crates[toColumn]...)
		crates[fromColumn] = crates[fromColumn][numCratesToMove:]
	}

	return getTopCrates(crates)
}

func solvePart2(lines []string) string {
	crates, linesRead := parseCrates(lines)

	for i := linesRead; i < len(lines); i++ {
		numCratesToMove, fromColumn, toColumn := parseAction(lines[i])

		top := make([]string, numCratesToMove)
		for j := 0; j < numCratesToMove; j++ {
			top[j] = crates[fromColumn][j]
		}

		crates[toColumn] = append(top, crates[toColumn]...)
		crates[fromColumn] = crates[fromColumn][numCratesToMove:]
	}

	return getTopCrates(crates)
}

func parseCrates(lines []string) ([][]string, int) {
	linesRead := 0
	for i := 0; len(lines[i]) > 0; i++ {
		linesRead++
	}
	totalLinesRead := linesRead + 1
	linesRead--
	numColumns := parseNumColumns(lines[linesRead])

	crates := make([][]string, numColumns)
	for i := 0; i < numColumns; i++ {
		crates[i] = make([]string, 0)
	}

	for i := 0; i < linesRead; i++ {
		for j := 0; j*4 < len(lines[i]); j++ {
			crate := lines[i][j*4 : j*4+3]
			if strings.ContainsRune(crate, '[') {
				crates[j] = append(crates[j], string(crate[1]))
			}
		}
	}

	return crates, totalLinesRead
}

func parseNumColumns(line string) int {
	return utils.ToInt(line[len(line)-1:])
}

func parseAction(line string) (int, int, int) {
	var numCratesToMove, fromColumn, toColumn int
	_, err := fmt.Sscanf(line, "move %d from %d to %d", &numCratesToMove, &fromColumn, &toColumn)
	if err != nil {
		panic(err)
	}

	// adjust for slices starting at 0
	fromColumn--
	toColumn--

	return numCratesToMove, fromColumn, toColumn
}

func getTopCrates(crates [][]string) string {
	topCrates := ""

	for i := 0; i < len(crates); i++ {
		topCrates = topCrates + crates[i][0]
	}

	return topCrates
}
