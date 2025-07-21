package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findCeil(t *testing.T) {
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
					Val: 8,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 16,
					},
				},
			},
			k: 12,
			want: &TreeNode{
				Val: 12,
			},
		},
		{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 16,
					},
				},
			},
			k: 13,
			want: &TreeNode{
				Val: 15,
			},
		},
		{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 16,
					},
				},
			},
			k:    17,
			want: nil,
		},
		{
			root: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
				Right: &TreeNode{
					Val: 12,
					Left: &TreeNode{
						Val: 10,
					},
					Right: &TreeNode{
						Val: 14,
						Left: &TreeNode{
							Val: 13,
						},
					},
				},
			},
			k: 11,
			want: &TreeNode{
				Val: 12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCeil(tt.root, tt.k)

			if tt.want != nil {
				assert.Equal(t, tt.want.Val, got.Val)
			} else {
				assert.Nil(t, got)
			}
		})
	}
}
