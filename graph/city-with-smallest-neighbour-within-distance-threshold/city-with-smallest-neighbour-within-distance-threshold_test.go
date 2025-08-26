package main

import "testing"

func Test_findTheCity(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n                 int
		edges             [][]int
		distanceThreshold int
		want              int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTheCity(tt.n, tt.edges, tt.distanceThreshold)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("findTheCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
