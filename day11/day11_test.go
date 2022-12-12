package day11

import (
	"advent-of-code-2022/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	lines := utils.ReadLines("day11_test-input.txt")

	assert.Equal(t, 10605, solvePart1(lines))
}

func TestPart2(t *testing.T) {
	lines := utils.ReadLines("day11_test-input.txt")

	//assert.Equal(t, 0, solvePart2(lines))
	assert.Equal(t, 2713310158, solvePart2(lines))
}
