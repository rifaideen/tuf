package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	queue := Constructor()
	assert.True(t, queue.Empty())
	queue.Push(1)
	assert.Equal(t, 1, queue.Peek())
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	assert.Equal(t, 1, queue.Pop())
	assert.Equal(t, 2, queue.Peek())
	assert.Equal(t, 2, queue.Pop())
	assert.Equal(t, 3, queue.Pop())
	assert.Equal(t, 4, queue.Pop())
	assert.Equal(t, 5, queue.Pop())

	t.Run("expect panic", func(t *testing.T) {
		defer func() {
			err := recover()

			assert.Contains(t, err, "queue is empty")
		}()
		queue.Pop()
	})
}
