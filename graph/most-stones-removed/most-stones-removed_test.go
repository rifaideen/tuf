package main

import "testing"

func Test_removeStones(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		stones [][]int
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeStones(tt.stones)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("removeStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
