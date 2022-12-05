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
	stacks, actions := utils.SplitFile(lines)
	crates := parseCrates(stacks)

	for i := 0; i < len(actions); i++ {
		numCratesToMove, fromColumn, toColumn := parseAction(actions[i])
		movingCrates := crates[fromColumn].popCratesReversed(numCratesToMove)
		crates[toColumn].pushCrates(movingCrates)
	}

	return getTopCrates(crates)
}

func solvePart2(lines []string) string {
	stacks, actions := utils.SplitFile(lines)
	crates := parseCrates(stacks)

	for i := 0; i < len(actions); i++ {
		numCratesToMove, fromColumn, toColumn := parseAction(actions[i])
		movingCrates := crates[fromColumn].popCratesInOrder(numCratesToMove)
		crates[toColumn].pushCrates(movingCrates)
	}

	return getTopCrates(crates)
}

type crate rune

func newCrateStack() crateStack {
	crates := make([]crate, 0)
	return crateStack{crates}
}

type crateStack struct {
	crates []crate
}

func (s *crateStack) pushCrate(c crate) {
	s.crates = append([]crate{c}, s.crates...)
}

func (s *crateStack) pushCrates(crates []crate) {
	s.crates = append(crates, s.crates...)
}

func (s *crateStack) popCratesInOrder(numCratesToMove int) []crate {
	top := make([]crate, numCratesToMove)
	for j := 0; j < numCratesToMove; j++ {
		top[j] = s.crates[j]
	}

	s.crates = s.crates[numCratesToMove:]

	return top
}

func (s *crateStack) popCratesReversed(numCratesToMove int) []crate {
	top := make([]crate, numCratesToMove)
	for j := 0; j < numCratesToMove; j++ {
		top[j] = s.crates[numCratesToMove-1-j]
	}

	s.crates = s.crates[numCratesToMove:]

	return top
}

func (s *crateStack) topCrate() crate {
	return s.crates[0]
}

func (s *crateStack) size() int {
	return len(s.crates)
}

func parseCrates(lines []string) []crateStack {
	numColumns := parseNumColumns(lines[len(lines)-1])

	crates := make([]crateStack, numColumns)
	for i := 0; i < numColumns; i++ {
		crates[i] = newCrateStack()
	}

	// read from bottom, above footer, and push crates on stack
	for i := len(lines) - 2; i >= 0; i-- {
		for j := 0; j*4 < len(lines[i]); j++ {
			entry := lines[i][j*4 : j*4+3]
			if strings.ContainsRune(entry, '[') {
				crates[j].pushCrate(crate(entry[1]))
			}
		}
	}

	return crates
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
		topCrates = topCrates + string(crates[i].topCrate())
	}

	return topCrates
}
