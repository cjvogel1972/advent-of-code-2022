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
		movingCrates := crates[fromColumn].popCratesReversed(numCratesToMove)
		crates[toColumn].pushCrates(movingCrates)
	}

	return getTopCrates(crates)
}

func solvePart2(lines []string) string {
	crates, linesRead := parseCrates(lines)

	for i := linesRead; i < len(lines); i++ {
		numCratesToMove, fromColumn, toColumn := parseAction(lines[i])
		movingCrates := crates[fromColumn].popCratesInOrder(numCratesToMove)
		crates[toColumn].pushCrates(movingCrates)
	}

	return getTopCrates(crates)
}

func newCrateStack() crateStack {
	crates := make([]string, 0)
	return crateStack{crates}
}

type crateStack struct {
	crates []string
}

func (c *crateStack) pushCrate(crate string) {
	c.crates = append([]string{crate}, c.crates...)
}

func (c *crateStack) pushCrates(crates []string) {
	c.crates = append(crates, c.crates...)
}

func (c *crateStack) popCratesInOrder(numCratesToMove int) []string {
	top := make([]string, numCratesToMove)
	for j := 0; j < numCratesToMove; j++ {
		top[j] = c.crates[j]
	}

	c.crates = c.crates[numCratesToMove:]

	return top
}

func (c *crateStack) popCratesReversed(numCratesToMove int) []string {
	top := make([]string, numCratesToMove)
	for j := 0; j < numCratesToMove; j++ {
		top[j] = c.crates[numCratesToMove-1-j]
	}

	c.crates = c.crates[numCratesToMove:]

	return top
}

func (c *crateStack) topCrate() string {
	return c.crates[0]
}

func (c *crateStack) size() int {
	return len(c.crates)
}

func parseCrates(lines []string) ([]crateStack, int) {
	linesRead := 0
	for i := 0; len(lines[i]) > 0; i++ {
		linesRead++
	}
	totalLinesRead := linesRead + 1
	linesRead--
	numColumns := parseNumColumns(lines[linesRead])

	crates := make([]crateStack, numColumns)
	for i := 0; i < numColumns; i++ {
		crates[i] = newCrateStack()
	}

	// read from bottom and push crates on crates
	for i := linesRead - 1; i >= 0; i-- {
		for j := 0; j*4 < len(lines[i]); j++ {
			crate := lines[i][j*4 : j*4+3]
			if strings.ContainsRune(crate, '[') {
				crates[j].pushCrate(string(crate[1]))
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

func getTopCrates(crates []crateStack) string {
	topCrates := ""

	for i := 0; i < len(crates); i++ {
		topCrates = topCrates + crates[i].topCrate()
	}

	return topCrates
}
