package main

import "testing"

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
