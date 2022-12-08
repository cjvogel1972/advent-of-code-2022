package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmptyBoolGrid(t *testing.T) {
	g := NewEmptyBoolGrid(5, 5)

	assert.Equal(t, 5, g.Width)
	assert.Equal(t, 5, g.Height)
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			assert.False(t, g.grid[y][x])
		}
	}
}

func TestCountTrue(t *testing.T) {
	g := NewEmptyBoolGrid(5, 5)
	for i := 0; i < 5; i++ {
		g.Set(i, i, true)
	}

	assert.Equal(t, 5, g.CountTrue())
}

func TestCountFalse(t *testing.T) {
	g := NewEmptyBoolGrid(5, 5)
	for i := 0; i < 5; i++ {
		g.Set(i, i, true)
	}

	assert.Equal(t, 20, g.CountFalse())
}
