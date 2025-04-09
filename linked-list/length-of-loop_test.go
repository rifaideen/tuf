package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLoop(t *testing.T) {
	fourth := &Node{
		data: 4,
		next: &Node{
			data: 5,
		},
	}

	fourth.next.next = fourth

	head := &Node{
		data: 1,
		next: &Node{
			data: 2,
			next: &Node{
				data: 3,
				next: fourth,
			},
		},
	}

	tests := []struct {
		name string // description of this test case
		node *Node
		want int
	}{
		{
			name: "test case 1",
			node: &Node{
				data: 1,
				next: &Node{
					data: 2,
					next: &Node{
						data: 3,
					},
				},
			},
			want: 0,
		},
		{
			name: "test case 2",
			node: head,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLoop(tt.node)
			assert.Equal(t, tt.want, got)
		})
	}
}
