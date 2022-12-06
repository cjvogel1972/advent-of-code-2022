package utils

// Range contains a range of integers
type Range struct {
	start int
	end   int
}

// NewRange creates a new range, making sure the start and end are in the correct order
func NewRange(a int, b int) Range {
	start := Min(a, b)
	end := Max(a, b)
	return Range{start, end}
}

// Contains checks if the given number is in the range
func (r Range) Contains(i int) bool {
	return Within(i, r.start, r.end)
}

// Within checks if another range is inside the current range
func (r Range) Within(r1 Range) bool {
	return Within(r1.start, r.start, r.end) && Within(r1.end, r.start, r.end)
}

// Overlap checks if another range overlaps the current range
func (r Range) Overlap(r1 Range) bool {
	return Within(r1.start, r.start, r.end) || Within(r1.end, r.start, r.end)
}
