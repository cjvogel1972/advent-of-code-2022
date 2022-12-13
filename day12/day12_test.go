package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var lines = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 31, solvePart1(lines))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 29, solvePart2(lines))
}
