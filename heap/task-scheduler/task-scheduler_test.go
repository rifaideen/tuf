package main

import "testing"

func Test_leastInterval(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		tasks []byte
		n     int
		want  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leastInterval(tt.tasks, tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
