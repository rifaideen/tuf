package main

import "testing"

func Test_solve(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		board [][]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.board)
		})
	}
}
