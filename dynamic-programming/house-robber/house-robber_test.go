package main

import "testing"

func Test_rob(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
