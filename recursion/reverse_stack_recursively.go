package main

func ReverseStack(s *Stack) {
	if s.IsEmpty() {
		return
	}

	top, _ := s.Pop()

	ReverseStack(s)
	insertStackBottom(s, top)
}

func insertStackBottom(s *Stack, value int) {
	if s.IsEmpty() {
		s.Push(value)
		return
	}

	top, _ := s.Pop()
	insertStackBottom(s, value)
	s.Push(top)
}
