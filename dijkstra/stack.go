package main

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	var none T
	if len(s.vals) == 0 {
		return none, false
	}
	first := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return first, true
}
