package main

import "testing"

func Test_isNStraightHand(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		hand      []int
		groupSize int
		want      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNStraightHand(tt.hand, tt.groupSize)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("isNStraightHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
