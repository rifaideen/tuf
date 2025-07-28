package main

import "testing"

func Test_largestBst(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestBst(tt.root)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("largestBst() = %v, want %v", got, tt.want)
			}
		})
	}
}
