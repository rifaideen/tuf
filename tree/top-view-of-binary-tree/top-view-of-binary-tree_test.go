package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_topView(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		want []int
	}{
		// TODO: Add test cases.
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
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: []int{4, 2, 1, 3, 7},
		},
		{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 40,
					},
					Right: &TreeNode{
						Val: 60,
					},
				},
				Right: &TreeNode{
					Val: 30,
					Left: &TreeNode{
						Val: 90,
					},
					Right: &TreeNode{
						Val: 100,
					},
				},
			},
			want: []int{40, 20, 10, 30, 100},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 6,
							},
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: []int{2, 1, 3, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topView(tt.root)
			assert.Equal(t, tt.want, got)
		})
	}
}
