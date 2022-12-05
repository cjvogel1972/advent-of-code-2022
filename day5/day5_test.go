package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var lines = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, "CMZ", solvePart1(lines))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "MCD", solvePart2(lines))
}

func TestNewCrateStack(t *testing.T) {
	stack := newCrateStack()

	assert.Equal(t, 0, len(stack.crates))
}

func TestStackPushCrate(t *testing.T) {
	stack := newCrateStack()
	stack.pushCrate('A')

	assert.Equal(t, 1, len(stack.crates))
	assert.Equal(t, crate('A'), stack.crates[0])
}

func TestStackPushCrates(t *testing.T) {
	stack := newCrateStack()
	stack.pushCrates([]crate{'A', 'B'})

	assert.Equal(t, 2, len(stack.crates))
	assert.Equal(t, crate('A'), stack.crates[0])
	assert.Equal(t, crate('B'), stack.crates[1])
}

func TestStackPopCratesInOrder(t *testing.T) {
	stack := newCrateStack()
	stack.pushCrates([]crate{'A', 'B', 'C', 'D'})
	top := stack.popCratesInOrder(2)

	assert.Equal(t, 2, len(stack.crates))
	assert.Equal(t, crate('C'), stack.crates[0])
	assert.Equal(t, crate('D'), stack.crates[1])
	assert.Equal(t, 2, len(top))
	assert.Equal(t, crate('A'), top[0])
	assert.Equal(t, crate('B'), top[1])
}

func TestStackPopCratesReversed(t *testing.T) {
	stack := newCrateStack()
	stack.pushCrates([]crate{'A', 'B', 'C', 'D'})
	top := stack.popCratesReversed(2)

	assert.Equal(t, 2, len(stack.crates))
	assert.Equal(t, crate('C'), stack.crates[0])
	assert.Equal(t, crate('D'), stack.crates[1])
	assert.Equal(t, 2, len(top))
	assert.Equal(t, crate('B'), top[0])
	assert.Equal(t, crate('A'), top[1])
}

func TestStackTopCrate(t *testing.T) {
	stack := newCrateStack()
	stack.pushCrates([]crate{'A', 'B', 'C', 'D'})

	assert.Equal(t, crate('A'), stack.topCrate())
}

func TestStackSize(t *testing.T) {
	stack := newCrateStack()
	stack.pushCrates([]crate{'A', 'B', 'C', 'D'})

	assert.Equal(t, 4, stack.size())
}

func TestParseCrates(t *testing.T) {
	crates, linesRead := parseCrates(lines)

	assert.Equal(t, 5, linesRead)
	assert.Equal(t, 3, len(crates))
	assert.Equal(t, 2, crates[0].size())
	assert.Equal(t, 3, crates[1].size())
	assert.Equal(t, 1, crates[2].size())
}

func TestParseAction(t *testing.T) {
	numCratesToMove, fromColumn, toColumn := parseAction(lines[5])

	assert.Equal(t, 1, numCratesToMove)
	assert.Equal(t, 1, fromColumn)
	assert.Equal(t, 0, toColumn)
}

func TestGetTopCrates(t *testing.T) {
	crates, _ := parseCrates(lines)

	assert.Equal(t, "NDP", getTopCrates(crates))
}
