package main

import "fmt"

// Stack represents a simple integer stack
type Stack struct {
	items []int
}

// NewStack creates a new empty stack
func NewStack() *Stack {
	return &Stack{items: []int{}}
}

// Push adds an element to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

// Peek returns the top element without removing it
func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// String returns a string representation of the stack
func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.items)
}

func SortStack(s *Stack) {
	// base case
	if s.IsEmpty() {
		return
	}

	temp, _ := s.Pop()

	SortStack(s)

	insertSorted(s, temp)
}

func insertSorted(s *Stack, value int) {
	// push directly when stack is empty
	if s.IsEmpty() {
		s.Push(value)
		return
	}

	// check and push if value is greater than top
	top, _ := s.Peek()

	if value > top {
		s.Push(value)
		return
	}

	// value is smaller, pop the highest values and insert them recursively
	//
	// this will continue popping the value from the stack until it becomes empty or value > top
	temp, _ := s.Pop()

	// recursively pop and insert at the right place
	insertSorted(s, value)

	// push the removed item on the top of the stack, this way stack gets sorted
	s.Push(temp)
}
