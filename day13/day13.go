package day13

import (
	"advent-of-code-2022/utils"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Puzzle struct{}

const (
	inOrder    = 1
	equal      = 2
	outOfOrder = 3
)

// Solve solves day 13's puzzles
func (Puzzle) Solve() {
	lines := utils.ReadLines("day13/day13-input.txt")

	fmt.Printf("Part 1: %d\n", solvePart1(lines))
	fmt.Printf("Part 2: %d\n", solvePart2(lines))
}

func solvePart1(lines []string) int {
	sumIndices := 0
	index := 0
	for i := 0; i < len(lines); i = i + 3 {
		index++
		left := newPacket(lines[i])
		right := newPacket(lines[i+1])
		result := listInOrder(left.values, right.values)
		if result == inOrder {
			sumIndices += index
		}
	}

	return sumIndices
}

func solvePart2(lines []string) int {
	divPacket1 := newPacket("[[2]]")
	divPacket2 := newPacket("[[6]]")

	packets := make([]*packet, 0)
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		p := newPacket(l)
		packets = append(packets, &p)
	}
	packets = append(packets, &divPacket1)
	packets = append(packets, &divPacket2)

	sort.Slice(packets, func(i, j int) bool {
		return listInOrder(packets[i].values, packets[j].values) == inOrder
	})

	var divIdx1 int
	var divIdx2 int

	for i, p := range packets {
		if p == &divPacket1 {
			divIdx1 = i + 1
		}
		if p == &divPacket2 {
			divIdx2 = i + 1
		}
	}

	return divIdx1 * divIdx2
}

func listInOrder(left []interface{}, right []interface{}) int {
	result := equal
	minSize := utils.Min(len(left), len(right))
	for i := 0; i < minSize; i++ {
		if reflect.TypeOf(left[i]) == reflect.TypeOf(right[i]) {
			if _, ok := left[i].(int); ok {
				intResult := compareInts(left[i].(int), right[i].(int))
				if intResult == inOrder || intResult == outOfOrder {
					result = intResult
					break
				}
			} else {
				listResult := listInOrder(left[i].([]interface{}), right[i].([]interface{}))
				if listResult == inOrder || listResult == outOfOrder {
					result = listResult
					break
				}
			}
		} else {
			l, r := createLists(left[i], right[i])
			listResult := listInOrder(l, r)
			if listResult == inOrder || listResult == outOfOrder {
				result = listResult
				break
			}
		}
	}

	if len(left) != len(right) && result == equal {
		if len(left) < len(right) {
			result = inOrder
		} else {
			result = outOfOrder
		}
	}

	return result
}

func compareInts(l int, r int) int {
	var result int

	if l == r {
		result = equal
	} else if l < r {
		result = inOrder
	} else {
		result = outOfOrder
	}

	return result
}

func createLists(left interface{}, right interface{}) ([]interface{}, []interface{}) {
	var l []interface{}
	var r []interface{}
	if _, ok := left.(int); ok {
		l = make([]interface{}, 1)
		l[0] = left
		r = right.([]interface{})
	} else {
		r = make([]interface{}, 1)
		r[0] = right
		l = left.([]interface{})
	}

	return l, r
}

type packet struct {
	values []interface{}
}

func newPacket(l string) packet {
	values := make([]interface{}, 0)
	l = l[1 : len(l)-1]
	for i := 0; i < len(l); i++ {
		if l[i] == ',' {
			continue
		}
		if l[i] == '[' {
			subList, idx := parseList(l, i+1)
			values = append(values, subList)
			i = idx
		} else {
			values = append(values, utils.ToInt(string(l[i])))
		}
	}

	return packet{values}
}

func parseList(l string, startIdx int) ([]interface{}, int) {
	list := make([]interface{}, 0)
	idx := startIdx
	for l[idx] != ']' {
		if l[idx] == ',' {
			idx++
			continue
		} else if l[idx] == '[' {
			subList, i := parseList(l, idx+1)
			idx = i
			list = append(list, subList)
		} else {
			n := make([]rune, 0)
			for l[idx] != ',' && l[idx] != ']' {
				n = append(n, rune(l[idx]))
				idx++
			}
			list = append(list, utils.ToInt(string(n)))
		}
	}

	return list, idx + 1
}
