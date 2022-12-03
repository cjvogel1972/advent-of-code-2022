package day2

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

type Puzzle struct{}

const winScore = 6
const tieScore = 3
const loseScore = 0

var shapes = map[string]play{
	"rock":     {"rock", "scissors", "paper", 1},
	"paper":    {"paper", "rock", "scissors", 2},
	"scissors": {"scissors", "paper", "rock", 3},
}

var elfShapeMap = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}

// Solve solves day 2's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day2/day2-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	var myShapeMap = map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	score := 0
	for _, line := range lines {
		play := strings.Split(line, " ")
		score += shapes[myShapeMap[play[1]]].score
		score += shapes[myShapeMap[play[1]]].play(elfShapeMap[play[0]])
	}

	return score
}

func solvePart2(lines []string) int {
	score := 0
	for _, line := range lines {
		entries := strings.Split(line, " ")

		elfShape := shapes[elfShapeMap[entries[0]]]
		switch entries[1] {
		case "X":
			score += loseScore
			score += shapes[elfShape.beats].score
		case "Y":
			score += tieScore
			score += shapes[elfShape.shape].score
		case "Z":
			score += winScore
			score += shapes[elfShape.losesTo].score
		}
	}

	return score
}

type play struct {
	shape   string
	beats   string
	losesTo string
	score   int
}

func (p play) play(otherShape string) int {
	if p.beats == otherShape {
		return winScore
	} else if p.shape == otherShape {
		return tieScore
	} else {
		return loseScore
	}
}
