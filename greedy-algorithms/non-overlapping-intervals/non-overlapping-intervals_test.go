package main

import "testing"

func Test_eraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		intervals [][]int
		want      int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := eraseOverlapIntervals(tt.intervals)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
