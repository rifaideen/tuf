package main

import "testing"

func Test_bstFromPreorder(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		preorder []int
		want     *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bstFromPreorder(tt.preorder)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
