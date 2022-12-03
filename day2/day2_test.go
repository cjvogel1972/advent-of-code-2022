package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var lines = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 15, solvePart1(lines))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 12, solvePart2(lines))
}
