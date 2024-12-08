package day16

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strings"
)

type Puzzle struct{}

// Solve solves day 16's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day16/day16-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	valves := parseValves(lines)

	totalPressure := dijkstra1(valves, "AA", 30)

	return totalPressure
}

func solvePart2(lines []string) int {
	valves := parseValves(lines)

	totalPressure := dijkstra2(valves, "AA", 26)

	return totalPressure
}

func parseValves(lines []string) map[string]valve {
	valves := make(map[string]valve)
	for _, l := range lines {
		v := newValve(l)
		valves[v.name] = v
	}
	return valves
}

func dijkstra1(valves map[string]valve, startValve string, maxTime int) int {
	distance := make(map[string]int)
	q := NewQueue()

	q.enqueue(path{startValve, 0, make(map[string]bool), 0})

	max := 0
	for !q.empty() {
		shortest := q.dequeue()
		if shortest.time == maxTime-1 {
			max = utils.Max(max, shortest.cost)
			continue
		}

		adjacent := computeAdjacents(*shortest, valves)
		for _, p := range adjacent {
			d, ok := distance[p.hash()]
			if !ok || p.cost > d {
				distance[p.hash()] = p.cost
				q.enqueue(p)
			}
		}
	}

	return max
}

func dijkstra2(valves map[string]valve, startValve string, maxTime int) int {
	workingValvesCnt := 0
	for _, v := range valves {
		if v.rate > 0 {
			workingValvesCnt++
		}
	}

	distance := make(map[string]int)
	double := make(map[string]int)
	q := NewQueue()

	q.enqueue(path{startValve, 0, make(map[string]bool), 0})

	for !q.empty() {
		shortest := q.dequeue()
		if shortest.time == maxTime-1 {
			pct := float32(len(shortest.openValves)) / float32(workingValvesCnt)
			if pct > .25 && pct < .75 {
				double[shortest.hashOpenValves()] = shortest.cost
			}
			continue
		}

		adjacent := computeAdjacents(*shortest, valves)
		for _, p := range adjacent {
			d, ok := distance[p.hash()]
			if !ok || p.cost > d {
				distance[p.hash()] = p.cost
				q.enqueue(p)
			}
		}
	}

	max := 0
	for a, aCost := range double {
	inner:
		for b, bCost := range double {
			if a == b {
				continue
			}

			for _, v := range strings.Split(a, " ") {
				if strings.Contains(b, v) {
					continue inner
				}
			}

			max = utils.Max(max, aCost+bCost)
		}
	}

	return max
}

func computeAdjacents(curr path, valves map[string]valve) []path {
	result := make([]path, 0)
	v := valves[curr.position]
	if _, ok := curr.openValves[curr.position]; !ok && v.rate != 0 {
		open := make(map[string]bool)
		for s, _ := range curr.openValves {
			open[s] = true
		}
		open[curr.position] = true
		cost := curr.cost
		for s, _ := range open {
			cost += valves[s].rate
		}
		next := path{curr.position, curr.time + 1, open, cost}
		result = append(result, next)
	}

	for _, t := range v.tunnels {
		open := make(map[string]bool)
		for s, _ := range curr.openValves {
			open[s] = true
		}
		cost := curr.cost
		for s, _ := range open {
			cost += valves[s].rate
		}
		next := path{t, curr.time + 1, open, cost}
		result = append(result, next)
	}

	return result
}

type valve struct {
	name    string
	rate    int
	tunnels []string
}

func newValve(l string) valve {
	var name string
	var rate int
	fmt.Sscanf(l, "Valve %s has flow rate=%d;", &name, &rate)

	idx := strings.LastIndex(l, "valves")
	var tunnels []string
	if idx == -1 {
		idx = strings.LastIndex(l, "valve")
		t := l[idx+6:]
		tunnels = []string{t}
	} else {
		t := l[idx+7:]
		tunnels = strings.Split(t, ", ")
	}

	return valve{name, rate, tunnels}
}

type path struct {
	position   string
	time       int
	openValves map[string]bool
	cost       int
}

func (p path) hash() string {
	return fmt.Sprintf("%s-%s-%d", p.position, p.hashOpenValves(), p.time)
}

func (p path) hashOpenValves() string {
	open := make([]string, 0, len(p.openValves))
	for s, _ := range p.openValves {
		open = append(open, s)
	}
	s := fmt.Sprintf("%v", open)

	return s[1 : len(s)-1]
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
		if p.cost > v.cost {
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
