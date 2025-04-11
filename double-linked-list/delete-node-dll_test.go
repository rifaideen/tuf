package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteNode(t *testing.T) {
	head := &Node{
		Data: 1,
	}

	// 2nd
	head.Next = &Node{
		Data: 2,
		Prev: head,
	}

	// 3rd
	head.Next.Next = &Node{
		Data: 3,
	}

	head.Next.Next.Prev = head.Next

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		x    int
		want []int
	}{
		{
			name: "test case 1",
			head: head,
			x:    3,
			want: []int{1, 2},
		},
		{
			name: "test case 2",
			head: head,
			x:    2,
			want: []int{1, 3},
		},
		{
			name: "test case 3",
			head: head,
			x:    1,
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := deleteNode(tt.head, tt.x)

			i := 0
			for head != nil {
				assert.Equal(t, tt.want[i], head.Data)

				i++
				head = head.Next
			}
		})
	}
}
