package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var lines = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestSplitFile(t *testing.T) {
	top, bottom := SplitFile(lines)

	assert.Equal(t, 4, len(top))
	assert.Equal(t, 4, len(bottom))
}
