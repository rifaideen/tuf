package main

import "testing"

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		l1   *Node
		l2   *Node
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.l1, tt.l2)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
