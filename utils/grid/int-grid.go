package grid

import "advent-of-code-2022/utils"

// IntGrid contains a multidimensional array of integers
type IntGrid struct {
	Grid[int]
}

// NewEmptyIntGrid creates a new grid of size width and height, with values initialized to 0
func NewEmptyIntGrid(width int, height int) IntGrid {
	g := New[int](width, height)
	g.Iterate(func(x int, y int) {
		_ = g.Set(x, y, 0)
	})

	return IntGrid{g}
}

// NewIntGridFromLines creates a new grid of single-digit integers from an array of lines
func NewIntGridFromLines(lines []string) IntGrid {
	width := len(lines[0])
	height := len(lines)

	g := NewEmptyIntGrid(width, height)
	g.Iterate(func(x int, y int) {
		_ = g.Set(x, y, utils.ToInt(string(lines[y][x])))
	})

	return g
}

// Max returns the maximum value in the grid
func (g IntGrid) Max() int {
	max := 0
	g.Iterate(func(x int, y int) {
		v, _ := g.Value(x, y)
		max = utils.Max(max, v)
	})

	return max
}
