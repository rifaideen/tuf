package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insertAt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head     *Node
		position int
		value    int
		want     []int
	}{
		{
			name:     "test case 1",
			head:     constructDll([]int{1, 2, 3, 4, 5}),
			position: 4,
			value:    6,
			want:     []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "test case 2",
			head:     constructDll([]int{1, 3, 4, 5}),
			position: 0,
			value:    2,
			want:     []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertAfter(tt.head, tt.position, tt.value)
			current := tt.head

			for _, v := range tt.want {
				assert.NotNil(t, current)
				assert.Equal(t, v, current.Data)
				current = current.Next
			}
		})
	}
}
