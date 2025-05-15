package main

import "testing"

func Test_countPrimes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPrimes(tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
