package main

import "testing"

func Test_myPow(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    float64
		n    int
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myPow(tt.x, tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
