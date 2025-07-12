package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_printPath(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		x    *TreeNode
		want []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			x: &TreeNode{
				Val: 7,
			},
			want: []int{1, 2, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := printPath(tt.root, tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}
