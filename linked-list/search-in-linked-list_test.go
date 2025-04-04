package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchKey(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		node *Node
		key  int
		want bool
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
			key:  3,
			want: true,
		},
		{
			name: "test case 2",
			node: &Node{
				data: 1,
				next: &Node{
					data: 2,
					next: &Node{
						data: 3,
					},
				},
			},
			key:  0,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchKey(tt.node, tt.key)
			assert.Equal(t, tt.want, got)
		})
	}
}
