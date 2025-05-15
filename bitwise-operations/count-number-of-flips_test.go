package main

import "testing"

func Test_minBitFlips(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		start int
		goal  int
		want  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minBitFlips(tt.start, tt.goal)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("minBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
