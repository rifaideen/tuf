package main

import "testing"

func Test_longestOnes3(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestOnes3(tt.nums, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("longestOnes3() = %v, want %v", got, tt.want)
			}
		})
	}
}
