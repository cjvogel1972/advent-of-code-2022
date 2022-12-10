package day10

import (
	"advent-of-code-2022/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	lines := utils.ReadLines("day10_test-input.txt")

	assert.Equal(t, 13140, solvePart1(lines))
}

func TestPart2(t *testing.T) {
	lines := utils.ReadLines("day10_test-input.txt")

	expected := []string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######.....",
	}
	result := solvePart2(lines)
	for i := 0; i < len(result); i++ {
		name := fmt.Sprintf("line %d", i)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expected[i], result[i])
		})
	}
}
