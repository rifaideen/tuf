package main

import "testing"

func Test_countGoodNumbers(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int64
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countGoodNumbers(tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("countGoodNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
