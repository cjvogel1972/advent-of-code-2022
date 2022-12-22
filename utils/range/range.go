package _range

import "advent-of-code-2022/utils"

// Range contains a range of integers
type Range struct {
	Start int
	End   int
}

// New creates a new range, making sure the start and end are in the correct order
func New(a int, b int) Range {
	start := utils.Min(a, b)
	end := utils.Max(a, b)
	return Range{start, end}
}

// Contains checks if the given number is in the range
func (r Range) Contains(i int) bool {
	return utils.Within(i, r.Start, r.End)
}

// Within checks if another range is inside the current range
func (r Range) Within(r1 Range) bool {
	return utils.Within(r1.Start, r.Start, r.End) && utils.Within(r1.End, r.Start, r.End)
}

// Overlap checks if another range overlaps the current range
func (r Range) Overlap(r1 Range) bool {
	return utils.Within(r1.Start, r.Start, r.End) || utils.Within(r1.End, r.Start, r.End)
}

func (r Range) Merge(r1 Range) Range {
	s := utils.Min(r.Start, r1.Start)
	e := utils.Max(r.End, r1.End)

	return New(s, e)
}
