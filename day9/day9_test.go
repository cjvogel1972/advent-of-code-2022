package day9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var lines = []string{
	"R 4",
	"U 4",
	"L 3",
	"D 1",
	"R 4",
	"D 1",
	"L 5",
	"R 2",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 13, solvePart1(lines))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 1, solvePart2(lines))
}

func TestPart2Big(t *testing.T) {
	var bigLines = []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}
	assert.Equal(t, 36, solvePart2(bigLines))
}
