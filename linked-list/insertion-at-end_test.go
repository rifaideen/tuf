package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insertEnd(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		num  int
		want []int
	}{
		{
			name: "test case 1",
			head: &Node{
				data: 1,
				next: &Node{
					data: 2,
					next: &Node{
						data: 3,
					},
				},
			},
			num:  4,
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := insertEnd(tt.head, tt.num)
			i := 0

			for head != nil {
				assert.Equal(t, tt.want[i], head.data)
				head = head.next
				i++
			}
		})
	}
}
