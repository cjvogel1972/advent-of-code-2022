package grid

import "advent-of-code-2022/utils"

// IntGrid contains a multidimensional array of integers
type IntGrid struct {
	Grid[int]
}

// NewEmptyIntGrid creates a new grid of size width and height, with values initialized to 0
func NewEmptyIntGrid(width int, height int) IntGrid {
	grid := New[int](width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid.Set(x, y, 0)
		}
	}

	return IntGrid{grid}
}

// NewIntGridFromLines creates a new grid of single-digit integers from an array of lines
func NewIntGridFromLines(lines []string) IntGrid {
	width := len(lines[0])
	height := len(lines)

	grid := NewEmptyIntGrid(width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid.Set(x, y, utils.ToInt(string(lines[y][x])))
		}
	}

	return grid
}

// Max returns the maximum value in the grid
func (g IntGrid) Max() int {
	max := 0
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			max = utils.Max(max, g.grid[j][i])
		}
	}

	return max
}
