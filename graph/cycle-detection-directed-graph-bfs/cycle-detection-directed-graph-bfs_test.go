package main

import "testing"

func Test_detectCycle(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		vertices int
		adjList  [][]int
		want     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectCycle(tt.vertices, tt.adjList)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
