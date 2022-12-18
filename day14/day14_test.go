package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var lines = []string{
	"498,4 -> 498,6 -> 496,6",
	"503,4 -> 502,4 -> 502,9 -> 494,9",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 24, solvePart1(lines))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 93, solvePart2(lines))
}
