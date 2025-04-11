package main

import "testing"

func Test_reverseList(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.head)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseListRecursive(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseListRecursive(tt.head)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("reverseListRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
