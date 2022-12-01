package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var lines = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 24000, solvePart1(lines))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 45000, solvePart2(lines))
}
