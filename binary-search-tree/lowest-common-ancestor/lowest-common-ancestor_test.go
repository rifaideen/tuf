package main

import "testing"

func Test_lowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowestCommonAncestor(tt.root, tt.p, tt.q)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
