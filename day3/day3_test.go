package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var lines = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 157, solvePart1(lines))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 70, solvePart2(lines))
}
