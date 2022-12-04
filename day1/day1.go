package day1

import (
	"advent-of-code-2022/utils"
	"fmt"
)

type Puzzle struct{}

// Solve solves day 1's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day1/calories.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	maxCalories := 0
	totalCalories := 0
	for _, line := range lines {
		if len(line) == 0 {
			if totalCalories > maxCalories {
				maxCalories = totalCalories
			}
			totalCalories = 0
		} else {
			totalCalories += utils.ToInt(line)
		}
	}

	return maxCalories
}

func solvePart2(lines []string) int {
	topThree := []int{0, 0, 0}
	totalCalories := 0
	for _, line := range lines {
		if len(line) == 0 {
			compareTopThree(totalCalories, topThree)
			totalCalories = 0
		} else {
			totalCalories += utils.ToInt(line)
		}
	}
	compareTopThree(totalCalories, topThree)

	return topThree[0] + topThree[1] + topThree[2]
}

func compareTopThree(totalCalories int, topThree []int) {
	if totalCalories > topThree[0] {
		topThree[2] = topThree[1]
		topThree[1] = topThree[0]
		topThree[0] = totalCalories
	} else if totalCalories > topThree[1] {
		topThree[2] = topThree[1]
		topThree[1] = totalCalories
	} else if totalCalories > topThree[2] {
		topThree[2] = totalCalories
	}
}
