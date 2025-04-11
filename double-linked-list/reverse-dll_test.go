package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseNode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		want []int
	}{
		{
			name: "test case 1",
			head: constructDll([]int{1, 2, 3, 4, 5}),
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "test case 2",
			head: constructDll([]int{1, 2}),
			want: []int{2, 1},
		},
		{
			name: "test case 3",
			head: constructDll([]int{1}),
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := reverseNode(tt.head)

			i := 0
			for head != nil {
				assert.Equal(t, tt.want[i], head.Data)

				i++
				head = head.Next
			}
		})
	}
}
