package day8

import (
	"advent-of-code-2022/utils"
	"advent-of-code-2022/utils/grid"
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
	g := grid.NewIntGridFromLines(lines)

	visible := grid.NewEmptyBoolGrid(g.Width, g.Height)
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			row, err := g.Row(i)
			if err != nil {
				panic(err)
			}
			visRow := treeVisible(row, j)

			col, err := g.Column(j)
			if err != nil {
				panic(err)
			}
			visColumn := treeVisible(col, i)

			err = visible.Set(j, i, visRow || visColumn)
			if err != nil {
				panic(err)
			}
		}
	}

	return visible.CountTrue()
}

func solvePart2(lines []string) int {
	g := grid.NewIntGridFromLines(lines)

	score := grid.NewEmptyIntGrid(g.Width, g.Height)
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			row, err := g.Row(i)
			if err != nil {
				panic(err)
			}
			scoreRow := treeScore(row, j)

			col, err := g.Column(j)
			if err != nil {
				panic(err)
			}
			scoreColumn := treeScore(col, i)

			err = score.Set(j, i, scoreRow*scoreColumn)
			if err != nil {
				panic(err)
			}
		}
	}

	return score.Max()
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
