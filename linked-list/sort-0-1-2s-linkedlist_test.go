package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortOneTwoThree(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		want []int
	}{
		{
			name: "test case 1",
			head: constructLL([]int{0, 1, 2, 2, 1, 0, 1, 2}),
			want: []int{0, 0, 1, 1, 1, 2, 2, 2},
		},
		{
			name: "test case 2",
			head: constructLL([]int{2, 1, 0}),
			want: []int{0, 1, 2},
		},
		{
			name: "test case 3",
			head: constructLL([]int{1, 0}),
			want: []int{0, 1},
		},
		{
			name: "test case 4",
			head: constructLL([]int{1}),
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sortOneTwoThree(tt.head)

			i := 0

			for head != nil {
				assert.Equal(t, tt.want[i], head.data)
				head = head.next
				i++
			}
		})
	}
}
