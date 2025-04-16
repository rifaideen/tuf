package main

import "testing"

func Test_rotateRight(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		k    int
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateRight(tt.head, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
