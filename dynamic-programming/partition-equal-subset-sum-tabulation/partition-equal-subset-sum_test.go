package main

import "testing"

func Test_canPartition(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canPartition(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
