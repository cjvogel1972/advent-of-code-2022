package day2

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

type Puzzle struct{}

// Solve solves day 2's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day2/day2-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	score := 0
	for _, line := range lines {
		play := strings.Split(line, " ")
		if play[1] == "X" {
			score += 1
		} else if play[1] == "Y" {
			score += 2
		} else if play[1] == "Z" {
			score += 3
		}

		if play[0] == "A" {
			if play[1] == "Y" {
				score += 6
			} else if play[1] == "X" {
				score += 3
			}
		} else if play[0] == "B" {
			if play[1] == "Z" {
				score += 6
			} else if play[1] == "Y" {
				score += 3
			}
		} else if play[0] == "C" {
			if play[1] == "X" {
				score += 6
			} else if play[1] == "Z" {
				score += 3
			}
		}
	}

	return score
}

func solvePart2(lines []string) int {
	score := 0
	for _, line := range lines {
		play := strings.Split(line, " ")

		if play[1] == "Y" {
			score += 3
		} else if play[1] == "Z" {
			score += 6
		}

		if play[0] == "A" {
			if play[1] == "X" {
				score += 3
			} else if play[1] == "Y" {
				score += 1
			} else if play[1] == "Z" {
				score += 2
			}
		} else if play[0] == "B" {
			if play[1] == "X" {
				score += 1
			} else if play[1] == "Y" {
				score += 2
			} else if play[1] == "Z" {
				score += 3
			}
		} else if play[0] == "C" {
			if play[1] == "X" {
				score += 1
			} else if play[1] == "Y" {
				score += 3
			} else if play[1] == "Z" {
				score += 2
			}
		}
	}

	return score
}
