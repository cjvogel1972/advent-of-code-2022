package _range

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		start int
		end   int
	}{
		{"in order", 10, 20, 10, 20},
		{"backwards", 20, 10, 10, 20},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := New(tt.a, tt.b)
			assert.Equal(t, tt.start, r.start)
			assert.Equal(t, tt.end, r.end)
		})
	}
}

func TestContains(t *testing.T) {
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
			r := Range{tt.min, tt.max}
			assert.Equal(t, tt.output, r.Contains(tt.i))
		})
	}

}

func TestWithin(t *testing.T) {
	tests := []struct {
		name    string
		r1Start int
		r1End   int
		r2Start int
		r2End   int
		output  bool
	}{
		{"before", 10, 11, 1, 3, false},
		{"front overlap", 3, 10, 1, 5, false},
		{"within", 1, 10, 3, 7, true},
		{"end overlap", 3, 10, 7, 15, false},
		{"after", 1, 10, 15, 20, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r1 := Range{tt.r1Start, tt.r1End}
			r2 := Range{tt.r2Start, tt.r2End}
			assert.Equal(t, tt.output, r1.Within(r2))
		})
	}

}

func TestOverlap(t *testing.T) {
	tests := []struct {
		name    string
		r1Start int
		r1End   int
		r2Start int
		r2End   int
		output  bool
	}{
		{"before", 10, 11, 1, 3, false},
		{"front overlap", 3, 10, 1, 5, true},
		{"within", 1, 10, 3, 7, true},
		{"end overlap", 3, 10, 7, 15, true},
		{"after", 1, 10, 15, 20, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r1 := Range{tt.r1Start, tt.r1End}
			r2 := Range{tt.r2Start, tt.r2End}
			assert.Equal(t, tt.output, r1.Overlap(r2))
		})
	}

}
