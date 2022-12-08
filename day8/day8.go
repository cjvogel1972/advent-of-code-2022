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
	width := len(lines[0])
	height := len(lines)

	trees := make([][]int, height)
	visible := make([][]bool, height)
	for i := 0; i < height; i++ {
		row := make([]int, width)
		visibleRow := make([]bool, width)
		for j := 0; j < width; j++ {
			row[j] = utils.ToInt(string(lines[i][j]))
			visibleRow[j] = false
		}
		trees[i] = row
		visible[i] = visibleRow
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				visible[i][j] = true
				continue
			}
			visLeft := true
			visRight := true
			visTop := true
			visBottom := true
			for k := j - 1; k >= 0; k-- {
				if trees[i][j] <= trees[i][k] {
					visLeft = false
					break
				}
			}
			for k := j + 1; k < width; k++ {
				if trees[i][j] <= trees[i][k] {
					visRight = false
					break
				}
			}
			for k := i - 1; k >= 0; k-- {
				if trees[i][j] <= trees[k][j] {
					visTop = false
					break
				}
			}
			for k := i + 1; k < height; k++ {
				if trees[i][j] <= trees[k][j] {
					visBottom = false
					break
				}
			}

			visible[i][j] = visLeft || visRight || visTop || visBottom
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
	width := len(lines[0])
	height := len(lines)

	trees := make([][]int, height)
	scenicScore := make([][]int, height)
	for i := 0; i < height; i++ {
		row := make([]int, width)
		scoreRow := make([]int, width)
		for j := 0; j < width; j++ {
			row[j] = utils.ToInt(string(lines[i][j]))
			scoreRow[j] = 0
		}
		trees[i] = row
		scenicScore[i] = scoreRow
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			scoreLeft := 0
			scoreRight := 0
			scoreTop := 0
			scoreBottom := 0
			for k := j - 1; k >= 0; k-- {
				scoreLeft++
				if trees[i][j] <= trees[i][k] {
					break
				}
			}
			for k := j + 1; k < width; k++ {
				scoreRight++
				if trees[i][j] <= trees[i][k] {
					break
				}
			}
			for k := i - 1; k >= 0; k-- {
				scoreTop++
				if trees[i][j] <= trees[k][j] {
					break
				}
			}
			for k := i + 1; k < height; k++ {
				scoreBottom++
				if trees[i][j] <= trees[k][j] {
					break
				}
			}

			scenicScore[i][j] = scoreLeft * scoreRight * scoreTop * scoreBottom
		}
	}

	maxScore := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			maxScore = utils.Max(maxScore, scenicScore[i][j])
		}
	}

	return maxScore
}
