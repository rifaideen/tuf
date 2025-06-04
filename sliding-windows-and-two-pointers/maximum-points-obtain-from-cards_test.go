package main

import "testing"

func Test_maxScore(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cardPoints []int
		k          int
		want       int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxScore(tt.cardPoints, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
