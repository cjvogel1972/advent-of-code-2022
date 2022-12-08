package grid

import (
	"fmt"
)

// Grid contains a multidimensional array of values
type Grid[T any] struct {
	grid   [][]T
	Width  int
	Height int
}

func New[T any](width int, height int) Grid[T] {
	grid := make([][]T, height)
	for y := 0; y < height; y++ {
		row := make([]T, width)
		grid[y] = row
	}

	return Grid[T]{grid, width, height}
}

// Row returns a row, returning an error if the row number is invalid
func (g *Grid[T]) Row(row int) ([]T, error) {
	if row < 0 || row >= g.Height {
		return nil, fmt.Errorf("invalid row number %d", row)
	}

	return g.grid[row], nil
}

// Column returns a column, returning an error if the row number is invalid
func (g *Grid[T]) Column(col int) ([]T, error) {
	if col < 0 || col >= g.Width {
		return nil, fmt.Errorf("invalid column number %d", col)
	}

	column := make([]T, g.Width)
	for i := 0; i < g.Height; i++ {
		column[i] = g.grid[i][col]
	}

	return column, nil
}

// Value returns the value at the given x, y coordinate, returning an error if the coordinate is invalid
func (g *Grid[T]) Value(x int, y int) (T, error) {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		var result T
		return result, fmt.Errorf("invalid coordinate [%d, %d]", x, y)
	}

	return g.grid[y][x], nil
}

// Set sets the value at the given x, y coordinate, returning an error if the coordinate is invalid
func (g *Grid[T]) Set(x int, y int, value T) error {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return fmt.Errorf("invalid coordinate [%d, %d]", x, y)
	}

	g.grid[y][x] = value
	return nil
}
