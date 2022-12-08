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
