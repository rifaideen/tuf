package main

import (
	"fmt"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	nodes := &Node{
		data: 1,
		next: &Node{
			data: 2,
			next: &Node{
				data: 3,
				next: &Node{
					data: 4,
					next: &Node{
						data: 5,
						next: nil,
					},
				},
			},
		},
	}

	third := nodes.next.next

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		node *Node
	}{
		{
			name: "test case 1",
			node: third,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.node)

			fmt.Printf("%+v\n", nodes)
		})
	}
}
