package day11

import (
	"advent-of-code-2022/utils"
	"fmt"
	"regexp"
	"strings"
)

type Puzzle struct{}

// Solve solves day 11's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day11/day11-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	monkeys := make([]monkey, 0)
	for i := 0; i < len(lines); i = i + 7 {
		m := parseMonkey(lines[i : i+7])
		monkeys = append(monkeys, m)
	}

	for i := 0; i < 20; i++ {
		for monkeyIdx := range monkeys {
			m := &monkeys[monkeyIdx]
			for _, item := range m.items {
				item = m.operation.do(item)
				item /= 3
				var newMonkey int
				if item%m.test == 0 {
					newMonkey = m.trueMonkey
				} else {
					newMonkey = m.falseMonkey
				}
				monkeys[newMonkey].items = append(monkeys[newMonkey].items, item)
				m.items = m.items[1:]
				m.itemsInspected++
			}
		}
	}

	topTwo := make([]int, 2)
	for _, m := range monkeys {
		if m.itemsInspected > topTwo[0] {
			topTwo[1] = topTwo[0]
			topTwo[0] = m.itemsInspected
		} else if m.itemsInspected > topTwo[1] {
			topTwo[1] = m.itemsInspected
		}
	}

	return topTwo[0] * topTwo[1]
}

func solvePart2(lines []string) int {
	monkeys := make([]monkey, 0)
	for i := 0; i < len(lines); i = i + 7 {
		m := parseMonkey(lines[i : i+7])
		monkeys = append(monkeys, m)
	}

	modulus := 1
	for _, m := range monkeys {
		modulus *= m.test
	}

	for i := 0; i < 10000; i++ {
		for monkeyIdx := range monkeys {
			m := &monkeys[monkeyIdx]
			for _, item := range m.items {
				item = m.operation.do(item)
				item %= modulus
				var newMonkey int
				if item%m.test == 0 {
					newMonkey = m.trueMonkey
				} else {
					newMonkey = m.falseMonkey
				}
				monkeys[newMonkey].items = append(monkeys[newMonkey].items, item)
				m.items = m.items[1:]
				m.itemsInspected++
			}
		}
	}

	topTwo := make([]int, 2)
	for _, m := range monkeys {
		if m.itemsInspected > topTwo[0] {
			topTwo[1] = topTwo[0]
			topTwo[0] = m.itemsInspected
		} else if m.itemsInspected > topTwo[1] {
			topTwo[1] = m.itemsInspected
		}
	}

	return topTwo[0] * topTwo[1]
}

type monkey struct {
	items          []int
	operation      operation
	test           int
	trueMonkey     int
	falseMonkey    int
	itemsInspected int
}

func parseMonkey(lines []string) monkey {
	var op operation
	var test int
	var trueMonkey int
	var falseMonkey int
	itemsRegEx := regexp.MustCompile(": (.*)")
	itemsInput := itemsRegEx.FindStringSubmatch(lines[1])
	opRegEx := regexp.MustCompile("new = old (.*)")
	opInput := opRegEx.FindStringSubmatch(lines[2])
	fmt.Sscanf(lines[2], "  Operation: new = old %s", &opInput)
	fmt.Sscanf(lines[3], "  Test: divisible by %d", &test)
	fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &trueMonkey)
	fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &falseMonkey)
	items := make([]int, 0)
	for _, s := range strings.Split(itemsInput[1], ",") {
		items = append(items, utils.ToInt(strings.Trim(s, " ")))
	}
	op = parseOperation(opInput[1])

	return monkey{items, op, test, trueMonkey, falseMonkey, 0}
}

func parseOperation(o string) operation {
	if o == "* old" {
		return operation{"^", 0}
	}
	
	e := strings.Split(o, " ")

	return operation{e[0], utils.ToInt(e[1])}
}

type operation struct {
	op      string
	operand int
}

func (o operation) do(i int) int {
	switch o.op {
	case "+":
		return i + o.operand
	case "-":
		return i - o.operand
	case "*":
		return i * o.operand
	case "/":
		return i / o.operand
	case "^":
		return i * i
	}

	panic(-1)
}
