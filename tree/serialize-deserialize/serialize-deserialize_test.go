package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec_Serialize(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		want string
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			want: "1,2,3,#,#,4,5,#,#,#,#",
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 6,
							Left: &TreeNode{
								Val: 8,
							},
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			want: "1,2,3,#,#,4,5,6,7,#,#,8,#,#,#,#,#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Constructor()
			got := c.serialize(tt.root)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCodec_Deserialize(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data string
	}{
		{
			data: "1,2,3,#,#,4,5,#,#,#,#",
		},
		{
			data: "1,2,3,#,#,4,5,6,7,#,#,8,#,#,#,#,#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Constructor()
			got := c.deserialize(tt.data)

			assert.Equal(t, tt.data, c.serialize(got))
		})
	}
}
