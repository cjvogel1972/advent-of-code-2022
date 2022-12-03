package day3

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
	"unicode"
)

type Puzzle struct{}

// Solve solves day 2's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day3/sacks.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(sacks []string) int {
	totalPriority := 0
	for _, sack := range sacks {
		numItems := len(sack)
		compartment1 := sack[0 : numItems/2]
		compartment2 := sack[(numItems / 2):numItems]
		var dupItem rune
		for _, item := range compartment1 {
			if strings.ContainsRune(compartment2, item) {
				dupItem = item
				break
			}
		}

		totalPriority += computePriority(dupItem)
	}

	return totalPriority
}

func solvePart2(sacks []string) int {
	totalPriority := 0
	for i := 0; i < len(sacks); i += 3 {
		var badge rune
		for _, item := range sacks[i] {
			if strings.ContainsRune(sacks[i+1], item) && strings.ContainsRune(sacks[i+2], item) {
				badge = item
				break
			}
		}

		totalPriority += computePriority(badge)
	}

	return totalPriority
}

func computePriority(item rune) int {
	if unicode.IsLower(item) {
		return int(item - 'a' + 1)
	} else {
		return int(item - 'A' + 27)
	}
}
