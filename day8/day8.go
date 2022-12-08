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
	g.Iterate(func(x int, y int) {
		row, err := g.Row(x)
		if err != nil {
			panic(err)
		}
		v := treeVisible(row, y)

		if !v {
			col, err := g.Column(y)
			if err != nil {
				panic(err)
			}
			v = treeVisible(col, x)
		}

		err = visible.Set(x, y, v)
		if err != nil {
			panic(err)
		}

	})

	return visible.CountTrue()
}

func solvePart2(lines []string) int {
	g := grid.NewIntGridFromLines(lines)

	score := grid.NewEmptyIntGrid(g.Width, g.Height)
	g.Iterate(func(x int, y int) {
		row, err := g.Row(x)
		if err != nil {
			panic(err)
		}
		s := treeScore(row, y)

		if s != 0 {
			col, err := g.Column(y)
			if err != nil {
				panic(err)
			}
			s *= treeScore(col, x)
		}

		err = score.Set(x, y, s)
		if err != nil {
			panic(err)
		}
	})

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
	if visBefore {
		return true
	}

	visAfter := true
	for i := treeIndex + 1; i < len(trees); i++ {
		if trees[treeIndex] <= trees[i] {
			visAfter = false
			break
		}
	}

	return visAfter
}

func treeScore(trees []int, treeIndex int) int {
	scoreBefore := 0
	for i := treeIndex - 1; i >= 0; i-- {
		scoreBefore++
		if trees[treeIndex] <= trees[i] {
			break
		}
	}
	if scoreBefore == 0 {
		return 0
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
