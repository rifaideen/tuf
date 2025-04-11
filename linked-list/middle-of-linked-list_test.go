package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_middleNode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		want int
	}{
		{
			name: "test case 1",
			head: constructLL([]int{1, 2, 3}),
			want: 2,
		},
		{
			name: "test case 2",
			head: constructLL([]int{1, 2, 3, 4}),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := middleNode(tt.head)
			assert.Equal(t, tt.want, head.data)
		})
	}
}
