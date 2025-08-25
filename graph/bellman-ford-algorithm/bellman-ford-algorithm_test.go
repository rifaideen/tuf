package main

import "testing"

func Test_bellmanFord(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		v     int
		edges [][]int
		src   int
		want  []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bellmanFord(tt.v, tt.edges, tt.src)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("bellmanFord() = %v, want %v", got, tt.want)
			}
		})
	}
}
