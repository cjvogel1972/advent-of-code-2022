package day14

import (
	"advent-of-code-2022/utils"
	_range "advent-of-code-2022/utils/range"
	"fmt"
	"strings"
)

type Puzzle struct{}

// Solve solves day 14's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day14/day14-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	rocks := make([]line, 0)
	lowest := 0

	for _, l := range lines {
		r := newRock(l)
		for _, s := range r {
			if s.horizontal {
				lowest = utils.Max(lowest, s.p1.Y)
			}
		}
		rocks = append(rocks, r...)
	}

	pile := make([]utils.Point, 0)
	fallingOff := false

	for !fallingOff {
		sand := utils.Point{500, 0}
		atRest := false
		for !atRest {
			copy := sand.Copy()
			copy.Y++
			if copy.Y > lowest {
				fallingOff = true
				break
			}
			if hitSomething(copy, pile, rocks) {
				copy.X--
				if hitSomething(copy, pile, rocks) {
					copy.X += 2
					if hitSomething(copy, pile, rocks) {
						pile = append(pile, sand)
						atRest = true
					}
				}
			}
			sand.X = copy.X
			sand.Y = copy.Y
		}
	}

	return len(pile)
}

func solvePart2(lines []string) int {
	rocks := make([]line, 0)
	lowest := 0

	for _, l := range lines {
		r := newRock(l)
		for _, s := range r {
			if s.horizontal {
				lowest = utils.Max(lowest, s.p1.Y)
			}
		}
		rocks = append(rocks, r...)
	}

	lowest += 2

	pile := make([]utils.Point, 0)
	plugged := false

	for !plugged {
		sand := utils.Point{500, 0}
		atRest := false
		for !atRest {
			copy := sand.Copy()
			copy.Y++
			if copy.Y == lowest {
				pile = append(pile, sand)
				atRest = true
				break
			}
			if hitSomething(copy, pile, rocks) {
				copy.X--
				if hitSomething(copy, pile, rocks) {
					copy.X += 2
					if hitSomething(copy, pile, rocks) {
						pile = append(pile, sand)
						atRest = true
						if sand.X == 500 && sand.Y == 0 {
							plugged = true
						}
					}
				}
			}
			sand.X = copy.X
			sand.Y = copy.Y
		}
	}

	return len(pile)
}

func hitSomething(p utils.Point, pile []utils.Point, rocks []line) bool {
	if hitPile(p, pile) {
		return true
	}

	return hitRock(p, rocks)
}

func hitPile(p utils.Point, pile []utils.Point) bool {
	for _, s := range pile {
		if s == p {
			return true
		}
	}

	return false
}

func hitRock(p utils.Point, rocks []line) bool {
	for _, rock := range rocks {
		if rock.hits(p) {
			return true
		}
	}

	return false
}

type line struct {
	p1         utils.Point
	p2         utils.Point
	horizontal bool
}

func newRock(l string) []line {
	rocks := make([]line, 0)
	entries := strings.Split(l, " ")
	var prev *utils.Point
	for _, e := range entries {
		if e == "->" {
			continue
		}
		coords := strings.Split(e, ",")
		p := utils.Point{utils.ToInt(coords[0]), utils.ToInt(coords[1])}
		if prev == nil {
			prev = &p
		} else {
			horizontal := false
			if prev.Y-p.Y == 0 {
				horizontal = true
			}
			rocks = append(rocks, line{*prev, p, horizontal})
			prev = &p
		}
	}

	return rocks
}

func (l line) hits(p utils.Point) bool {
	if l.horizontal {
		if p.Y == l.p1.Y {
			x := _range.New(l.p1.X, l.p2.X)
			if x.Contains(p.X) {
				return true
			}
		}
	} else {
		if p.X == l.p1.X {
			y := _range.New(l.p1.Y, l.p2.Y)
			if y.Contains(p.Y) {
				return true
			}
		}
	}

	return false
}
