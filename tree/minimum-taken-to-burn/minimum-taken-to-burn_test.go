package main

import "testing"

func Test_minTimeToBurn(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root   *TreeNode
		target *TreeNode
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minTimeToBurn(tt.root, tt.target)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("minTimeToBurn() = %v, want %v", got, tt.want)
			}
		})
	}
}
