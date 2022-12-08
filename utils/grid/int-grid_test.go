package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var lines = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func TestNewEmptyIntGrid(t *testing.T) {
	g := NewEmptyIntGrid(5, 5)

	assert.Equal(t, 5, g.Width)
	assert.Equal(t, 5, g.Height)
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			assert.Equal(t, 0, g.grid[y][x])
		}
	}
}

func TestNewFromLines(t *testing.T) {
	g := NewIntGridFromLines(lines)

	expected := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	assert.Equal(t, 5, g.Width)
	assert.Equal(t, 5, g.Height)
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			assert.Equal(t, expected[y][x], g.grid[y][x])
		}
	}
}

func TestMax(t *testing.T) {
	g := NewIntGridFromLines(lines)

	assert.Equal(t, 9, g.Max())
}
