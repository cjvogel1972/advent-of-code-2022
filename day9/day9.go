package day9

import (
	"advent-of-code-2022/utils"
	"fmt"
	"math"
	"strings"
)

type Puzzle struct{}

// Solve solves day 9's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day9/day9-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	head := &utils.Point{}
	tail := &utils.Point{}
	tailLocs := make(map[utils.Point]int)
	tailLocs[tail.Copy()] = 1

	for _, l := range lines {
		action := strings.Split(l, " ")
		dir := action[0]
		dist := utils.ToInt(action[1])
		if dir == "R" {
			for i := 0; i < dist; i++ {
				head.X++
				fixDiff(head, tail, true, tailLocs)
			}
		} else if dir == "L" {
			for i := 0; i < dist; i++ {
				head.X--
				fixDiff(head, tail, true, tailLocs)
			}
		} else if dir == "U" {
			for i := 0; i < dist; i++ {
				head.Y++
				fixDiff(head, tail, true, tailLocs)
			}
		} else if dir == "D" {
			for i := 0; i < dist; i++ {
				head.Y--
				fixDiff(head, tail, true, tailLocs)
			}
		}
	}

	return len(tailLocs)
}

func fixDiff(knot1 *utils.Point, knot2 *utils.Point, tail bool, tailLocs map[utils.Point]int) {
	yDiff := knot1.Y - knot2.Y
	xDiff := knot1.X - knot2.X
	moved := false
	if yDiff == 0 {
		if math.Abs(float64(xDiff)) == 2 {
			if xDiff > 0 {
				knot2.X++
				moved = true
			} else {
				knot2.X--
				moved = true
			}
		}
	} else if xDiff == 0 {
		if math.Abs(float64(yDiff)) == 2 {
			if yDiff > 0 {
				knot2.Y++
				moved = true
			} else {
				knot2.Y--
				moved = true
			}
		}
	} else if xDiff*yDiff > 1 || xDiff*yDiff < -1 {
		moved = true
		if xDiff > 0 {
			knot2.X++
		} else {
			knot2.X--
		}
		if yDiff > 0 {
			knot2.Y++
		} else {
			knot2.Y--
		}
	}

	if tail && moved {
		key := knot2.Copy()
		cnt, ok := tailLocs[key]
		if !ok {
			cnt = 0
		}
		cnt++
		tailLocs[key] = cnt
	}
}

func solvePart2(lines []string) int {
	knots := make([]*utils.Point, 10)
	for i := 0; i < 10; i++ {
		knots[i] = &utils.Point{}
	}
	tailLocs := make(map[utils.Point]int)
	tailLocs[knots[9].Copy()] = 1

	for _, l := range lines {
		action := strings.Split(l, " ")
		dir := action[0]
		dist := utils.ToInt(action[1])
		if dir == "R" {
			for i := 0; i < dist; i++ {
				knots[0].X++
				for j := 1; j < 10; j++ {
					tail := j == 9
					fixDiff(knots[j-1], knots[j], tail, tailLocs)
				}
			}
		} else if dir == "L" {
			for i := 0; i < dist; i++ {
				knots[0].X--
				for j := 1; j < 10; j++ {
					tail := j == 9
					fixDiff(knots[j-1], knots[j], tail, tailLocs)
				}
			}
		} else if dir == "U" {
			for i := 0; i < dist; i++ {
				knots[0].Y++
				for j := 1; j < 10; j++ {
					tail := j == 9
					fixDiff(knots[j-1], knots[j], tail, tailLocs)
				}
			}
		} else if dir == "D" {
			for i := 0; i < dist; i++ {
				knots[0].Y--
				for j := 1; j < 10; j++ {
					tail := j == 9
					fixDiff(knots[j-1], knots[j], tail, tailLocs)
				}
			}
		}
	}

	return len(tailLocs)
}
