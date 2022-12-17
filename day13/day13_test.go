package day13

import (
	"advent-of-code-2022/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	lines := utils.ReadLines("day13_test-input.txt")

	assert.Equal(t, 13, solvePart1(lines))
}

func TestPart2(t *testing.T) {
	lines := utils.ReadLines("day13_test-input.txt")

	assert.Equal(t, 140, solvePart2(lines))
}
