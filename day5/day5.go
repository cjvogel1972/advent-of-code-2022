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
	crateConfig, actions := utils.SplitFile(lines)
	stacks := parseCrates(crateConfig)

	for i := 0; i < len(actions); i++ {
		numCratesToMove, fromStack, toStack := parseAction(actions[i])
		for j := 0; j < numCratesToMove; j++ {
			stacks[toStack].Push(stacks[fromStack].Pop())
		}
	}

	return getTopCrates(stacks)
}

func solvePart2(lines []string) string {
	crateConfig, actions := utils.SplitFile(lines)
	stacks := parseCrates(crateConfig)

	for i := 0; i < len(actions); i++ {
		numCratesToMove, fromStack, toStack := parseAction(actions[i])
		tempStack := utils.Stack{}
		for j := 0; j < numCratesToMove; j++ {
			tempStack.Push(stacks[fromStack].Pop())
		}
		for j := 0; j < numCratesToMove; j++ {
			stacks[toStack].Push(tempStack.Pop())
		}
	}

	return getTopCrates(stacks)
}

func parseCrates(lines []string) []utils.Stack {
	numColumns := parseNumColumns(lines[len(lines)-1])

	stacks := make([]utils.Stack, numColumns)
	for i := 0; i < numColumns; i++ {
		stacks[i] = utils.Stack{}
	}

	// read from bottom, above footer, and push crates on stack
	for i := len(lines) - 2; i >= 0; i-- {
		for j := 0; j*4 < len(lines[i]); j++ {
			entry := lines[i][j*4 : j*4+3]
			if strings.ContainsRune(entry, '[') {
				stacks[j].Push(rune(entry[1]))
			}
		}
	}

	return stacks
}

func parseNumColumns(line string) int {
	return utils.ToInt(line[len(line)-1:])
}

func parseAction(line string) (int, int, int) {
	var numCratesToMove, fromStack, toStack int
	_, err := fmt.Sscanf(line, "move %d from %d to %d", &numCratesToMove, &fromStack, &toStack)
	if err != nil {
		panic(err)
	}

	// adjust for slices starting at 0
	fromStack--
	toStack--

	return numCratesToMove, fromStack, toStack
}

func getTopCrates(crates []utils.Stack) string {
	topCrates := ""

	for i := 0; i < len(crates); i++ {
		topCrates = topCrates + string(crates[i].Peek().(rune))
	}

	return topCrates
}
