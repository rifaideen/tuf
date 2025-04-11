package main

import "testing"

func Test_hasCycle(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasCycle(tt.head)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
