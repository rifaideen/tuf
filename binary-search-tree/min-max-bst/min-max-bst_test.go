package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMinMax(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root  *TreeNode
		want  int
		want2 int
	}{
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			want:  1,
			want2: 7,
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 10,
					},
				},
			},
			want:  1,
			want2: 10,
		},
		{
			root: &TreeNode{
				Val: 10,
			},
			want:  10,
			want2: 10,
		},
		{
			root:  nil,
			want:  -1,
			want2: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := findMinMax(tt.root)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want2, got2)
		})
	}
}
