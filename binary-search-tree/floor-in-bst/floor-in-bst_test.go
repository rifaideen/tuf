package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findFloor(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		k    int
		want *TreeNode
	}{
		{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 15,
				},
			},
			k: 7,
			want: &TreeNode{
				Val: 6,
			},
		},
		{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 15,
				},
			},
			k: 6,
			want: &TreeNode{
				Val: 6,
			},
		},
		{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 15,
				},
			},
			k:    1,
			want: nil,
		},
		{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 15,
				},
			},
			k: 10,
			want: &TreeNode{
				Val: 10,
			},
		},
		{
			root: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 12,
					Right: &TreeNode{
						Val: 13,
					},
				},
			},
			k: 10,
			want: &TreeNode{
				Val: 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findFloor(tt.root, tt.k)

			if tt.want != nil {
				assert.Equal(t, tt.want.Val, got.Val)
			} else {
				assert.Nil(t, got)
			}
		})
	}
}
