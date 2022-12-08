package grid

// BoolGrid contains a multidimensional array of booleans
type BoolGrid struct {
	Grid[bool]
}

// NewEmptyBoolGrid creates a new grid of size width and height, with values initialized to false
func NewEmptyBoolGrid(width int, height int) BoolGrid {
	grid := New[bool](width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid.Set(x, y, false)
		}
	}

	return BoolGrid{grid}
}

// CountTrue counts the number of true items in the grid
func (g BoolGrid) CountTrue() int {
	count := 0
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			if g.grid[i][j] {
				count++
			}
		}
	}

	return count
}

// CountFalse counts the number of false items in the grid
func (g BoolGrid) CountFalse() int {
	count := 0
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			if !g.grid[i][j] {
				count++
			}
		}
	}

	return count
}
