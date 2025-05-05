package main

import "testing"

func Test_exist(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		board [][]byte
		word  string
		want  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exist(tt.board, tt.word)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
