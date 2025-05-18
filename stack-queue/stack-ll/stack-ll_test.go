package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	assert.Equal(t, 4, s.Pop())
	s.Push(5)
	assert.Equal(t, 5, s.Top())
}
