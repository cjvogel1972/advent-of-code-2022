package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInBetween(t *testing.T) {
	tests := []struct {
		name   string
		i      int
		min    int
		max    int
		output bool
	}{
		{"before", 0, 1, 10, false},
		{"front range", 1, 1, 10, true},
		{"mid range", 5, 1, 10, true},
		{"end range", 10, 1, 10, true},
		{"after", 11, 1, 10, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, InBetween(tt.i, tt.min, tt.max))
		})
	}
}

func TestHasDuplicates(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"abcdef", false},
		{"aaaa", true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.output, HasDuplicates(tt.input))
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name   string
		a      int
		b      int
		output int
	}{
		{"b bigger", 10, 20, 20},
		{"a bigger", 30, 20, 30},
		{"equal", 5, 5, 5},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, Max(tt.a, tt.b))
		})
	}
}
