package main

import (
	"fmt"
	"testing"
)

func Test_changeTree(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
	}{
		{
			root: &TreeNode{
				Val: 160,
				Left: &TreeNode{
					Val: 10,
				},
				Right: &TreeNode{
					Val: 60,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			changeTree(tt.root)
			fmt.Printf("%+v\n", tt.root)
		})
	}
}
