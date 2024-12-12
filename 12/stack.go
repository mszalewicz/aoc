package main

import "errors"

type Stack[T comparable] struct {
	elements []T
}

func (s *Stack[T]) Add(item T) {
	s.elements = append(s.elements, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}

func (s *Stack[T]) Remove(value T) (T, error) {
	for i, v := range s.elements {
		if v == value {
			s.elements = append(s.elements[:i], s.elements[i+1:]...)
			return v, nil
		}
	}
	var zero T
	return zero, errors.New("value not found")
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) Contains(value T) bool {
	for _, v := range s.elements {
		if v == value {
			return true
		}
	}
	return false
}
