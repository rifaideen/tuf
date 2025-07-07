package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_printBoundary(t *testing.T) {
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
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 5,
							},
							Right: &TreeNode{
								Val: 6,
							},
						},
					},
				},
				Right: &TreeNode{
					Val: 7,
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 9,
							Left: &TreeNode{
								Val: 10,
							},
							Right: &TreeNode{
								Val: 11,
							},
						},
					},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 10, 11, 9, 8, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := printBoundary(tt.root)
			assert.Equal(t, tt.want, got)
		})
	}
}
