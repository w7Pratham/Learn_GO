package main

import (
	"fmt"
)

type Stack struct {
	items []rune
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(data rune) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Peek() (rune, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}
