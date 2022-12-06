package utils

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(i interface{}) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	i := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return i
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	return s.items[len(s.items)-1]
}

func (s *Stack) Size() int {
	return len(s.items)
}
