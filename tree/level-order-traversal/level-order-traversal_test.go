package main

import "testing"

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(tt.root)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
