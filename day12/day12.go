package day12

import (
	"advent-of-code-2022/utils"
	"fmt"
	"math"
)

type Puzzle struct{}

// Solve solves day 12's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day12/day12-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	var start = utils.Point{}
	var end = utils.Point{}
	g := make(map[utils.Point]rune)
	for y, line := range lines {
		for x, r := range line {
			if r == 'S' {
				start.X = x
				start.Y = y
				r = 'a'
			}
			if r == 'E' {
				end.X = x
				end.Y = y
				r = 'z'
			}
			g[utils.Point{x, y}] = r
		}
	}

	moves := dijkstra(g, start, end)

	return moves
}

func solvePart2(lines []string) int {
	var end = utils.Point{}
	g := make(map[utils.Point]rune)
	for y, line := range lines {
		for x, r := range line {
			if r == 'S' {
				r = 'a'
			}
			if r == 'E' {
				end.X = x
				end.Y = y
				r = 'z'
			}
			g[utils.Point{x, y}] = r
		}
	}

	starts := make([]utils.Point, 0)
	for p, r := range g {
		if r == 'a' {
			starts = append(starts, p)
		}
	}

	minDist := math.MaxInt
	for _, start := range starts {
		dist := dijkstra(g, start, end)
		minDist = utils.Min(minDist, dist)
	}

	return minDist
}

func dijkstra(m map[utils.Point]rune, start utils.Point, end utils.Point) int {
	visited := make(map[utils.Point]bool)

	distance := make(map[utils.Point]int)
	for p, _ := range m {
		distance[p] = math.MaxInt
	}
	q := NewQueue()

	distance[start] = 0
	q.enqueue(path{start, 0})
	visited[start] = true

	for !q.empty() {
		shortest := q.dequeue()
		adjacent := computeAdjacents(shortest.coord)
		for _, p := range adjacent {
			if _, ok := m[p]; ok {
				if canMove(m, shortest.coord, p) {
					if _, ok = visited[p]; !ok {
						coord := path{p, shortest.cost + 1}
						distance[coord.coord] = coord.cost
						visited[coord.coord] = true
						q.enqueue(coord)
					}
				}
			}
		}
	}

	return distance[end]
}

func canMove(m map[utils.Point]rune, p1 utils.Point, p2 utils.Point) bool {
	if m[p1]+1 >= m[p2] {
		return true
	}

	return false
}

func computeAdjacents(p utils.Point) []utils.Point {
	adjacent := make([]utils.Point, 4)
	adjacent[0] = utils.Point{p.X - 1, p.Y}
	adjacent[1] = utils.Point{p.X, p.Y - 1}
	adjacent[2] = utils.Point{p.X + 1, p.Y}
	adjacent[3] = utils.Point{p.X, p.Y + 1}
	return adjacent
}

type path struct {
	coord utils.Point
	cost  int
}

type queue struct {
	items []path
}

func NewQueue() queue {
	items := make([]path, 0)
	return queue{items}
}

func (q *queue) enqueue(p path) {
	if len(q.items) == 0 {
		q.items = append(q.items, p)
		return
	}
	var insertFlag bool
	for k, v := range q.items {
		if p.cost < v.cost {
			if k > 0 {
				q.items = append(q.items[:k+1], q.items[k:]...)
				q.items[k] = p
				insertFlag = true
			} else {
				q.items = append([]path{p}, q.items...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		q.items = append(q.items, p)
	}
}

func (q *queue) dequeue() *path {
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	return &item
}

func (q *queue) empty() bool {
	return len(q.items) == 0
}
