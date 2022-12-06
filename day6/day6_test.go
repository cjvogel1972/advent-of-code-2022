package day6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var lines = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 7, solvePart1(lines[0]))
	assert.Equal(t, 5, solvePart1(lines[1]))
	assert.Equal(t, 6, solvePart1(lines[2]))
	assert.Equal(t, 10, solvePart1(lines[3]))
	assert.Equal(t, 11, solvePart1(lines[4]))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 19, solvePart2(lines[0]))
	assert.Equal(t, 23, solvePart2(lines[1]))
	assert.Equal(t, 23, solvePart2(lines[2]))
	assert.Equal(t, 29, solvePart2(lines[3]))
	assert.Equal(t, 26, solvePart2(lines[4]))
}
