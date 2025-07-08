package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bottomView(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
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
							Val: 8,
						},
						Right: &TreeNode{
							Val: 9,
						},
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
			want: []int{4, 8, 6, 9, 7},
		},
		{
			root: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 22,
					Left: &TreeNode{
						Val: 34,
					},
					Right: &TreeNode{
						Val: 25,
					},
				},
			},
			want: []int{5, 8, 34, 22, 25},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bottomView(tt.root)
			assert.Equal(t, tt.want, got)
		})
	}
}
