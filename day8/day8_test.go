package day8

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

func TestPart1(t *testing.T) {
	assert.Equal(t, 21, solvePart1(lines))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 8, solvePart2(lines))
}

func TestTreeVisible(t *testing.T) {
	tests := []struct {
		name   string
		row    []int
		i      int
		result bool
	}{
		{"left outer visible", []int{2, 5, 5, 1, 2}, 0, true},
		{"right outer visible", []int{2, 5, 5, 1, 2}, 4, true},
		{"visible only right", []int{3, 3, 5, 4, 9}, 2, true},
		{"visible only left", []int{6, 5, 3, 3, 2}, 1, true},
		{"not visible", []int{3, 3, 5, 4, 9}, 3, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.result, treeVisible(tt.row, tt.i))
		})
	}
}

func TestTreeScore(t *testing.T) {
	tests := []struct {
		name   string
		row    []int
		i      int
		result int
	}{
		{"left outer", []int{2, 5, 5, 1, 2}, 0, 0},
		{"right outer", []int{2, 5, 5, 1, 2}, 4, 0},
		{"tree block left, clear right", []int{2, 5, 5, 1, 2}, 2, 2},
		{"clear left, block two down on right", []int{3, 3, 5, 4, 9}, 2, 4},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.result, treeScore(tt.row, tt.i))
		})
	}
}
