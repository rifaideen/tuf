package main

import "testing"

func Test_deleteNode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		key  int
		want *TreeNode
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteNode(tt.root, tt.key)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
