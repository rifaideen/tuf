package main

import "testing"

func Test_findTarget(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		k    int
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTarget(tt.root, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
