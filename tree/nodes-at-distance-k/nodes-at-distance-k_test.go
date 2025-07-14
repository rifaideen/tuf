package main

import "testing"

func Test_distanceK(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root   *TreeNode
		target *TreeNode
		k      int
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := distanceK(tt.root, tt.target, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("distanceK() = %v, want %v", got, tt.want)
			}
		})
	}
}
