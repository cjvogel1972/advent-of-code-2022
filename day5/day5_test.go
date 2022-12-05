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

func TestParseCrates(t *testing.T) {
	crates, linesRead := parseCrates(lines)

	assert.Equal(t, 5, linesRead)
	assert.Equal(t, 3, len(crates))
	assert.Equal(t, 2, len(crates[0]))
	assert.Equal(t, 3, len(crates[1]))
	assert.Equal(t, 1, len(crates[2]))
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
