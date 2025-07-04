package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_traverse(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nodes *TreeNode
		want  [][]int
	}{
		{
			nodes: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: [][]int{
				{1, 2, 3},
				{2, 1, 3},
				{2, 3, 1},
			},
		},
		{
			nodes: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: [][]int{
				{
					1, 2, 4, 5, 3, 6, 7,
				},
				{
					4, 2, 5, 1, 6, 3, 7,
				},
				{
					4, 5, 2, 6, 7, 3, 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := traverse(tt.nodes)
			assert.Equal(t, tt.want, got)
		})
	}
}
