package main

import "testing"

func Test_searchBST(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		val  int
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchBST(tt.root, tt.val)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
