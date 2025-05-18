package main

import (
	"testing"
)

func TestQueueUsingStack(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
}
