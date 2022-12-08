package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	g := New[int](5, 5)

	assert.Equal(t, 5, g.Width)
	assert.Equal(t, 5, g.Height)
	assert.Equal(t, 5, len(g.grid[0]))
	assert.Equal(t, 5, len(g.grid))
}

func TestRow(t *testing.T) {
	g := createGrid()

	tests := []struct {
		name   string
		i      int
		result []int
		error  bool
	}{
		{"good row", 1, []int{2, 5, 5, 1, 2}, false},
		{"negative row", -1, nil, true},
		{"bad row", 5, nil, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r, err := g.Row(tt.i)
			assert.Equal(t, tt.result, r)
			assert.Equal(t, tt.error, err != nil)
		})
	}
}

func TestColumn(t *testing.T) {
	g := createGrid()

	tests := []struct {
		name   string
		i      int
		result []int
		error  bool
	}{
		{"good column", 1, []int{0, 5, 5, 3, 5}, false},
		{"negative column", -1, nil, true},
		{"bad column", 5, nil, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r, err := g.Column(tt.i)
			assert.Equal(t, tt.result, r)
			assert.Equal(t, tt.error, err != nil)
		})
	}
}

func TestValue(t *testing.T) {
	g := createGrid()

	tests := []struct {
		name   string
		x      int
		y      int
		result int
		error  bool
	}{
		{"good coordinate", 2, 2, 3, false},
		{"negative x", -1, 2, -1, true},
		{"bad x", 5, 2, -1, true},
		{"negative y", 2, -1, -1, true},
		{"bad x", 2, 5, -1, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r, err := g.Value(tt.x, tt.y)
			if !tt.error {
				assert.Equal(t, tt.result, r)
			}
			assert.Equal(t, tt.error, err != nil)
		})
	}
}

func TestSet(t *testing.T) {
	g := createGrid()

	tests := []struct {
		name  string
		x     int
		y     int
		value int
		error bool
	}{
		{"good coordinate", 2, 2, 3, false},
		{"negative x", -1, 2, -1, true},
		{"bad x", 5, 2, -1, true},
		{"negative y", 2, -1, -1, true},
		{"bad x", 2, 5, -1, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := g.Set(tt.x, tt.y, tt.value)
			if !tt.error {
				assert.Equal(t, tt.value, g.grid[tt.y][tt.x])
			}
			assert.Equal(t, tt.error, err != nil)
		})
	}
}

func createGrid() Grid[int] {
	g := New[int](5, 5)

	var grid = [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			g.grid[y][x] = grid[y][x]
		}
	}

	return g
}
