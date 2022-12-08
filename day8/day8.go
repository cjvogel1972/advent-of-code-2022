package day8

import (
	"advent-of-code-2022/utils"
	"fmt"
)

type Puzzle struct{}

// Solve solves day 8's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day8/day8-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	trees, width, height := parseTrees(lines)

	visible := make([][]bool, height)
	for i := 0; i < height; i++ {
		visible[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			visRow := treeVisible(trees[i], j)
			visColumn := treeVisible(getColumn(trees, j), i)
			visible[i][j] = visRow || visColumn
		}
	}

	visibleCount := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if visible[i][j] {
				visibleCount++
			}
		}
	}

	return visibleCount
}

func solvePart2(lines []string) int {
	trees, width, height := parseTrees(lines)

	score := make([][]int, height)
	for i := 0; i < height; i++ {
		score[i] = make([]int, width)
		for j := 0; j < width; j++ {
			scoreRow := treeScore(trees[i], j)
			scoreColumn := treeScore(getColumn(trees, j), i)
			score[i][j] = scoreRow * scoreColumn
		}
	}

	maxScore := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			maxScore = utils.Max(maxScore, score[i][j])
		}
	}

	return maxScore
}

func parseTrees(lines []string) ([][]int, int, int) {
	width := len(lines[0])
	height := len(lines)

	trees := make([][]int, height)
	for i := 0; i < height; i++ {
		row := make([]int, width)
		for j := 0; j < width; j++ {
			row[j] = utils.ToInt(string(lines[i][j]))
		}
		trees[i] = row
	}

	return trees, width, height
}

func getColumn(trees [][]int, columnNo int) []int {
	column := make([]int, len(trees))
	for i := 0; i < len(trees); i++ {
		column[i] = trees[i][columnNo]
	}

	return column
}

func treeVisible(trees []int, treeIndex int) bool {
	visBefore := true
	for i := treeIndex - 1; i >= 0; i-- {
		if trees[treeIndex] <= trees[i] {
			visBefore = false
			break
		}
	}

	visAfter := true
	for i := treeIndex + 1; i < len(trees); i++ {
		if trees[treeIndex] <= trees[i] {
			visAfter = false
			break
		}
	}

	return visBefore || visAfter
}

func treeScore(trees []int, treeIndex int) int {
	scoreBefore := 0
	for i := treeIndex - 1; i >= 0; i-- {
		scoreBefore++
		if trees[treeIndex] <= trees[i] {
			break
		}
	}

	scoreAfter := 0
	for i := treeIndex + 1; i < len(trees); i++ {
		scoreAfter++
		if trees[treeIndex] <= trees[i] {
			break
		}
	}

	return scoreBefore * scoreAfter
}
