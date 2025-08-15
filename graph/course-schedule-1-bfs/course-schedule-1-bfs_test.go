package main

import "testing"

func Test_canFinish(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canFinish(tt.numCourses, tt.prerequisites)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
